package contractTemplate

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContractTemplate struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID  primitive.ObjectID `json:"business_id" bson:"businessId,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Html        string             `json:"html" bson:"html"`
	Remark      string             `json:"remark" bson:"remark"`
	CreateBy    primitive.ObjectID `json:"create_by" bson:"createBy"`
	IsActive    bool               `json:"is_active" bson:"isActive"`
	IsPublisher bool               `json:"is_publisher" bson:"isPublisher"` // สามารถให้คนอื่นใช้ฟอร์มได้
	CreateAt    time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt    time.Time          `json:"update_at" bson:"updateAt"`
}

type CreateContractTemplateRequest struct {
	ContractTemplate
}

type GetContractTemplateRequest struct {
	BusinessID primitive.ObjectID
	MyChannel  bool
	IsRoot     bool
	Page       int64 `uri:"page"`
	Limit      int64 `uri:"limit"`
}

type UpdateContractTemplateRequest struct {
	IDStr       string `uri:"id" json:"id"`
	ID          primitive.ObjectID
	BusinessID  primitive.ObjectID
	CreateBy    primitive.ObjectID
	IsRoot      bool
	Name        string `json:"name"`
	Html        string `json:"html"`
	Remark      string `json:"remark"`
	IsActive    bool   `json:"is_active"`
	IsPublisher bool   `json:"is_publisher"`
	CreateAt    time.Time
	UpdateAt    time.Time
}

type GetContractTemplateDetailRequest struct {
	ID         primitive.ObjectID
	BusinessID primitive.ObjectID
	IsRoot     bool
	IDStr      string `uri:"id"`
	//BusinessIDStr string `uri:"business_id"`
}
