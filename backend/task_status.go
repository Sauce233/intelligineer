package main

import "gorm.io/gorm"

func ListTaskStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type TaskStatus struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		Color     string `json:"color"`
		TaskCount uint   `json:"taskCount"`
	}

	taskCounts, err := gorm.G[TaskStatus](c.DB).Table("task_statuses ts").Select(
		"*, (select count(*) from tasks t where t.status_id = ts.id) as task_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "獲取任務狀態列表失敗", nil
	}

	return "", taskCounts
}
