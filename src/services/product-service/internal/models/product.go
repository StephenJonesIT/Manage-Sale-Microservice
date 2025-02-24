/*
* @File: models.product.go
* @Description: Defines products information will be returned to the clients
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package models

import "product-service/common"

type Product struct {
	Product_ID     string  `json:"product_id" gorm:"column:product_id;primaryKey"`
	Product_Name   string  `json:"product_name" gorm:"column:product_name"`
	Price          float64 `json:"price" gorm:"column:price"`
	Discount       float64 `json:"discount,omitempty" gorm:"column:discount"`
	Plant_Type     int     `json:"plant_type" gorm:"column:plant_type"`
	Uint           string  `json:"unit,omitempty" gorm:"column:unit"`
	Image_Url      string  `json:"image_url,omitempty" gorm:"column:image_url"`
	Description    string  `json:"description,omitempty" gorm:"column:description"`
	Product_Status *Status `json:"status,omitempty" gorm:"status"`
	Category_ID    string  `json:"category_id,omitempty" gorm:"column:category_id"`
	Supplier_ID	   string  `json:"supplier_id,omitempty" gorm:"column:supplier_id"`
	common.CommonTime
}

func(Product) TableName() string {
	return "products"
}