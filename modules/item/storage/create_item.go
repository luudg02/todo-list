package storage

import (
	"social-todo-list/modules/item/model"

	"golang.org/x/net/context"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
