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

// API tạo một item mới
// Hàm này trả về một hàm khác, hàm này sẽ được gọi khi có một request POST đến /v1/items
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// * Để có thể tạo ra 1 item mới chúng ta cần truyền/gọi request body từ client lên cho server rồi Gin sẽ parse nó thành dữ liệu rồi mới xuống DB
		// ! Như vậy là chúng ta cần 1 bước trung gian là lấy dữ liệu từ request body chuyển 1 struct nào đó để chứa dữ liệu đó. Ở đây thì chúng ta có thể chuyển trực tiếp luôn struct TodoItem của chúng ta tuy nhiên trên thực tế thì struct chính có rất nhiều field name, thậm chí nó sẽ có những cái Object/struct lồng vào nhau nữa => nó sẽ rất là lớn
		// * => Giải pháp là chúng ta sử dụng struct nhỏ hơn. Ở dòng 40 chúng ta đã khai báo một struct nhỏ hơn để chứa dữ liệu từ request body

		// Đây là cách để lấy dữ liệu từ request body
		var data model.TodoItemCreation // Khai báo một biến kiểu ToItemCreation để chứa dữ liệu từ request body

		// Xử lý lỗi nếu có
		if err := c.ShouldBind(&data); err != nil { // ở trong hàm ShouldBind diễn ra quá trình Unmarshal từ JSON sang struct
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)        // Khởi tạo một biến storage để chứa connection đến DB
		business := biz.NewCreateItemBiz(store) // Khởi tạo một biến biz để chứa business logic

		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil { // Gọi hàm CreateNewItem từ business logic
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Trả về lỗi nếu có
			return
		}

		// Trả về dữ liệu đã được parse từ request body
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id)) // Tại sao lại trả về data.Id mà không phải full struct TodoItemCreation vì chúng ta đang ở API Create nên là chúng ta chỉ cần làm đúng nhiệm vụ create của nó. Nếu như muốn trả về full struct thì chúng ta sẽ phải tốn thêm 1 bước nữa là phải query dữ liệu lên để lấy được full, tại vì đôi khi dữ liệu chúng ta insert vào (TodoItemCreation) và dữ liệu trả ra (TodoItem) có khả năng là sẽ có sự chênh lệch rất lớn, nhiều khi chúng ta sẽ phải query đi rất nhiều bảng nữa để mới lấy được full dữ liệu đó => Cho nên không nên làm vậy => Tóm lại: API nào thì làm đúng nhiệm vụ của nó, với lại là có convention là nếu như client muốn lấy full dữ liệu mới nhất thì họ sẽ gọi API "GET detail by id"

		// Insert dữ liệu từ struct TodoItemCreation xuống DB. Để làm được điều đó thì chúng ta sẽ cần 1 hàm để cho GORM biết được là struct TodoItemCreation này sẽ map với bảng nào trong DB và những field nào trong struct này sẽ map với những field nào trong bảng đó
		// Sau đó chúng ta cần cho GORM biết clà file name của struct TodoItemCreation sẽ xuống cột nào trong dưới DB bằng cách thêm tag ở trong json tag của struct TodoItemCreation (xem dòng 41, 42 và 43 đoạn có gorm:"column:...")
		// Sau khi xong định nghĩa để xuống DB thì bây giờ sau khi parse thành công thì chúng ta sử dụng connection của DB để insert dữ liệu xuống DB; thêm tham số "db *gorm.DB" vào hàm CreateItem(), tại dùng 140 chúng ta sẽ thêm db.Create(&data)
	}
}
