// @File models.warehouse.go
// @Description Defines warehouses information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package models

type Warehouses struct {
	Warehouse_ID 	int 	`gorm:"column:warehouse_id" json:"Warehouse_ID"`
	Warehouses_Name string 	`gorm:"column:warehouse_name" json:"Warehouse_Name"`
	Location 		string 	`gorm:"column:location" json:"Location"`
}

func(Warehouses) TableName() string{
	return "warehouses"
}