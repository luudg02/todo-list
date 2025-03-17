package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// API tạo một item mới
// Hàm này trả về một hàm khác, hàm này sẽ được gọi khi có một request POST đến /v1/items
func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// /v1/items/:id
		id, err := strconv.Atoi(c.Param("id")) // Chúng ta sẽ lấy ra giá trị của tham số id từ URL, ở đây là lấy ra id từ URL "/v1/items/:id" => Giá trị giả về về là "1" (kiểu string). Nếu muốn đưa về kiểu số nguyên thì chúng ta cần phải parse nó về kiểu số nguyên, lúc này chúng ta sử dụng hàm strconv.Atoi(Param("id")). Hàm này sẽ trả về 2 giá trị, giá trị đầu tiên là giá trị số nguyên của id, giá trị thứ 2 là lỗi nếu có
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)     // Khởi tạo một biến storage để chứa connection đến DB
		business := biz.NewGetItemBiz(store) // Khởi tạo một biến biz để chứa business logic

		data, err := business.GetItemById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Trả về dữ liệu đã lấy được
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data)) // Trả về dữ liệu đã lấy được
	}
}
