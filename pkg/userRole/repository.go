package userrole

import (
	"context"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	/*
		userCollectionName = "userRolePermission"

		fieldID                = "_id"
		fieldPath              = "path"
		fieldName              = "name"
		fieldSlug              = "slug"
		fieldAllowGetMethod    = "allowGetMethod"
		fieldAllowPostMethod   = "allowPostMethod"
		fieldAllowPutMethod    = "allowPutMethod"
		fieldAllowDeleteMethod = "allowDeleteMethod"
	*/

	userCollectionName = "UserRole"

	fieldID             = "_id"
	fieldName           = "name"
	fieldPermissionID   = "permissionId"
	fieldSlug           = "slug"
	fieldSlugIdentityID = "slugIdentityId"
	fieldIsActive       = "isActive"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req UserRoleRequest) error
	// get
	GetUserRoleListBySlug(ctx context.Context, req GetUserRoleListRequest) ([]UserRole, *PaginatedData, error)

	// root admin
	GetUserRoleListAll(ctx context.Context, req GetUserRoleListRequest) ([]UserRole, *PaginatedData, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req UserRoleRequest) error {
	_, err := r.db.Collection(userCollectionName).InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserRoleListBySlug(ctx context.Context, req GetUserRoleListRequest) ([]UserRole, *PaginatedData, error) {
	filter := bson.M{fieldSlugIdentityID: req.Slug, fieldIsActive: true}
	collection := r.db.Collection(userCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldName, Value: 1},
		{Key: fieldPermissionID, Value: 1},
		{Key: fieldSlug, Value: 1},
		{Key: fieldIsActive, Value: 1},
	}

	var users []UserRole
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, nil, err
	}

	return users, paginatedData, nil
}

// Root Admin
func (r *repository) GetUserRoleListAll(ctx context.Context, req GetUserRoleListRequest) ([]UserRole, *PaginatedData, error) {
	filter := bson.M{}
	collection := r.db.Collection(userCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldName, Value: 1},
		{Key: fieldPermissionID, Value: 1},
		{Key: fieldSlug, Value: 1},
		{Key: fieldIsActive, Value: 1},
	}

	var users []UserRole
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, nil, err
	}

	return users, paginatedData, nil
}
