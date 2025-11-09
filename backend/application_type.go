package main

import (
	"gorm.io/gorm"
)

func ListApplicationType(c *Context, u *struct{}, r *struct{}) (string, any) {

	type ApplicationType struct {
		ID               uint   `json:"id"`
		Name             string `json:"name"`
		Color            string `json:"color"`
		ApplicationCount uint   `json:"applicationCount"`
	}

	applicationTypes, err := gorm.G[ApplicationType](c.DB).Table("application_types at").Select(
		"*, (select count(*) from applications a where a.type_id = at.id) as application_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取申请类型列表失败", nil
	}

	return "", applicationTypes
}
