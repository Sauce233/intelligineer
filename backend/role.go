package main

import (
	"errors"

	"gorm.io/gorm"
)

func ListRole(c *Context, u *struct{}, r *struct{}) (string, any) {

	type Role struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		Color     string `json:"color"`
		UserCount uint   `json:"userCount"`
	}

	roles, err := gorm.G[Role](c.DB).Table("roles r").Select(
		"*, (select count(*) from user_roles ur where r.id = ur.role_id) as user_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取角色列表失败", nil
	}

	return "", roles
}

func AddRole(c *Context, u *struct {
	ID uint `uri:"id"`
}, r *struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}) (string, bool) {

	role := Role{
		ID:    u.ID,
		Name:  r.Name,
		Color: r.Color,
	}

	if row, err := gorm.G[Role](c.DB).Where("id = ?", u.ID).Updates(c, role); row == 0 {
		if err := gorm.G[Role](c.DB).Create(c, &role); err != nil {
			c.Error(err)
			return "角色创建失败", false
		}
	} else if err != nil {
		c.Error(err)
		return "角色更新失败", false
	}

	return "角色创建成功", true
}

func DeleteRole(c *Context, u *struct {
	ID string `uri:"roleId"`
}, r *struct{}) (string, bool) {

	if row, err := gorm.G[Role](c.DB).Where("id = ?", u.ID).Delete(c); row == 0 {
		return "找不到这个角色", false
	} else if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "请先删除这个状态下的所有角色", false
	} else if err != nil {
		c.Error(err)
		return "角色删除失败", false
	}

	return "角色删除成功", true
}
