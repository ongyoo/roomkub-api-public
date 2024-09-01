package user

import (
	"context"
	"errors"
	"strings"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-business-channel-api/businessChannel"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user/userRolePermission"

	rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"

	userModel "github.com/ongyoo/roomkub-api/pkg/model/user"
	validate "github.com/ongyoo/roomkub-api/pkg/validate"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Crypto
	crypto "github.com/ongyoo/roomkub-api/pkg/crypto"
	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
)

type Service interface {
	// Create
	Create(ctx context.Context, req userModel.UserCreateFormRequest) (*string, error)
	// Login
	Login(ctx context.Context, req userModel.UserLoginRequest) (*string, error)
	// Update
	UpdateProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateMyProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdatePassword(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateMyPassword(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateMyProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateBan(ctx context.Context, req userModel.UpdateUserFormRequest) error
	UpdateVerify(ctx context.Context, req userModel.UpdateUserFormRequest) error
	// get
	FindEmail(ctx context.Context, req userModel.UserLoginRequest) (*userModel.User, error)
	GetById(ctx context.Context, req userModel.GetUserByIdRequest) (*userModel.User, error)
	GetUserAllListByBusinessChannel(ctx context.Context, req userModel.GetUserListByChannelIdRequest) ([]GetUserItemList, *PaginatedData, error)
	GetUserAllListByIDs(ctx context.Context, req userModel.GetUserListByIDsRequest) ([]GetUserItemList, *PaginatedData, error)
	// root admin
	GetUserAllList(ctx context.Context, req userModel.GetUserListRequest) ([]GetUserItemList, *PaginatedData, error)
}

type service struct {
	repository             Repository
	businessChannelService businessChannel.Service
	userRoleService        userRolePermission.Service
}

func NewService(businessChannelService businessChannel.Service, userRoleService userRolePermission.Service, repo Repository) *service {
	return &service{repository: repo, businessChannelService: businessChannelService, userRoleService: userRoleService}
}

func (s service) Create(ctx context.Context, req userModel.UserCreateFormRequest) (*string, error) {
	if !validate.ValidEmail(req.Email.String()) {
		return nil, errors.New("อีเมลไม่ถูกต้อง กรุณาตรวจสอบอีกครั้ง !!")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("กรุณาใส่รหัสผ่านผู้ใช้งาน หรือต้องใส่มากกว่า 6 อักษรขึ้นไป")
	}

	emailHash := crypto.Hash(req.Email.String())
	_, err := s.FindEmail(ctx, userModel.UserLoginRequest{Email: emailHash})
	if err == nil {
		return nil, errors.New("มีผู้ใช้อีเมลนี้แล้ว กรุณาตรวจสอบอีกครั้ง")
	}

	// Hash id
	req.EmailIdentityID = emailHash
	req.FirstNameIdentityID = crypto.Hash(req.FirstName.String())
	req.NIdentityID = crypto.Hash(req.NID.String())
	req.PhoneIdentityID = crypto.Hash(req.Phone.String())
	req.SlugIdentityID = crypto.Hash(req.Slug.String())
	// update flag
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	req.IsActive = true
	return s.repository.Insert(ctx, req)
}

func (s service) Login(ctx context.Context, req userModel.UserLoginRequest) (*string, error) {
	if !validate.ValidEmail(req.Email) {
		return nil, errors.New("อีเมลไม่ถูกต้อง กรุณาตรวจสอบอีกครั้ง !!")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("กรุณาใส่รหัสผ่านผู้ใช้งาน หรือต้องใส่มากกว่า 6 อักษรขึ้นไป")
	}

	req.Email = crypto.Hash(req.Email)

	res, err := s.repository.FindEmail(ctx, req)
	if err != nil {
		return nil, errors.New("user not match [1011] (ข้อมูลยูสเซอร์ไม่ถูกต้อง)")
	}

	if res.IsBaned {
		return nil, errors.New("You have been banned from the system. (คุณถูกแบนออกจากระบบ สอบถามเพิ่มเติมที่ผู้ดูแลระบบ)")
	}

	isPasswordMatch := crypto.CheckPasswordHash(req.Password, res.Password)
	if !isPasswordMatch {
		return nil, errors.New("user not match [1012] (ข้อมูลยูสเซอร์ไม่ถูกต้อง)")
	}

	roleRes, roleErr := s.userRoleService.GetRootUserRoleByID(ctx, rolePermission.GetUserRoleByIdRequest{IDStr: res.RoleID.Hex()})
	if roleErr != nil {
		return nil, roleErr
	}

	payload := jwt.UserPayload{
		ID:           res.ID.Hex(),
		Email:        res.Email.String(),
		FirstName:    res.FirstName.String(),
		LastName:     res.LastName.String(),
		NickName:     res.NickName.String(),
		ThumbnailURl: res.ThumbnailURl,
		RoleID:       res.RoleID.Hex(),
		RoleName:     roleRes.Name,
		BusinessID:   "652a430cf598856ca659e2e0", // chanel id
	}

	jwtToken, err := jwt.GenerateJWT(payload)
	if err != nil {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	return &jwtToken, nil
}

// Update
func (s service) UpdateProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.User.ID = objID
	req.EmailIdentityID = crypto.Hash(req.Email.String())
	req.FirstNameIdentityID = crypto.Hash(req.FirstName.String())
	req.NIdentityID = crypto.Hash(req.NID.String())
	req.UpdateAt = time.Now()
	return s.repository.UpdateProfileInfo(ctx, req)
}

func (s service) UpdateMyProfileInfo(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.UserID = userClaims.Payload.ID
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.EmailIdentityID = crypto.Hash(req.Email.String())
	req.FirstNameIdentityID = crypto.Hash(req.FirstName.String())
	req.NIdentityID = crypto.Hash(req.NID.String())
	req.UpdateAt = time.Now()
	return s.repository.UpdateProfileInfo(ctx, req)
}

func (s service) UpdatePassword(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	if len(req.Password) < 6 {
		return errors.New("กรุณาใส่รหัสผ่านผู้ใช้งาน หรือต้องใส่มากกว่า 6 อักษรขึ้นไป")
	}

	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdatePassword(ctx, req)
}

func (s service) UpdateMyPassword(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	if len(req.Password) < 6 {
		return errors.New("กรุณาใส่รหัสผ่านผู้ใช้งาน หรือต้องใส่มากกว่า 6 อักษรขึ้นไป")
	}

	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.UserID = userClaims.Payload.ID
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdatePassword(ctx, req)
}

func (s service) UpdateProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdateProfileImage(ctx, req)
}

func (s service) UpdateMyProfileImage(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.UserID = userClaims.Payload.ID
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdateProfileImage(ctx, req)
}

func (s service) UpdateBan(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdateBan(ctx, req)
}

func (s service) UpdateVerify(ctx context.Context, req userModel.UpdateUserFormRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.User.ID = objID
	req.UpdateAt = time.Now()
	return s.repository.UpdateVerify(ctx, req)
}

// get
func (s service) FindEmail(ctx context.Context, req userModel.UserLoginRequest) (*userModel.User, error) {
	return s.repository.FindEmail(ctx, req)
}

func (s service) GetById(ctx context.Context, req userModel.GetUserByIdRequest) (*userModel.User, error) {
	objID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.ID = objID
	return s.repository.GetById(ctx, req)
}

func (s service) GetUserAllListByBusinessChannel(ctx context.Context, req userModel.GetUserListByChannelIdRequest) ([]GetUserItemList, *PaginatedData, error) {
	channelRes, err := s.businessChannelService.GetBusinessChannelById(ctx, req.ChannelID)
	if err != nil {
		return nil, nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	idItems := req.IDs
	for _, item := range channelRes.Members {
		idItems = append(idItems, item.ID)
	}

	res, paginated, err := s.repository.GetUserAllListByIDs(ctx, userModel.GetUserListByIDsRequest{IDs: idItems, Page: req.Page, Limit: req.Limit})
	if err != nil {
		return nil, nil, err
	}

	var resItems []GetUserItemList
	for _, item := range res {
		resItems = append(resItems, GetUserItemList{
			ID:           item.ID.Hex(),
			Email:        item.Email.String(),
			FirstName:    item.FirstName.String(),
			LastName:     item.LastName.String(),
			NickName:     item.NickName.String(),
			ThumbnailURl: item.ThumbnailURl,
			IsActive:     item.IsActive,
			IsBaned:      item.IsBaned,
			IsVerify:     item.IsVerify,
			Slug:         item.Slug.String(),
			CreateAt:     item.CreateAt,
			UpdateAt:     item.UpdateAt,
		})
	}
	return resItems, paginated, err
}

func (s service) GetUserAllListByIDs(ctx context.Context, req userModel.GetUserListByIDsRequest) ([]GetUserItemList, *PaginatedData, error) {
	idItems := req.IDs
	items := strings.Split(req.UserIDStr, ",")
	for _, itemID := range items {
		objID, err := primitive.ObjectIDFromHex(itemID)
		if err == nil {
			idItems = append(idItems, objID)
		}
	}

	res, paginated, err := s.repository.GetUserAllListByIDs(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []GetUserItemList
	for _, item := range res {
		resItems = append(resItems, GetUserItemList{
			ID:           item.ID.Hex(),
			Email:        item.Email.String(),
			FirstName:    item.FirstName.String(),
			LastName:     item.LastName.String(),
			NickName:     item.NickName.String(),
			ThumbnailURl: item.ThumbnailURl,
			IsActive:     item.IsActive,
			IsBaned:      item.IsBaned,
			IsVerify:     item.IsVerify,
			Slug:         item.Slug.String(),
			CreateAt:     item.CreateAt,
			UpdateAt:     item.UpdateAt,
		})
	}
	return resItems, paginated, err
}

// root admin
func (s service) GetUserAllList(ctx context.Context, req userModel.GetUserListRequest) ([]GetUserItemList, *PaginatedData, error) {
	res, paginated, err := s.repository.GetUserAllList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []GetUserItemList
	for _, item := range res {
		resItems = append(resItems, GetUserItemList{
			ID:           item.ID.Hex(),
			Email:        item.Email.String(),
			FirstName:    item.FirstName.String(),
			LastName:     item.LastName.String(),
			NickName:     item.NickName.String(),
			ThumbnailURl: item.ThumbnailURl,
			IsActive:     item.IsActive,
			IsBaned:      item.IsBaned,
			IsVerify:     item.IsVerify,
			Slug:         item.Slug.String(),
			CreateAt:     item.CreateAt,
			UpdateAt:     item.UpdateAt,
		})
	}
	return resItems, paginated, err
}
