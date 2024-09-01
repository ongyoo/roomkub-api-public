package packageList

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PackageTag string

const (
	PackageTagRecommend PackageTag = "RECOMMEND"
	PackageTagNew       PackageTag = "NEW"
	PackageTagSave      PackageTag = "SAVE"
	PackageTagHot       PackageTag = "HOT"
	PackageTagDiscount  PackageTag = "DISCOUNT"
)

type Package struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             mongo.Encrypted    `json:"name" bson:"name"`
	NameIdentityID   string             `json:"name_identity_id" bson:"nameIdentityId"`
	Detail           mongo.Encrypted    `json:"detail" bson:"detail"`
	Meta             PackageMeta        `json:"meta" bson:"meta"`
	Price            float64            `json:"price" bson:"price" default:"0.0"`
	DiscountRate     float64            `json:"discount_rate" bson:"discountRate" default:"0.0"` // 10% = 10
	UseLimitDay      int16              `json:"use_limit_day" bson:"useLimitDay" default:"0"`    // 30 = 30 day
	IsShowStoreFront bool               `json:"is_show_store_front" bson:"isShowStoreFront"`
	IsDiscount       bool               `json:"is_discount" bson:"isDiscount"`
	IsActive         bool               `json:"is_active" bson:"updateAt"`
	Tags             []PackageTag       `json:"tags" bson:"tags"`
	CreateAt         time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt         time.Time          `json:"update_at" bson:"updateAt"`
}

type PackageMeta struct { // this set limit package
	MemberLimit int16 `json:"member_limit" bson:"memberLimit" default:"2"`
	RoomLimit   int16 `json:"room_limit" bson:"roomLimit" default:"10"`
}
