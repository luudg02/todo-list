package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// API lấy danh sách các item
func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Đây là cách để lấy dữ liệu từ request body
		var paging common.Paging // Khai báo một biến kiểu Paging để chứa dữ liệu từ request body

		// Xử lý lỗi nếu có
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Nếu mà ở dưới không có truyền lên thì chúng ta nên có 1 defer limit cũng như là trang bắt đầu từ trang số 1 nếu như mà không ai làm gì nó
		paging.Process()

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)      // Khởi tạo một biến storage để chứa connection đến DB
		business := biz.NewListItemBiz(store) // Khởi tạo một biến biz để chứa business logic

		result, err := business.ListItem(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Trả về dữ liệu đã được parse từ request body
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter)) // Trả về dữ liệu đã lấy được
	}
}
