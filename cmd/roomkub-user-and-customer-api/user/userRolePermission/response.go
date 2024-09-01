package userRolePermission

import (
	"time"

	rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
)

type UserRoleResponse struct {
	ID          string                        `json:"id"`
	BusinessID  string                        `json:"business_id"`
	Name        string                        `json:"name"`
	Slug        string                        `json:"slug"`
	Type        []rolePermission.UserRoleType `json:"type"` // ROOT_ADMIN_USER | USER | SUB_USER
	Permissions []UserPermissionResponse      `json:"permissions"`
	IsActive    bool                          `json:"is_active"`
	CreateAt    time.Time                     `json:"create_at"`
	UpdateAt    time.Time                     `json:"update_at"`
}

type UserPermissionResponse struct {
	ID                string                        `json:"id"`
	BusinessID        string                        `json:"business_id"`
	Path              string                        `json:"path"`
	Name              string                        `json:"name"`
	Type              []rolePermission.UserRoleType `json:"type"` // ROOT_ADMIN_USER | USER | SUB_USER
	AllowGetMethod    bool                          `json:"allow_get_Method"`
	AllowPostMethod   bool                          `json:"allow_post_Method"`
	AllowPutMethod    bool                          `json:"allow_put_Method"`
	AllowDeleteMethod bool                          `json:"allow_delete_Method"`
	IsActive          bool                          `json:"is_active"`
	CreateAt          time.Time                     `json:"create_at"`
	UpdateAt          time.Time                     `json:"update_at"`
}

/*
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
*/
