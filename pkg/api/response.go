package api

type APIResponse[T any] struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"not found"`
	Result  T      `json:"result"`
}

type APIMessage struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"not found"`
}

type APIErrorMessage struct {
	ErrorCode uint   `json:"error_code" example:"404"`
	Message   string `json:"message" example:"not found"`
}

type PaginatedContent[T any] struct {
	APIResponse[T]
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}
