package businessChannel

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	role "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BusinessChannelType string

const (
	BusinessChannelTypeRootSystem BusinessChannelType = "ROOT_SYSTEM" // ดูแลระบบ
	BusinessChannelTypeResidence  BusinessChannelType = "RESIDENCE"   // หอพัก
	BusinessChannelTypeApartment  BusinessChannelType = "APARTMENT"   // อพาร์ทเม้น
	BusinessChannelTypeHotel      BusinessChannelType = "HOTEL"       // โรงแรม
	BusinessChannelTypeHostel     BusinessChannelType = "HOSTEL"      // โฮสเทล
	BusinessChannelTypeAgent      BusinessChannelType = "AGENT"       // นายหน้า
)

type BusinessChannel struct {
	ID          primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	ChannelID   primitive.ObjectID      `json:"channel_id" bson:"channelId,omitempty"`
	Members     []BusinessChannelMember `json:"members" bson:"members"`
	Owner       primitive.ObjectID      `json:"owner" bson:"owner"`
	Setting     BusinessChannelSetting  `json:"setting" bson:"setting"`
	Type        BusinessChannelType     `json:"type" bson:"type"`
	IsActive    bool                    `json:"is_active" bson:"isActive"`
	IsBaned     bool                    `json:"is_ban" bson:"isBan"`
	MyPackageID primitive.ObjectID      `json:"my_package_id" bson:"myPackageId"`
	CreateAt    time.Time               `json:"create_at" bson:"createAt"`
	UpdateAt    time.Time               `json:"update_at" bson:"updateAt"`
	Remark      string                  `json:"remark" bson:"remark"`
}

type BusinessChannelMember struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Role       role.UserRole      `json:"role" bson:"role"`
	IsBaned    bool               `json:"is_ban" bson:"isBan"`
	IsSeparate bool               `json:"is_separate" bson:"isSeparate"`
	Remark     string             `json:"remark" bson:"remark"`
}

type BusinessChannelSetting struct {
	Name             mongo.Encrypted `json:"name" bson:"name"`
	LogoThumbnailURl string          `json:"logo_thumbnail_url" bson:"logoThumbnailURl"`
}

// Request
type BusinessChannelCreateFormRequest struct {
	BusinessChannel
}

type UpdateDeleteBusinessFormRequest struct {
	ID     string `json:"id"`
	Remark string `json:"remark"`
}

type UpdateBanBusinessFormRequest struct {
	ID     string `json:"id"`
	Remark string `json:"remark"`
}

type UpdateBusinessChannelMyPackageFormRequest struct {
	ID             string `json:"id"`
	MyPackageIDStr string `json:"package_id"`
	MyPackageID    primitive.ObjectID
}

type UpdateBusinessChannelSettingFormRequest struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	LogoThumbnailURl string `json:"logo_thumbnail_url"`
}

type GetBusinessChannelListRequest struct {
	Page  int64 `uri:"page"`
	Limit int64 `uri:"limit"`
}

type GetBusinessChannelListByUserIdRequest struct {
	Page   int64 `uri:"page"`
	Limit  int64 `uri:"limit"`
	UserID primitive.ObjectID
}
