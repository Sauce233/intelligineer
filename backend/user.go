package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ListUser(c *Context, u *struct{}, r *struct {
	RoleID uint   `form:"roleId"`
	Name   string `form:"name"`
	Email  string `form:"email"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}) (message string, resp any) {

	type Application struct {
		UserID uint
		Count  uint
	}

	type Role struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type User struct {
		ID               uint   `json:"id"`
		Name             string `json:"name"`
		Email            string `json:"email"`
		Profile          string `json:"profile"`
		ApplicationCount uint   `json:"applicationCount"`
		RecordCount      uint   `json:"recordCount"`
		Roles            []Role `json:"roles" gorm:"many2many:user_roles"`
	}

	query := gorm.G[User](c.DB).Table("users u").Preload("Roles", nil).Select(`*,
		(select count(*) from applications a where u.id = a.applicant_id) as application_count,
		(select count(*) from records r where u.id = r.user_id) as record_count`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.Email != "" {
		query = query.Where("email LIKE ?", "%"+r.Email+"%")
	}

	if r.RoleID != 0 {
		query = query.Where("id IN (?)", gorm.G[UserRole](c.DB).Select("user_id").Where("role_id = ?", r.RoleID))
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计员工数量失败", nil
	}

	users, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取员工列表失败", nil
	}

	return "", NewList(users, total)
}

func AddUser(c *Context, u *struct{}, r *struct {
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Profile  string `json:"profile"`
}) (message string, resp uint) {

	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		return "密码加密失败", 0
	}

	user := User{
		Password: string(password),
		Name:     r.Name,
		Email:    r.Email,
		Profile:  r.Profile,
	}

	if err := gorm.G[User](c.DB).Create(c, &user); errors.Is(err, gorm.ErrDuplicatedKey) {
		return "该邮箱已被使用", 0
	} else if err != nil {
		c.Error(err)
		return "用户创建失败", 0
	}

	return "用户创建成功", user.ID
}

func GetUser(c *Context, u *struct {
	ID uint `uri:"userId"`
}, r *struct{}) (message string, resp any) {

	type Role struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type User struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Profile   string    `json:"profile"`
		Roles     []Role    `json:"roles" gorm:"many2many:user_roles"`
	}

	user, err := gorm.G[User](c.DB).Preload("Roles", nil).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到对应的用户", nil
	} else if err != nil {
		c.Error(err)
		return "查找用户失败", nil
	}

	return "", &user
}

func ListUserRecord(c *Context, u *struct {
	UserID uint `uri:"userId"`
}, r *struct {
	Title string        `form:"title"`
	Time  *[2]time.Time `form:"time[]"`
	Page  int           `form:"page"`
	Size  int           `form:"size"`
}) (message string, resp any) {

	type Project struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Task struct {
		ID        uint    `json:"id"`
		Name      string  `json:"name"`
		ProjectID uint    `json:"-"`
		Project   Project `json:"project"`
	}

	type Record struct {
		ID         uint      `json:"id"`
		CreatedAt  time.Time `json:"createdAt"`
		TaskID     uint      `json:"-"`
		Task       Task      `json:"task"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		Minutes    uint      `json:"minutes"`
		PhotoCount uint      `json:"photoCount"`
	}

	query := gorm.G[Record](c.DB).Table("records r").Scopes(PreloadAll(2)).Select(`*,
		(select count(*) from photos p where p.record_id = r.id) as photo_count`).Where("user_id = ?", u.UserID)

	if r.Title != "" {
		query = query.Where("title LIKE ?", "%"+r.Title+"%")
	}

	if r.Time != nil {
		query = query.Where("created_at >= ? AND created_at <= ?", r.Time[0], r.Time[1])
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计施工记录数量失败", nil
	}

	records, err := query.Order("id desc").Scopes(Paginate(r.Page, r.Size, 9)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取施工记录列表失败", nil
	}

	return "", NewList(records, total)
}

func ListUserAttendance(c *Context, u *struct {
	UserID uint `uri:"userId"`
}, r *struct {
	StatusID uint `form:"statusId"`
	Page     int  `form:"page"`
	Size     int  `form:"size"`
}) (string, *List) {

	type Attendance struct {
		ID     uint           `json:"id"`
		Name   string         `json:"name"`
		TypeID uint           `json:"-"`
		Type   AttendanceType `json:"type"`
	}

	type UserAttendance struct {
		AttendanceID uint                 `json:"-"`
		Attendance   Attendance           `json:"attendance"`
		Time         *time.Time           `json:"time"`
		StatusID     uint                 `json:"-"`
		Status       UserAttendanceStatus `json:"status"`
	}

	q := gorm.G[UserAttendance](c.DB).Table("user_attendances ua").Where("user_id = ?", u.UserID)

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

func EditUser(c *Context, u *struct {
	ID uint `uri:"userId"`
}, r *struct {
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Profile  string `json:"profile"`
}) (message string, resp *User) {

	password, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		return "密码加密失败", nil
	}

	user := User{
		Password: string(password),
		Name:     r.Name,
		Email:    r.Email,
		Profile:  r.Profile,
	}

	if row, err := gorm.G[User](c.DB).Where("id = ?", u.ID).Select(
		"password, name, email, profile",
	).Updates(c, user); errors.Is(err, gorm.ErrDuplicatedKey) {
		return "这个邮箱已被其他用户使用", nil
	} else if err != nil {
		c.Error(err)
		return "用户更新失败", nil
	} else if row == 0 {
		return "找不到这个用户", nil
	}

	return "用户信息更新成功", &user
}

func DeleteUser(c *Context, u *struct {
	ID uint `uri:"userId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[User](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "用户删除失败", false
	} else if row == 0 {
		return "找不到对应的用户", false
	}

	return "用户删除成功", true
}

func StatUserHours(c *Context, u *struct {
	UserID uint `uri:"userId"`
	Days   int  `uri:"days"`
}, r *struct{}) (string, any) {

	var data []struct {
		Date  string `json:"date"`
		Total uint   `json:"total"`
	}

	err := c.DB.Select("w.date as date, COALESCE(w.total, 0) as total").Table(GetSQLUintList(u.Days)+" d").Joins(
		"LEFT JOIN (?) w ON w.date = DATE_SUB(CURDATE(), INTERVAL d.n DAY)",
		c.DB.Table("records").Select("DATE(created_at) as date, SUM(minutes) AS total").Where("user_id = ?", u.UserID).Group("date"),
	).Find(&data)
	if err != nil {
		fmt.Println(err)
	}

	return "", data
}
