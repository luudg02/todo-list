package main

import (
	"log"
	"net/http"
	"os"
	ginitem "social-todo-list/modules/item/transport/gin"

	"github.com/gin-gonic/gin" // Đây là Gin, một web framework cho Go để xây dựng REST API
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv" // godotenv giúp chúng ta đọc file .env
)

func main() {
	// Tải các biến từ tệp .env
	if err := godotenv.Load(); err != nil {
		log.Println("Cảnh báo: không thể tải tệp .env:", err)
	}

	// Kết nối với DB (MySQL)
	// DSN (Data Source Name) là một chuỗi kết nối đến DB
	dsn := os.Getenv("DB_CONN_STR") // Lấy giá trị của biến môi trường DB_CONN_STR
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err) // Fatalln sẽ in ra lỗi và kết thúc chương trình
	}

	// Tạo một web server với Gin
	// Đây là một web framework cho Go để xây dựng REST API
	r := gin.Default() // Khởi tạo một web server với Gin

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items => Đây là API để tạo một item mới (Create a new item)
	// GET 	 /v1/items => Đây là API để lấy danh sách các item (Read all items). Trong trường hợp muốn phân trang thì /v1/items?page=1 ...
	// GET 	 /v1/items/:id => Đây là API để lấy thông tin của một item (Get item detail by id). NOTE: id là số nguyên và tự tăng lên
	// (PUT  || PATCH) /v1/items/:id => Đây là API để cập nhật thông tin của một item (Update item by id). NOTE: PUT là thay đổi cả Object đó luôn, có nghĩa là chúng ta "đẩy" value Object như thế nào thì toàn bộ Object trên server của chúng ta. Còn PATCH là thay đổi từng thành phần, ví dụ như chúng ta chỉ muốn thay đổi "Title" thì truyền "TItle" lên, nếu là PUT thì nó sẽ xóa hết chỉ còn "Title" mà thôi, còn nếu là PATCH thì nó hiểu là chỉ thay đổi "Title" mà thôi
	// DELETE /v1/items/:id => Đây là API để xóa một item (Delete item by id)

	// Dưới đây là cách định nghĩa một route với các HTTP method khác nhau
	v1 := r.Group("/v1") // Định nghĩa một nhóm route với path prefix là "/v1"
	// Đây là một cách để nhóm các route cùng loại lại với nhau, giúp cho việc quản lý route trở nên dễ dàng hơn, lưu ý là cách thụt lề để định nghĩa route bên trong nhóm route
	{
		// Định nghĩa các routes liên quan đến TodoItems
		items := v1.Group("/items") // Định nghĩa một nhóm route với path prefix là "/v1/items"
		{
			items.POST("", ginitem.CreateItem(db))       // Định nghĩa một route với method POST và path đầy đủ là "/v1/items" - Tạo mới một item
			items.GET("", ginitem.ListItem(db))          // Định nghĩa một route với method GET và path đầy đủ là "/v1/items" - Lấy danh sách items
			items.GET("/:id", ginitem.GetItem(db))       // Định nghĩa một route với method GET và path đầy đủ là "/v1/items/:id" - Lấy chi tiết một item theo id (id là tham số)
			items.PATCH("/:id", ginitem.UpdateItem(db))  // Định nghĩa một route với method PATCH và path đầy đủ là "/v1/items/:id" - Cập nhật một item theo id
			items.DELETE("/:id", ginitem.DeleteItem(db)) // Định nghĩa một route với method DELETE và path đầy đủ là "/v1/items/:id" - Xóa một item theo id
		}
	}

	// Định nghĩa một route với method GET và path là "/ping"
	r.GET("/ping", func(c *gin.Context) { // c là một biến kiểu Context, nó chứa thông tin của request và response
		// Trả về một JSON object với key là message và value là pong
		c.JSON(http.StatusOK, gin.H{ // http.StatusOK = 200
			"message": "pong", // Trả về một JSON object với key là message và value là struct item
		})
	})
	// Chạy web server ở cổng 8080
	r.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
