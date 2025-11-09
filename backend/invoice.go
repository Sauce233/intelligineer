package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListInvoice(c *Context, u *struct{}, r *struct {
	Title    string `form:"title"`
	StatusID uint   `form:"statusId"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}) (message string, resp any) {

	type Client struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Project struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		ClientID uint   `json:"-"`
		Client   Client `json:"client"`
	}

	type Invoice struct {
		ID        uint          `json:"id"`
		CreatedAt time.Time     `json:"createdAt"`
		ProjectID uint          `json:"projectId"`
		Project   Project       `json:"project"`
		Name      string        `json:"name"`
		Number    string        `json:"number"`
		Profile   string        `json:"profile"`
		Amount    float64       `json:"amount"`
		IssueDate time.Time     `json:"issueDate"`
		DueDate   time.Time     `json:"dueDate"`
		StatusID  uint          `json:"-"`
		Status    InvoiceStatus `json:"status"`
	}

	query := gorm.G[Invoice](c.DB).Scopes(PreloadAll(2))

	if r.Title != "" {
		query = query.Where("title LIKE ?", "%"+r.Title+"%")
	}

	if r.StatusID != 0 {
		query = query.Where("status_id = ?", r.StatusID)
	}

	total, err := query.Count(c.Context, "*")
	if err != nil {
		c.Error(err)
		return "统计发票总数失败", nil
	}

	invoices, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Order("id desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取发票列表失败", nil
	}

	return "", NewList(invoices, total)
}

func AddInvoice(c *Context, u *struct {
	ProjectID uint `uri:"projectId"`
}, r *struct {
	Title   string    `json:"title"`
	Profile string    `json:"profile"`
	Amount  float64   `json:"amount"`
	Paid    bool      `json:"paid"`
	Expiry  time.Time `json:"expiry"`
}) (message string, resp bool) {

	if err := gorm.G[Invoice](c.DB).Create(c.Context, &Invoice{
		Profile: r.Profile,
		Amount:  r.Amount,
	}); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在这个合同", false
	} else if err != nil {
		c.Error(err)
		return "添加发票失败", false
	}

	return "添加发票成功", true
}

func GetInvoice(c *Context, u *struct {
	ID uint `uri:"invoiceId"`
}, r *struct{}) (message string, resp any) {

	type Client struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Project struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		ClientID uint   `json:"-"`
		Client   Client `json:"client"`
	}

	type InvoiceStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Invoice struct {
		ID        uint          `json:"id"`
		CreatedAt time.Time     `json:"createdAt"`
		ProjectID uint          `json:"-"`
		Project   Project       `json:"project"`
		Title     string        `json:"title"`
		Number    string        `json:"number"`
		Profile   string        `json:"profile"`
		Amount    float64       `json:"amount"`
		IssueDate time.Time     `json:"issueDate"`
		DueDate   time.Time     `json:"dueDate"`
		StatusID  uint          `json:"-"`
		Status    InvoiceStatus `json:"status"`
	}

	invoice, err := gorm.G[Invoice](c.DB).Scopes(PreloadAll(2)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到这个发票", nil
	} else if err != nil {
		c.Error(err)
		return "查询发票失败", nil
	}

	return "", &invoice
}

func EditInvoice(c *Context, u *struct {
	ID uint `uri:"invoiceId"`
}, r *struct {
	Title   string    `json:"title"`
	Profile string    `json:"profile"`
	Amount  float64   `json:"amount"`
	Paid    bool      `json:"paid"`
	Expiry  time.Time `json:"expiry"`
}) (message string, resp bool) {

	if row, err := gorm.G[Invoice](c.DB).Where("id = ?", u.ID).Select(
		"title, profile, amount, paid, expiry",
	).Updates(c, Invoice{
		Profile: r.Profile,
		Amount:  r.Amount,
	}); err != nil {
		c.Error(err)
		return "更新发票失败", false
	} else if row == 0 {
		return "找不到这个发票", false
	}

	return "更新发票成功", true
}

func DeleteInvoice(c *Context, u *struct {
	ID uint `uri:"invoiceId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Invoice](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "删除发票失败", false
	} else if row == 0 {
		return "找不到这个发票", false
	}

	return "删除发票成功", false
}
