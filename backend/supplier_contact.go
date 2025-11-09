package main

import (
	"errors"

	"gorm.io/gorm"
)

func AddSupplierContact(c *Context, u *struct {
	SupplierID uint `uri:"supplierId"`
}, r *struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}) (string, bool) {

	if err := gorm.G[SupplierContact](c.DB).Create(c, &SupplierContact{
		Name:       r.Name,
		Phone:      r.Phone,
		Email:      r.Email,
		Profile:    r.Profile,
		SupplierID: u.SupplierID,
	}); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在这个供应商", false
	} else if err != nil {
		c.Error(err)
		return "联系人添加失败", false
	}

	return "联系人添加成功", true
}

func EditSupplierContact(c *Context, u *struct {
	ID uint `uri:"supplierContactId"`
}, r *struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Profile string `json:"profile"`
}) (string, bool) {

	if row, err := gorm.G[SupplierContact](c.DB).Select(
		"name, phone, email, profile",
	).Updates(
		c, SupplierContact{Name: r.Name, Phone: r.Phone, Email: r.Email, Profile: r.Profile},
	); err != nil {
		c.Error(err)
		return "更新联系人失败", false
	} else if row == 0 {
		return "不存在这个联系人", false
	}

	return "更新联系人成功", true
}

func DeleteSupplierContact(c *Context, u *struct {
	ID uint `uri:"supplierContactId"`
}, r *struct{}) (string, bool) {

	if row, err := gorm.G[SupplierContact](c.DB).Where("id = ?", u.ID).Delete(c); row == 0 {
		return "不存在这个联系人", false
	} else if err != nil {
		c.Error(err)
		return "删除联系人失败", false
	}

	return "删除联系人成功", true
}
