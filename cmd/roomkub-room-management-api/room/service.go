package room

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/life4/genesis/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Upload
	"github.com/ongyoo/roomkub-api/pkg/upload"

	// Crypto

	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
	// Utils
	"github.com/ongyoo/roomkub-api/pkg/utlis"
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
	// Insert
	Create(ctx context.Context, req CreateRoomRequest) (*string, error)
	// Get
	GetRoomAllList(ctx context.Context, req GetRoomListRequest) ([]RoomResponse, *PaginatedData, error)
	Find(ctx context.Context, req FindRoomRequest) (*Room, error)
	// Update
	UpdateInfo(ctx context.Context, req UpdateRoomRequest) error
	UpdateThumbnail(ctx context.Context, req UpdateRoomRequest, fileHeader multipart.FileHeader) error
	UpdateImage(ctx context.Context, req UpdateRoomRequest, fileHeader multipart.FileHeader) error
	UpdateImages(ctx context.Context, req UpdateRoomRequest) error
	UpdateActive(ctx context.Context, req UpdateRoomRequest) error
	UpdateUnActive(ctx context.Context, req UpdateRoomRequest) error
	UpdatePublisher(ctx context.Context, req UpdateRoomRequest) error
	UpdateUnPublisher(ctx context.Context, req UpdateRoomRequest) error
	UpdatePaymentStatus(ctx context.Context, req UpdateRoomPaymentStatusRequest) error
	UpdateCurrentMeter(ctx context.Context, req UpdateRoomRequest) error
	// Delete
	DeleteImage(ctx context.Context, req DeleteImageRoomRequest) error
}

type service struct {
	repository    Repository
	uploadService upload.Service
}

func NewService(repo Repository, uploadService upload.Service) *service {
	return &service{repository: repo, uploadService: uploadService}
}

// Insert
func (s service) Create(ctx context.Context, req CreateRoomRequest) (*string, error) {
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

	if req.Name == "" {
		return nil, errors.New("required")
	}

	if req.Description == "" {
		return nil, errors.New("required")
	}

	if req.Price == 0 {
		return nil, errors.New("required")
	}

	if req.Price == 0 {
		return nil, errors.New("required")
	}

	req.Status = RoomStatusAvailable
	req.PaymentStatus = RoomPaymentStatusNone
	req.IsActive = true
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.Insert(ctx, req)
}

// Get
func (s service) GetRoomAllList(ctx context.Context, req GetRoomListRequest) ([]RoomResponse, *PaginatedData, error) {
	if !req.IsRoot {
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
	}

	res, paginatedData, err := s.repository.GetRoomAllList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []RoomResponse
	for _, item := range res {
		resItems = append(resItems, RoomResponse{
			ID:                 item.ID.Hex(),
			BusinessID:         item.BusinessID.Hex(),
			Name:               item.Name.String(),
			Description:        utlis.SubString(item.Description.String(), 128),
			Floor:              item.Floor,
			Price:              item.Price,
			SpecialPrice:       item.SpecialPrice,
			IsShowSpecialPrice: item.IsShowSpecialPrice,
			Tags:               item.Tags,
			ThumbnailUrl:       item.ThumbnailURl,
			Status:             item.Status,
			PaymentStatus:      item.PaymentStatus,
			Type:               item.Type,
			SubType:            item.SubType,
			IsPublisher:        item.IsPublisher,
			IsActive:           item.IsActive,
			CreateAt:           item.CreateAt,
			UpdateAt:           item.UpdateAt,
		})
	}
	return resItems, paginatedData, err
}

func (s service) Find(ctx context.Context, req FindRoomRequest) (*Room, error) {
	if !req.IsRoot {
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
	}

	if req.ID.IsZero() {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}
	return s.repository.Find(ctx, req)
}

// Update
func (s service) UpdateInfo(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.UpdateAt = time.Now()
	return s.repository.UpdateInfo(ctx, req)
}

func (s service) UpdateThumbnail(ctx context.Context, req UpdateRoomRequest, fileHeader multipart.FileHeader) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}

	key := "room"
	resUpload, err := s.uploadService.Upload(ctx, key, fileHeader)
	if err != nil {
		return err
	}

	req.ThumbnailURlPublic = resUpload.PublicUrl
	req.ThumbnailURl = resUpload.PrivateUrl
	req.UpdateAt = time.Now()
	return s.repository.UpdateThumbnail(ctx, req)
}

func (s service) UpdateImages(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}

	/*
		key := "room"
		resUpload, err := s.uploadService.Upload(ctx, key, fileHeader)
		if err != nil {
			return err
		}
	*/

	req.UpdateAt = time.Now()
	return s.repository.UpdateImages(ctx, req)
}

func (s service) UpdateImage(ctx context.Context, req UpdateRoomRequest, fileHeader multipart.FileHeader) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}

	key := "room"
	resUpload, err := s.uploadService.Upload(ctx, key, fileHeader)
	if err != nil {
		return err
	}

	imageNewItem := RoomImage{
		ID: primitive.NewObjectID(),
		Name:                req.ImageName,
		ThumbnailPrivateURl: resUpload.PrivateUrl,
		ThumbnailPublicURl:  resUpload.PublicUrl,
		IsActive:            true,
		UpdateAt:            time.Now(),
	}

	res, err := s.Find(ctx, FindRoomRequest{ID: req.ID, BusinessID: req.BusinessID})
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}

	imageItems := res.Images
	imageItems = append(imageItems, imageNewItem)
	req.Images = imageItems
	req.UpdateAt = time.Now()
	return s.repository.UpdateImages(ctx, req)
}

func (s service) UpdateActive(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.IsActive = true
	req.UpdateAt = time.Now()
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdateUnActive(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.IsActive = false
	req.UpdateAt = time.Now()
	return s.repository.UpdateActive(ctx, req)
}

func (s service) UpdatePublisher(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.IsPublisher = true
	req.UpdateAt = time.Now()
	return s.repository.UpdatePublisher(ctx, req)
}

func (s service) UpdateUnPublisher(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.IsPublisher = false
	req.UpdateAt = time.Now()
	return s.repository.UpdatePublisher(ctx, req)
}

func (s service) UpdatePaymentStatus(ctx context.Context, req UpdateRoomPaymentStatusRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.PaymentStatus = RoomPaymentStatus(req.PaymentStatusStr)
	req.UpdateAt = time.Now()
	return s.repository.UpdatePaymentStatus(ctx, req)
}

func (s service) UpdateCurrentMeter(ctx context.Context, req UpdateRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if !s.verifyBusinessID(ctx, req.ID, req.BusinessID) {
		return errors.New("เกิดข้อผิดพลาด กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.UpdateAt = time.Now()
	return s.repository.UpdateCurrentMeter(ctx, req)
}

// Delete
func (s service) DeleteImage(ctx context.Context, req DeleteImageRoomRequest) error {
	if !req.IsRoot {
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
	}

	if req.IDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = objID
	}

	if req.ImageIDStr != "" {
		objID, err := primitive.ObjectIDFromHex(req.ImageIDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ImageID = objID
	}

	res, err := s.Find(ctx, FindRoomRequest{ID: req.ID, BusinessID: req.BusinessID})
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}

	imageItems := res.Images
	findIndex := slices.FindIndex(
		imageItems,
		func(el RoomImage) bool { return el.ID == req.ImageID },
	)

	if findIndex > -1 {
		imageItems = utlis.RemoveIndex(imageItems, findIndex)
	} else {
		// not found
		return errors.New("not found for delete image")
	}

	updateReq := UpdateRoomRequest{IDStr: req.ID.Hex(), BusinessIDStr: req.BusinessID.Hex()}
	updateReq.Images = imageItems
	return s.UpdateImages(ctx, updateReq)
}

// private
func (s service) verifyBusinessID(ctx context.Context, ID primitive.ObjectID, businessID primitive.ObjectID) bool {
	res, err := s.Find(ctx, FindRoomRequest{ID: ID, IsRoot: true})
	if err != nil {
		return false
	}

	if res == nil {
		return false
	}

	if res.ID != ID && res.BusinessID != businessID {
		return false
	}
	return true
}
