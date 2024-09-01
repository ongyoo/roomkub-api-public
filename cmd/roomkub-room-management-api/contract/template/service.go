package contractTemplate

import (
	"context"
	"errors"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	// Upload
	//"github.com/ongyoo/roomkub-api/pkg/upload"
	// Crypto
	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// Utils
)

type Service interface {
	// insert
	Insert(ctx context.Context, req CreateContractTemplateRequest) error

	// get
	GetContractTemplateList(ctx context.Context, req GetContractTemplateRequest) ([]ContractTemplateResponse, *PaginatedData, error)
	Find(ctx context.Context, req GetContractTemplateDetailRequest) (*ContractTemplate, error)

	// update
	UpdateContract(ctx context.Context, req UpdateContractTemplateRequest) error
	UpdateActive(ctx context.Context, req UpdateContractTemplateRequest) error
	UpdatePublisher(ctx context.Context, req UpdateContractTemplateRequest) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

// Insert
func (s service) Insert(ctx context.Context, req CreateContractTemplateRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	if userClaims.Payload.ID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.ID)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.CreateBy = objID
	}

	req.IsActive = true
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.Insert(ctx, req)
}

func (s service) GetContractTemplateList(ctx context.Context, req GetContractTemplateRequest) ([]ContractTemplateResponse, *PaginatedData, error) {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return nil, nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	res, paginatedData, err := s.repository.GetContractTemplateList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []ContractTemplateResponse
	for _, item := range res {
		resItems = append(resItems, ContractTemplateResponse{
			ID:          item.ID.Hex(),
			Name:        item.Name,
			IsPublisher: item.IsPublisher,
			CreateAt:    item.CreateAt,
			UpdateAt:    item.UpdateAt,
		})
	}
	return resItems, paginatedData, err
}

func (s service) Find(ctx context.Context, req GetContractTemplateDetailRequest) (*ContractTemplate, error) {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if req.IDStr == "" {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.ID = objID

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}
	return s.repository.Find(ctx, req)
}

func (s service) UpdateContract(ctx context.Context, req UpdateContractTemplateRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}
	return s.repository.UpdateContract(ctx, req)
}

func (s service) UpdateActive(ctx context.Context, req UpdateContractTemplateRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdatePublisher(ctx context.Context, req UpdateContractTemplateRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}
	return s.repository.UpdatePublisher(ctx, req)
}
