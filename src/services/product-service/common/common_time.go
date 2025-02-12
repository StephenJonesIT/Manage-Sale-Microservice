package common

import "time"

type CommonTime struct {
	Create_At *time.Time  `json:"create_at,omitempty" gorm:"column:create_at"`
	Update_At *time.Time  `json:"update_at,omitempty" gorm:"column:update_at"`
}