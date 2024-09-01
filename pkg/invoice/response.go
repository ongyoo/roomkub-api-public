package Invoice

import (
	"time"

	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	userModel "github.com/ongyoo/roomkub-api/pkg/model/user"
)

//"github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/pipeline"

type BaseInvoiceItemResponse[T any] struct {
	Customer T `json:"invoice"`
}

type BaseInvoiceItemsResponse[T any] struct {
	Customer T `json:"invoice_list"`
}

type CreateRoomResponse struct {
	RoomID *string `json:"invoice_id"`
}

type InvoiceResponse struct {
	ID             string        `json:"id"`
	BusinessID     string        `json:"business_id"`
	RoomName       string        `json:"room_name"`
	CustomerName   string        `json:"customer_name"`
	CreateUserName string        `json:"create_by_user_name"`
	TotalPrice     float32       `json:"total_price"`
	Status         InvoiceStatus `json:"status"`
	Type           InvoiceType   `json:"type"`
	IsSendNotify   bool          `json:"is_send_notify"`
	IsPrinted      bool          `json:"is_printed"`
	CreateAt       time.Time     `json:"create_at"`
	UpdateAt       time.Time     `json:"update_at"`
}

type InvoiceDetailResponse struct {
	ID           string                `json:"id"`
	BusinessID   string                `json:"business_id"`
	RoomName     string                `json:"room_name"`
	Customer     UserOrCustomerRespons `json:"customer"`
	CreateByUser userModel.User        `json:"create_by_user"`
	Items        []InvoiceItem         `json:"items"`
	TotalPrice   float32               `json:"total_price"`
	Status       InvoiceStatus         `json:"status"`
	Type         InvoiceType           `json:"type"`
	IsSendNotify bool                  `json:"is_send_notify"`
	IsPrinted    bool                  `json:"is_printed"`
	CreateAt     time.Time             `json:"create_at"`
	UpdateAt     time.Time             `json:"update_at"`
	SendNotifyAt time.Time             `json:"send_notify_at"`
	PrintedAt    time.Time             `json:"printed_at"`
}

type UserOrCustomerResponseData interface {
	isResponseData()
}

type UserData struct {
	*userModel.User
}

func (i UserData) isResponseData() {}

type CustomerData struct {
	*customer.CustomerDetailResponse
}

func (s CustomerData) isResponseData() {}

type UserOrCustomerRespons struct {
	Data UserOrCustomerResponseData
}
