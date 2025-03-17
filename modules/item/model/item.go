package model

import (
	"errors"
	"social-todo-list/common"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

/*
 * Đây là struct tương ứng với bảng TodoItem trong DB
 * Sử dụng tag json bởi vì khi khi lập trình REST API chúng ra sẽ giao tiếp với client thông qua ngôn ngữ trung gian là JSON (JavaScript Object Notation)
 */
type TodoItem struct {
	common.SQLModel        // Embedding SQLModel vào struct TodoItem (Không được hiểu lầm là kế thừa trong OOP). Đặc tính của embedding là các field và các method của struct được embedded sẽ được đưa vào struct chứa nó
	Title           string `json:"title" gorm:"column:title;"`
	// Image : có thể bỏ qua vì là NULL ở dưới db
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status"` // status trong DB là enum, bản chất khi đưa lên struct thì có dạng là string
	/*
	 * Khi sử dụng omitempty thì sẽ bỏ qua khi các giá trị dưới đây ở kiểu dữ liệu mặc định của nó
	 * int: 0
	 * float: 0.0
	 * string: ""
	 * bool: false
	 * struct: nil
	 * pointer: nil
	 */
}

// Hàm có receiver TodoItem với tên là TableName giả về tên bảng trong DB mà struct này sẽ map tới
func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int         `json:"id" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`

	// ! Chỉ sử dụng  4 trường trên bởi vì: "CreatedAt" và "UpdatedAt" là do DB đã sử dụng CURRENT_TIMESTAMP để tự động thêm vào nên là không cần lấy
}

// Hàm có receiver TodoItemCreation với tên là TableName giả về tên bảng trong DB mà struct này sẽ map tới
func (TodoItemCreation) TableName() string {
	return "todo_items"
}

// Struct hõ trợ update dữ liệu dưới DB
type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status"`
}

// Hàm có receiver TodoItemUpdate với tên là TableName giả về tên bảng trong DB mà struct này sẽ map tới
func (TodoItemUpdate) TableName() string {
	return "todo_items"
}
