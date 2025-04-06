package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	In Type = iota // 0
	Out            // 1
)

var allTypeTransaction = [2]string{"in", "out"}

func (t *Type) String() string {
	if t == nil {
		return ""
	}
	return allTypeTransaction[*t]
}

func ParseTransactionType(s string) (Type, error) {
	for i, v := range allTypeTransaction {
		if strings.EqualFold(v, s) { // So sánh không phân biệt hoa thường
			return Type(i), nil
		}
	}
	return Type(0), fmt.Errorf("invalid transaction type: %s", s)
}

func (t *Type) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("fail to scan data from database")
	}

	v, err := ParseTransactionType(string(bytes))
	if err != nil {
		return fmt.Errorf("fail to scan data from database: %v", err)
	}

	*t = v
	return nil
}

func (t *Type) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	str := string(data)

	// Xử lý giá trị null
	if str == "null" {
		return nil
	}

	// Xử lý giá trị số (ví dụ: 0 hoặc 1)
	if len(str) > 0 && str[0] != '"' {
		// Chuyển đổi giá trị số thành kiểu Type
		val, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("invalid JSON value: %s", str)
		}

		// Kiểm tra giá trị hợp lệ
		if val < 0 || val >= len(allTypeTransaction) {
			return fmt.Errorf("invalid transaction type value: %d", val)
		}

		*t = Type(val)
		return nil
	}

	// Xử lý giá trị chuỗi (ví dụ: "in" hoặc "out")
	if len(str) < 2 || str[0] != '"' || str[len(str)-1] != '"' {
		return fmt.Errorf("invalid JSON string: %s", str)
	}

	// Cắt chuỗi để loại bỏ dấu ngoặc kép
	trimmedStr := str[1 : len(str)-1]

	// Phân tích chuỗi thành kiểu TransactionType
	v, err := ParseTransactionType(trimmedStr)
	if err != nil {
		return err
	}

	// Gán giá trị cho biến t
	*t = v
	return nil
}

func (t *Type) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return t.String(), nil
}