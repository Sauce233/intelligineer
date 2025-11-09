package main

import "gorm.io/gorm"

func ListInvoiceStatus(c *Context, u *struct{}, r *struct{}) (string, any) {

	type InvoiceStatus struct {
		ID           uint    `json:"id"`
		Name         string  `json:"name"`
		Color        string  `json:"color"`
		InvoiceCount uint    `json:"invoiceCount"`
		TotalAmount  float64 `json:"totalAmount"`
	}

	invoiceStatuses, err := gorm.G[InvoiceStatus](c.DB).Table("invoice_statuses iss").Select(`*, 
		(select count(*) from invoices i where i.status_id = iss.id) as invoice_count,
		(select sum(amount) from invoices i where i.status_id = iss.id) as total_amount`,
	).Find(c)
	if err != nil {
		c.Error(err)
		return "获取发票状态列表失败", nil
	}

	return "", invoiceStatuses
}
