package common

type Response struct {
	Code  	int 	    	`json:"code"`
	Message interface{}     `json:"message,omitempty"`
	Data  	interface{}   	`json:"data"`
	Filters interface{} 	`json:"filter,omitempty"`
	Paging  interface{} 	`json:"paging,omitempty"`
}

func NewResponse(code int, message string, data, filter, paging interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
		Filters: filter,
		Paging:  paging,
	}
}