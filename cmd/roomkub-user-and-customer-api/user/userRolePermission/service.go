package userRolePermission

import (
	"context"
	"strings"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
	"github.com/ongyoo/roomkub-api/pkg/utlis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// Crypto
	// JWT
)

type Service interface {
	// Create
	CreateUserRole(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	CreateUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// Update
	UpdateUserRoleInfo(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	UpdateUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	UpdateUserPermissionInfo(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// Delete
	DeleteUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error
	DeleteUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error
	// Get
	GetUserRoleAllList(ctx context.Context, req rolePermission.GetUserRoleListRequest) ([]rolePermission.UserRole, *PaginatedData, error)
	GetUserPermissionAllList(ctx context.Context, req rolePermission.GetUserPermissionListRequest) ([]rolePermission.UserPermission, *PaginatedData, error)
	GetUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*UserRoleResponse, error)
	GetUserRoleByIDs(ctx context.Context, req rolePermission.GetUserRoleByIdsRequest) ([]rolePermission.UserRole, *PaginatedData, error)
	GetUserPermissionByID(ctx context.Context, req rolePermission.GetRolePermissionByIdRequest) (*rolePermission.UserPermission, error)
	GetUserPermissionByIDs(ctx context.Context, req rolePermission.GetRolePermissionByIdsRequest) ([]rolePermission.UserPermission, *PaginatedData, error)

	// Root
	GetRootUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*UserRoleResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

// Create
func (s service) CreateUserRole(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	if req.BusinessIDStr == "" {
		req.BusinessID = primitive.NilObjectID
	} else {
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return err
		}
		req.BusinessID = businessID
	}
	req.Slug = utlis.ToSnakeCase(req.Name.String())
	req.IsActive = true
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.InsertUserRole(ctx, req)
}

func (s service) CreateUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	if req.BusinessIDStr == "" {
		req.BusinessID = primitive.NilObjectID
	} else {
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return err
		}
		req.BusinessID = businessID
	}
	req.IsActive = true
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.InsertUserPermission(ctx, req)
}

// Update
func (s service) UpdateUserRoleInfo(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return err
	}
	_, err = s.GetUserRoleByID(ctx, rolePermission.GetUserRoleByIdRequest{IDStr: req.IDStr, BusinessIDStr: req.BusinessIDStr})
	if err != nil {
		return err
	}
	req.ID = objID
	req.BusinessID = businessID
	req.UpdateAt = time.Now()
	return s.repository.UpdateUserRoleInfo(ctx, req)
}

func (s service) UpdateUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return err
	}
	_, err = s.GetUserRoleByID(ctx, rolePermission.GetUserRoleByIdRequest{IDStr: req.IDStr, BusinessIDStr: req.BusinessIDStr})
	if err != nil {
		return err
	}

	permissionIDs := []primitive.ObjectID{}
	for _, itemID := range req.PermissionStrIDs {
		pID, err := primitive.ObjectIDFromHex(itemID)
		if err == nil {
			permissionIDs = append(permissionIDs, pID)
		}
	}

	req.ID = objID
	req.BusinessID = businessID
	req.PermissionIDs = permissionIDs
	req.UpdateAt = time.Now()
	return s.repository.UpdateUserRolePermission(ctx, req)
}

func (s service) UpdateUserPermissionInfo(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return err
	}
	_, err = s.GetUserPermissionByID(ctx, rolePermission.GetRolePermissionByIdRequest{IDStr: req.IDStr, BusinessIDStr: req.BusinessIDStr})
	if err != nil {
		return err
	}
	req.ID = objID
	req.BusinessID = businessID
	req.UpdateAt = time.Now()
	return s.repository.UpdateUserPermissionInfo(ctx, req)
}

// Delete
func (s service) DeleteUserRolePermission(ctx context.Context, req rolePermission.UserRoleCreateFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return err
	}
	_, err = s.GetUserPermissionByID(ctx, rolePermission.GetRolePermissionByIdRequest{ID: objID, BusinessID: businessID})
	if err != nil {
		return err
	}
	req.ID = objID
	req.BusinessID = businessID
	req.IsActive = false
	req.UpdateAt = time.Now()
	return s.repository.DeleteUserRolePermission(ctx, req)
}

func (s service) DeleteUserPermission(ctx context.Context, req rolePermission.UserPermissionCreateFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return err
	}
	_, err = s.GetUserPermissionByID(ctx, rolePermission.GetRolePermissionByIdRequest{ID: objID, BusinessID: businessID})
	if err != nil {
		return err
	}
	req.ID = objID
	req.BusinessID = businessID
	req.IsActive = false
	req.UpdateAt = time.Now()
	return s.repository.DeleteUserPermission(ctx, req)
}

// Get
func (s service) GetUserRoleAllList(ctx context.Context, req rolePermission.GetUserRoleListRequest) ([]rolePermission.UserRole, *PaginatedData, error) {
	typeItems := req.Type
	typeStrList := strings.Split(req.TypeStr, ",")
	for _, typeItem := range typeStrList {
		typeItems = append(typeItems, rolePermission.UserRoleType(strings.ToUpper(typeItem)))
	}
	req.Type = typeItems
	return s.repository.GetUserRoleAllList(ctx, req)
}

func (s service) GetUserPermissionAllList(ctx context.Context, req rolePermission.GetUserPermissionListRequest) ([]rolePermission.UserPermission, *PaginatedData, error) {
	typeItems := req.Type
	if req.TypeStr != "" {
		typeStrList := strings.Split(req.TypeStr, ",")
		for _, typeItem := range typeStrList {
			typeItems = append(typeItems, rolePermission.UserRoleType(strings.ToUpper(typeItem)))
		}
		req.Type = typeItems
	} else {
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return nil, nil, err
		}
		req.BusinessID = businessID
		req.Active = 1
		req.Type = []rolePermission.UserRoleType{rolePermission.UserRoleTypeUser, rolePermission.UserRoleTypeSubUser}
	}
	return s.repository.GetUserPermissionAllList(ctx, req)
}

func (s service) GetUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*UserRoleResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return nil, err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return nil, err
	}

	req.ID = objID
	req.BusinessID = businessID

	resRole, err := s.repository.GetUserRoleByID(ctx, req)
	if err != nil {
		return nil, err
	}

	resPermission, _, err := s.repository.GetUserPermissionByIDs(ctx, rolePermission.GetRolePermissionByIdsRequest{BusinessID: req.BusinessID, IDs: resRole.PermissionIDs, Page: 1, Limit: 50, IsRoot: req.IsRoot})
	if err != nil {
		return nil, err
	}

	permissionItems := []UserPermissionResponse{}
	for _, item := range resPermission {
		permissionItems = append(permissionItems, UserPermissionResponse{
			ID:                item.ID.Hex(),
			BusinessID:        item.BusinessID.Hex(),
			Path:              item.Path.String(),
			Name:              item.Name.String(),
			Type:              item.Type,
			AllowGetMethod:    item.AllowGetMethod,
			AllowPostMethod:   item.AllowPostMethod,
			AllowPutMethod:    item.AllowPutMethod,
			AllowDeleteMethod: item.AllowDeleteMethod,
			IsActive:          item.IsActive,
			CreateAt:          item.CreateAt,
			UpdateAt:          item.UpdateAt,
		})
	}

	res := UserRoleResponse{
		ID:          resRole.ID.Hex(),
		BusinessID:  resRole.BusinessID.Hex(),
		Name:        resRole.Name.String(),
		Slug:        resRole.Slug,
		Type:        resRole.Type,
		Permissions: permissionItems,
		IsActive:    resRole.IsActive,
		CreateAt:    resRole.CreateAt,
		UpdateAt:    resRole.UpdateAt,
	}
	return &res, nil
}

func (s service) GetRootUserRoleByID(ctx context.Context, req rolePermission.GetUserRoleByIdRequest) (*UserRoleResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return nil, err
	}

	if req.BusinessIDStr != "" {
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return nil, err
		}
		req.BusinessID = businessID
	}

	req.ID = objID
	req.IsRoot = true
	resRole, err := s.repository.GetUserRoleByID(ctx, req)
	if err != nil {
		return nil, err
	}

	resPermission, _, err := s.repository.GetUserPermissionByIDs(ctx, rolePermission.GetRolePermissionByIdsRequest{BusinessID: req.BusinessID, IDs: resRole.PermissionIDs, Page: 1, Limit: 50, IsRoot: req.IsRoot})
	if err != nil {
		return nil, err
	}

	permissionItems := []UserPermissionResponse{}
	for _, item := range resPermission {
		permissionItems = append(permissionItems, UserPermissionResponse{
			ID:                item.ID.Hex(),
			BusinessID:        item.BusinessID.Hex(),
			Path:              item.Path.String(),
			Name:              item.Name.String(),
			Type:              item.Type,
			AllowGetMethod:    item.AllowGetMethod,
			AllowPostMethod:   item.AllowPostMethod,
			AllowPutMethod:    item.AllowPutMethod,
			AllowDeleteMethod: item.AllowDeleteMethod,
			IsActive:          item.IsActive,
			CreateAt:          item.CreateAt,
			UpdateAt:          item.UpdateAt,
		})
	}

	res := UserRoleResponse{
		ID:          resRole.ID.Hex(),
		BusinessID:  resRole.BusinessID.Hex(),
		Name:        resRole.Name.String(),
		Slug:        resRole.Slug,
		Type:        resRole.Type,
		Permissions: permissionItems,
		IsActive:    resRole.IsActive,
		CreateAt:    resRole.CreateAt,
		UpdateAt:    resRole.UpdateAt,
	}
	return &res, nil
}

func (s service) GetUserRoleByIDs(ctx context.Context, req rolePermission.GetUserRoleByIdsRequest) ([]rolePermission.UserRole, *PaginatedData, error) {
	idItems := req.IDs
	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return nil, nil, err
	}
	req.BusinessID = businessID
	idStrList := strings.Split(req.IDsStr, ",")
	for _, idItem := range idStrList {
		objID, err := primitive.ObjectIDFromHex(idItem)
		if err == nil {
			idItems = append(idItems, objID)
		}
	}
	return s.repository.GetUserRoleByIDs(ctx, req)
}

func (s service) GetUserPermissionByID(ctx context.Context, req rolePermission.GetRolePermissionByIdRequest) (*rolePermission.UserPermission, error) {
	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return nil, err
	}

	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return nil, err
	}
	req.ID = objID
	req.BusinessID = businessID
	return s.repository.GetUserPermissionByID(ctx, req)
}

func (s service) GetUserPermissionByIDs(ctx context.Context, req rolePermission.GetRolePermissionByIdsRequest) ([]rolePermission.UserPermission, *PaginatedData, error) {
	idItems := req.IDs
	businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
	if err != nil {
		return nil, nil, err
	}
	req.BusinessID = businessID

	idStrList := strings.Split(req.IDsStr, ",")
	for _, idItem := range idStrList {
		objID, err := primitive.ObjectIDFromHex(idItem)
		if err == nil {
			idItems = append(idItems, objID)
		}
	}
	return s.repository.GetUserPermissionByIDs(ctx, req)
}
