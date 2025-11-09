package main

import (
	"gorm.io/gorm"
)

func ListProfile(c *Context, u *struct{}, r *struct {
	Name   string `form:"name"`
	TypeID uint   `form:"typeId"`
	UserID uint   `form:"userId"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}) (string, *List) {

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Profile struct {
		ID       uint        `json:"id"`
		Name     string      `json:"name"`
		TypeID   uint        `json:"-"`
		Type     ProfileType `json:"type"`
		UserID   uint        `json:"-"`
		User     User        `json:"user"`
		Profile  string      `json:"profile"`
		Filename string      `json:"filename"`
	}

	q := gorm.G[Profile](c.DB).Scopes()

	if r.Name != "" {
		q = q.Where("MATCH(name) AGAINST(?)", r.Name)
	}

	if r.TypeID != 0 {
		q = q.Where("type_id = ?", r.TypeID)
	}

	if r.UserID != 0 {
		q = q.Where("user_id = ?", r.UserID)
	}

	total, err := q.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "獲取資料數量失敗", nil
	}

	profiles, err := q.Scopes(PreloadAll(1), Paginate(r.Page, r.Size, 9)).Find(c)
	if err != nil {
		c.Error(err)
		return "獲取資料列表失敗", nil
	}

	return "", NewList(profiles, total)
}

func AddProfile(c *Context, u *struct {
	UserID uint `uri:"id"`
}, r *struct {
}) (string, uint) {

	return "資料添加成功", 0
}

func DeleteProfile(c *Context, u *struct{}, r *struct{}) (string, bool) {

	return "", true
}
