package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota // * Đặt ItemStatus ở sau cùng => rằng buộc có các biến const thuộc kiểu ItemStatus mà ItemStatus là kiểu int nên là nó sẽ thỏa mãn "iota"
	ItemStatusDone
	ItemStatusDeleted
)

// Biến allItemStatuses là một mảng chứa các trạng thái của Item
var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

// Chuyển các biến const ItemStatus thành string
func (item *ItemStatus) string() string {
	return allItemStatuses[*item] // Trả về string tương ứng với ItemStatus
}

// Hàm này dùng để chuyển từ kiểu string sang kiểu ItemStatus
func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New(fmt.Sprintf("Invalid status string: %s", s))
}

// ĐỌc SQL từ dưới DB lên struct
// Hàm này sẽ được gọi khi scan dữ liệu từ DB lên struct
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("Fail to scan data from sql: %s", value))
	}

	*item = v

	return nil
}

// Từ ngược từ struct xuống DB
// Hàm này sẽ được gọi khi insert dữ liệu từ struct xuống DB
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.string(), nil
}

// chuyển từ struct sang JSON (JSON Encoding)
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil { // Nếu mà item thực sự là nil thì hàm item.string() sẽ bị crash, nên chúng ta cần phải xử lý riêng
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.string())), nil
}

// lấy từ JSON sang struct (JSON Decoding)
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStr2ItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
