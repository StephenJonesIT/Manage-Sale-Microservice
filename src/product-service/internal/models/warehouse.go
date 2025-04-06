// @File models.warehouse.go
// @Description Defines warehouses information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package models

type Warehouses struct {
	Warehouse_ID    int    `gorm:"column:warehouse_id" json:"warehouse_id,omitempty"`
	Warehouses_Name string `gorm:"column:warehouse_name" json:"warehouse_name"`
	Location        string `gorm:"column:location" json:"location"`
}

func (Warehouses) TableName() string {
	return "warehouses"
}
