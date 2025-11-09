package main

import (
	"errors"

	"gorm.io/gorm"
)

func ListCategory(c *Context, u *struct{}, r *struct {
	Name string `form:"name"`
	Page int    `form:"page"`
	Size int    `form:"size"`
}) (string, *List) {

	type Material struct {
		CategoryID uint   `json:"-"`
		Name       string `json:"name"`
	}

	type Category struct {
		ID            uint       `json:"id"`
		Name          string     `json:"name"`
		Profile       string     `json:"profile"`
		MaterialCount uint       `json:"materialCount"`
		Materials     []Material `json:"materials"`
	}

	query := gorm.G[Category](c.DB).Table("categories c").Select(`*,
	(select count(*) from materials m where m.category_id = c.id) as material_count`).Preload("Materials", func(db gorm.PreloadBuilder) error {
		db.LimitPerRecord(3)
		return nil
	})

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计材料类型数量失败", nil
	}

	categories, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Order("id desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取材料类型列表失败", nil
	}

	return "", NewList(categories, total)
}

func AddCategory(c *Context, u *struct{}, r *struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
}) (string, uint) {

	category := Category{
		Name:    r.Name,
		Profile: r.Profile,
	}

	if err := gorm.G[Category](c.DB).Create(c, &category); err != nil {
		c.Error(err)
		return "材料类型创建失败", 0
	}

	return "材料类型创建成功", category.ID
}

func GetCategory(c *Context, u *struct {
	ID uint `uri:"categoryId"`
}, r *struct{}) (string, any) {

	type Category struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Profile string `json:"profile"`
	}

	category, err := gorm.G[Category](c.DB).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到这个材料类型", nil
	} else if err != nil {
		c.Error(err)
		return "查找材料类型失败", nil
	}

	return "", &category
}

func EditCategory(c *Context, u *struct {
	ID uint `uri:"categoryId"`
}, r *struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
}) (string, *Category) {

	category := Category{
		Name:    r.Name,
		Profile: r.Profile,
	}

	if row, err := gorm.G[Category](c.DB).Where("id = ?", u.ID).Select(
		"name, color, profile",
	).Updates(c, category); err != nil {
		c.Error(err)
		return "材料类型更新失败", nil
	} else if row == 0 {
		return "没有这个材料类型", nil
	}

	return "材料类型更新成功", &category
}

func DeleteCategory(c *Context, u *struct {
	ID uint `uri:"categoryId"`
}, r *struct{}) (string, bool) {

	if row, err := gorm.G[Category](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "请先删除这个材料类型下的材料", false
	} else if err != nil {
		c.Error(err)
		return "材料类型删除失败", false
	} else if row == 0 {
		return "找不到这个材料类型", false
	}

	return "材料类型删除成功", true
}
