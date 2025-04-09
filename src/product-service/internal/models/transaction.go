// @File models.inventory_transaction.go
// @Description Defines inventory transaction information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package models

import "time"

type InventoryTransaction struct{
	Transaction_ID 		string 		`json:"transaction_id" gorm:"column:transaction_id"`
	Product_ID 			string 		`json:"product_id" gorm:"column:product_id"`
	Warehouse_ID		int			`json:"warehouse_id" gorm:"column:warehouse_id"`
	Quantity 			int 		`json:"quantity" gorm:"column:quantity"`
	Transaction_Type 	*Type 		`json:"transaction_type,omitempty" gorm:"column:transaction_type"`
	Transaction_Date 	*time.Time 	`json:"transaction_date,omitempty" gorm:"column:transaction_date"`
}

func(InventoryTransaction) TableName() string{
	return "inventory_transactions"
}

type Transaction struct{
	Transaction_ID 		string 		`json:"transaction_id" gorm:"column:transaction_id"`
	Product_ID 			string 		`json:"product_id" gorm:"column:product_id"`
	Warehouse_ID		int			`json:"warehouse_id" gorm:"column:warehouse_id"`
	Quantity 			int 		`json:"quantity" gorm:"column:quantity"`
	Transaction_Type 	*Type 		`json:"transaction_type,omitempty" gorm:"column:transaction_type"`
	Transaction_Date 	*time.Time 	`json:"transaction_date,omitempty" gorm:"column:transaction_date"`
	Product_Name    	string
	Warehouses_Name 	string
}