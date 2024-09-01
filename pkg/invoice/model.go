package Invoice

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceStatus string

const (
	InvoiceStatusPaid InvoiceStatus = "PAID"
	InvoiceStatusWait InvoiceStatus = "WAIT"
)

type InvoiceType string

const (
	InvoiceTypeRoomFeeDaily   InvoiceType = "ROOM_FEE_D"  // รายวัน
	InvoiceTypeRoomFeeMonthly InvoiceType = "ROOM_FEE_M"  // รายเดือน
	InvoiceTypeServiceFee     InvoiceType = "SERVICE_FEE" // ค่าบริการใช้เซอร์วิส
	InvoiceTypeInstallment    InvoiceType = "INSTALLMENT" // ผ่อนชำระ
	InvoiceTypeOther          InvoiceType = "OTHER"       // อื่นๆ เช่น เบิก จ่ายบิล สำหรับพนักงานเท่านั้น
)

type Invoice struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID     primitive.ObjectID `json:"business_id" bson:"businessId,omitempty"`
	RoomPipelineID primitive.ObjectID `json:"room_pipeline_id" bson:"roomPipelineId"`
	CustomerID     primitive.ObjectID `json:"customer_id" bson:"customerId"`
	Items          []InvoiceItem      `json:"items" bson:"items"`
	TotalPrice     float32            `json:"total_price" bson:"totalPrice"`
	Remark         string             `json:"remark" bson:"remark"`
	Status         InvoiceStatus      `json:"status" bson:"status"`
	Type           InvoiceType        `json:"type" bson:"type"`
	IsActive       bool               `json:"is_active" bson:"isActive"`
	IsSendNotify   bool               `json:"is_send_notify" bson:"isSendNotify"`
	IsPrinted      bool               `json:"is_printed" bson:"isPrinted"`
	CreateAt       time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt       time.Time          `json:"update_at" bson:"updateAt"`
	SendNotifyAt   time.Time          `json:"send_notify_at" bson:"sendNotifyAt"`
	PrintedAt      time.Time          `json:"printed_at" bson:"printedAt"`
	CreateUserID   primitive.ObjectID `json:"create_user_id" bson:"createUserId"` // ผู้สร้าง
}

type InvoiceItem struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Price      float32            `json:"price" bson:"price"`
	Unit       int                `json:"unit" bson:"unit"`
	TotalPrice float32            `json:"total_price" bson:"totalPrice"`
	Ref        string             `json:"ref" bson:"ref"`
	Remark     string             `json:"remark" bson:"remark"`
	IsActive   bool               `json:"is_active" bson:"isActive"`
	IsApproved bool               `json:"is_approved" bson:"isApproved"`
	CreateAt   time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt   time.Time          `json:"update_at" bson:"updateAt"`
}

// Request
type CreateInvoiceRequest struct {
	Invoice
}

type GetInvoiceListRequest struct {
	RoomPipelineID    primitive.ObjectID
	BusinessID        primitive.ObjectID
	CustomerID        primitive.ObjectID
	Status            *InvoiceStatus
	Type              *InvoiceType
	IsRoot            bool
	BusinessIDStr     string `uri:"business_id"`
	RoomPipelineIDStr string `uri:"room_pipeline_id"`
	CustomerIDStr     string `uri:"customer_id"`
	Page              int64  `uri:"page"`
	Limit             int64  `uri:"limit"`
}

type FindInvoiceRequest struct {
	IDStr      string `uri:"id"`
	BusinessID primitive.ObjectID
	ID         primitive.ObjectID
	IsRoot     bool
}

type UpdateInvoiceRequest struct {
	IDStr      string `uri:"id"`
	BusinessID primitive.ObjectID
	ID         primitive.ObjectID
	IsRoot     bool
	Invoice
}
