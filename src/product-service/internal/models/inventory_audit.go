// @File models.inventory_transaction.go
// @Description Defines inventory audit information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package models

import "time"

type InventoryAudit struct{
	Audit_ID 			int 		`json:"audit_id" gorm:"column:audit_id"`
	Product_ID 			string 		`json:"product_id" gorm:"column:product_id"`
	Warehouses 			int 		`json:"warehouse_id" gorm:"column:warehouse_id"` 
	Quantity 			int 		`json:"quatity" gorm:"column:actual_quatity"`
	Transaction_Type 	*Type 		`json:"transaction_type" gorm:"column:transaction_type"`
	Audit_Date 			*time.Time 	`json:"audit_date" gorm:"column:audit_date"`
	Auditor 			string 		`json:"auditor" gorm:"column:auditor"`
}

func(InventoryAudit) TableName() string{
	return "inventory_audit"
}
