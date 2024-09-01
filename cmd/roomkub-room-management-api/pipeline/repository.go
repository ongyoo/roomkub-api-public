package roomPipeline

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "roomPipeline"

	fieldID             = "_id"
	fieldBusinessID     = "businessId"
	fieldContractID     = "contractId"
	fieldRoomName       = "roomName"
	fieldMeter          = "meter"
	fieldAdminRemark    = "adminRemark"
	fieldRemark         = "remark"
	fieldIsCurrent      = "isCurrent"
	fieldIsActive       = "isActive"
	fieldStatus         = "status"
	fieldTimeLineStatus = "timeLineStatus"
	fieldCreateBy       = "createBy"
	fieldRecorderBy     = "recorderBy"
	fieldCreateAt       = "createAt"
	fieldUpdateAt       = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateroomPipelineRequest) (*string, error)
	// get
	GetRoomPipelineList(ctx context.Context, req GetroomPipelineListRequest) ([]roomPipeline, *PaginatedData, error)
	Find(ctx context.Context, req FindroomPipelineRequest) (*roomPipeline, error)
	// update
	UpdateByAdmin(ctx context.Context, req UpdateroomPipelineRequest) error
	UpdateTimeLineStatus(ctx context.Context, req UpdateroomPipelineRequest) error
	UpdateCurrent(ctx context.Context, req UpdateroomPipelineRequest) error
	UpdateActive(ctx context.Context, req UpdateroomPipelineRequest) error
	UpdateMeter(ctx context.Context, req UpdateroomPipelineRequest) error
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

// Insert
func (r *repository) Insert(ctx context.Context, req CreateroomPipelineRequest) (*string, error) {
	result, err := r.db.Collection(collectionName).InsertOne(ctx, req.roomPipeline)
	if err != nil {
		return nil, err
	}
	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		id := insertedID.Hex()
		return &id, nil
	}
	return nil, nil
}

// Get
func (r *repository) GetRoomPipelineList(ctx context.Context, req GetroomPipelineListRequest) ([]roomPipeline, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldContractID] = req.ContractID
		filter[fieldBusinessID] = req.BusinessID
		filter[fieldIsActive] = true

		if req.Status != nil {
			filter[fieldStatus] = req.Status
		}
	}
	collection := r.db.Collection(collectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldContractID, Value: 1},
		{Key: fieldMeter, Value: 1},
		{Key: fieldAdminRemark, Value: 1},
		{Key: fieldRemark, Value: 1},
		{Key: fieldIsCurrent, Value: 1},
		{Key: fieldStatus, Value: 1},
		{Key: fieldCreateBy, Value: 1},
		{Key: fieldRecorderBy, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var roomPipelineItems []roomPipeline
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&roomPipelineItems).Find()
	if err != nil {
		return nil, nil, err
	}

	return roomPipelineItems, paginatedData, nil
}

func (r *repository) Find(ctx context.Context, req FindroomPipelineRequest) (*roomPipeline, error) {
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

	res := r.db.Collection(collectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var roomPipeline roomPipeline
	if err := res.Decode(&roomPipeline); err != nil {
		return nil, res.Err()
	}

	return &roomPipeline, nil
}

// Update
/*
res, err := s.db.Collection(merchantCollection).UpdateOne(ctx, bson.M{"partnerMerchantID": partnerMerchantID}, bson.M{
		"$set": bson.M{
			"merchantStatus": status,
		},
	})
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}
	return nil
*/
func (r *repository) UpdateByAdmin(ctx context.Context, req UpdateroomPipelineRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldRemark:      req.Remark,
		fieldAdminRemark: req.AdminRemark,
		fieldStatus:      req.Status,
		fieldRecorderBy:  req.RecorderBy,
		fieldUpdateAt:    req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateTimeLineStatus(ctx context.Context, req UpdateroomPipelineRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldTimeLineStatus: req.TimeLineStatus,
		fieldUpdateAt:       req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateCurrent(ctx context.Context, req UpdateroomPipelineRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsCurrent: req.IsCurrent,
		fieldUpdateAt:  req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateActive(ctx context.Context, req UpdateroomPipelineRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsActive: req.IsActive,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateMeter(ctx context.Context, req UpdateroomPipelineRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldMeter:    req.Meter,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}
