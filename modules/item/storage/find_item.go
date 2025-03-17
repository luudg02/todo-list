package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	// Lấy dữ liệu ra với id tương ứng
	// Lấy dữ liệu từ bảng todo_items với id tương ứng với id truyền vào
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
