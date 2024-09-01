package user

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	userModel "github.com/ongyoo/roomkub-api/pkg/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollectionName = "user"

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
	fieldAddress             = "address"
	fieldProvince            = "province"
	fieldPostCode            = "postCode"
	fieldPhone               = "phone"
	fieldPhoneId             = "phoneId"
	fieldThumbnailURl        = "thumbnailURl"
	fieldRoleID              = "roleId"
	fieldIsActive            = "isActive"
	fieldIsBaned             = "isBaned"
	fieldIsVerify            = "isVerify"
	fieldSlug                = "slug"
	fieldSlugIdentityId      = "slugIdentityId"
	fieldCreateAt            = "createAt"
	fieldUpdateAt            = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req userModel.UserCreateFormRequest) (*string, error)
	// set
	UpdateProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdatePassword(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateBan(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateVerify(ctx context.Context, req userModel.UpdateUserFormRequest) error
	// get
	FindEmail(ctx context.Context, req userModel.UserLoginRequest) (*userModel.User, error)
	GetById(ctx context.Context, req userModel.GetUserByIdRequest) (*userModel.User, error)
	GetUserAllListByIDs(ctx context.Context, req userModel.GetUserListByIDsRequest) ([]userModel.User, *PaginatedData, error)
	// root admin
	GetUserAllList(ctx context.Context, req userModel.GetUserListRequest) ([]userModel.User, *PaginatedData, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req userModel.UserCreateFormRequest) (*string, error) {
	result, err := r.db.Collection(userCollectionName).InsertOne(ctx, req.User)
	if err != nil {
		return nil, err
	}
	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		id := insertedID.Hex()
		return &id, nil
	}
	return nil, nil
}

func (r *repository) FindEmail(ctx context.Context, req userModel.UserLoginRequest) (*userModel.User, error) {
	filter := bson.M{fieldEmailIdentityId: req.Email, fieldIsActive: true}
	res := r.db.Collection(userCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var user userModel.User
	if err := res.Decode(&user); err != nil {
		return nil, res.Err()
	}

	return &user, nil
}

// Update
func (r *repository) UpdateProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	//.UpdateByID(ctx, id, bson.M{"$set": bson.M{"status": StatusCancelled}})
	//.ReplaceOne(ctx, bson.M{"shopId": req.ShopId, "status": StatusPending}, req.User)
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
		fieldPhoneId:             req.PhoneIdentityID,
		fieldAddress:             req.Address,
		fieldProvince:            req.Province,
		fieldPostCode:            req.PostCode,
		fieldUpdateAt:            req.UpdateAt,
	}}
	res, err := r.db.Collection(userCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdatePassword(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldPassword: req.Password,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateBan(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsBaned:  req.IsBaned,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateVerify(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsVerify: req.IsVerify,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(userCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldThumbnailURl: req.ThumbnailURl,
		fieldUpdateAt:     req.UpdateAt,
	}}
	res, err := r.db.Collection(userCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) GetUserAllListByIDs(ctx context.Context, req userModel.GetUserListByIDsRequest) ([]userModel.User, *PaginatedData, error) {
	filter := bson.M{fieldID: bson.M{"$in": req.IDs}}
	collection := r.db.Collection(userCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldEmail, Value: 1},
		{Key: fieldFirstName, Value: 1},
		{Key: fieldLastName, Value: 1},
		{Key: fieldNickName, Value: 1},
		{Key: fieldAddress, Value: 1},
		{Key: fieldProvince, Value: 1},
		{Key: fieldPostCode, Value: 1},
		{Key: fieldPhone, Value: 1},
		{Key: fieldThumbnailURl, Value: 1},
		{Key: fieldRoleID, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldIsBaned, Value: 1},
		{Key: fieldIsVerify, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var users []userModel.User
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, nil, err
	}

	return users, paginatedData, nil
}

// get all users for root admin
func (r *repository) GetUserAllList(ctx context.Context, req userModel.GetUserListRequest) ([]userModel.User, *PaginatedData, error) {
	filter := bson.M{}
	collection := r.db.Collection(userCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldEmail, Value: 1},
		{Key: fieldFirstName, Value: 1},
		{Key: fieldLastName, Value: 1},
		{Key: fieldNickName, Value: 1},
		{Key: fieldRoleID, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldThumbnailURl, Value: 1},
		{Key: fieldIsBaned, Value: 1},
		{Key: fieldIsVerify, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var users []userModel.User
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&users).Find()
	if err != nil {
		return nil, nil, err
	}

	return users, paginatedData, nil
}

func (r *repository) GetById(ctx context.Context, req userModel.GetUserByIdRequest) (*userModel.User, error) {
	filter := bson.M{fieldID: req.ID}
	res := r.db.Collection(userCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var user userModel.User
	if err := res.Decode(&user); err != nil {
		return nil, res.Err()
	}

	return &user, nil
}
