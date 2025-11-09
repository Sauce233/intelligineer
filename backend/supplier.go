package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListSupplier(c *Context, u *struct{}, r *struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Page    int    `form:"page"`
	Size    int    `form:"size"`
}) (message string, resp *List) {

	type Supplier struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Address       string `json:"address"`
		ContactCount  uint   `json:"contactCount"`
		PurchaseCount uint   `json:"purchaseCount"`
	}

	query := gorm.G[Supplier](c.DB).Table("suppliers s").Select(`*,
		(select count(*) from supplier_contacts as sc where sc.supplier_id = s.id) as contact_count,
		(select count(*) from purchases as p where p.supplier_id = s.id) as purchase_count`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.Address != "" {
		query = query.Where("address LIKE ?", "%"+r.Address+"%")
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计供应商数量失败", nil
	}

	suppliers, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Order("id desc").Find(c.Context)
	if err != nil {
		c.Error(err)
		return "获取供应商列表失败", nil
	}

	return "", NewList(suppliers, total)
}

func AddSupplier(c *Context, u *struct{}, r *struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Profile string `json:"profile"`
}) (message string, resp bool) {

	if err := gorm.G[Supplier](c.DB).Create(c.Context, &Supplier{
		Name:    r.Name,
		Address: r.Address,
		Profile: r.Profile,
	}); err != nil {
		c.Error(err)
		return "添加供应商失败", false
	}

	return "添加供应商成功", true
}

func GetSupplier(c *Context, u *struct {
	ID uint `uri:"supplierId"`
}, r *struct{}) (message string, resp any) {

	type Purchase struct {
		ID         uint           `json:"id"`
		Name       string         `json:"name"`
		SupplierID uint           `json:"-"`
		StatusID   uint           `json:"-"`
		Status     PurchaseStatus `json:"status"`
	}

	type SupplierContact struct {
		ID         uint   `json:"id"`
		Name       string `json:"name"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Profile    string `json:"profile"`
		SupplierID uint   `json:"-"`
	}

	type Supplier struct {
		ID        uint              `json:"id"`
		Name      string            `json:"name"`
		Address   string            `json:"address"`
		Profile   string            `json:"profile"`
		Contacts  []SupplierContact `json:"contacts"`
		Purchases []Purchase        `json:"purchases"`
	}

	supplier, err := gorm.G[Supplier](c.DB).Scopes(PreloadAll(2)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "没有这个供应商", nil
	} else if err != nil {
		c.Error(err)
		return "查询供应商失败", nil
	}

	return "", &supplier
}

func ListSupplierContact(c *Context, u *struct {
	SupplierID uint `uri:"supplierId"`
}, r *struct {
	Page int `form:"page"`
	Size int `form:"size"`
}) (string, any) {

	type SupplierContact struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Profile string `json:"profile"`
	}

	query := gorm.G[SupplierContact](c.DB).Scopes()

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计供应商联系人数量失败", nil
	}

	contacts, err := query.Scopes(Paginate(r.Page, r.Size, 12)).Where("supplier_id = ?", u.SupplierID).Find(c)
	if err != nil {
		c.Error(err)
		return "获取供应商联系人列表失败", nil
	}

	return "", NewList(contacts, total)
}

func ListSupplierPurchase(c *Context, u *struct {
	SupplierID uint `uri:"supplierId"`
}, r *struct {
	Name       string `form:"name"`
	MaterialID uint   `form:"materialId"`
	StatusID   uint   `form:"statusId"`
	Sort       string `form:"sort"`
	Order      string `form:"order"`
	Page       int    `form:"page"`
	Size       int    `form:"size"`
}) (message string, resp *List) {

	type Purchase struct {
		ID            uint           `json:"id"`
		CreatedAt     time.Time      `json:"createdAt"`
		UpdatedAt     time.Time      `json:"updatedAt"`
		Name          string         `json:"name"`
		Profile       string         `json:"profile"`
		MaterialCount uint           `json:"materialCount"`
		TotalPrice    string         `json:"totalPrice"`
		StatusID      uint           `json:"-"`
		Status        PurchaseStatus `json:"status"`
	}

	query := gorm.G[Purchase](c.DB).Table("purchases p").Where("supplier_id = ?", u.SupplierID).Select(`*,
	(select count(*) from purchase_materials pm where pm.purchase_id = p.id) as material_count,
	ROUND((SELECT SUM(pm.quantity * pm.price) FROM purchase_materials pm WHERE pm.purchase_id = p.id), 2) AS total_price`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.StatusID != 0 {
		query = query.Where("status_id = ?", r.StatusID)
	}

	if r.MaterialID != 0 {
		query = query.Where("id IN (?)", gorm.G[PurchaseMaterial](c.DB).Select("purchase_id").Where("material_id = ?", r.MaterialID))
	}

	if r.Order != "" && r.Sort != "" {
		query = query.Order(r.Sort + " " + r.Order)
	}

	total, err := query.Count(c.Context, "*")
	if err != nil {
		c.Error(err)
		return "统计订单数量失败", nil
	}

	purchases, err := query.Scopes(Paginate(r.Page, r.Size, 9), PreloadAll(1)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取供应商订单列表失败", nil
	}

	return "", NewList(purchases, total)
}

func DeleteSupplier(c *Context, u *struct {
	ID uint `uri:"supplierId"`
}, r *struct{}) (message string, resp any) {

	if row, err := gorm.G[Supplier](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "这个供应商存在未清空的资源", false
	} else if row == 0 {
		return "没有这个供应商", false
	} else if err != nil {
		c.Error(err)
		return "供应商删除失败", false
	}

	return "删除成功", true
}

func EditSupplier(c *Context, u *struct {
	ID uint `uri:"supplierId"`
}, r *struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Profile string `json:"profile"`
}) (message string, resp bool) {

	if row, err := gorm.G[Supplier](c.DB).Where("id = ?", u.ID).Select(
		"name, address, profile",
	).Updates(
		c, Supplier{Name: r.Name, Address: r.Address, Profile: r.Profile},
	); err != nil {
		c.Error(err)
		return "更新供应商失败", false
	} else if row == 0 {
		return "找不到这个供应商", false
	}

	return "更新供应商成功", true
}
