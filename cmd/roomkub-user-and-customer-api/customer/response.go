package customer

import (
	"time"

	"github.com/ongyoo/roomkub-api/pkg/document"
)

type BaseCustomerItemResponse[T any] struct {
	Customer T `json:"customer"`
}

type BaseCustomerItemsResponse[T any] struct {
	Customer T `json:"customer_list"`
}

type CustomerResponse struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	NickName     string    `json:"nick_name"`
	Gender       string    `json:"gender_type"`
	Phone        string    `json:"phone"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	IsActive     bool      `json:"is_active"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}

type CustomerDetailResponse struct {
	ID           string              `json:"id"`
	Email        string              `json:"email"`
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	NickName     string              `json:"nick_name"`
	NID          string              `json:"n_id"`
	Gender       string              `json:"gender_type"`
	Address      string              `json:"address"`
	Province     string              `json:"province"`
	PostCode     string              `json:"post_code"`
	Phone        string              `json:"phone"`
	ThumbnailUrl string              `json:"thumbnail_url"`
	Documents    []document.Document `json:"documents"`
	CreateAt     time.Time           `json:"create_at"`
	UpdateAt     time.Time           `json:"update_at"`
}
