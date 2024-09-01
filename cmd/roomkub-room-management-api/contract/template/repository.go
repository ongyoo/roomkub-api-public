package contractTemplate

import (
	"context"
	"fmt"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "contractTemplate"

	fieldID          = "_id"
	fieldBusinessID  = "businessId"
	fieldName        = "name"
	fieldHtml        = "html"
	fieldRemark      = "remark"
	fieldCreateBy    = "createBy"
	fieldIsActive    = "isActive"
	fieldIsPublisher = "isPublisher"
	fieldCreateAt    = "createAt"
	fieldUpdateAt    = "updateAt"
)

type Repository interface {
	// insert
	Insert(ctx context.Context, req CreateContractTemplateRequest) error
	// get
	GetContractTemplateList(ctx context.Context, req GetContractTemplateRequest) ([]ContractTemplate, *PaginatedData, error)
	Find(ctx context.Context, req GetContractTemplateDetailRequest) (*ContractTemplate, error)
	// update
	UpdateContract(ctx context.Context, req UpdateContractTemplateRequest) error
	UpdateActive(ctx context.Context, req UpdateContractTemplateRequest) error
	UpdatePublisher(ctx context.Context, req UpdateContractTemplateRequest) error
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, req CreateContractTemplateRequest) error {
	result, err := r.db.Collection(collectionName).InsertOne(ctx, req.ContractTemplate)
	if err != nil {
		return err
	}
	if _, ok := result.InsertedID.(primitive.ObjectID); ok {
		//id := insertedID.Hex()
		return nil
	}
	return nil
}

// Get
func (r *repository) GetContractTemplateList(ctx context.Context, req GetContractTemplateRequest) ([]ContractTemplate, *PaginatedData, error) {
	filter := bson.M{}
	if !req.IsRoot {
		// filter["$AND"] = []bson.M{
		// 	{fieldIsPublisher: bson.M{"$eq": true}},
		// }
		filter[fieldIsActive] = true
		if req.MyChannel {
			filter[fieldBusinessID] = req.BusinessID
		} else {
			// all Publisher
			filter[fieldIsPublisher] = true
		}
	}

	collection := r.db.Collection(collectionName)
	projection := bson.D{
		{Key: fieldID, Value: 1},
		{Key: fieldBusinessID, Value: 1},
		{Key: fieldName, Value: 1},
		{Key: fieldIsPublisher, Value: 1},
		{Key: fieldCreateAt, Value: 1},
		{Key: fieldUpdateAt, Value: 1},
	}

	var contractTemplateItems []ContractTemplate
	paginatedData, err := New(collection).Context(ctx).Limit(req.Limit).Page(req.Page).Sort(fieldID, -1).Select(projection).Filter(filter).Decode(&contractTemplateItems).Find()
	if err != nil {
		return nil, nil, err
	}

	return contractTemplateItems, paginatedData, nil
}

// Get
func (r *repository) Find(ctx context.Context, req GetContractTemplateDetailRequest) (*ContractTemplate, error) {
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

	var contractTemplate ContractTemplate
	if err := res.Decode(&contractTemplate); err != nil {
		return nil, res.Err()
	}

	return &contractTemplate, nil
}

// Update
func (r *repository) UpdateContractTemplateInfo(ctx context.Context, req UpdateContractTemplateRequest) error {
	updateField := bson.M{"$set": bson.M{
		fieldName:     req.Name,
		fieldHtml:     req.Html,
		fieldRemark:   req.Remark,
		fieldCreateBy: req.CreateBy,
		fieldUpdateAt: req.UpdateAt,
	}}
	res, err := r.db.Collection(collectionName).UpdateByID(ctx, req.ID, updateField)
	if err != nil {
		return err
	}

	if res.MatchedCount != 1 {
		err := fmt.Errorf("invalid update count: %d documents are matched", res.MatchedCount)
		return err
	}

	return nil
}

func (r *repository) UpdateContract(ctx context.Context, req UpdateContractTemplateRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldName:     req.Name,
		fieldHtml:     req.Html,
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

func (r *repository) UpdateActive(ctx context.Context, req UpdateContractTemplateRequest) error {
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

func (r *repository) UpdatePublisher(ctx context.Context, req UpdateContractTemplateRequest) error {
	filter := bson.M{fieldID: req.ID}
	if !req.IsRoot {
		filter[fieldBusinessID] = req.BusinessID
	}

	updateField := bson.M{"$set": bson.M{
		fieldIsPublisher: req.IsPublisher,
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
