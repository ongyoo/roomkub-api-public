package document

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	documentCollectionName = "document"

	fieldID           = "_id"
	fieldBusinessID   = "businessId"
	fieldRefID        = "refId"
	fieldName         = "name"
	fieldDescription  = "description"
	fieldRemark       = "remark"
	fieldFileName     = "fileName"
	fieldContentType  = "contentType"
	fieldDocumentType = "documentType"
	fieldPrivateUrl   = "privateUrl"
	fieldPublicUrl    = "publicUrl"
	fieldIsActive     = "isActive"
	fieldUpdateAt     = "updateAt"
	fieldCreateAt     = "createAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateDocumentRequest) (*string, error)
	// // Get
	Find(ctx context.Context, req FindDocumentRequest) (*Document, error)
	GetDocumentList(ctx context.Context, req FindDocumentRequest) ([]Document, *PaginatedData, error)
	// // Update
	// UpdateProfileInfo(ctx context.Context, req UpdateCustomerFormRequest) error
	// UpdatePassword(ctx context.Context, req UpdateCustomerFormRequest) error
	// UpdateActive(ctx context.Context, req UpdateCustomerFormRequest) error
	// UpdateLineRef(ctx context.Context, req UpdateCustomerFormRequest) error
	// UpdateProfileImage(ctx context.Context, req UpdateCustomerFormRequest) error
	// // Get
	// GetCustomerAllList(ctx context.Context, req GetCustomerListRequest) ([]Customer, *PaginatedData, error)
	// GetById(ctx context.Context, req GetCustomerByIdRequest) (*Customer, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req CreateDocumentRequest) (*string, error) {
	result, err := r.db.Collection(documentCollectionName).InsertOne(ctx, req.Document)
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
func (r *repository) Find(ctx context.Context, req FindDocumentRequest) (*Document, error) {
	filter := bson.M{}
	if !req.IsRoot {
		if !req.ID.IsZero() {
			filter[fieldID] = req.ID
		}

		if !req.BusinessID.IsZero() {
			filter[fieldBusinessID] = req.BusinessID
		}

		if !req.RefID.IsZero() {
			filter[fieldRefID] = req.RefID
		}

		filter[fieldIsActive] = true
	}

	res := r.db.Collection(documentCollectionName).FindOne(ctx, filter)
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	var document Document
	if err := res.Decode(&document); err != nil {
		return nil, res.Err()
	}

	return &document, nil
}

func (r *repository) GetDocumentList(ctx context.Context, req FindDocumentRequest) ([]Document, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldRefID] = req.RefID
		filter[fieldBusinessID] = req.BusinessID
		filter[fieldIsActive] = true
	}
	collection := r.db.Collection(documentCollectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldRefID, Value: 1},
		{Key: fieldName, Value: 1},
		{Key: fieldContentType, Value: 1},
		{Key: fieldDocumentType, Value: 1},
		{Key: fieldPrivateUrl, Value: 1},
		{Key: fieldPublicUrl, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var roomContractItems []Document
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&roomContractItems).Find()
	if err != nil {
		return nil, nil, err
	}

	return roomContractItems, paginatedData, nil
}

// Update
func (r *repository) UpdateActive(ctx context.Context, req UpdateDocumentRequest) error {
	filter := bson.M{}
	if !req.IsRoot {
		filter[fieldID] = req.ID
		filter[fieldRefID] = req.RefID
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsActive: req.IsActive,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(documentCollectionName).UpdateOne(ctx, filter, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}
