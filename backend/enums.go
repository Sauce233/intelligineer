package main

import "gorm.io/gorm"

func GetEnums(c *Context, u *struct{}, r *struct{}) (string, any) {

	type Enum struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	tables := [10]string{
		"project_statuses",
		"invoice_statuses",
		"task_statuses",
		"material_statuses",
		"purchase_statuses",
		"attendance_types",
		"profile_types",
		"roles",
		"application_types",
		"application_statuses",
	}

	enumsMap := make(map[string][]Enum)

	for _, table := range tables {
		enums, err := gorm.G[Enum](c.DB).Table(table).Find(c)
		if err != nil {
			c.Error(err)
			return "获取枚举失败", nil
		}
		enumsMap[table] = enums
	}

	return "", enumsMap
}
