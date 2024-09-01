package room

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomCollectionName = "room"

	fieldID                 = "_id"
	fieldBusinessID         = "businessId"
	fieldName               = "name"
	fieldDescription        = "description"
	fieldThumbnailURl       = "thumbnailURl"
	fieldThumbnailPublicURl = "thumbnailPublicURl"
	fieldImages             = "images"
	fieldFloor              = "floor"
	fieldPrice              = "price"
	fieldSpecialPrice       = "specialPrice"
	fieldIsSpecialPrice     = "isSpecialPrice"
	fieldTags               = "tags"
	fieldLocation           = "location"
	fieldIsActive           = "isActive"
	fieldIsPublisher        = "isPublisher"
	fieldStatus             = "status"
	fieldType               = "type"
	fieldSubType            = "subType"
	fieldPaymentStatus      = "paymentStatus"
	fieldCurrentMeter       = "currentMeter"
	fieldRemark             = "remark"
	fieldCreateBy           = "createBy"
	fieldCreateAt           = "createAt"
	fieldUpdateAt           = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateRoomRequest) (*string, error)
	// Get
	GetRoomAllList(ctx context.Context, req GetRoomListRequest) ([]Room, *PaginatedData, error)
	GetDetail(ctx context.Context, req FindRoomRequest) (*Room, error)
	Find(ctx context.Context, req FindRoomRequest) (*Room, error)
	// Update
	UpdateInfo(ctx context.Context, req UpdateRoomRequest) error
	UpdateThumbnail(ctx context.Context, req UpdateRoomRequest) error
	UpdateImages(ctx context.Context, req UpdateRoomRequest) error
	UpdateActive(ctx context.Context, req UpdateRoomRequest) error
	UpdatePublisher(ctx context.Context, req UpdateRoomRequest) error
	UpdatePaymentStatus(ctx context.Context, req UpdateRoomPaymentStatusRequest) error
	UpdateCurrentMeter(ctx context.Context, req UpdateRoomRequest) error
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req CreateRoomRequest) (*string, error) {
	result, err := r.db.Collection(roomCollectionName).InsertOne(ctx, req.Room)
	if err != nil {
		return nil, err
	}
	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		id := insertedID.Hex()
		return &id, nil
	}
	return nil, nil
}

// Update
func (r *repository) UpdateInfo(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldName:           req.Name,
		fieldDescription:    req.Description,
		fieldFloor:          req.Floor,
		fieldPrice:          req.Price,
		fieldSpecialPrice:   req.SpecialPrice,
		fieldIsSpecialPrice: req.IsShowSpecialPrice,
		fieldTags:           req.Tags,
		fieldLocation:       req.Location,
		fieldStatus:         req.Status,
		fieldType:           req.Type,
		fieldSubType:        req.SubType,
		fieldRemark:         req.Remark,
		fieldUpdateAt:       req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateThumbnail(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldThumbnailURl:       req.ThumbnailURl,
		fieldThumbnailPublicURl: req.ThumbnailURlPublic,
		fieldUpdateAt:           req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateImages(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldImages:   req.Images,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateActive(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsActive: req.IsActive,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdatePublisher(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldIsPublisher: req.IsPublisher,
		fieldUpdateAt:    req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateCurrentMeter(ctx context.Context, req UpdateRoomRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldCurrentMeter: req.CurrentMeter,
		fieldUpdateAt:     req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdatePaymentStatus(ctx context.Context, req UpdateRoomPaymentStatusRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldPaymentStatus: req.PaymentStatus,
		fieldUpdateAt:      req.UpdateAt,
	}}
	res, err := r.db.Collection(roomCollectionName).UpdateByID(ctx, req.ID, updateField)
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
func (r *repository) GetRoomAllList(ctx context.Context, req GetRoomListRequest) ([]Room, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldIsActive] = true
		filter[fieldBusinessID] = req.BusinessID
	}
	collection := r.db.Collection(roomCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldName, Value: 1},
		{Key: fieldFloor, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldDescription, Value: 1},
		{Key: fieldTags, Value: 1},
		{Key: fieldPrice, Value: 1},
		{Key: fieldSpecialPrice, Value: 1},
		{Key: fieldIsSpecialPrice, Value: 1},
		{Key: fieldLocation, Value: 1},
		{Key: fieldStatus, Value: 1},
		{Key: fieldPaymentStatus, Value: 1},
		{Key: fieldType, Value: 1},
		{Key: fieldSubType, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldIsPublisher, Value: 1},
		{Key: fieldThumbnailURl, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var rooms []Room
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&rooms).Find()
	if err != nil {
		return nil, nil, err
	}

	return rooms, paginatedData, nil
}

func (r *repository) GetDetail(ctx context.Context, req FindRoomRequest) (*Room, error) {
	return r.Find(ctx, req)
}

func (r *repository) Find(ctx context.Context, req FindRoomRequest) (*Room, error) {
	filter := bson.M{}
	if !req.IsRoot {
		if !req.ID.IsZero() {
			filter[fieldID] = req.ID
		}

		if !req.BusinessID.IsZero() {
			filter[fieldBusinessID] = req.BusinessID
		}

		filter[fieldIsActive] = true
	}

	res := r.db.Collection(roomCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var room Room
	if err := res.Decode(&room); err != nil {
		return nil, res.Err()
	}

	return &room, nil
}
