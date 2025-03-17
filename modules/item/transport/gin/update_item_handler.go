package ginitem

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/modules/item/biz"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// API edit 1 item by id
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		// /v1/items/:id
		id, err := strconv.Atoi(c.Param("id")) // Chúng ta sẽ lấy ra giá trị của tham số id từ URL, ở đây là lấy ra id từ URL "/v1/items/:id" => Giá trị giả về về là "1" (kiểu string). Nếu muốn đưa về kiểu số nguyên thì chúng ta cần phải parse nó về kiểu số nguyên, lúc này chúng ta sử dụng hàm strconv.Atoi(Param("id")). Hàm này sẽ trả về 2 giá trị, giá trị đầu tiên là giá trị số nguyên của id, giá trị thứ 2 là lỗi nếu có
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// parse dữ liệu từ request body
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Trả về dữ liệu đã lấy được
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true)) // Trả về true nếu cập nhật thành công
	}
}
