package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListMaterial(c *Context, u *struct{}, r *struct {
	Name       string `form:"name"`
	CategoryID uint   `form:"categoryId"`
	StatusID   uint   `form:"statusId"`
	Page       int    `form:"page"`
	Size       int    `form:"size"`
}) (message string, resp any) {

	type Category struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Material struct {
		ID            uint           `json:"id"`
		UpdatedAt     time.Time      `json:"updatedAt"`
		CategoryID    uint           `json:"-"`
		Category      Category       `json:"category"`
		StatusID      uint           `json:"-"`
		Status        MaterialStatus `json:"status"`
		Name          string         `json:"name"`
		Quantity      float64        `json:"quantity"`
		Unit          string         `json:"unit"`
		Profile       string         `json:"profile"`
		TaskCount     uint           `json:"taskCount"`
		PurchaseCount uint           `json:"purchaseCount"`
	}

	query := gorm.G[Material](c.DB).Table("materials m").Scopes(PreloadAll(1)).Select(`*,
	 (select count(*) from purchase_materials pm where pm.material_id = m.id) as purchase_count,
	 (select count(*) from task_materials tm where tm.material_id = m.id) as task_count`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.StatusID != 0 {
		query = query.Where("status_id = ?", r.StatusID)
	}

	if r.CategoryID != 0 {
		query = query.Where("category_id = ?", r.CategoryID)
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计库存数量失败", nil
	}

	materials, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Order("id desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取库存列表失败", nil
	}

	return "", NewList(materials, total)
}

func AddMaterial(c *Context, u *struct {
	CategoryID uint `uri:"categoryId"`
}, r *struct {
	StatusID uint    `json:"statusId"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Profile  string  `json:"profile"`
}) (message string, resp bool) {

	if err := gorm.G[Material](c.DB).Create(c, &Material{
		CategoryID: u.CategoryID,
		StatusID:   r.StatusID,
		Name:       r.Name,
		Quantity:   r.Quantity,
		Unit:       r.Unit,
		Profile:    r.Profile,
	}); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在这个材料类型或材料状态", false
	} else if err != nil {
		c.Error(err)
		return "添加材料失败", false
	}

	return "添加材料成功", true
}

func GetMaterial(c *Context, u *struct {
	ID uint `uri:"materialId"`
}, r *struct{}) (message string, resp any) {

	type Category struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Profile string `json:"profile"`
	}

	type MaterialStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Material struct {
		ID         uint           `json:"id"`
		UpdatedAt  time.Time      `json:"updatedAt"`
		CategoryID uint           `json:"-"`
		Category   Category       `json:"category"`
		StatusID   uint           `json:"-"`
		Status     MaterialStatus `json:"status"`
		Name       string         `json:"name"`
		Quantity   float64        `json:"quantity"`
		Unit       string         `json:"unit"`
		Profile    string         `json:"profile"`
	}

	material, err := gorm.G[Material](c.DB).Scopes(PreloadAll(1)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到这个材料", nil
	} else if err != nil {
		c.Error(err)
		return "查找材料失败", nil
	}

	return "", &material
}

func EditMaterial(c *Context, u *struct {
	ID uint `uri:"materialId"`
}, r *struct {
	CategoryID uint    `json:"categoryId"`
	StatusID   uint    `json:"statusId"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	Unit       string  `json:"unit"`
	Profile    string  `json:"profile"`
}) (message string, resp bool) {

	if row, err := gorm.G[Material](c.DB).Where("id = ?", u.ID).Select(
		"material_status_id, name, quantity, unit, profile",
	).Updates(c, Material{
		StatusID: r.StatusID,
		Name:     r.Name,
		Quantity: r.Quantity,
		Unit:     r.Unit,
		Profile:  r.Profile,
	}); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "这个材料状态不存在", false
	} else if err != nil {
		c.Error(err)
		return "修改材料失败", false
	} else if row == 0 {
		return "找不到对应的材料", false
	}

	return "修改材料成功", true
}

func DeleteMaterial(c *Context, u *struct {
	ID uint `uri:"materialId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Material](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "请先删除此材料与相关项目的关联", false
	} else if err != nil {
		c.Error(err)
		return "材料删除失败", false 
	} else if row == 0 {
		return "找不到这个材料", false
	}

	return "材料删除成功", true
}
