/*
* @File: common.filter.go
* @Description: Defines filter information of the service
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package common

type Filter struct {
	Status string `form:"status" json:"status"`
}