package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListClient(c *Context, u *struct{}, r *struct {
	Name    string `form:"name"`
	Address string `form:"address"`
	Page    int    `form:"page"`
	Size    int    `form:"size"`
}) (message string, resp *List) {

	type ClientContact struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		ClientID uint   `json:"-"`
	}

	type Client struct {
		ID           uint            `json:"id"`
		Name         string          `json:"name"`
		Address      string          `json:"address"`
		ContactCount uint            `json:"contactCount"`
		ProjectCount uint            `json:"projectCount"`
		Contacts     []ClientContact `json:"contacts"`
	}

	query := gorm.G[Client](c.DB).Table("clients c").Select(`*,
		(select count(*) from client_contacts cc where cc.client_id = c.id) as contact_count,
		(select count(*) from projects p where p.client_id = c.id) as project_count`)

	if r.Name != "" {
		query = query.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if r.Address != "" {
		query = query.Where("address LIKE ?", "%"+r.Address+"%")
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计客户数量失败", nil
	}

	clients, err := query.Scopes(Paginate(r.Page, r.Size, 9), PreloadAll(1)).Order("id desc").Find(c.Context)
	if err != nil {
		c.Error(err)
		return "获取客户列表失败", nil
	}

	return "", NewList(clients, total)
}

func ListClientProject(c *Context, u *struct {
	ClientID uint `uri:"clientId"`
}, r *struct {
	Name     string `form:"name"`
	StatusID uint   `form:"statusId"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}) (message string, resp *List) {

	type Project struct {
		ID              uint          `json:"id"`
		CreatedAt       time.Time     `json:"createdAt"`
		UpdatedAt       time.Time     `json:"updatedAt"`
		Name            string        `json:"name"`
		Profile         string        `json:"profile"`
		Offer           float64       `json:"offer"`
		Budget          float64       `json:"budget"`
		StatusID        uint          `json:"-"`
		Status          ProjectStatus `json:"status"`
		AttachmentCount uint          `json:"attachmentCount"`
		InvoiceCount    uint          `json:"invoiceCount"`
		TaskCount       uint          `json:"taskCount"`
	}

	query := gorm.G[Project](c.DB).Table("projects p").Scopes(PreloadAll(2)).Select(`*, 
		(select count(*) from invoices i where i.project_id = p.id) as invoice_count, 
		(select count(*) from tasks t where t.project_id = p.id) as task_count, 
		(select count(*) from attachments a where a.project_id = p.id) as attachment_count`).Where("p.client_id = ?", u.ClientID)

	if r.Name != "" {
		query = query.Where("MATCH(p.name) AGAINST(?)", r.Name)
	}

	if r.StatusID != 0 {
		query = query.Where("p.status_id = ?", r.StatusID)
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计项目数量失败", nil
	}

	projects, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Order("p.id desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取项目列表失败", nil
	}

	return "", NewList(projects, total)
}

func AddClient(c *Context, u *struct{}, r *struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Profile string `json:"profile"`
}) (message string, resp uint) {

	client := Client{
		Name:    r.Name,
		Address: r.Address,
		Profile: r.Profile,
	}

	if err := gorm.G[Client](c.DB).Create(c.Context, &client); err != nil {
		c.Error(err)
		return "添加客户失败", 0
	}

	return "添加客户成功", client.ID
}

func GetClient(c *Context, u *struct {
	ID uint `uri:"clientId"`
}, r *struct{}) (message string, resp any) {

	type ClientContact struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		Name      string    `json:"name"`
		Phone     string    `json:"phone"`
		Email     string    `json:"email"`
		Profile   string    `json:"profile"`
		ClientID  uint      `json:"-"`
	}

	type ProjectStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Client struct {
		ID        uint            `json:"id"`
		CreatedAt time.Time       `json:"createdAt"`
		UpdatedAt time.Time       `json:"updatedAt"`
		Name      string          `json:"name"`
		Address   string          `json:"address"`
		Profile   string          `json:"profile"`
		Contacts  []ClientContact `json:"contacts"`
	}

	client, err := gorm.G[Client](c.DB).Preload("Contacts", nil).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "没有这个客户", nil
	} else if err != nil {
		c.Error(err)
		return "查询客户失败", nil
	}

	return "", &client
}

func EditClient(c *Context, u *struct {
	ID uint `uri:"clientId"`
}, r *struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Profile string `json:"profile"`
}) (message string, resp bool) {

	if row, err := gorm.G[Client](c.DB).Where("id = ?", u.ID).Select(
		"name", "address", "profile",
	).Updates(
		c, Client{Name: r.Name, Address: r.Address, Profile: r.Profile},
	); err != nil {
		c.Error(err)
		return "更新客户失败", false
	} else if row == 0 {
		return "找不到这个客户", false
	}

	return "更新客户成功", true
}

func DeleteClient(c *Context, u *struct {
	ID uint `uri:"clientId"`
}, r *struct{}) (message string, resp any) {

	if row, err := gorm.G[Client](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "这个客户存在未清空的资源", false
	} else if row == 0 {
		return "没有这个客户", false
	} else if err != nil {
		c.Error(err)
		return "客户删除失败", false
	}

	return "删除成功", true
}
