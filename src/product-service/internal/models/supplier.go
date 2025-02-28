// @File models.supplier.go
// @Description Defines suppplier information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package models

import "product-service/common"

type Supplier struct {
	Supplier_ID   int    `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier_Name string `gorm:"column:supplier_name" json:"supplier_name"`
	Phone         string `gorm:"column:contact_phone" json:"phone"`
	Email         string `gorm:"column:contact_email" json:"email"`
	Address       string `gorm:"column:address" json:"address"`
	City          string `gorm:"column:city" json:"city"`
	Country       string `gorm:"column:country" json:"country"`
	common.CommonTime
}

func (Supplier) TableName() string {
	return "suppliers"
}

type AddSupplier struct {
	Supplier_Name string `gorm:"column:supplier_name" json:"supplier_name"`
	Phone         string `gorm:"column:contact_phone" json:"phone"`
	Email         string `gorm:"column:contact_email" json:"email"`
	Address       string `gorm:"column:address" json:"address"`
	City          string `gorm:"column:city" json:"city"`
	Country       string `gorm:"column:country" json:"country"`
}

func(AddSupplier) TableName() string {
	return Supplier{}.TableName()
}

type UpdateSupplier struct {
	Supplier_ID   int    `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier_Name string `gorm:"column:supplier_name" json:"supplier_name"`
	Phone         string `gorm:"column:contact_phone" json:"phone"`
	Email         string `gorm:"column:contact_email" json:"email"`
	Address       string `gorm:"column:address" json:"address"`
	City          string `gorm:"column:city" json:"city"`
	Country       string `gorm:"column:country" json:"country"`
}

func(UpdateSupplier) TableName() string{
	return Supplier{}.TableName()
}