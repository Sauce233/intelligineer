package main

import "gorm.io/gorm"

func ListProjectStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type ProjectStatus struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		ProjectCount uint   `json:"projectCount"`
	}

	projectStatuses, err := gorm.G[ProjectStatus](c.DB).Table("project_statuses ps").Select(
		"*, (select count(*) from projects p where ps.id = p.status_id) as project_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取项目状态失败", nil
	}

	return "", projectStatuses
}
