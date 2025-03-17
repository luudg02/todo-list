package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`                 // Tag này có nghĩa là khi chúng ta parse struct TodoItem thành JSON thì key của JSON sẽ là "id" và gorm:"column:id;" có nghĩa là field Id trong struct TodoItem sẽ map với cột id trong bảng todo_items trong DB
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"` // Sử dụng con trỏ vì nếu con trỏ là nil thì ở dưới DB sẽ là NULL
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"` // * Sử dụng "omitempty" để bỏ qua nếu giá trị là nil (nếu mà không sử dụng omitempty đồng thời giá trị trước khi parse là nil thì khi parse thành JSON sẽ là NULL)
}
