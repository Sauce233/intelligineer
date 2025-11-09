package main

/*

func EditProjectMaterial(c *Context, u *struct {
	ProjectID  uint `uri:"projectId"`
	MaterialID uint `uri:"materialId"`
}, r *struct{}) (message string, resp any) {

	var req struct {
		Quantity float64 `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return "读取材料数量失败", false
	}

	if err := c.DB.Save(&ProjectMaterial{
		ProjectID:  u.ProjectID,
		MaterialID: u.MaterialID,
		Quantity:   req.Quantity,
	}).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
		return "这个材料已经添加过了", false
	} else if err != nil {
		c.Error(err)
		return "编辑项目材料失败", false
	}

	return "编辑项目材料成功", true
}

func DeleteProjectMaterial(c *Context, u *struct {
	ProjectID  uint `uri:"projectId"`
	MaterialID uint `uri:"materialId"`
}, r *struct{}) (message string, resp any) {

	if result := c.DB.Where(
		"project_id = ? AND material_id = ?", u.ProjectID, u.MaterialID,
	).Delete(new(ProjectMaterial)); result.RowsAffected == 0 {
		return "这个项目没有分配这个材料", false
	} else if result.Error != nil {
		c.Error(result.Error)
		return "删项目材料失败", false
	}

	return "删除项目材料成功", true
}
*/
