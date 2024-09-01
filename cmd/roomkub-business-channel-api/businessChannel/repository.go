package businessChannel

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	businessChannelCollectionName = "businessChannel"

	fieldID          = "_id"
	fieldChannelId   = "channelId"
	fieldMembers     = "members"
	fieldOwner       = "owner"
	fieldType        = "type"
	fieldSetting     = "setting"
	fieldIsActive    = "isActive"
	fieldIsBan       = "isBan"
	fieldMyPackageId = "myPackageId"
	fieldCreateAt    = "createAt"
	fieldUpdateAt    = "updateAt"
	fieldRemark      = "remark"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req BusinessChannelCreateFormRequest) error
	// update
	UpdateBan(ctx context.Context, req BusinessChannelCreateFormRequest) error
	UpdateMyPackageId(ctx context.Context, req BusinessChannelCreateFormRequest) error
	UpdateSetting(ctx context.Context, req BusinessChannelCreateFormRequest) error
	// delete
	UpdateDelete(ctx context.Context, req BusinessChannelCreateFormRequest) error
	// get
	GetBusinessChannelById(ctx context.Context, channelId primitive.ObjectID) (*BusinessChannel, error)
	GetBusinessChannelListByUserId(ctx context.Context, req GetBusinessChannelListByUserIdRequest) ([]BusinessChannel, *PaginatedData, error)
	// root
	GetBusinessChannelAllList(ctx context.Context, req GetBusinessChannelListRequest) ([]BusinessChannel, *PaginatedData, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	_, err := r.db.Collection(businessChannelCollectionName).InsertOne(ctx, req.BusinessChannel)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateBan(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsBan:  req.IsBaned,
		fieldRemark: req.Remark,
	}}
	res, err := r.db.Collection(businessChannelCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateDelete(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldRemark:   req.Remark,
		fieldIsActive: req.IsActive,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(businessChannelCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateMyPackageId(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldMyPackageId: req.MyPackageID,
		fieldUpdateAt:    req.UpdateAt,
	}}
	res, err := r.db.Collection(businessChannelCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateSetting(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldSetting:  req.Setting,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(businessChannelCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

// Get
func (r *repository) GetBusinessChannelById(ctx context.Context, channelId primitive.ObjectID) (*BusinessChannel, error) {
	filter := bson.M{fieldChannelId: channelId, fieldIsActive: true}
	res := r.db.Collection(businessChannelCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var businessChannel BusinessChannel
	if err := res.Decode(&businessChannel); err != nil {
		return nil, res.Err()
	}
	return &businessChannel, nil
}

func (r *repository) GetBusinessChannelListByUserId(ctx context.Context, req GetBusinessChannelListByUserIdRequest) ([]BusinessChannel, *PaginatedData, error) {
	//filter := bson.M{fieldMembers: bson.M{"$in": req.UserID}, fieldIsActive: true}
	filter := bson.M{fieldMembers: bson.M{"$elemMatch": bson.M{fieldID: req.UserID}}, fieldIsActive: true}
	collection := r.db.Collection(businessChannelCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldChannelId, Value: 1},
		{Key: fieldMembers, Value: 1},
		{Key: fieldOwner, Value: 1},
		{Key: fieldType, Value: 1},
		{Key: fieldSetting, Value: 1},
		{Key: fieldMyPackageId, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldIsBan, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var businessChannels []BusinessChannel
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&businessChannels).Find()
	if err != nil {
		return nil, nil, err
	}

	return businessChannels, paginatedData, nil
}

// Root
func (r *repository) GetBusinessChannelAllList(ctx context.Context, req GetBusinessChannelListRequest) ([]BusinessChannel, *PaginatedData, error) {
	filter := bson.M{}
	collection := r.db.Collection(businessChannelCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldChannelId, Value: 1},
		{Key: fieldMembers, Value: 1},
		{Key: fieldOwner, Value: 1},
		{Key: fieldType, Value: 1},
		{Key: fieldSetting, Value: 1},
		{Key: fieldMyPackageId, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldIsBan, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var businessChannels []BusinessChannel
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&businessChannels).Find()
	if err != nil {
		return nil, nil, err
	}

	return businessChannels, paginatedData, nil
}
