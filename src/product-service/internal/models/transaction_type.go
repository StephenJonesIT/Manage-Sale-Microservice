/*
* @File: models.transaction_type.go
* @Description: Implements transaction type enumeration and related functions
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
 */

package models

import (
	"database/sql/driver"
	"errors"
)

type Type int

const (
	In Type = iota
	Out
)

var allTypeTransaction = [2]string{"In", "Out"}

func(t *Type) String() string{
	return allTypeTransaction[*t]
}

func parseTransactionType(s string) (Type, error){
	for i := range allTypeTransaction {
		if allTypeTransaction[i] == s {
			return Type(i), nil;
		}
	}
	return Type(0), errors.New("invalid type string")
}

func(t *Type) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if  !ok {
		return errors.New("fail to scan data from database")
	}

	v, err := parseTransactionType(string(bytes))

	if err != nil {
		return errors.New("fail to scan data from database")
	}

	*t = v
	return nil
}

func (t *Type) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	v, err := parseTransactionType(str[1 : len(str)-1])
	if err != nil {
		return err
	}
	*t = v
	return nil
}

func (t *Type) Value() (driver.Value, error) {
	return t.String(), nil
}