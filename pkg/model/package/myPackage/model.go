package mypackage

import (
	"time"

	PackageItem "github.com/ongyoo/roomkub-api/pkg/model/package"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MyPackage struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChannelId    primitive.ObjectID `json:"channel_id" bson:"channelId"`
	PackageItems *[]MyPackageItem   `json:"package_items" bson:"packageItems"`
	CreateAt     time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt     time.Time          `json:"update_at" bson:"updateAt"`
}

type MyPackageItem struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Package     PackageItem.Package `json:"package" bson:"package"`
	UpdateAt    time.Time           `json:"update_at" bson:"updateAt"`
	StartAt     time.Time           `json:"start_at" bson:"startAt"`
	ExpiredAt   time.Time           `json:"expired_at" bson:"expiredAt"`
	IsActive    bool                `json:"is_active" bson:"updateAt"`
	IsExpired   bool                `json:"is_expired" bson:"isExpired"`
	IsSubscribe bool                `json:"is_subscribe" bson:"isSubscribe" default:"false"`
}
