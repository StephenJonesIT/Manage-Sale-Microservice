/*
* @File: common.response.go
* @Description: Defines Response information will be returned to the clients
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package common

type Response struct {
	Code  	int 	    	`json:"code,omitempty"`
	Message interface{}     `json:"message,omitempty"`
	Data  	interface{}   	`json:"data"`
	Filters interface{} 	`json:"filter,omitempty"`
	Paging  interface{} 	`json:"paging,omitempty"`
}

func NewResponse(data interface{}) *Response{
	return &Response{
		Data: data,
	}
}

func NewDetailResponse(code int, message string, data, filter, paging interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Filters: filter,
		Paging:  paging,
	}
}