package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListTask(c *Context, u *struct {
	MaterialID uint `uri:"materialId"`
	ProjectID  uint `uri:"projectId"`
}, r *struct {
	Name      string        `form:"name"`
	StatusID  uint          `form:"statusId"`
	TimeRange *[2]time.Time `form:"timeRange[]"`
	Page      int           `form:"page"`
	Size      int           `form:"size"`
}) (message string, resp *List) {

	type Project struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type TaskStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		ID            uint       `json:"id"`
		CreatedAt     time.Time  `json:"createdAt"`
		Name          string     `json:"name"`
		Profile       string     `json:"profile"`
		ProjectID     uint       `json:"-"`
		Project       Project    `json:"project"`
		StartTime     time.Time  `json:"startTime"`
		EndTime       time.Time  `json:"endTime"`
		StatusID      uint       `json:"-"`
		Status        TaskStatus `json:"status"`
		MaterialCount uint       `json:"materialCount"`
		RecordCount   uint       `json:"recordCount"`
	}

	q := gorm.G[Task](c.DB).Table("tasks t").Select(`*,
		(select count(*) from task_materials tm where tm.task_id = t.id) as material_count,
		(select count(*) from records r where r.task_id = t.id) as record_count`)

	if r.Name != "" {
		q = q.Where("name LIKE ?", "%"+r.Name+"%")
	}

	if u.ProjectID != 0 {
		q = q.Where("project_id = ?", u.ProjectID)
	}

	if r.StatusID != 0 {
		q = q.Where("status_id = ?", r.StatusID)
	}

	if u.MaterialID != 0 {
		q = q.Where("id IN (?)", gorm.G[TaskMaterial](c.DB).Select("task_id").Where("material_id = ?", u.MaterialID))
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

func AddTask(c *Context, u *struct {
	ProjectID uint `uri:"projectId"`
}, r *struct {
	Name     string       `json:"name"`
	Profile  string       `json:"profile"`
	Time     [2]time.Time `json:"time"`
	StatusID uint         `json:"statusId"`
}) (message string, resp uint) {

	task := Task{
		Name:      r.Name,
		Profile:   r.Profile,
		StartTime: r.Time[0],
		EndTime:   r.Time[1],
		StatusID:  r.StatusID,
		ProjectID: u.ProjectID,
	}

	if err := gorm.G[Task](c.DB).Create(c, &task); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "不存在对应的專案", 0
	} else if err != nil {
		c.Error(err)
		return "任务创建失败", 0
	}

	return "任务创建成功", task.ID
}

func GetTask(c *Context, u *struct {
	ID uint `uri:"taskId"`
}, r *struct{}) (message string, resp any) {

	type Project struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Material struct {
		ID       uint    `json:"id"`
		Name     string  `json:"name"`
		Quantity float64 `json:"quantity"`
		Unit     string  `json:"unit"`
	}

	type TaskMaterial struct {
		TaskID     uint     `json:"-"`
		MaterialID uint     `json:"-"`
		Material   Material `json:"material"`
		Quantity   float64  `json:"quantity"`
		Profile    string   `json:"profile"`
	}

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Record struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		TaskID    uint      `json:"-"`
		UserID    uint      `json:"-"`
		User      User      `json:"user"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Minutes   string    `json:"minutes"`
	}

	type TaskStatus struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	type Task struct {
		ID        uint           `json:"id"`
		CreatedAt time.Time      `json:"createdAt"`
		UpdatedAt time.Time      `json:"updatedAt"`
		Name      string         `json:"name"`
		Profile   string         `json:"profile"`
		ProjectID uint           `json:"-"`
		Project   Project        `json:"project"`
		StartTime time.Time      `json:"startTime"`
		EndTime   time.Time      `json:"endTime"`
		StatusID  uint           `json:"-"`
		Status    TaskStatus     `json:"status"`
		Materials []TaskMaterial `json:"materials"`
		Records   []Record       `json:"records"`
	}

	project, err := gorm.G[Task](c.DB).Scopes(PreloadAll(2)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到对应的任务", nil
	} else if err != nil {
		c.Error(err)
		return "查找任务失败", nil
	}

	return "", &project
}

func ListTaskMaterial(c *Context, u *struct {
	TaskID uint `uri:"taskId"`
}, r *struct {
	Name     string `form:"name"`
	StatusID uint   `form:"statusId"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}) (message string, resp any) {

	type Material struct {
		ID       uint    `json:"id"`
		Name     string  `json:"name"`
		Quantity float64 `json:"quantity"`
		Unit     string  `json:"unit"`
	}

	type TaskMaterial struct {
		MaterialID uint     `json:"materialId"`
		Material   Material `json:"material"`
		Quantity   float64  `json:"quantity"`
		Profile    string   `json:"profile"`
	}

	query := gorm.G[TaskMaterial](c.DB).Table("task_materials tm").Scopes(PreloadAll(1)).Where("tm.task_id = ?", u.TaskID)

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计库存数量失败", nil
	}

	materials, err := query.Scopes(Paginate(r.Page, r.Size, 9)).Find(c)
	if err != nil {
		c.Error(err)
		return "获取库存列表失败", nil
	}

	return "", NewList(materials, total)
}

func ListTaskRecord(c *Context, u *struct {
	TaskID uint `uri:"taskId"`
}, r *struct {
	Title string        `form:"title"`
	Time  *[2]time.Time `form:"time[]"`
	Page  int           `form:"page"`
	Size  int           `form:"size"`
}) (message string, resp any) {

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Record struct {
		ID         uint      `json:"id"`
		CreatedAt  time.Time `json:"createdAt"`
		UserID     uint      `json:"userId"`
		User       User      `json:"user"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		Minutes    uint      `json:"minutes"`
		PhotoCount uint      `json:"photoCount"`
	}

	query := gorm.G[Record](c.DB).Table("records r").Scopes(PreloadAll(2)).Select(`*,
		(select count(*) from photos p where p.record_id = r.id) as photo_count`).Where("task_id = ?", u.TaskID)

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

func EditTask(c *Context, u *struct {
	ID uint `uri:"taskId"`
}, r *struct {
	Name     string       `json:"name"`
	Profile  string       `json:"profile"`
	Time     [2]time.Time `json:"time"`
	StatusID uint         `json:"statusId"`
}) (message string, resp *Task) {

	task := Task{
		Name:      r.Name,
		Profile:   r.Profile,
		StartTime: r.Time[0],
		EndTime:   r.Time[1],
		StatusID:  r.StatusID,
	}

	if rows, err := gorm.G[Task](c.DB).Where("id = ?", u.ID).Select(
		"name, profile, start_time, end_time, status_id",
	).Updates(c, task); err != nil {
		c.Error(err)
		return "数据更新失败", nil
	} else if rows == 0 {
		return "找不到对应的任务", nil
	}

	return "任务更新成功", &task
}

func DeleteTask(c *Context, u *struct {
	ID uint `uri:"taskId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Task](c.DB).Where("id = ?", u.ID).Delete(c); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "请先删除关联的数据", false
	} else if err != nil {
		c.Error(err)
		return "删除任务失败", false
	} else if row == 0 {
		return "找不到对应的任务", false
	}

	return "删除任务成功", true
}
