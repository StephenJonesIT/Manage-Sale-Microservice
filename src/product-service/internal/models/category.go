// @File models.category.go
// @Description Defines category product information will be returned for clients
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package models

type Category struct{
	Category_ID  int `gorm:"column:category_id" json:"category_id"`
	Category_Name string `gorm:"column:category_name" json:"category_name"`
	Description string `gorm:"column:description" json:"description"`
}

func(Category) TableName() string{
	return "category_product"
}

type AddCategory struct{
	Name string `gorm:"column:category_name" json:"category_name"`
	Description string `gorm:"column:description" json:"description"`
}

func(AddCategory) TableName() string{
	return Category{}.TableName()
}