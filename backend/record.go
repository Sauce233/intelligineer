package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListRecord(c *Context, u *struct{}, r *struct {
	TaskID uint          `form:"taskId"`
	UserID uint          `form:"userId"`
	Title  string        `form:"title"`
	Time   *[2]time.Time `form:"time[]"`
	Page   int           `form:"page"`
	Size   int           `form:"size"`
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

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Record struct {
		ID         uint      `json:"id"`
		CreatedAt  time.Time `json:"createdAt"`
		TaskID     uint      `json:"-"`
		Task       Task      `json:"task"`
		UserID     uint      `json:"-"`
		User       User      `json:"user"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		Minutes    uint      `json:"minutes"`
		PhotoCount uint      `json:"photoCount"`
	}

	query := gorm.G[Record](c.DB).Table("records r").Scopes(PreloadAll(2)).Select(`*,
		(select count(*) from photos p where p.record_id = r.id) as photo_count`)

	if r.TaskID != 0 {
		query = query.Where("task_id = ?", r.TaskID)
	}

	if r.UserID != 0 {
		query = query.Where("user_id = ?", r.UserID)
	}

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

func AddRecord(c *Context, u *struct {
	TaskID uint `uri:"taskId"`
}, r *struct {
	CreatedAt time.Time `json:"createdAt"`
	ProjectID uint      `json:"projectId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Minutes   uint      `json:"minutes"`
}) (message string, resp uint) {

	record := Record{
		UserID:  c.User.ID,
		TaskID:  u.TaskID,
		Title:   r.Title,
		Content: r.Content,
		Minutes: r.Minutes,
	}

	if err := gorm.G[Record](c.DB).Create(c, &record); errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "找不到对应的项目", 0
	} else if err != nil {
		c.Error(err)
		return "创建施工记录失败", 0
	}

	return "施工记录创建成功", record.ID
}

func GetRecord(c *Context, u *struct {
	ID uint `uri:"recordId"`
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

	type Task struct {
		ID        uint    `json:"id"`
		Name      string  `json:"name"`
		ProjectID uint    `json:"-"`
		Project   Project `json:"project"`
	}

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Record struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		TaskID    uint      `json:"-"`
		Task      Task      `json:"task"`
		UserID    uint      `json:"-"`
		User      User      `json:"user"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Minutes   uint      `json:"minutes"`
	}

	record, err := gorm.G[Record](c.DB).Scopes(PreloadAll(3)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到对应的施工记录", nil
	} else if err != nil {
		c.Error(err)
		return "查找施工记录失败", nil
	}

	return "", &record
}

func EditRecord(c *Context, u *struct {
	ID uint `uri:"recordId"`
}, r *struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Minutes uint   `json:"minutes"`
}) (message string, resp *Record) {

	record := Record{
		Title:   r.Title,
		Content: r.Content,
		Minutes: r.Minutes,
		UserID:  c.User.ID,
	}

	if row, err := gorm.G[Record](c.DB).Where("id = ?", u.ID).Select(
		"title, content, minutes, user_id",
	).Updates(c, record); err != nil {
		c.Error(err)
		return "数据更新失败", nil
	} else if row == 0 {
		return "找不到对应的施工记录", nil
	}

	return "施工记录修改成功", &record
}

func DeleteRecord(c *Context, u *struct {
	ID uint `uri:"recordId"`
}, r *struct{}) (message string, resp bool) {

	if row, err := gorm.G[Record](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "删除施工记录失败", false 
	} else if row == 0 {
		return "找不到这个施工记录", false
	}

	return "施工记录删除成功", true
}
