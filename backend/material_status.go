package main

import (
	"errors"

	"gorm.io/gorm"
)

func ListMaterialStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type MaterialStatus struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Color         string `json:"color"`
		MaterialCount uint   `json:"materialCount"`
	}

	materialStatuses, err := gorm.G[MaterialStatus](c.DB).Table("material_statuses ms").Select(
		"*, (select count(*) from materials m where m.status_id = ms.id) as material_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取材料状态列表失败", nil
	}

	return "", materialStatuses
}

func AddMaterialStatus(c *Context, u *struct{}, r *struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}) (string, bool) {

	if err := gorm.G[MaterialStatus](c.DB).Create(c, &MaterialStatus{
		Name:  r.Name,
		Color: r.Color,
	}); err != nil {
		c.Error(err)
		return "材料状态添加失败", false
	}

	return "材料状态添加成功", true
}

func EditMaterialStatus(c *Context, u *struct {
	ID uint `uri:"id"`
}, r *struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}) (string, bool) {

	if row, err := gorm.G[MaterialStatus](c.DB).Where("id = ?", u.ID).Select(
		"name, color",
	).Updates(c, MaterialStatus{
		Name: r.Name, Color: r.Color,
	}); err != nil {
		c.Error(err)
		return "材料状态更新失败", false
	} else if row == 0 {
		return "找不到这个材料状态", false
	}

	return "材料状态创建成功", true
}

func DeleteMaterialStatus(c *Context, u *struct {
	ID string `uri:"id"`
}, r *struct{}) (string, bool) {

	if row, err := gorm.G[MaterialStatus](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "请先删除这个状态下的所有材料", false
	} else if err != nil {
		c.Error(err)
		return "材料状态删除失败", false
	} else if row == 0 {
		return "找不到这个材料状态", false
	}

	return "材料状态删除成功", true
}
