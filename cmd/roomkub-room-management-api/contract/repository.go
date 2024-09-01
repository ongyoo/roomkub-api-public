package roomContract

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	roomContractCollectionName = "roomContract"

	fieldID               = "_id"
	fieldBusinessID       = "businessId"
	fieldCustomerID       = "customerId"
	fieldRoomID           = "roomId"
	fieldContractTemplate = "contractTemplate"
	fieldDocuments        = "documents"
	fieldRemark           = "remark"
	fieldDepositTotal     = "depositTotal"
	fieldPaidTotal        = "paidTotal"
	fieldIsInstallment    = "isInstallment"
	fieldIsActive         = "isActive"
	fieldIsReportOut      = "isReportOut"
	fieldReportOutAt      = "reportOutAt"
	fieldIsMigrate        = "is_migrate"
	fieldCreateBy         = "createBy"
	fieldCreateAt         = "createAt"
	fieldUpdateAt         = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateRoomContractRequest) (*string, error)
	// get
	Find(ctx context.Context, req FindRoomContractRequest) (*RoomContract, error)
	GetRoomContractList(ctx context.Context, req GetRoomContractRequest) ([]RoomContract, *PaginatedData, error)
	// update
	Update(ctx context.Context, req UpdateRoomContractRequest) error
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req CreateRoomContractRequest) (*string, error) {
	result, err := r.db.Collection(roomContractCollectionName).InsertOne(ctx, req.RoomContract)
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
func (r *repository) Find(ctx context.Context, req FindRoomContractRequest) (*RoomContract, error) {
	filter := bson.M{}
	if !req.IsRoot {
		if !req.ID.IsZero() {
			filter[fieldID] = req.ID
		}

		if !req.BusinessID.IsZero() {
			filter[fieldBusinessID] = req.BusinessID
		}

		//filter[fieldIsActive] = true
	}

	res := r.db.Collection(roomContractCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var roomContract RoomContract
	if err := res.Decode(&roomContract); err != nil {
		return nil, res.Err()
	}

	return &roomContract, nil
}

// Update
func (r *repository) Update(ctx context.Context, req UpdateRoomContractRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldRemark:      req.Remark,
		fieldPaidTotal:   req.PaidTotal,
		fieldIsReportOut: req.IsReportOut,
		fieldReportOutAt: req.ReportOutAt,
		fieldIsActive:    req.IsActive,
		fieldUpdateAt:    req.UpdateAt,
	}}
	res, err := r.db.Collection(roomContractCollectionName).UpdateByID(ctx, req.ID, updateField)
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
func (r *repository) GetRoomContractList(ctx context.Context, req GetRoomContractRequest) ([]RoomContract, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldRoomID] = req.RoomID
		filter[fieldBusinessID] = req.BusinessID
	}
	collection := r.db.Collection(roomContractCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldCustomerID, Value: 1},
		{Key: fieldRoomID, Value: 1},
		{Key: fieldIsActive, Value: 1},
		{Key: fieldDepositTotal, Value: 1},
		{Key: fieldPaidTotal, Value: 1},
		{Key: fieldIsInstallment, Value: 1},
		{Key: fieldIsReportOut, Value: 1},
		{Key: fieldReportOutAt, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var roomContractItems []RoomContract
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&roomContractItems).Find()
	if err != nil {
		return nil, nil, err
	}

	return roomContractItems, paginatedData, nil
}
