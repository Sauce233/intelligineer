package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListAttendance(c *Context, u *struct{}, r *struct {
	Name      string        `form:"name"`
	TypeID    uint          `form:"typeId"`
	TimeRange *[2]time.Time `form:"timeRange[]"`
	Page      int           `form:"page"`
	Size      int           `form:"size"`
}) (string, *List) {

	type Attendance struct {
		ID        uint           `json:"id"`
		Name      string         `json:"name"`
		Location  string         `json:"location"`
		TypeID    uint           `json:"-"`
		Type      AttendanceType `json:"type"`
		StartTime time.Time      `json:"startTime"`
		EndTime   time.Time      `json:"endTime"`
		UserCount uint           `json:"userCount"`
	}

	q := gorm.G[Attendance](c.DB).Table("attendances a").Select(
		"*, (select count(*) from user_attendances ua where ua.attendance_id = a.id) as user_count",
	)

	if r.Name != "" {
		q = q.Where("MATCH(name) AGAINST(?)", r.Name)
	}

	if r.TypeID != 0 {
		q = q.Where("type_id = ?", r.TypeID)
	}

	if r.TimeRange != nil {
		q = q.Where("start_time <= ? AND end_time >= ?", r.TimeRange[1], r.TimeRange[0])
	}

	total, err := q.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "获取签到数量失败", nil
	}

	attendances, err := q.Scopes(PreloadAll(1), Paginate(r.Page, r.Size, 9)).Order("start_time desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取签到列表失败", nil
	}

	return "", NewList(attendances, total)
}

func GetAttendance(c *Context, u *struct {
	ID uint `uri:"attendanceId"`
}, r *struct{}) (string, any) {

	type Attendance struct {
		ID        uint           `json:"id"`
		CreatedAt time.Time      `json:"createdAt"`
		Name      string         `json:"name"`
		Location  string         `json:"location"`
		TypeID    uint           `json:"-"`
		Type      AttendanceType `json:"type"`
		StartTime time.Time      `json:"startTime"`
		EndTime   time.Time      `json:"endTime"`
	}

	attendance, err := gorm.G[Attendance](c.DB).Scopes(PreloadAll(1)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到这个签到", nil
	} else if err != nil {
		c.Error(err)
		return "获取签到失败", nil
	}

	return "", &attendance
}

func ListAttendanceUser(c *Context, u *struct {
	AttendanceID uint `uri:"attendanceId"`
}, r *struct {
	StatusID uint `form:"statusId"`
	Page     int  `form:"page"`
	Size     int  `form:"size"`
}) (string, *List) {

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type UserAttendance struct {
		UserID   uint                 `json:"userId"`
		User     User                 `json:"user"`
		Time     *time.Time           `json:"time"`
		StatusID uint                 `json:"-"`
		Status   UserAttendanceStatus `json:"status"`
	}

	q := gorm.G[UserAttendance](c.DB).Table("user_attendances ua").Where("attendance_id = ?", u.AttendanceID)

	if r.StatusID != 0 {
		q = q.Where("status_id = ?", r.StatusID)
	}

	total, err := q.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "获取用户签到数量失败", nil
	}

	userAttendances, err := q.Scopes(PreloadAll(2), Paginate(r.Page, r.Size, 12)).Order("time desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取用户签到列表失败", nil
	}

	return "", NewList(userAttendances, total)
}
