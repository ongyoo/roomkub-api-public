package room

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	meterModel "github.com/ongyoo/roomkub-api/pkg/model/meter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomStatus string

const (
	RoomStatusAvailable RoomStatus = "available"  // ว่าง
	RoomStatusBusy      RoomStatus = "busy"       // ไม่ว่าง
	RoomStatusReportOut RoomStatus = "report_out" // แจ้งออก
	RoomStatusRepair    RoomStatus = "repair"     // ซ่อม
)

type RoomType string

const (
	RoomTypeHotel     RoomType = "hotel"     // โรงแรม
	RoomTypeApartment RoomType = "apartment" // apartment
	RoomTypeHouse     RoomType = "house"     // บ้าน
)

type RoomSubType string

const (
	RoomSubTypeOneBedroom RoomSubType = "one_bedroom" // ห้อง 1 เตียง
	RoomSubTypeOther      RoomSubType = "other"       // อื่นๆ
)

type RoomPaymentStatus string

const (
	RoomPaymentStatusPaid              RoomPaymentStatus = "paid"                  // ชำระแล้ว
	RoomPaymentStatusWaitingForPayment RoomPaymentStatus = "waiting_for_payment"   // รอชำระ
	RoomPaymentStatusWaitingForApprove RoomPaymentStatus = "waiting_for_approve"   // รอตรวจสอบชำระ
	RoomPaymentStatusCashPayment       RoomPaymentStatus = "cash_payment"          // ชำระเงินสด
	RoomPaymentStatusBankTransfer      RoomPaymentStatus = "payment_bank_transfer" // ชำระด้วยการโอนเงิน
	RoomPaymentStatusOverdue           RoomPaymentStatus = "overdue"               // ค้างชำระ
	RoomPaymentStatusNone              RoomPaymentStatus = "none"                  // ไม่มีสถานะ
)

type Room struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID         primitive.ObjectID `json:"business_id" bson:"businessId,omitempty"`
	Name               mongo.Encrypted    `json:"name" bson:"name"`
	Description        mongo.Encrypted    `json:"description" bson:"description"` // max 256 char
	ThumbnailURlPublic string             `json:"thumbnail_public_url" bson:"thumbnailPublicURl"`
	ThumbnailURl       string             `json:"thumbnail_url" bson:"thumbnailURl"`
	Images             []RoomImage        `json:"images" bson:"images"`
	Floor              int16              `json:"floor" bson:"floor"`
	Price              float32            `json:"price" bson:"price"`                // ราคาปกติ เช่น 2500
	SpecialPrice       float32            `json:"special_price" bson:"specialPrice"` // ต้องแสดงราคาพิเศษ เช่น 2300
	DailyPrice         float32            `json:"daily_price" bson:"dailyPrice"`     // ราคารายวัน
	IsShowSpecialPrice bool               `json:"is_special_price" bson:"isSpecialPrice"`
	Tags               []string           `json:"tags" bson:"tags"`
	Location           RoomLocation       `json:"location" bson:"location"`
	IsActive           bool               `json:"is_active" bson:"isActive"`
	IsPublisher        bool               `json:"is_publisher" bson:"isPublisher"`
	Status             RoomStatus         `json:"status" bson:"status"`
	Type               RoomType           `json:"type" bson:"type"`
	SubType            RoomSubType        `json:"sub_type" bson:"subType"`
	PaymentStatus      RoomPaymentStatus  `json:"payment_status" bson:"paymentStatus"`
	CurrentMeter       meterModel.Meter   `json:"current_meter" bson:"currentMeter"`
	Remark             string             `json:"remark" bson:"remark"`
	CreateBy           primitive.ObjectID `json:"create_by" bson:"createBy"`
	CreateAt           time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt           time.Time          `json:"update_at" bson:"updateAt"`
}

type RoomLocation struct {
	Name      string    `json:"name" bson:"name"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   `json:"longitude" bson:"longitude"`
	CreateAt  time.Time `json:"create_at" bson:"createAt"`
	UpdateAt  time.Time `json:"update_at" bson:"updateAt"`
}

type RoomImage struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                string             `json:"name" bson:"name"`
	ThumbnailPrivateURl string             `json:"thumbnail_private_url" bson:"thumbnailPrivateURl"`
	ThumbnailPublicURl  string             `json:"thumbnail_public_url" bson:"thumbnailPublicURl"`
	IsActive            bool               `json:"is_active" bson:"isActive"`
	UpdateAt            time.Time          `json:"update_at" bson:"updateAt"`
}

// Request
type CreateRoomRequest struct {
	Room
}

type UpdateRoomRequest struct {
	IDStr         string `json:"id"`
	BusinessIDStr string `json:"business_id"`
	ImageName     string `json:"name"`
	BusinessID    primitive.ObjectID
	IsRoot        bool
	Room
}

type DeleteImageRoomRequest struct {
	IDStr         string `json:"id"`
	BusinessIDStr string `json:"business_id"`
	ImageIDStr    string `json:"image_id"`
	ID            primitive.ObjectID
	BusinessID    primitive.ObjectID
	ImageID       primitive.ObjectID
	UpdateAt      time.Time
	IsRoot        bool
}

type UpdateRoomPaymentStatusRequest struct {
	IDStr            string `json:"id"`
	BusinessIDStr    string `json:"business_id"`
	PaymentStatusStr string `json:"payment_status"`
	ID               primitive.ObjectID
	PaymentStatus    RoomPaymentStatus
	BusinessID       primitive.ObjectID
	IsRoot           bool
	UpdateAt         time.Time
}

type GetRoomListRequest struct {
	BusinessID primitive.ObjectID
	IsRoot     bool
	Page       int64 `uri:"page"`
	Limit      int64 `uri:"limit"`
}

type FindRoomRequest struct {
	IDStr         string `uri:"id"`
	BusinessIDStr string `uri:"business_id"`
	BusinessID    primitive.ObjectID
	ID            primitive.ObjectID
	IsRoot        bool
}
