package user

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email               mongo.Encrypted    `json:"email" bson:"email"`
	EmailIdentityID     string             `json:"email_identity_id" bson:"emailIdentityId"`
	Password            string             `json:"password" bson:"password"`
	FirstName           mongo.Encrypted    `json:"first_name" bson:"firstName"`
	FirstNameIdentityID string             `json:"first_name_identity_id" bson:"firstNameIdentityId"`
	LastName            mongo.Encrypted    `json:"last_name" bson:"lastName"`
	NickName            mongo.Encrypted    `json:"nick_name" bson:"nickName"`
	NID                 mongo.Encrypted    `json:"n_id" bson:"nId"`
	NIdentityID         string             `json:"n_identity_id" bson:"nIdentityId"`
	Address             mongo.Encrypted    `json:"address" bson:"address"`
	Province            string             `json:"province" bson:"province"`
	PostCode            string             `json:"post_code" bson:"postCode"`
	Phone               mongo.Encrypted    `json:"phone" bson:"phone"`
	PhoneIdentityID     string             `json:"phone_id" bson:"phoneId"`
	ThumbnailURl        string             `json:"thumbnail_url" bson:"thumbnailURl"`
	RoleID              primitive.ObjectID `json:"role_id" bson:"roleId"`
	IsActive            bool               `json:"is_active" bson:"isActive" default:"true"`
	IsBaned             bool               `json:"is_baned" bson:"isBaned" default:"false"`
	IsVerify            bool               `json:"is_verify" bson:"isVerify" default:"false"`
	Slug                mongo.Encrypted    `json:"slug" bson:"slug"`
	SlugIdentityID      string             `json:"slug_identity_id" bson:"slugIdentityId"`
	CreateAt            time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt            time.Time          `json:"update_at" bson:"updateAt"`
}

// Request
type UserCreateFormRequest struct {
	User
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserListRequest struct {
	Page  int64 `uri:"page"`
	Limit int64 `uri:"limit"`
}

type GetUserListByIDsRequest struct {
	Page      int64  `uri:"page"`
	Limit     int64  `uri:"limit"`
	UserIDStr string `uri:"ids"`
	UserID    []string
	IDs       []primitive.ObjectID
}

type GetUserListByChannelIdRequest struct {
	Page      int64  `uri:"page"`
	Limit     int64  `uri:"limit"`
	ChannelID string `uri:"channel_id"`
	IDs       []primitive.ObjectID
}

type GetUserByIdRequest struct {
	UserID string `uri:"id"`
	ID     primitive.ObjectID
}

// update
type UpdateUserFormRequest struct {
	UserID string `json:"id"`
	User
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password"`
}
