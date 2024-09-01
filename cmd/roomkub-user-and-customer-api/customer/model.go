package customer

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"github.com/ongyoo/roomkub-api/pkg/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerStatus string

const (
	CustomerStatusBad         CustomerStatus = "bad"          // แย่ point 1
	CustomerStatusOverdue     CustomerStatus = "overdue"      // ค้างชำระ point 2
	CustomerStatusLatePayment CustomerStatus = "late_payment" // ชำระล่าช้า point 3
	CustomerStatusNormal      CustomerStatus = "normal"       // ปกติ point 4
	CustomerStatusGood        CustomerStatus = "good"         // ดีทุกอย่าง point 5
)

type CustomerGenderType string

const (
	CustomerGenderTypeMen   CustomerGenderType = "men"   // ผู้ชาย
	CustomerGenderTypeWomen CustomerGenderType = "women" // ผู้หญิง
	CustomerGenderTypeLGBTQ CustomerGenderType = "lgbt"  // LGBTQ
	CustomerGenderTypeNone  CustomerGenderType = "none"  // ไม่ระบุเพศ
)

type Customer struct {
	ID                  primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Email               mongo.Encrypted     `json:"email" bson:"email"`
	EmailIdentityID     string              `json:"email_identity_id" bson:"emailIdentityId"`
	Password            string              `json:"password" bson:"password"`
	FirstName           mongo.Encrypted     `json:"first_name" bson:"firstName"`
	FirstNameIdentityID string              `json:"first_name_identity_id" bson:"firstNameIdentityId"`
	LastName            mongo.Encrypted     `json:"last_name" bson:"lastName"`
	NickName            mongo.Encrypted     `json:"nick_name" bson:"nickName"`
	NID                 mongo.Encrypted     `json:"n_id" bson:"nId"`
	NIdentityID         string              `json:"n_identity_id" bson:"nIdentityId"`
	GenderType          CustomerGenderType  `json:"gender_type" bson:"genderType"`
	Address             mongo.Encrypted     `json:"address" bson:"address"`
	Province            string              `json:"province" bson:"province"`
	PostCode            string              `json:"post_code" bson:"postCode"`
	Phone               mongo.Encrypted     `json:"phone" bson:"phone"`
	PhoneIdentityID     string              `json:"phone_identity_id" bson:"phoneIdentityId"`
	ThumbnailURl        string              `json:"thumbnail_url" bson:"thumbnailURl"`
	RoleID              primitive.ObjectID  `json:"role_id" bson:"roleId"`
	Documents           []document.Document `json:"documents" bson:"documents"`
	IsActive            bool                `json:"is_active" bson:"isActive" default:"true"`
	LineRef             string              `json:"line_ref" bson:"lineRef"`
	Line                LineModel           `json:"line" bson:"line"`
	IsLineLiff          bool                `json:"is_line_liff" bson:"isLineLiff" default:"false"`
	CreateAt            time.Time           `json:"create_at" bson:"createAt"`
	UpdateAt            time.Time           `json:"update_at" bson:"updateAt"`
}

type LineModel struct {
	LineToken mongo.Encrypted `json:"line_token" bson:"lineToken"`
	LineRef   mongo.Encrypted `json:"line_ref" bson:"lineRef"`
}

// toDo new pipeline
/*
Status              CustomerStatus       `json:"status" bson:"status"`
	VoteScore           int                  `json:"vote_score" bson:"vote_score"`
*/

// Request
type CreateCustomerRequest struct {
	Customer
}

type FindCustomerRequest struct {
	IDStr     string `uri:"id"`
	FirstName string `uri:"first_name"`
	Email     string `uri:"email"`
	Password  string `uri:"password"`
	NID       string `uri:"n_id"`
	Phone     string `uri:"phone"`
	LineRef   string `uri:"line_ref"`
	ID        primitive.ObjectID
}

type UpdateCustomerFormRequest struct {
	IDStr string `json:"id"`
	Customer
}

type GetCustomerListRequest struct {
	Page  int64 `uri:"page"`
	Limit int64 `uri:"limit"`
}

type GetCustomerByIdRequest struct {
	IDStr string `uri:"id"`
	ID    primitive.ObjectID
}

type GetCustomerListByIDsRequest struct {
	Page      int64  `uri:"page"`
	Limit     int64  `uri:"limit"`
	UserIDStr string `uri:"ids"`
	UserID    []string
	IDs       []primitive.ObjectID
}
