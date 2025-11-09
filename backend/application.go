package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

func ListApplication(c *Context, u *struct{}, r *struct {
	Title    string        `form:"title"`
	TypeID   uint          `form:"typeId"`
	StatusID uint          `form:"statusId"`
	Time     *[2]time.Time `form:"time[]"`
	Page     int           `form:"page"`
	Size     int           `form:"size"`
}) (message string, resp any) {

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Application struct {
		ID          uint              `json:"id"`
		CreatedAt   time.Time         `json:"createdAt"`
		UpdatedAt   time.Time         `json:"updatedAt"`
		ApplicantID uint              `json:"-"`
		Applicant   User              `json:"applicant"`
		ApproverID  *uint             `json:"-"`
		Approver    *User             `json:"approver"`
		Title       string            `json:"title"`
		Content     string            `json:"content"`
		TypeID      uint              `json:"-"`
		Type        ApplicationType   `json:"type"`
		StatusID    uint              `json:"-"`
		Status      ApplicationStatus `json:"status"`
	}

	query := gorm.G[Application](c.DB).Scopes()

	if r.TypeID != 0 {
		query = query.Where("type_id = ?", r.TypeID)
	}

	if r.Title != "" {
		query = query.Where("MATCH(title) AGAINST(?)", r.Title)
	}

	if r.StatusID != 0 {
		query = query.Where("status_id = ?", r.StatusID)
	}

	if r.Time != nil {
		query = query.Where("created_at >= ? AND created_at <= ?", r.Time[0], r.Time[1])
	}

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "获取申请总数失败", nil
	}

	applications, err := query.Scopes(PreloadAll(1), Paginate(r.Page, r.Size, 9)).Order("id desc").Find(c)
	if err != nil {
		c.Error(err)
		return "获取申请列表失败", nil
	}

	return "", NewList(applications, total)
}

func AddApplication(c *Context, u *struct{}, r *struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TypeID  uint   `json:"typeId"`
}) (message string, resp bool) {

	if err := c.DB.Create(&Application{
		ApplicantID: c.User.ID,
		Title:       r.Title,
		Content:     r.Content,
		TypeID:      r.TypeID,
	}).Error; errors.Is(err, gorm.ErrForeignKeyViolated) {
		return "找不到对应的申请类型", false
	} else if err != nil {
		c.Error(err)
		return "创建申请失败", false
	}

	return "申请创建成功", true
}

func GetApplication(c *Context, u *struct {
	ID uint `uri:"applicationId"`
}, r *struct{}) (message string, resp any) {

	type User struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	type Application struct {
		ID          uint              `json:"id"`
		CreatedAt   time.Time         `json:"createdAt"`
		UpdatedAt   time.Time         `json:"updatedAt"`
		ApplicantID uint              `json:"-"`
		Applicant   User              `json:"applicant"`
		ApproverID  uint              `json:"-"`
		Approver    User              `json:"approver"`
		Title       string            `json:"title"`
		Content     string            `json:"content"`
		Reply       string            `json:"reply"`
		TypeID      uint              `json:"-"`
		Type        ApplicationType   `json:"type"`
		StatusID    uint              `json:"-"`
		Status      ApplicationStatus `json:"status"`
	}

	application, err := gorm.G[Application](c.DB).Scopes(PreloadAll(1)).Where("id = ?", u.ID).Take(c)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "找不到对应的申请", false
	} else if err != nil {
		c.Error(err)
		return "查找申请失败", false
	}

	return "", &application
}

func ReplyApplication(c *Context, u *struct {
	ID uint `uri:"applicationId"`
}, r *struct {
	Reply  string `json:"reply" binding:"required"`
	Status bool   `json:"status"`
}) (message string, resp any) {

	if result := c.DB.Model(new(Application)).Where(u.ID).Updates(map[string]any{
		"reply":         r.Reply,
		"status":        r.Status,
		"reply_user_id": c.User.ID,
	}); result.RowsAffected == 0 {
		return "找不到对应的申请", false
	} else if result.Error != nil {
		c.Error(result.Error)
		return "批复申请失败", false
	}

	return "申请批复成功", true
}

func ApproveApplication(c *Context, u *struct {
	ID uint `uri:"applicationId"`
}, r *struct{}) (message string, resp bool) {

	if result := c.DB.Model(&Application{}).Where(u.ID).Updates(map[string]any{
		"status":        true,
		"reply":         "已批准",
		"reply_user_id": c.User.ID,
	}); result.RowsAffected == 0 {
		return "找不到对应的申请", false
	} else if result.Error != nil {
		c.Error(result.Error)
		return "批准申请失败", false
	}

	return "申请已批准", true
}

func RejectApplication(c *Context, u *struct {
	ID string `uri:"applicationId"`
}, r *struct {
	Reply string `json:"reply"`
}) (message string, resp bool) {

	reply := r.Reply
	if reply == "" {
		reply = "已拒绝"
	}

	if result := c.DB.Model(&Application{}).Where(u.ID).Updates(&map[string]interface{}{
		"status":        false,
		"reply":         reply,
		"reply_user_id": c.User.ID,
	}); result.RowsAffected == 0 {
		return "找不到对应的申请", false
	} else if result.Error != nil {
		c.Error(result.Error)
		return "拒绝申请失败", false
	}

	return "申请已拒绝", true
}
func DeleteApplication(c *Context, u *struct {
	ID uint `uri:"applicationId"`
}, r *struct{}) (message string, resp any) {

	if row, err := gorm.G[Application](c.DB).Where("id = ?", u.ID).Delete(c); err != nil {
		c.Error(err)
		return "删除申请失败", false
	} else if row == 0 {
		return "找不到这个申请", false
	}

	return "申请删除成功", true
}
