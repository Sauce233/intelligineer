package main

import (
	"errors"

	"gorm.io/gorm"
)

func ListClientContact(c *Context, u *struct {
	ClientID uint `uri:"clientId"`
}, r *struct {
	Page int `form:"page"`
	Size int `form:"size"`
}) (string, any) {

	type ClientContact struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Profile string `json:"profile"`
	}

	query := gorm.G[ClientContact](c.DB).Where("client_id = ?", u.ClientID)

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计客户联系人数量失败", nil
	}

	clientContacts, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取客户联系人列表失败", nil
	}

	return "", NewList(clientContacts, total)
}

func AddClientContact(c *Context, u *struct {
	ClientID uint `uri:"clientId"`
}, r *struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}) (string, uint) {

	clientContact := ClientContact{
		Name:     r.Name,
		Phone:    r.Phone,
		Email:    r.Email,
		Profile:  r.Profile,
		ClientID: u.ClientID,
	}

	if err := gorm.G[ClientContact](c.DB).Create(c, &clientContact); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在这个客户", 0
	} else if err != nil {
		c.Error(err)
		return "联系人添加失败", 0
	}

	return "联系人添加成功", clientContact.ID
}

func EditClientContact(c *Context, u *struct {
	ID uint `uri:"clientContactId"`
}, r *struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}) (string, *ClientContact) {

	clientContact := ClientContact{
		Name:    r.Name,
		Phone:   r.Phone,
		Email:   r.Email,
		Profile: r.Profile,
	}

	if row, err := gorm.G[ClientContact](c.DB).Select(
		"name, phone, email, profile",
	).Updates(c, clientContact); err != nil {
		c.Error(err)
		return "更新联系人失败", nil
	} else if row == 0 {
		return "不存在这个联系人", nil
	}

	return "更新联系人成功", &clientContact
}

func DeleteClientContact(c *Context, u *struct {
	ID uint `uri:"clientContactId"`
}, r *struct{}) (string, bool) {

	if row, err := gorm.G[ClientContact](c.DB).Where("id = ?", u.ID).Delete(c); row == 0 {
		return "不存在这个联系人", false
	} else if err != nil {
		c.Error(err)
		return "删除联系人失败", false
	}

	return "删除联系人成功", true
}
