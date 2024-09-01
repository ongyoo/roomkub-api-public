package Invoice

import (
	"context"
	"errors"
	"fmt"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"

	roomPipeline "github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/pipeline"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user"
	userModel "github.com/ongyoo/roomkub-api/pkg/model/user"

	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
)

type Service interface {
	// insert
	Create(ctx context.Context, req CreateInvoiceRequest) (*string, error)
	// get
	GetInvoiceList(ctx context.Context, req GetInvoiceListRequest) ([]InvoiceResponse, *PaginatedData, error)
	Find(ctx context.Context, req FindInvoiceRequest) (*InvoiceDetailResponse, error)
	/*
		// update
		Update(ctx context.Context, req UpdateInvoiceRequest) error
		UpdateItems(ctx context.Context, req UpdateInvoiceRequest) error
		UpdatePrinted(ctx context.Context, req UpdateInvoiceRequest) error
		UpdateSendNotify(ctx context.Context, req UpdateInvoiceRequest) error
	*/
}

type service struct {
	repository          Repository
	userService         user.Service
	customerService     customer.Service
	roomPipelineService roomPipeline.Service
}

func NewService(repo Repository, userService user.Service, customerService customer.Service, roomPipelineService roomPipeline.Service) *service {
	return &service{repository: repo,
		userService:         userService,
		customerService:     customerService,
		roomPipelineService: roomPipelineService,
	}
}

// Insert
func (s service) Create(ctx context.Context, req CreateInvoiceRequest) (*string, error) {
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
		req.CreateUserID = objID
	}

	var totalPrice float32 = 0.0
	editItems := req.Items
	for index, item := range req.Items {
		editItem := item
		editItem.TotalPrice = item.Price * float32(item.Unit)
		editItems[index] = editItem
		totalPrice += editItem.TotalPrice
	}

	req.Items = editItems
	req.TotalPrice = totalPrice
	req.IsActive = true
	req.Status = InvoiceStatusWait
	req.IsSendNotify = false
	req.IsPrinted = false
	req.CreateAt = time.Now()
	req.UpdateAt = time.Now()
	return s.repository.Insert(ctx, req)
}

// get
func (s service) GetInvoiceList(ctx context.Context, req GetInvoiceListRequest) ([]InvoiceResponse, *PaginatedData, error) {
	res, paginatedData, err := s.repository.GetInvoiceList(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	var resItems []InvoiceResponse
	for _, item := range res {
		// RoomName
		roomName := "ไม่มีข้อมูล"
		if !item.RoomPipelineID.IsZero() {
			res, err := s.roomPipelineService.Find(ctx, roomPipeline.FindroomPipelineRequest{IDStr: item.RoomPipelineID.Hex()})
			if err == nil {
				roomName = res.RoomName
			}
		}
		// CustomerID
		customerName := "ไม่มีข้อมูล"
		if !item.CustomerID.IsZero() {
			// สำหรับ get พนักงาน
			if item.Type == InvoiceTypeOther {
				res, err := s.userService.GetById(ctx, userModel.GetUserByIdRequest{UserID: item.CustomerID.Hex()})
				if err == nil {
					customerName = fmt.Sprintf("%s %s (%s)", res.FirstName.String(), res.LastName.String(), res.NickName.String())
				}
			} else {
				res, err := s.customerService.Find(ctx, customer.FindCustomerRequest{IDStr: item.CustomerID.Hex()})
				if err == nil {
					customerName = fmt.Sprintf("%s %s (%s)", res.FirstName, res.LastName, res.NickName)
				}
			}
		}

		// CreateByID
		createUserName := "ไม่มีข้อมูล"
		if !item.CreateUserID.IsZero() {
			res, err := s.userService.GetById(ctx, userModel.GetUserByIdRequest{UserID: item.CreateUserID.Hex()})
			if err == nil {
				createUserName = fmt.Sprintf("%s %s (%s)", res.FirstName.String(), res.LastName.String(), res.NickName.String())
			}
		}

		resItems = append(resItems, InvoiceResponse{
			ID:             item.ID.Hex(),
			BusinessID:     item.BusinessID.Hex(),
			RoomName:       roomName,
			CustomerName:   customerName,
			CreateUserName: createUserName,
			TotalPrice:     item.TotalPrice,
			Status:         item.Status,
			Type:           item.Type,
			IsSendNotify:   item.IsSendNotify,
			IsPrinted:      item.IsPrinted,
			CreateAt:       item.CreateAt,
			UpdateAt:       item.UpdateAt,
		})
	}
	return resItems, paginatedData, err
}

func (s service) Find(ctx context.Context, req FindInvoiceRequest) (*InvoiceDetailResponse, error) {
	res, err := s.repository.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	// RoomName
	roomName := "ไม่มีข้อมูล"
	if !res.RoomPipelineID.IsZero() {
		res, err := s.roomPipelineService.Find(ctx, roomPipeline.FindroomPipelineRequest{IDStr: res.RoomPipelineID.Hex()})
		if err == nil {
			roomName = res.RoomName
		}
	}
	// CustomerID
	customerObj := UserOrCustomerRespons{}
	if !res.CustomerID.IsZero() {
		// สำหรับ get พนักงาน
		if res.Type == InvoiceTypeOther {
			res, err := s.userService.GetById(ctx, userModel.GetUserByIdRequest{UserID: res.CustomerID.Hex()})
			if err == nil && res != nil {
				customerObj = UserOrCustomerRespons{Data: UserData{res}}
			}
		} else {
			res, err := s.customerService.Find(ctx, customer.FindCustomerRequest{IDStr: res.CustomerID.Hex()})
			if err == nil && res != nil {
				customerObj = UserOrCustomerRespons{Data: CustomerData{res}}
			}
		}
	}

	// CreateByID
	createUser := userModel.User{}
	if !res.CreateUserID.IsZero() {
		userRes, err := s.userService.GetById(ctx, userModel.GetUserByIdRequest{UserID: res.CreateUserID.Hex()})
		if err == nil && userRes != nil {
			createUser = *userRes
		}
	}

	item := InvoiceDetailResponse{
		ID:           res.ID.Hex(),
		BusinessID:   res.BusinessID.Hex(),
		RoomName:     roomName,
		Customer:     customerObj,
		CreateByUser: createUser,
		Items:        res.Items,
		TotalPrice:   res.TotalPrice,
		Status:       res.Status,
		Type:         res.Type,
		IsSendNotify: res.IsSendNotify,
		IsPrinted:    res.IsPrinted,
		CreateAt:     res.CreateAt,
		UpdateAt:     res.UpdateAt,
		SendNotifyAt: res.SendNotifyAt,
		PrintedAt:    res.PrintedAt,
	}
	return &item, nil
}
