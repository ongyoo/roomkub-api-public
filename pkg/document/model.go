package document

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentType string

const (
	DocumentTypeIdCard           DocumentType = "id-card"
	DocumentTypePaymentSlip      DocumentType = "payment_slip"
	DocumentTypeWaterMeter       DocumentType = "water_meter"       // มิเตอร์น้ำ
	DocumentTypeElectricityMeter DocumentType = "electricity_meter" // มิเตอร์ไฟฟ้า
	DocumentTypeContract         DocumentType = "contract"          // เอกสารสัญญา
	DocumentTypeSignature        DocumentType = "signature"         // ลายเซ็น
	DocumentTypeOther            DocumentType = "other"
)

type Document struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID   primitive.ObjectID `json:"business_id" bson:"businessId"`
	RefID        primitive.ObjectID `json:"ref_id" bson:"refId"`
	Name         mongo.Encrypted    `json:"name" bson:"name"`
	Description  mongo.Encrypted    `json:"description" bson:"description"`
	Remark       mongo.Encrypted    `json:"remark" bson:"remark"`
	FileName     mongo.Encrypted    `json:"file_name" bson:"fileName"`
	ContentType  mongo.Encrypted    `json:"content_type" bson:"contentType"`
	DocumentType DocumentType       `json:"document_type" bson:"documentType"`
	PrivateURL   mongo.Encrypted    `json:"private_url" bson:"privateUrl"`
	PublicURL    mongo.Encrypted    `json:"public_url" bson:"publicUrl"`
	IsActive     bool               `json:"is_active" bson:"isActive"`
	CreateAt     time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt     time.Time          `json:"update_at" bson:"updateAt"`
}

// Request
type CreateDocumentRequest struct {
	Document
}

type FindDocumentRequest struct {
	IDStr         string `json:"id"`
	RefIDStr      string `json:"ref_id"`
	Page          int64  `uri:"page"`
	Limit         int64  `uri:"limit"`
	ID            primitive.ObjectID
	BusinessID    primitive.ObjectID
	RefID         primitive.ObjectID
	IsRoot        bool
}

type UpdateDocumentRequest struct {
	IDStr         string `json:"id"`
	RefIDStr      string `json:"ref_id"`
	ID            primitive.ObjectID
	BusinessID    primitive.ObjectID
	RefID         primitive.ObjectID
	IsRoot        bool
	IsActive     bool
	UpdateAt     time.Time
}
