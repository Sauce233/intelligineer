package main

import "gorm.io/gorm"

func ListApplicationStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type ApplicationStatus struct {
		ID               uint   `json:"id"`
		Name             string `json:"name"`
		Color            string `json:"color"`
		ApplicationCount uint   `json:"applicationCount"`
	}

	applicationStatuses, err := gorm.G[ApplicationStatus](c.DB).Table("application_statuses `as`").Select(
		"*, (select count(*) from applications a where a.status_id = `as`.id) as application_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "獲取申請狀態列表失敗", nil
	}

	return "", applicationStatuses
}
