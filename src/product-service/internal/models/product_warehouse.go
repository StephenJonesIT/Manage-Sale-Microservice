// @File models.product_warehouse.go
// @Description Defines product warehouse information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package models

import "time"

type ProductWarehouses struct {
	Product_ID   	string 		`json:"product_id" gorm:"column:product_id"`
	WareHouse_ID 	int    		`json:"warehouse_id" gorm:"column:warehouse_id"`
	Quantity     	int    		`json:"quantity" gorm:"column:quantity"`
	Last_Updated  	*time.Time	`json:"last_updated,omitempty" gorm:"column:last_updated"`
}

func(ProductWarehouses) TableName() string{
	return "product_warehouse"
}