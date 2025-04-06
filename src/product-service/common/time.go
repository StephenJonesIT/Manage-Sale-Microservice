/*
* @File: common.time.go
* @Description: Defines time information for the service
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/

package common

import "time"

type CommonTime struct {
	Create_At *time.Time  `json:"create_at,omitempty" gorm:"column:created_at"`
	Update_At *time.Time  `json:"update_at,omitempty" gorm:"column:updated_at"`
}