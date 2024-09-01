package room

import (
	"time"
)

type BaseRoomItemResponse[T any] struct {
	Customer T `json:"room"`
}

type BaseRoomItemsResponse[T any] struct {
	Customer T `json:"room_list"`
}

type CreateRoomResponse struct {
	RoomID *string `json:"room_id"`
}

type RoomResponse struct {
	ID                 string            `json:"id"`
	BusinessID         string            `json:"business_id"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	ThumbnailUrl       string            `json:"thumbnail_url"`
	Floor              int16             `json:"floor"`
	Price              float32           `json:"price"`
	SpecialPrice       float32           `json:"special_price"`
	IsShowSpecialPrice bool              `json:"is_special_price"`
	Tags               []string          `json:"tags"`
	Location           string            `json:"location"`
	Status             RoomStatus        `json:"status"`
	PaymentStatus      RoomPaymentStatus `json:"payment_status"`
	Type               RoomType          `json:"type"`
	SubType            RoomSubType       `json:"sub_type"`
	IsPublisher        bool              `json:"is_publisher"`
	IsActive           bool              `json:"is_active"`
	CreateAt           time.Time         `json:"create_at"`
	UpdateAt           time.Time         `json:"update_at"`
}
