package businessChannel

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/life4/genesis/slices"
	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
)

type Service interface {
	// Create
	CreateBusinessChannel(ctx context.Context, req BusinessChannelCreateFormRequest) error
	// Update
	UpdateMyPackageId(ctx context.Context, req UpdateBusinessChannelMyPackageFormRequest) error
	UpdateSetting(ctx context.Context, req UpdateBusinessChannelSettingFormRequest) error
	// Delete
	UpdateDelete(ctx context.Context, req UpdateDeleteBusinessFormRequest) error
	// Get
	GetBusinessChannelById(ctx context.Context, id string) (*BusinessChannel, error)
	GetBusinessChannelListByUserId(ctx context.Context, req GetBusinessChannelListByUserIdRequest) ([]GetBusinessChannelItem, *PaginatedData, error)
	// Root
	GetBusinessChannelAllList(ctx context.Context, req GetBusinessChannelListRequest) ([]GetRootBusinessChannelItem, *PaginatedData, error)
	UpdateBan(ctx context.Context, req UpdateBanBusinessFormRequest) error
	UpdateUnBan(ctx context.Context, req UpdateBanBusinessFormRequest) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

// Create
func (s service) CreateBusinessChannel(ctx context.Context, req BusinessChannelCreateFormRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	objID, err := primitive.ObjectIDFromHex(userClaims.Payload.ID)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.Owner = objID
	req.Members = []BusinessChannelMember{{ID: objID, IsBaned: false, IsSeparate: false}}
	req.ChannelID = primitive.NewObjectID()
	req.IsActive = true
	req.IsBaned = false
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.Insert(ctx, req)
}

// Update
func (s service) UpdateMyPackageId(ctx context.Context, req UpdateBusinessChannelMyPackageFormRequest) error {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}

	myPackageID, err := primitive.ObjectIDFromHex(req.MyPackageIDStr)
	if err != nil {
		return errors.New(err.Error() + " [Package ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}
	newReq := BusinessChannelCreateFormRequest{}
	newReq.ID = id
	newReq.MyPackageID = myPackageID
	newReq.UpdateAt = time.Now()
	return s.repository.UpdateMyPackageId(ctx, newReq)
}

func (s service) UpdateSetting(ctx context.Context, req UpdateBusinessChannelSettingFormRequest) error {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}
	newReq := BusinessChannelCreateFormRequest{}
	newReq.ID = id
	newReq.Setting = BusinessChannelSetting{
		Name:             mongo.Encrypted(req.Name),
		LogoThumbnailURl: req.LogoThumbnailURl,
	}
	newReq.UpdateAt = time.Now()
	return s.repository.UpdateSetting(ctx, newReq)
}

// Delete
func (s service) UpdateDelete(ctx context.Context, req UpdateDeleteBusinessFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}

	newReq := BusinessChannelCreateFormRequest{}
	newReq.ID = objID
	newReq.IsActive = false
	newReq.Remark = req.Remark
	newReq.UpdateAt = time.Now()
	return s.repository.UpdateDelete(ctx, newReq)
}

// Get
func (s service) GetBusinessChannelById(ctx context.Context, channelId string) (*BusinessChannel, error) {
	objID, err := primitive.ObjectIDFromHex(channelId)
	if err != nil {
		return nil, errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}
	return s.repository.GetBusinessChannelById(ctx, objID)
}

func (s service) GetBusinessChannelListByUserId(ctx context.Context, req GetBusinessChannelListByUserIdRequest) ([]GetBusinessChannelItem, *PaginatedData, error) {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	objID, err := primitive.ObjectIDFromHex(userClaims.Payload.ID)
	if err != nil {
		return nil, nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.UserID = objID
	res, paginated, err := s.repository.GetBusinessChannelListByUserId(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []GetBusinessChannelItem
	for _, item := range res {
		findIndex := slices.FindIndex(
			item.Members,
			func(el BusinessChannelMember) bool { return el.ID.Hex() == req.UserID.Hex() },
		)
		isSeparate := false
		isBaned := item.IsBaned
		if findIndex > -1 {
			memberItem := item.Members[findIndex]
			isSeparate = memberItem.IsSeparate
			isBaned = item.IsBaned || memberItem.IsBaned
		}
		resItems = append(resItems, GetBusinessChannelItem{
			ID:               item.ID.Hex(),
			ChannelID:        item.ChannelID.Hex(),
			Name:             item.Setting.Name.String(),
			LogoThumbnailURl: item.Setting.LogoThumbnailURl,
			IsActive:         item.IsActive,
			IsBaned:          isBaned,
			IsSeparate:       isSeparate,
			CreateAt:         item.CreateAt,
			UpdateAt:         item.UpdateAt,
		})
	}
	return resItems, paginated, err
}

// Root
func (s service) GetBusinessChannelAllList(ctx context.Context, req GetBusinessChannelListRequest) ([]GetRootBusinessChannelItem, *PaginatedData, error) {
	res, paginated, err := s.repository.GetBusinessChannelAllList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []GetRootBusinessChannelItem
	for _, item := range res {
		resItems = append(resItems, GetRootBusinessChannelItem{
			ID:               item.ID.Hex(),
			ChannelID:        item.ChannelID.Hex(),
			MyPackageID:      item.MyPackageID.Hex(),
			Name:             item.Setting.Name.String(),
			LogoThumbnailURl: item.Setting.LogoThumbnailURl,
			IsActive:         item.IsActive,
			IsBaned:          item.IsBaned,
			CreateAt:         item.CreateAt,
			UpdateAt:         item.UpdateAt,
		})
	}
	return resItems, paginated, err
}

func (s service) UpdateBan(ctx context.Context, req UpdateBanBusinessFormRequest) error {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}
	newReq := BusinessChannelCreateFormRequest{}
	newReq.ID = id
	newReq.Remark = req.Remark
	newReq.IsBaned = true
	newReq.UpdateAt = time.Now()
	return s.repository.UpdateBan(ctx, newReq)
}

func (s service) UpdateUnBan(ctx context.Context, req UpdateBanBusinessFormRequest) error {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return errors.New(err.Error() + " [ID fail] กรุณาติดต่อผู้ดูแลระบบ")
	}
	newReq := BusinessChannelCreateFormRequest{}
	newReq.ID = id
	newReq.Remark = req.Remark
	newReq.IsBaned = false
	newReq.UpdateAt = time.Now()
	return s.repository.UpdateBan(ctx, newReq)
}
