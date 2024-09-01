package roomPipeline

import (
	"context"
	"errors"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Customer
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	// Upload
	//"github.com/ongyoo/roomkub-api/pkg/upload"
	// Crypto
	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
	// Utils
)

type Service interface {
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

type service struct {
	repository Repository
}

func NewService(repo Repository, customerService customer.Service) *service {
	return &service{repository: repo}
}

func (s service) Create(ctx context.Context, req CreateroomPipelineRequest) (*string, error) {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}

	if userClaims.Payload.ID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.ID)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.CreateBy = objID
	}

	return s.repository.Insert(ctx, req)
}

// get
func (s service) GetRoomPipelineList(ctx context.Context, req GetroomPipelineListRequest) ([]roomPipeline, *PaginatedData, error) {
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
	return s.repository.GetRoomPipelineList(ctx, req)
}

func (s service) Find(ctx context.Context, req FindroomPipelineRequest) (*roomPipeline, error) {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = objID
	}
	return s.repository.Find(ctx, req)
}

// Update
func (s service) UpdateByAdmin(ctx context.Context, req UpdateroomPipelineRequest) error {
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
	return s.repository.UpdateByAdmin(ctx, req)
}

func (s service) UpdateTimeLineStatus(ctx context.Context, req UpdateroomPipelineRequest) error {
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
	return s.repository.UpdateTimeLineStatus(ctx, req)
}

func (s service) UpdateCurrent(ctx context.Context, req UpdateroomPipelineRequest) error {
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
	return s.repository.UpdateCurrent(ctx, req)
}

func (s service) UpdateActive(ctx context.Context, req UpdateroomPipelineRequest) error {
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
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdateMeter(ctx context.Context, req UpdateroomPipelineRequest) error {
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

	return s.repository.UpdateMeter(ctx, req)
}
