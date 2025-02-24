/*
* @File: models.product_status.go
* @Description: Implements product status enumeration and related functions
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/

package models

import (
	"database/sql/driver"
	"errors"
)

type Status int

const (
	Available Status = iota
	OutOfStock
	Discontinued
	Pre_Order
	BackOrdered
	Reserved
	OnSale
	NewArrival
	Damaged
	Pending
)

var allProductStatuses = [10]string{
	"Available",
	"Out of Stock",
	"Discontinued",
	"Pre-Order",
	"Back-Ordered",
	"Reserved",
	"On Sale",
	"New Arrival",
	"Damaged",
	"Pending",
}

func (status *Status) String() string {
	return allProductStatuses[*status]
}

func parseStrProductStatus(s string) (Status, error) {
	for i := range allProductStatuses {
		if allProductStatuses[i] == s {
			return Status(i), nil
		}
	}
	return Status(0), errors.New("invalid status string")
}

func (status *Status) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New("fail to scan data from sql")
	}

	v, err := parseStrProductStatus(string(bytes))

	if err != nil {
		return errors.New("fail to scan data from sql")
	}

	*status = v

	return nil
}

func (status *Status) MarshalJSON() ([]byte, error) {
	if status == nil {
		return nil, nil
	}
	return []byte(`"` + status.String() + `"`), nil
}

func (status *Status) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	v, err := parseStrProductStatus(str[1 : len(str)-1])
	if err != nil {
		return err
	}
	*status = v
	return nil
}

func (status *Status) Value() (driver.Value, error) {
	return status.String(), nil
}