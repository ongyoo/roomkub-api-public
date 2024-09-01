package Invoice

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "invoice"

	fieldID             = "_id"
	fieldBusinessID     = "businessId"
	fieldRoomPipelineID = "roomPipelineId"
	fieldCustomerID     = "customerId"
	fieldItems          = "items"
	fieldTotalPrice     = "totalPrice"
	fieldRemark         = "remark"
	fieldStatus         = "status"
	fieldType           = "type"
	fieldIsActive       = "isActive"
	fieldIsSendNotify   = "isSendNotify"
	fieldIsPrinted      = "isPrinted"
	fieldPrintedAt      = "printedAt"
	fieldCreateUserID   = "create_user_id"
	fieldCreateAt       = "createAt"
	fieldUpdateAt       = "updateAt"
	fieldSendNotifyAt   = "sendNotifyAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateInvoiceRequest) (*string, error)
	// get
	GetInvoiceList(ctx context.Context, req GetInvoiceListRequest) ([]Invoice, *PaginatedData, error)
	Find(ctx context.Context, req FindInvoiceRequest) (*Invoice, error)
	// update
	Update(ctx context.Context, req UpdateInvoiceRequest) error
	UpdateItems(ctx context.Context, req UpdateInvoiceRequest) error
	UpdatePrinted(ctx context.Context, req UpdateInvoiceRequest) error
	UpdateSendNotify(ctx context.Context, req UpdateInvoiceRequest) error
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

// Insert
func (r *repository) Insert(ctx context.Context, req CreateInvoiceRequest) (*string, error) {
	result, err := r.db.Collection(collectionName).InsertOne(ctx, req.Invoice)
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
func (r *repository) GetInvoiceList(ctx context.Context, req GetInvoiceListRequest) ([]Invoice, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldRoomPipelineID] = req.RoomPipelineID
		filter[fieldBusinessID] = req.BusinessID
		filter[fieldIsActive] = true
	}

	if !req.CustomerID.IsZero() {
		filter[fieldCreateUserID] = req.CustomerID
	}

	if req.Status != nil {
		filter[fieldStatus] = req.Status
	}

	if req.Type != nil {
		filter[fieldType] = req.Type
	}

	collection := r.db.Collection(collectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldCustomerID, Value: 1},
		{Key: fieldTotalPrice, Value: 1},
		{Key: fieldStatus, Value: 1},
		{Key: fieldType, Value: 1},
		{Key: fieldIsPrinted, Value: 1},
		{Key: fieldIsSendNotify, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var invoiceItems []Invoice
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&invoiceItems).Find()
	if err != nil {
		return nil, nil, err
	}

	return invoiceItems, paginatedData, nil
}

func (r *repository) Find(ctx context.Context, req FindInvoiceRequest) (*Invoice, error) {
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

	var invoice Invoice
	if err := res.Decode(&invoice); err != nil {
		return nil, res.Err()
	}

	return &invoice, nil
}

// Update
func (r *repository) Update(ctx context.Context, req UpdateInvoiceRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldRemark:   req.Remark,
		fieldStatus:   req.Status,
		fieldType:     req.Type,
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

func (r *repository) UpdateItems(ctx context.Context, req UpdateInvoiceRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldItems:    req.Items,
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

func (r *repository) UpdatePrinted(ctx context.Context, req UpdateInvoiceRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsPrinted: req.IsPrinted,
		fieldPrintedAt: req.PrintedAt,
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

func (r *repository) UpdateSendNotify(ctx context.Context, req UpdateInvoiceRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsSendNotify: req.IsSendNotify,
		fieldSendNotifyAt: req.SendNotifyAt,
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
