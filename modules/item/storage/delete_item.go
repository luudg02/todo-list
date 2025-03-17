package storage

import (
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	// Xóa dữ liệu từ bảng todo_items với id tương ứng với id truyền vào
	/*	// Cách 1: Đây là cách "Hard delete", có nghĩa là chúng ta xóa thẳng dòng dữ liệu đó trong DB
		if err := db.Table("todo_items").Where("id = ?", id).Delete(nil).Error; err != nil { // Truyền nil vào hàm Delete bởi vì chúng ta đã xác định được id là gì rồi thì ko cần để value "&data" vào
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	*/
	// Trong hầu hết các trường hợp thì không nên sử dụng ""Hard delete" mà nên sử dụng "Soft delete" (có nghĩa là chúng ta sẽ không xóa dòng dữ liệu đó mà chỉ đánh dấu nó là đã bị xóa (Deleted))
	// Cách 2: Soft delete
	if err := s.db.Table("todo_items").
		Where(cond).
		Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil { // Truyền nil vào hàm Delete bởi vì chúng ta đã xác định được id là gì rồi thì ko cần để value "&data" vào
		return err
	}

	return nil
}
