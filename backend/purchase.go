package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListPurchase(c *Context, u *struct{}, r *struct {
	Name       string `form:"name"`
	MaterialID uint   `form:"materialId"`
	SupplierID uint   `form:"supplierId"`
	StatusID   uint   `form:"statusId"`
	Sort       string `form:"sort"`
	Order      string `form:"order"`
	Page       int    `form:"page"`
	Size       int    `form:"size"`
}) (message string, resp *List) {

	type Supplier struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Purchase struct {
		ID            uint           `json:"id"`
		CreatedAt     time.Time      `json:"createdAt"`
		UpdatedAt     time.Time      `json:"updatedAt"`
		Name          string         `json:"name"`
		Profile       string         `json:"profile"`
		MaterialCount uint           `json:"materialCount"`
		TotalPrice    string         `json:"totalPrice"`
		SupplierID    uint           `json:"-"`
		Supplier      Supplier       `json:"supplier"`
		StatusID      uint           `json:"-"`
		Status        PurchaseStatus `json:"status"`
	}

	query := gorm.G[Purchase](c.DB).Table("purchases p").Select(`*,
	(select count(*) from purchase_materials pm where pm.purchase_id = p.id) as material_count,
	ROUND((SELECT SUM(pm.quantity * pm.price) FROM purchase_materials pm WHERE pm.purchase_id = p.id), 2) AS total_price`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.StatusID != 0 {
		query = query.Where("status_id = ?", r.StatusID)
	}

	if r.SupplierID != 0 {
		query = query.Where("supplier_id = ?", r.SupplierID)
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

func AddPurchase(c *Context, u *struct {
	SupplierID uint `uri:"supplierId"`
}, r *struct {
	Name     string `json:"name"`
	Profile  string `json:"profile"`
	StatusID uint   `json:"statusId"`
}) (message string, resp bool) {

	if err := gorm.G[Purchase](c.DB).Create(c, &Purchase{
		SupplierID: u.SupplierID,
		Name:       r.Name,
		Profile:    r.Profile,
		StatusID:   r.StatusID,
	}); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "没有这个供应商", false
	} else if err != nil {
		c.Error(err)
		return "添加订单失败", false
	}

	return "添加订单成功", true
}

func GetPurchase(c *Context, u *struct {
	ID uint `uri:"purchaseId"`
}, r *struct{}) (message string, resp any) {

	type Category struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Material struct {
		ID         uint     `json:"id"`
		Name       string   `json:"name"`
		Quantity   float64  `json:"quantity"`
		Unit       string   `json:"unit"`
		CategoryID uint     `json:"-"`
		Category   Category `json:"category"`
	}

	type PurchaseMaterial struct {
		PurchaseID uint     `json:"-"`
		MaterialID uint     `json:"-"`
		Material   Material `json:"material"`
		Quantity   float64  `json:"quantity"`
		Price      float64  `json:"price"`
	}

	type Supplier struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Purchase struct {
		ID         uint               `json:"id"`
		CreatedAt  time.Time          `json:"createdAt"`
		UpdatedAt  time.Time          `json:"updatedAt"`
		Name       string             `json:"name"`
		Profile    string             `json:"profile"`
		SupplierID uint               `json:"-"`
		Supplier   Supplier           `json:"supplier"`
		StatusID   uint               `json:"-"`
		Status     PurchaseStatus     `json:"status"`
		Materials  []PurchaseMaterial `json:"materials"`
	}

	purchase, err := gorm.G[Purchase](c.DB).Scopes(PreloadAll(3)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "不存在此供应商订单", nil
	} else if err != nil {
		c.Error(err)
		return "查找供应商订单失败", nil
	}

	return "", &purchase
}

func EditPurchase(c *Context, u *struct {
	ID uint `uri:"purchaseId"`
}, r *struct {
	Name     string `json:"name"`
	Profile  string `json:"profile"`
	StatusID uint   `json:"statusId"`
}) (message string, resp bool) {

	if row, err := gorm.G[Purchase](c.DB).Where("id = ?", u.ID).Select(
		"name, profile, paid, received",
	).Updates(c, Purchase{
		Name:     r.Name,
		Profile:  r.Profile,
		StatusID: r.StatusID,
	}); err != nil {
		c.Error(err)
		return "更新供应商订单失败", false
	} else if row == 0 {
		return "找不到这个供应商订单", false
	}

	return "更新供应商订单成功", true
}

func DeletePurchase(c *Context, u *struct {
	ID uint `uri:"purchaseId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Purchase](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "删除供应商订单失败", false
	} else if row == 0 {
		return "找不到这个供应商订单", false
	}

	return "供应商订单删除成功", true
}
