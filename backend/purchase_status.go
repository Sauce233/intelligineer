package main

import "gorm.io/gorm"

func ListPurchaseStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type PurchaseStatus struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Color         string `json:"color"`
		PurchaseCount uint   `json:"purchaseCount"`
	}

	purchaseStatuses, err := gorm.G[PurchaseStatus](c.DB).Table("purchase_statuses ps").Select(
		"*, (select count(*) from purchases p where p.status_id = ps.id) as purchase_count",
	).Find(c)
	if err != nil {
		c.Error(err)
		return "獲取訂單狀態列表失敗", nil
	}

	return "", purchaseStatuses
}
