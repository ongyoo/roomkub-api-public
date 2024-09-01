package customer

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	customerCollectionName = "customer"

	fieldID                  = "_id"
	fieldEmail               = "email"
	fieldEmailIdentityId     = "emailIdentityId"
	fieldPassword            = "password"
	fieldFirstName           = "firstName"
	fieldFirstNameIdentityId = "firstNameIdentityId"
	fieldLastName            = "lastName"
	fieldNickName            = "nickName"
	fieldNId                 = "nId"
	fieldNIdentityId         = "nIdentityId"
	fieldGenderType          = "genderType"
	fieldAddress             = "address"
	fieldProvince            = "province"
	fieldPostCode            = "postCode"
	fieldPhone               = "phone"
	fieldPhoneIdentityId     = "phoneIdentityId"
	fieldThumbnailURl        = "thumbnailURl"
	fieldRoleID              = "roleId"
	fieldIsActive            = "isActive"
	fieldLineRef             = "lineRef"
	fieldIsLineLiff          = "isLineLiff"
	fieldCreateAt            = "createAt"
	fieldUpdateAt            = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateCustomerRequest) (*string, error)
	// Get
	Find(ctx context.Context, req FindCustomerRequest) (*Customer, error)
	// Update
	UpdateProfileInfo(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdatePassword(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateActive(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateLineRef(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateProfileImage(ctx context.Context, req UpdateCustomerFormRequest) error
	// Get
	GetCustomerAllList(ctx context.Context, req GetCustomerListRequest) ([]Customer, *PaginatedData, error)
	GetById(ctx context.Context, req GetCustomerByIdRequest) (*Customer, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req CreateCustomerRequest) (*string, error) {
	result, err := r.db.Collection(customerCollectionName).InsertOne(ctx, req.Customer)
	if err != nil {
		return nil, err
	}
	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		id := insertedID.Hex()
		return &id, nil
	}
	return nil, nil
}

func (r *repository) Find(ctx context.Context, req FindCustomerRequest) (*Customer, error) {
	filter := bson.M{}
	if !req.ID.IsZero() {
		filter[fieldID] = req.ID
	}

	if req.Email != "" {
		filter[fieldEmailIdentityId] = req.Email
	}

	if req.Password != "" {
		filter[fieldPassword] = req.Password
	}

	if req.FirstName != "" {
		filter[fieldFirstNameIdentityId] = req.FirstName
	}

	if req.NID != "" {
		filter[fieldNIdentityId] = req.NID
	}

	if req.Phone != "" {
		filter[fieldPhoneIdentityId] = req.Phone
	}

	if req.LineRef != "" {
		filter[fieldLineRef] = req.LineRef
	}

	res := r.db.Collection(customerCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var customer Customer
	if err := res.Decode(&customer); err != nil {
		return nil, res.Err()
	}

	return &customer, nil
}

// Update
func (r *repository) UpdateProfileInfo(ctx context.Context, req UpdateCustomerFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldEmail:               req.Email,
		fieldEmailIdentityId:     req.EmailIdentityID,
		fieldFirstName:           req.FirstName,
		fieldFirstNameIdentityId: req.FirstNameIdentityID,
		fieldLastName:            req.LastName,
		fieldNickName:            req.NickName,
		fieldNId:                 req.NID,
		fieldNIdentityId:         req.NIdentityID,
		fieldPhone:               req.Phone,
		fieldPhoneIdentityId:     req.PhoneIdentityID,
		fieldGenderType:          req.GenderType,
		fieldAddress:             req.Address,
		fieldProvince:            req.Province,
		fieldPostCode:            req.PostCode,
		fieldUpdateAt:            req.UpdateAt,
	}}
	res, err := r.db.Collection(customerCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdatePassword(ctx context.Context, req UpdateCustomerFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldPassword: req.Password,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(customerCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateActive(ctx context.Context, req UpdateCustomerFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsActive: req.IsActive,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(customerCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateLineRef(ctx context.Context, req UpdateCustomerFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsLineLiff: req.IsLineLiff,
		fieldLineRef:    req.LineRef,
		fieldUpdateAt:   req.UpdateAt,
	}}
	res, err := r.db.Collection(customerCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateProfileImage(ctx context.Context, req UpdateCustomerFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldThumbnailURl: req.ThumbnailURl,
		fieldUpdateAt:     req.UpdateAt,
	}}
	res, err := r.db.Collection(customerCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

// Root
func (r *repository) GetCustomerAllList(ctx context.Context, req GetCustomerListRequest) ([]Customer, *PaginatedData, error) {
	filter := bson.M{}
	collection := r.db.Collection(customerCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldEmail, Value: 1},
		{Key: fieldFirstName, Value: 1},
		{Key: fieldLastName, Value: 1},
		{Key: fieldNickName, Value: 1},
		{Key: fieldGenderType, Value: 1},
		{Key: fieldPhone, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldThumbnailURl, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var customers []Customer
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&customers).Find()
	if err != nil {
		return nil, nil, err
	}

	return customers, paginatedData, nil
}

func (r *repository) GetById(ctx context.Context, req GetCustomerByIdRequest) (*Customer, error) {
	filter := bson.M{fieldID: req.ID}
	res := r.db.Collection(customerCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var customer Customer
	if err := res.Decode(&customer); err != nil {
		return nil, res.Err()
	}

	return &customer, nil
}
