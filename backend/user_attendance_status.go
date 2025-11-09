package main

import "gorm.io/gorm"

func ListUserAttendanceStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type UserAttendanceStatus struct {
		ID                  uint   `json:"id"`
		Name                string `json:"name"`
		Color               string `json:"color"`
		UserAttendanceCount uint   `json:"userAttendanceCount"`
	}

	userAttendanceStatuses, err := gorm.G[UserAttendanceStatus](c.DB).Table("user_attendance_statuses uas").Select(
		"*, (select count(*) from user_attendances ua where ua.status_id = uas.id) as user_attendance_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取项目状态失败", nil
	}

	return "", userAttendanceStatuses
}
