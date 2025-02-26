// @File models.inventory_transaction.go
// @Description Defines inventory transaction information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package models

import "time"

type InventoryTransaction struct{
	Transaction_ID 		int 		`json:"transaction_id" gorm:"column:transaction_id"`
	Product_ID 			string 		`json:"product_id" gorm:"column:product_id"`
	Quantity 			int 		`json:"quatity" gorm:"column:quatity"`
	Transaction_Type 	*Type 		`json:"transaction_type" gorm:"column:transaction_type"`
	Transaction_Date 	*time.Time 	`json:"transaction_date" gorm:"column:transaction_date"`
	Supplier_ID 		int 		`json:"supplier_id" gorm:"column:supplier_id"`
}

func(InventoryTransaction) TableName() string{
	return "inventory_transactions"
}
