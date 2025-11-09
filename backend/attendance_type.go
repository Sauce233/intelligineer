package main

import "gorm.io/gorm"

func ListAttendanceType(c *Context, u *struct{}, r *struct{}) (string, any) {

	type AttendanceType struct {
		ID              uint   `json:"id"`
		Name            string `json:"name"`
		Color           string `json:"color"`
		AttendanceCount uint   `json:"attendanceCount"`
	}

	attendanceTypes, err := gorm.G[AttendanceType](c.DB).Table("attendance_types at").Select(
		"*, (select count(*) from attendances a where a.type_id = at.id) as attendance_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "獲取打卡列表失敗", nil
	}

	return "", attendanceTypes
}
