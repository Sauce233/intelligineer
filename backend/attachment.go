package main

import "gorm.io/gorm"

func ListProjectAttachment(c *Context, u *struct {
	ProjectID uint `uri:"projectId"`
}, r *struct{}) (string, any) {

	type Attachment struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Filename string `json:"filename"`
	}

	query := gorm.G[Attachment](c.DB).Where("project_id = ?", u.ProjectID)

	total, err := query.Count(c, "*")
	if err != nil {
		c.Error(err)
		return "统计附件数量失败", nil
	}

	attachments, err := query.Find(c)
	if err != nil {
		c.Error(err)
		return "获取附件列表失败", nil
	}

	return "", NewList(attachments, total)
}

func AddProjectAttachment(c *Context, u *struct{}, r *struct{}) (string, uint) {

	return "", 1
}
