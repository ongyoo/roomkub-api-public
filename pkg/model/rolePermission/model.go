package rolePermission

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRoleType string

const (
	UserRoleTypeRootAdminUser UserRoleType = "ROOT_ADMIN_USER"
	UserRoleTypeUser          UserRoleType = "USER"
	UserRoleTypeSubUser       UserRoleType = "SUB_USER"
)

type UserRole struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	BusinessID    primitive.ObjectID   `json:"business_id" bson:"businessId"`
	Name          mongo.Encrypted      `json:"name" bson:"name"`
	Slug          string               `json:"slug" bson:"slug"`
	Type          []UserRoleType       `json:"type" bson:"type"` // ROOT_ADMIN_USER | USER | SUB_USER
	PermissionIDs []primitive.ObjectID `json:"permission_ids" bson:"permissionIds"`
	IsActive      bool                 `json:"is_active" bson:"isActive"`
	CreateAt      time.Time            `json:"create_at" bson:"createAt"`
	UpdateAt      time.Time            `json:"update_at" bson:"updateAt"`
}

type UserPermission struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID        primitive.ObjectID `json:"business_id" bson:"businessId"`
	Path              mongo.Encrypted    `json:"path" bson:"path"`
	Name              mongo.Encrypted    `json:"name" bson:"name"`
	Type              []UserRoleType     `json:"type" bson:"type"` // ROOT_ADMIN_USER | USER | SUB_USER
	AllowGetMethod    bool               `json:"allow_get_Method" bson:"allowGetMethod"`
	AllowPostMethod   bool               `json:"allow_post_Method" bson:"allowPostMethod"`
	AllowPutMethod    bool               `json:"allow_put_Method" bson:"allowPutMethod"`
	AllowDeleteMethod bool               `json:"allow_delete_Method" bson:"allowDeleteMethod"`
	IsActive          bool               `json:"is_active" bson:"isActive"`
	CreateAt          time.Time          `json:"create_at" bson:"createAt"`
	UpdateAt          time.Time          `json:"update_at" bson:"updateAt"`
}

// Payment `bson:",inline"`
// Request
type UserRoleCreateFormRequest struct {
	UserRole
	IDStr            string   `json:"id"`
	BusinessIDStr    string   `json:"business_id"`
	PermissionStrIDs []string `json:"permission_ids"`
	IsRoot           bool
}

type UserPermissionCreateFormRequest struct {
	UserPermission
	IDStr         string `json:"id"`
	BusinessIDStr string `json:"business_id"`
	IsRoot        bool
}

type GetUserRoleListRequest struct {
	Page          int64          `uri:"page"`
	Limit         int64          `uri:"limit"`
	Active        int8           `uri:"active"` // 0 = false | 1 = true | 2 = all
	TypeStr       string         `uri:"type"`
	Type          []UserRoleType // ROOT_ADMIN_USER | USER | SUB_USER
	BusinessIDStr string         `uri:"business_id"`
	BusinessID    primitive.ObjectID
	IsRoot        bool
}

type GetUserPermissionListRequest struct {
	Page          int64          `uri:"page"`
	Limit         int64          `uri:"limit"`
	Active        int8           `uri:"active"` // 0 = false | 1 = true | 2 = all
	TypeStr       string         `uri:"type"`
	Type          []UserRoleType // ROOT_ADMIN_USER | USER | SUB_USER
	BusinessIDStr string         `uri:"business_id"`
	BusinessID    primitive.ObjectID
	IsRoot        bool
}

type GetUserRoleByIdRequest struct {
	IDStr         string `uri:"id"`
	BusinessIDStr string `uri:"business_id"`
	ID            primitive.ObjectID
	BusinessID    primitive.ObjectID
	IsRoot        bool
}

type GetRolePermissionByIdRequest struct {
	IDStr         string `uri:"id"`
	BusinessIDStr string `uri:"business_id"`
	BusinessID    primitive.ObjectID
	ID            primitive.ObjectID
	IsRoot        bool
}

type GetUserRoleByIdsRequest struct {
	Page          int64  `uri:"page"`
	Limit         int64  `uri:"limit"`
	IDsStr        string `uri:"id"` // id,id
	BusinessIDStr string `uri:"business_id"`
	BusinessID    primitive.ObjectID
	IDs           []primitive.ObjectID
	IsRoot        bool
}

type GetRolePermissionByIdsRequest struct {
	Page          int64  `uri:"page"`
	Limit         int64  `uri:"limit"`
	IDsStr        string `uri:"id"`
	BusinessIDStr string `uri:"business_id"`
	BusinessID    primitive.ObjectID
	IDs           []primitive.ObjectID
	IsRoot        bool
}
