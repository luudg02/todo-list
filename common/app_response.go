package common

type successRes struct {
	// Dữ liệu
	Data interface{} `json:"data"`

	// Thông tin phân trang
	Paging interface{} `json:"paging,omitempty"`

	// Thông tin filter để đối ứng với client để xem là client truyền filter đó thì bên Backend có nhận đủ những filter đó không
	Filter interface{} `json:"filter,omitempty"`
}

// Hàm này giả về một response thành công
func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

// Có một số trường hợp không có paging hoặc filter, nên cần có hàm này để tạo response
func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
