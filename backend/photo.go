package main

func AddPhoto(c *Context, u *struct {
	RecordID uint `uri:"recordId"`
}, r *struct{}) (message string, resp bool) {
	return "", true
}

func DeletePhoto(c *Context, u *struct {
	ID uint `uri:"photoId"`
}, r *struct{}) (message string, resp bool) {

	if res := c.DB.Delete(new(Photo), u.ID); res.RowsAffected == 0 {
		return "没有这个照片", false
	} else if res.Error != nil {
		c.Error(res.Error)
		return "照片删除失败", false
	}

	return "照片删除成功", true
}
