package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var result []model.TodoItem

	// Lọc ra những dòng mà status khác "Deleted"
	// ! Nếu như không có dòng này thì nó sẽ lấy ra tất cả các dòng trong bảng todo_items bất kể status là gì
	db := s.db.Where("status <> ?", "Deleted") // Nên để ngay trước Count bởi vì nếu không nó sẽ đi đếm hết cả cái bảng đó

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	// Tính tổng số dòng mà đáp ứng query
	if err := db.Table("todo_items").Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&result).Error; err != nil {

		return nil, err
	}

	return result, nil
}
