package common

// Struct hỗ trợ cho việc phân trang
type Paging struct {
	// Page: trang hiện tại
	Page int `json:"page" form:"page"` // form: là để cho Gin biết là lấy giá trị từ query param
	// Limit: số lượng item trên một trang
	Limit int `json:"limit" form:"limit"`
	// Total: tổng số dòng mà đáp ứng query
	Total int64 `json:"total" form:"-"` // "-" là để cho Gin biết là không cần lấy giá trị từ query param, lý do là vì giá trị này sẽ được tính toán từ dưới DB
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
}
