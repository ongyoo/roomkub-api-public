package roomContract

import (
	"time"

	"github.com/ongyoo/roomkub-api/pkg/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomContract struct {
	ID               primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	BusinessID       primitive.ObjectID      `json:"business_id" bson:"businessId,omitempty"`
	CustomerID       primitive.ObjectID      `json:"customer_id" bson:"customerId"`
	RoomID           primitive.ObjectID      `json:"room_id" bson:"roomId"`
	ContractTemplate ContractTemplateContent `json:"contract_template" bson:"contractTemplate"`
	Documents        []document.Document     `json:"documents" bson:"documents"`
	Remark           string                  `json:"remark" bson:"remark"`
	DepositTotal     float32                 `json:"deposit_total" bson:"depositTotal"`    // ยอดเงินมัดจำ
	PaidTotal        float32                 `json:"paid_total" bson:"paidTotal"`          // ยอดเงินมัดจำที่จ่ายแล้ว
	IsInstallment    bool                    `json:"is_installment " bson:"isInstallment"` // ขอผ่อนชำระ
	IsActive         bool                    `json:"is_active" bson:"isActive"`
	IsReportOut      bool                    `json:"is_report_out" bson:"isReportOut"` // แจ้งออก
	ReportOutAt      time.Time               `json:"report_out_at" bson:"reportOutAt"` // ออกวันที่
	CreateBy         primitive.ObjectID      `json:"create_by" bson:"createBy"`        // ผู้ออกสัญญา
	IsMigrate        bool                    `json:"is_migrate" bson:"isMigrate"`      // สำหรับย้ายเอกสารระบบเก่า
	CreateAt         time.Time               `json:"create_at" bson:"createAt"`
	UpdateAt         time.Time               `json:"update_at" bson:"updateAt"`
}

type ContractTemplateContent struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RefID    primitive.ObjectID `json:"ref_id" bson:"refId"`
	Name     string             `json:"name" bson:"name"`
	Html     string             `json:"html" bson:"html"`
	Remark   string             `json:"remark" bson:"remark"`
	CreateAt time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt time.Time          `json:"update_at" bson:"updateAt"`
}

// Request
type CreateRoomContractRequest struct {
	RoomContract
	CustomerIDStr         string `json:"customer_id"`
	ContractTemplateIDStr string `json:"contract_template_id"`
	RoomIDStr             string `json:"room_id"`
	CreateDateStr         string `json:"create_date"`
}

type UpdateRoomContractRequest struct {
	IDStr         string `json:"id"`
	BusinessIDStr string `json:"business_id"`
	BusinessID    primitive.ObjectID
	IsRoot        bool
	RoomContract
}

type FindRoomContractRequest struct {
	IDStr         string `uri:"id"`
	BusinessIDStr string `uri:"business_id"`
	BusinessID    primitive.ObjectID
	ID            primitive.ObjectID
	IsRoot        bool
}

type GetRoomContractRequest struct {
	RoomID        primitive.ObjectID
	BusinessID    primitive.ObjectID
	IsRoot        bool
	BusinessIDStr string `uri:"business_id"`
	RoomIDStr     string `uri:"room_id"`
	Page          int64  `uri:"page"`
	Limit         int64  `uri:"limit"`
}
