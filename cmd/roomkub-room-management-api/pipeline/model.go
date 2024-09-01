package roomPipeline

import (
	"time"

	meterModel "github.com/ongyoo/roomkub-api/pkg/model/meter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type roomPipelineStatus string

const (
	roomPipelineStatusBordCast           roomPipelineStatus = "BORDCAST"             // แสดงงานออกให้กดรับงานจดมิเตอร์
	roomPipelineStatusAssigned           roomPipelineStatus = "ASSIGNED"             // มีคนรับมอบหมายแล้ว
	roomPipelineStatusInProgress         roomPipelineStatus = "INPROGRESS"           // กำลังทำงาน
	roomPipelineStatusWaitingForApproval roomPipelineStatus = "WAITING_FOR_APPROVAL" // รอตรวจสอบข้อมูล
	roomPipelineStatusApproved           roomPipelineStatus = "APPROVED"             // อนุมัติข้อมูล
	roomPipelineStatusReject             roomPipelineStatus = "REJECT"               // ปฏิเสธ
	roomPipelineStatusCancel             roomPipelineStatus = "CANCEL"               // ยกเลิก
)

type roomPipeline struct {
	ID             primitive.ObjectID           `json:"id" bson:"_id,omitempty"`
	BusinessID     primitive.ObjectID           `json:"business_id" bson:"businessId,omitempty"`
	ContractID     primitive.ObjectID           `json:"contract_id" bson:"contractId"`
	RoomName       string                       `json:"room_name" bson:"roomName"`
	Meter          meterModel.Meter             `json:"meter" bson:"meter"`
	AdminRemark    string                       `json:"admin_remark" bson:"adminRemark"`
	Remark         string                       `json:"remark" bson:"remark"`
	IsCurrent      bool                         `json:"is_current" bson:"isCurrent"`
	IsActive       bool                         `json:"is_active" bson:"isActive"`
	Status         roomPipelineStatus           `json:"status" bson:"status"`
	TimeLineStatus []roomPipelineStatusTimeLine `json:"time_line_status" bson:"timeLineStatus"`
	CreateAt       time.Time                    `json:"create_at" bson:"createAt"`
	UpdateAt       time.Time                    `json:"update_at" bson:"updateAt"`
	CreateBy       primitive.ObjectID           `json:"create_by" bson:"createBy"`     // admin
	RecorderBy     primitive.ObjectID           `json:"recorder_by" bson:"recorderBy"` // userId
}

type roomPipelineStatusTimeLine struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status   roomPipelineStatus `json:"status" bson:"status"`
	CreateAt time.Time          `json:"create_at" bson:"createAt"`
	CreateBy primitive.ObjectID `json:"create_by" bson:"createBy"` // admin
}

// Request
type CreateroomPipelineRequest struct {
	roomPipeline
}

type GetroomPipelineListRequest struct {
	ContractID    primitive.ObjectID
	BusinessID    primitive.ObjectID
	Status        *roomPipelineStatus
	IsRoot        bool
	BusinessIDStr string `uri:"business_id"`
	ContractIDStr string `uri:"contract_id"`
	Page          int64  `uri:"page"`
	Limit         int64  `uri:"limit"`
}

type FindroomPipelineRequest struct {
	IDStr      string `uri:"id"`
	BusinessID primitive.ObjectID
	ID         primitive.ObjectID
	IsRoot     bool
}

type UpdateroomPipelineRequest struct {
	IDStr      string `json:"id"`
	BusinessID primitive.ObjectID
	IsRoot     bool
	roomPipeline
}
