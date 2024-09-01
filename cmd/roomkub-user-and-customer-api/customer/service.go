package customer

import (
	"context"
	"errors"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/ongyoo/roomkub-api/pkg/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Crypto
	crypto "github.com/ongyoo/roomkub-api/pkg/crypto"
	// JWT
	/*
		"context"
		"errors"
		"strings"
		"time"

		. "github.com/gobeam/mongo-go-pagination"
		"github.com/ongyoo/roomkub-api/cmd/roomkub-business-channel-api/businessChannel"
		"github.com/ongyoo/roomkub-api/cmd/roomkub-user-customer-api/user/userRolePermission"

		rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"

		userModel "github.com/ongyoo/roomkub-api/pkg/model/user"
		validate "github.com/ongyoo/roomkub-api/pkg/validate"
		"go.mongodb.org/mongo-driver/bson/primitive"

		// Crypto
		crypto "github.com/ongyoo/roomkub-api/pkg/crypto"
		// JWT
		jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
	*/)

type Service interface {
	// insert
	Create(ctx context.Context, req CreateCustomerRequest) (*string, error)
	// Get
	Find(ctx context.Context, req FindCustomerRequest) (*CustomerDetailResponse, error)
	// Update
	UpdateProfileInfo(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdatePassword(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateActive(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateUnActive(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateLineRef(ctx context.Context, req UpdateCustomerFormRequest) error
	UpdateProfileImage(ctx context.Context, req UpdateCustomerFormRequest) error
	// Get
	GetCustomerAllList(ctx context.Context, req GetCustomerListRequest) ([]CustomerResponse, *PaginatedData, error)
	GetById(ctx context.Context, req GetCustomerByIdRequest) (*Customer, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s service) Create(ctx context.Context, req CreateCustomerRequest) (*string, error) {
	if req.Password != "" {
		if len(req.Password) < 6 {
			return nil, errors.New("กรุณาใส่รหัสผ่านผู้ใช้งาน หรือต้องใส่มากกว่า 6 อักษรขึ้นไป")
		}

		hashPassword, err := crypto.HashPassword(req.Password)
		if err == nil {
			return nil, errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบบ")
		}
		req.Password = hashPassword
	}

	if req.Email != "" {
		if !validate.ValidEmail(req.Email.String()) {
			return nil, errors.New("อีเมลไม่ถูกต้อง กรุณาตรวจสอบอีกครั้ง !!")
		}
		emailHash := crypto.Hash(req.Email.String())
		_, err := s.Find(ctx, FindCustomerRequest{Email: emailHash})
		if err == nil {
			return nil, errors.New("มีผู้ใช้อีเมลนี้แล้ว กรุณาตรวจสอบอีกครั้ง")
		}
		req.EmailIdentityID = emailHash
	}

	// Hash id
	req.FirstNameIdentityID = crypto.Hash(req.FirstName.String())
	req.NIdentityID = crypto.Hash(req.NID.String())
	req.PhoneIdentityID = crypto.Hash(req.Phone.String())
	// update flag
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	req.IsActive = true
	return s.repository.Insert(ctx, req)
}

// Get
func (s service) Find(ctx context.Context, req FindCustomerRequest) (*CustomerDetailResponse, error) {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if req.FirstName != "" {
		req.FirstName = crypto.Hash(req.FirstName)
	}

	if req.Email != "" {
		req.Email = crypto.Hash(req.Email)
	}

	if req.Password != "" {
		hashPassword, err := crypto.HashPassword(req.Password)
		if err == nil {
			return nil, errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบบ")
		}
		req.Password = hashPassword
	}

	if req.NID != "" {
		req.NID = crypto.Hash(req.NID)
	}

	if req.Phone != "" {
		req.Phone = crypto.Hash(req.Phone)
	}

	res, err := s.repository.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	resItem := CustomerDetailResponse{
		ID:           res.ID.Hex(),
		Email:        res.Email.String(),
		FirstName:    res.FirstName.String(),
		LastName:     res.LastName.String(),
		NickName:     res.NickName.String(),
		NID:          res.NID.String(),
		Gender:       string(res.GenderType),
		Address:      res.Address.String(),
		Province:     res.Province,
		PostCode:     res.PostCode,
		Phone:        res.Phone.String(),
		ThumbnailUrl: res.ThumbnailURl,
		Documents:    res.Documents,
		CreateAt:     res.CreateAt,
		UpdateAt:     res.UpdateAt,
	}
	return &resItem, nil
}

// Update
func (s service) UpdateProfileInfo(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if req.Email.String() != "" {
		if !validate.ValidEmail(req.Email.String()) {
			return errors.New("อีเมลไม่ถูกต้อง กรุณาตรวจสอบอีกครั้ง !!")
		}
		req.EmailIdentityID = crypto.Hash(req.Email.String())
	}

	if req.FirstName.String() != "" {
		req.FirstNameIdentityID = crypto.Hash(req.FirstName.String())
	}

	if req.NID.String() != "" {
		req.NIdentityID = crypto.Hash(req.NID.String())
	}

	if req.Phone.String() != "" {
		req.PhoneIdentityID = crypto.Hash(req.Phone.String())
	}

	genderStr := string(req.GenderType)
	if genderStr != "" {
		genderType := CustomerGenderType(genderStr)
		if genderType == "" {
			return errors.New("กรุณาเช็คข้อมูลเพศอีกครั้ง")
		}
	}

	req.UpdateAt = time.Now()

	return s.repository.UpdateProfileInfo(ctx, req)
}

func (s service) UpdatePassword(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	hashPassword, err := crypto.HashPassword(req.Password)
	if err == nil {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบบ")
	}
	req.Password = hashPassword
	return s.repository.UpdatePassword(ctx, req)
}

func (s service) UpdateActive(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	req.IsActive = true
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdateUnActive(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	req.IsActive = false
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdateLineRef(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if req.LineRef == "" {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบบ")
	}
	req.IsLineLiff = true
	return s.repository.UpdateLineRef(ctx, req)
}

func (s service) UpdateProfileImage(ctx context.Context, req UpdateCustomerFormRequest) error {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if req.ThumbnailURl == "" {
		return errors.New("เกิดข้อผิดพลาด กรุณาใส่ข้อมูลให้ครบถ้วน")
	}
	return s.repository.UpdateProfileImage(ctx, req)
}

// Get
func (s service) GetCustomerAllList(ctx context.Context, req GetCustomerListRequest) ([]CustomerResponse, *PaginatedData, error) {
	res, paginatedData, err := s.repository.GetCustomerAllList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []CustomerResponse
	for _, item := range res {
		resItems = append(resItems, CustomerResponse{
			ID:           item.ID.Hex(),
			Email:        item.Email.String(),
			FirstName:    item.FirstName.String(),
			LastName:     item.LastName.String(),
			NickName:     item.NickName.String(),
			Gender:       string(item.GenderType),
			ThumbnailUrl: item.ThumbnailURl,
			IsActive:     item.IsActive,
			Phone:        item.Phone.String(),
			CreateAt:     item.CreateAt,
			UpdateAt:     item.UpdateAt,
		})
	}
	return resItems, paginatedData, err
}

func (s service) GetById(ctx context.Context, req GetCustomerByIdRequest) (*Customer, error) {
	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}
	return s.repository.GetById(ctx, req)
}
