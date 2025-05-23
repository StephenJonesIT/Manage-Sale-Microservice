/*
* @File: common.paging.go
* @Description: Defines Paging information of the service
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package common

type Paging struct{
	Page 	int 	`json:"page" form:"page"`
	Limit 	int 	`json:"limit" form:"limit"`
	Total 	int64 	`json:"total" form:"-"`
}

func(paging *Paging) Process(){
	if paging.Page < 1{
		paging.Page = 1
	}
	if paging.Limit <= 0 || paging.Limit > 100 {
		paging.Limit = 10
	}
}