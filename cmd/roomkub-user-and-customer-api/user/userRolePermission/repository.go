package userRolePermission

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userRoleCollectionName       = "userRole"
	userPermissionCollectionName = "userPermission"

	// User Role
	userRoleFieldID            = "_id"
	userRoleFieldBusinessId    = "businessId"
	userRoleFieldName          = "name"
	userRoleFieldSlug          = "slug"
	userRoleFieldType          = "type"
	userRoleFieldPermissionIDs = "permissionIds"
	userRoleFieldIsActive      = "isActive"
	userRoleFieldCreateAt      = "createAt"
	userRoleFieldUpdateAt      = "updateAt"

	// User Permission
	userPermissionFieldID                = "_id"
	userPermissionBusinessId             = "businessId"
	userPermissionFieldPath              = "path"
	userPermissionFieldName              = "name"
	userPermissionFieldType              = "type"
	userPermissionFieldAllowGetMethod    = "allowGetMethod"
	userPermissionFieldAllowPostMethod   = "allowPostMethod"
	userPermissionFieldAllowPutMethod    = "allowPutMethod"
	userPermissionFieldAllowDeleteMethod = "allowDeleteMethod"
	userPermissionFieldIsActive          = "isActive"
	userPermissionFieldCreateAt          = "createAt"
	userPermissionFieldUpdateAt          = "updateAt"
)

type Repository interface {
	// insert
	InsertUserRole(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	InsertUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// update
	UpdateUserRoleInfo(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	UpdateUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	UpdateUserPermissionInfo(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// delete
	DeleteUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	DeleteUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// get
	GetUserRoleAllList(ctx context.Context, req rolePermission.GetUserRoleListRequest) ([]rolePermission.UserRole, *PaginatedData, error)
	GetUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*rolePermission.UserRole, error)
	GetUserRoleByIDs(ctx context.Context, req rolePermission.GetUserRoleByIdsRequest) ([]rolePermission.UserRole, *PaginatedData, error)
	GetUserPermissionAllList(ctx context.Context, req rolePermission.GetUserPermissionListRequest) ([]rolePermission.UserPermission, *PaginatedData, error)
	GetUserPermissionByID(ctx context.Context, req rolePermission.GetRolePermissionByIdRequest) (*rolePermission.UserPermission, error)
	GetUserPermissionByIDs(ctx context.Context, req rolePermission.GetRolePermissionByIdsRequest) ([]rolePermission.UserPermission, *PaginatedData, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

// insert
func (r *repository) InsertUserRole(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	_, err := r.db.Collection(userRoleCollectionName).InsertOne(ctx, req.UserRole)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) InsertUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	_, err := r.db.Collection(userPermissionCollectionName).InsertOne(ctx, req.UserPermission)
	if err != nil {
		return err
	}
	return nil
}

// update
func (r *repository) UpdateUserRoleInfo(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		userRoleFieldName:     req.Name,
		userRoleFieldSlug:     req.Slug,
		userRoleFieldType:     req.Type,
		userRoleFieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userRoleCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
}

func (r *repository) UpdateUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		userRoleFieldPermissionIDs: req.PermissionIDs,
		userRoleFieldUpdateAt:      req.UpdateAt,
	}}
	res, err := r.db.Collection(userRoleCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
}

func (r *repository) UpdateUserPermissionInfo(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		userPermissionFieldPath:              req.Path,
		userPermissionFieldName:              req.Name,
		userPermissionFieldType:              req.Type,
		userPermissionFieldAllowGetMethod:    req.AllowGetMethod,
		userPermissionFieldAllowPostMethod:   req.AllowPostMethod,
		userPermissionFieldAllowPutMethod:    req.AllowPutMethod,
		userPermissionFieldAllowDeleteMethod: req.AllowDeleteMethod,
		userPermissionFieldUpdateAt:          req.UpdateAt,
	}}
	res, err := r.db.Collection(userPermissionCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
}

// delete
func (r *repository) DeleteUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		userRoleFieldIsActive: req.IsActive,
		userRoleFieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userRoleCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
}

func (r *repository) DeleteUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		userRoleFieldIsActive: req.IsActive,
		userRoleFieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userPermissionCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
}

// get
func (r *repository) GetUserRoleAllList(ctx context.Context, req rolePermission.GetUserRoleListRequest) ([]rolePermission.UserRole, *PaginatedData, error) {
	filter := bson.M{}
	if req.Active == 0 {
		filter[userRoleFieldIsActive] = false
	}

	if req.Active == 1 {
		filter[userRoleFieldIsActive] = true
	}

	if !req.IsRoot && !req.BusinessID.IsZero() {
		filter[userRoleFieldBusinessId] = req.BusinessID
	}

	if len(req.Type) > 0 {
		filter[userRoleFieldType] = bson.M{"$in": req.Type}
	}

	collection := r.db.Collection(userRoleCollectionName)
	projection := bson.D{
		{Key: userRoleFieldID, Value: 1},
		{Key: userRoleFieldName, Value: 1},
		{Key: userRoleFieldSlug, Value: 1},
		{Key: userRoleFieldType, Value: 1},
		{Key: userRoleFieldPermissionIDs, Value: 1},
		{Key: userRoleFieldIsActive, Value: 1},
		{Key: userRoleFieldCreateAt, Value: 1},
		{Key: userRoleFieldUpdateAt, Value: 1},
	}

	var userRoles []rolePermission.UserRole
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(userRoleFieldID, -1).Select(projection).Filter(filter).Decode(&userRoles).Find()
	if err != nil {
		return nil, nil, err
	}

	return userRoles, paginatedData, nil
}

func (r *repository) GetUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*rolePermission.UserRole, error) {
	filter := bson.M{userRoleFieldID: req.ID}
	if !req.IsRoot && !req.BusinessID.IsZero() {
		filter[userRoleFieldBusinessId] = bson.M{"$in": []primitive.ObjectID{req.BusinessID, primitive.NilObjectID}}
	}

	if !req.IsRoot {
		filter[userRoleFieldIsActive] = true
	}

	res := r.db.Collection(userRoleCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var userRole rolePermission.UserRole
	if err := res.Decode(&userRole); err != nil {
		return nil, res.Err()
	}

	return &userRole, nil
}

func (r *repository) GetUserRoleByIDs(ctx context.Context, req rolePermission.GetUserRoleByIdsRequest) ([]rolePermission.UserRole, *PaginatedData, error) {
	filter := bson.M{userRoleFieldID: bson.M{"$in": req.IDs}}
	if !req.IsRoot && !req.BusinessID.IsZero() {
		filter[userRoleFieldBusinessId] = req.BusinessID
	}

	if !req.IsRoot {
		filter[userRoleFieldIsActive] = true
	}

	collection := r.db.Collection(userRoleCollectionName)
	projection := bson.D{
		{Key: userRoleFieldID, Value: 1},
		{Key: userRoleFieldName, Value: 1},
		{Key: userRoleFieldSlug, Value: 1},
		{Key: userRoleFieldType, Value: 1},
		{Key: userRoleFieldPermissionIDs, Value: 1},
		{Key: userRoleFieldBusinessId, Value: 1},
		{Key: userRoleFieldIsActive, Value: 1},
		{Key: userRoleFieldCreateAt, Value: 1},
		{Key: userRoleFieldUpdateAt, Value: 1},
	}

	var userRoles []rolePermission.UserRole
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(userRoleFieldID, -1).Select(projection).Filter(filter).Decode(&userRoles).Find()
	if err != nil {
		return nil, nil, err
	}
	return userRoles, paginatedData, nil
}

func (r *repository) GetUserPermissionAllList(ctx context.Context, req rolePermission.GetUserPermissionListRequest) ([]rolePermission.UserPermission, *PaginatedData, error) {
	filter := bson.M{}
	if !req.BusinessID.IsZero() {
		filter[userRoleFieldBusinessId] = bson.M{"$in": []primitive.ObjectID{req.BusinessID, primitive.NilObjectID}}
	} else {
		filter[userRoleFieldBusinessId] = primitive.NilObjectID
	}

	if req.Active == 0 {
		filter[userRoleFieldIsActive] = false
	}

	if req.Active == 1 {
		filter[userRoleFieldIsActive] = true
	}

	if len(req.Type) > 0 {
		filter[userRoleFieldType] = bson.M{"$in": req.Type}
	}

	collection := r.db.Collection(userPermissionCollectionName)
	projection := bson.D{
		{Key: userPermissionFieldID, Value: 1},
		{Key: userPermissionFieldPath, Value: 1},
		{Key: userPermissionFieldName, Value: 1},
		{Key: userPermissionFieldAllowGetMethod, Value: 1},
		{Key: userPermissionFieldAllowPostMethod, Value: 1},
		{Key: userPermissionFieldAllowPutMethod, Value: 1},
		{Key: userPermissionFieldAllowDeleteMethod, Value: 1},
		{Key: userPermissionFieldType, Value: 1},
		{Key: userPermissionFieldIsActive, Value: 1},
		{Key: userPermissionFieldCreateAt, Value: 1},
		{Key: userPermissionFieldUpdateAt, Value: 1},
	}

	var permissions []rolePermission.UserPermission
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(userRoleFieldID, -1).Select(projection).Filter(filter).Decode(&permissions).Find()
	if err != nil {
		return nil, nil, err
	}

	return permissions, paginatedData, nil
}

func (r *repository) GetUserPermissionByID(ctx context.Context, req rolePermission.GetRolePermissionByIdRequest) (*rolePermission.UserPermission, error) {
	filter := bson.M{userPermissionFieldID: req.ID}
	if !req.IsRoot && !req.BusinessID.IsZero() {
		//filter[userPermissionBusinessId] = req.BusinessID
		filter[userPermissionBusinessId] = bson.M{"$in": []primitive.ObjectID{req.BusinessID, primitive.NilObjectID}}
	}

	if !req.IsRoot {
		filter[userPermissionFieldIsActive] = true
	}
	res := r.db.Collection(userPermissionCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var userPermission rolePermission.UserPermission
	if err := res.Decode(&userPermission); err != nil {
		return nil, res.Err()
	}

	return &userPermission, nil
}
func (r *repository) GetUserPermissionByIDs(ctx context.Context, req rolePermission.GetRolePermissionByIdsRequest) ([]rolePermission.UserPermission, *PaginatedData, error) {
	filter := bson.M{userPermissionFieldID: bson.M{"$in": req.IDs}}
	if !req.IsRoot && !req.BusinessID.IsZero() {
		filter[userPermissionBusinessId] = bson.M{"$in": []primitive.ObjectID{req.BusinessID, primitive.NilObjectID}}
		//userPermissionFieldType
	}

	if !req.IsRoot {
		//filter[userPermissionFieldType] = bson.M{"$in": []string{"USER", "SUB_USER"}}
		filter[userPermissionFieldIsActive] = true
	}

	collection := r.db.Collection(userPermissionCollectionName)
	projection := bson.D{
		{Key: userPermissionFieldID, Value: 1},
		{Key: userPermissionFieldPath, Value: 1},
		{Key: userPermissionFieldName, Value: 1},
		{Key: userPermissionFieldAllowGetMethod, Value: 1},
		{Key: userPermissionFieldAllowPostMethod, Value: 1},
		{Key: userPermissionFieldAllowPutMethod, Value: 1},
		{Key: userPermissionFieldAllowDeleteMethod, Value: 1},
		{Key: userPermissionFieldType, Value: 1},
		{Key: userPermissionFieldIsActive, Value: 1},
		{Key: userPermissionFieldCreateAt, Value: 1},
		{Key: userPermissionFieldUpdateAt, Value: 1},
	}

	var permissions []rolePermission.UserPermission
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(userRoleFieldID, -1).Select(projection).Filter(filter).Decode(&permissions).Find()
	if err != nil {
		return nil, nil, err
	}

	return permissions, paginatedData, nil
}
