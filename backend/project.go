package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListProject(c *Context, u *struct{}, r *struct {
	Name     string `form:"name"`
	StatusID uint   `form:"statusId"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}) (message string, resp *List) {

	type Client struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type ProjectStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type TaskStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		ID        uint       `json:"id"`
		Name      string     `json:"name"`
		StatusID  uint       `json:"-"`
		Status    TaskStatus `json:"status"`
		ProjectID uint       `json:"-"`
	}

	type Project struct {
		ID              uint          `json:"id"`
		CreatedAt       time.Time     `json:"createdAt"`
		UpdatedAt       time.Time     `json:"updatedAt"`
		Name            string        `json:"name"`
		ClientID        uint          `json:"-"`
		Client          Client        `json:"client"`
		Profile         string        `json:"profile"`
		Offer           float64       `json:"offer"`
		Budget          float64       `json:"budget"`
		StatusID        uint          `json:"-"`
		Status          ProjectStatus `json:"status"`
		AttachmentCount uint          `json:"attachmentCount"`
		InvoiceCount    uint          `json:"invoiceCount"`
		TaskCount       uint          `json:"taskCount"`
		Tasks           []Task        `json:"tasks"`
	}

	query := gorm.G[Project](c.DB).Table("projects p").Scopes(PreloadAll(2)).Select(`*, 
		(select count(*) from invoices i where i.project_id = p.id) as invoice_count, 
		(select count(*) from tasks t where t.project_id = p.id) as task_count, 
		(select count(*) from attachments a where a.project_id = p.id) as attachment_count`)

	if r.Name != "" {
		query = query.Where("p.name LIKE ?", "%"+r.Name+"%")
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

func AddProject(c *Context, u *struct {
	ClientID uint `uri:"clientId"`
}, r *struct {
	Name     string  `json:"name"`
	Profile  string  `json:"profile"`
	Offer    float64 `json:"offer"`
	Budget   float64 `json:"budget"`
	StatusID uint    `json:"statusId"`
}) (message string, resp uint) {

	project := Project{
		Name:     r.Name,
		Profile:  r.Profile,
		Offer:    r.Offer,
		Budget:   r.Budget,
		StatusID: r.StatusID,
		ClientID: u.ClientID,
	}

	if err := gorm.G[Project](c.DB).Create(c, &project); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在这个客户", 0
	} else if err != nil {
		c.Error(err)
		return "创建项目失败", 0
	}

	return "项目创建成功", project.ID
}

func GetProject(c *Context, u *struct {
	ID uint `uri:"projectId"`
}, r *struct{}) (message string, resp any) {

	type Client struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type ProjectStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Attachment struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		Filename  string `json:"filename"`
		ProjectID uint   `json:"projectId"`
	}

	type TaskStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		ID        uint       `json:"id"`
		Name      string     `json:"name"`
		Profile   string     `json:"profile"`
		ProjectID uint       `json:"projectId"`
		StartTime time.Time  `json:"startTime"`
		EndTime   time.Time  `json:"endTime"`
		StatusID  uint       `json:"statusId"`
		Status    TaskStatus `json:"status"`
	}

	type InvoiceStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Invoice struct {
		ID        uint          `json:"id"`
		ProjectID uint          `json:"projectId"`
		Name      string        `json:"name"`
		Number    string        `json:"number"`
		Amount    float64       `json:"amount"`
		IssueDate time.Time     `json:"issueDate"`
		DueDate   time.Time     `json:"dueDate"`
		StatusID  uint          `json:"statusId"`
		Status    InvoiceStatus `json:"status"`
	}

	type Project struct {
		ID          uint          `json:"id"`
		CreatedAt   time.Time     `json:"createdAt"`
		UpdatedAt   time.Time     `json:"updatedAt"`
		Name        string        `json:"name"`
		ClientID    uint          `json:"clientId"`
		Client      Client        `json:"client"`
		Profile     string        `json:"profile"`
		Offer       float64       `json:"offer"`
		Budget      float64       `json:"budget"`
		StatusID    uint          `json:"statusId"`
		Status      ProjectStatus `json:"status"`
		Attachments []Attachment  `json:"attachments"`
		Tasks       []Task        `json:"tasks"`
		Invoices    []Invoice     `json:"invoices"`
	}

	project, err := gorm.G[Project](c.DB).Table("projects p").Scopes(PreloadAll(2)).Where("p.id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到对应的项目", nil
	} else if err != nil {
		c.Error(err)
		return "查找项目失败", nil
	}

	return "", &project
}

func ListProjectTask(c *Context, u *struct {
	ProjectID uint `uri:"projectId"`
}, r *struct {
	Name      string        `form:"name"`
	StatusID  uint          `form:"statusId"`
	TimeRange *[2]time.Time `form:"timeRange[]"`
	Page      int           `form:"page"`
	Size      int           `form:"size"`
}) (message string, resp *List) {

	type Task struct {
		ID            uint       `json:"id"`
		CreatedAt     time.Time  `json:"createdAt"`
		Name          string     `json:"name"`
		Profile       string     `json:"profile"`
		StartTime     time.Time  `json:"startTime"`
		EndTime       time.Time  `json:"endTime"`
		StatusID      uint       `json:"-"`
		Status        TaskStatus `json:"status"`
		MaterialCount uint       `json:"materialCount"`
		RecordCount   uint       `json:"recordCount"`
	}

	q := gorm.G[Task](c.DB).Table("tasks t").Where("project_id = ?", u.ProjectID).Select(`*,
		(select count(*) from task_materials tm where tm.task_id = t.id) as material_count,
		(select count(*) from records r where r.task_id = t.id) as record_count`)

	if r.Name != "" {
		q = q.Where("MATCH(name) AGAINST(?)", r.Name)
	}

	if r.StatusID != 0 {
		q = q.Where("status_id = ?", r.StatusID)
	}

	if r.TimeRange != nil {
		q = q.Where("start_time <= ? AND end_time >= ?", r.TimeRange[1], r.TimeRange[0])
	}

	total, err := q.Count(c.Context, "*")
	if err != nil {
		c.Error(err)
		return "统计任务数量失败", nil
	}

	projects, err := q.Scopes(Paginate(r.Page, r.Size, 9), PreloadAll(1)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取任务列表失败", nil
	}

	return "", NewList(projects, total)
}

func ListProjectInvoice(c *Context, u *struct {
	ProjectID uint `uri:"projectId"`
}, r *struct {
	Title    string `form:"title"`
	StatusID uint   `form:"statusId"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}) (message string, resp any) {

	type Invoice struct {
		ID        uint          `json:"id"`
		CreatedAt time.Time     `json:"createdAt"`
		Name      string        `json:"name"`
		Number    string        `json:"number"`
		Profile   string        `json:"profile"`
		Amount    float64       `json:"amount"`
		IssueDate time.Time     `json:"issueDate"`
		DueDate   time.Time     `json:"dueDate"`
		StatusID  uint          `json:"-"`
		Status    InvoiceStatus `json:"status"`
	}

	query := gorm.G[Invoice](c.DB).Scopes(PreloadAll(2)).Where("project_id = ?", u.ProjectID)

	if r.Title != "" {
		query = query.Where("MATCH(title) AGAINST(?)", r.Title)
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

func EditProject(c *Context, u *struct {
	ID uint `uri:"projectId"`
}, r *struct {
	Name     string  `json:"name"`
	Profile  string  `json:"profile"`
	Offer    float64 `json:"offer"`
	Budget   float64 `json:"budget"`
	StatusID uint    `json:"statusId"`
}) (message string, resp *Project) {

	project := Project{
		Name:     r.Name,
		Profile:  r.Profile,
		Offer:    r.Offer,
		Budget:   r.Budget,
		StatusID: r.StatusID,
	}

	if row, err := gorm.G[Project](c.DB).Where("id = ?", u.ID).Select(
		"name, profile, offer, budget, status_id",
	).Updates(c, project); err != nil {
		c.Error(err)
		return "数据更新失败", nil
	} else if row == 0 {
		return "找不到对应的项目", nil
	}

	return "编辑项目成功", &project
}

func DeleteProject(c *Context, u *struct {
	ID uint `uri:"projectId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Project](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "删除项目失败", false
	} else if row == 0 {
		return "找不到对应的项目", false
	}

	return "项目删除成功", true
}
