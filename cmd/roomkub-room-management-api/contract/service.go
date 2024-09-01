package roomContract

import (
	"context"
	"errors"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Customer
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/customer"
	// Contract template
	contractTemplate "github.com/ongyoo/roomkub-api/cmd/roomkub-room-management-api/contract/template"
	// Upload
	//"github.com/ongyoo/roomkub-api/pkg/upload"
	// Crypto
	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
	// Utils
)

type Service interface {
	// Insert
	Create(ctx context.Context, req CreateRoomContractRequest) (*string, error)
	// get
	Find(ctx context.Context, req FindRoomContractRequest) (*RoomContract, error)
	GetRoomContractList(ctx context.Context, req GetRoomContractRequest) ([]RoomContract, *PaginatedData, error)
	// update
	Update(ctx context.Context, req UpdateRoomContractRequest) error
}

type service struct {
	repository              Repository
	customerService         customer.Service
	contractTemplateService contractTemplate.Service
}

func NewService(repo Repository, customerService customer.Service, contractTemplateService contractTemplate.Service) *service {
	return &service{repository: repo,
		customerService:         customerService,
		contractTemplateService: contractTemplateService}
}

// Insert
func (s service) Create(ctx context.Context, req CreateRoomContractRequest) (*string, error) {
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

	if req.RoomIDStr != "" {
		roomID, err := primitive.ObjectIDFromHex(req.RoomIDStr)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.RoomID = roomID
	}

	// CustomerIDStr empty
	if req.CustomerIDStr == "" {
		return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}

	customerID, err := primitive.ObjectIDFromHex(req.CustomerIDStr)
	if err != nil {
		return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.CustomerID = customerID

	// check document
	if len(req.Documents) == 0 {
		// Find Customer
		res, err := s.customerService.Find(ctx, customer.FindCustomerRequest{IDStr: req.CustomerIDStr})
		if err != nil {
			return nil, errors.New(err.Error() + " [Find Customer] กรุณาติดต่อผู้ดูแลระบบ")
		}

		if len(res.Documents) == 0 {
			return nil, errors.New(" [Find Customer] ไม่มีข้อมูลเอกสารกรุณาอัพโหลดเอกสาร")
		}
		req.Documents = res.Documents
	}

	// check template
	res, err := s.contractTemplateService.Find(ctx, contractTemplate.GetContractTemplateDetailRequest{IDStr: req.ContractTemplateIDStr})
	if err != nil {
		return nil, errors.New(err.Error() + " [Find Customer] กรุณาติดต่อผู้ดูแลระบบ")
	}

	req.ContractTemplate = ContractTemplateContent{
		RefID:    res.ID,
		Name:     res.Name,
		Html:     res.Html,
		Remark:   res.Remark,
		CreateAt: res.CreateAt,
		UpdateAt: res.UpdateAt,
	}

	if req.IsMigrate {
		createDate, err := time.Parse("2006-01-02", req.CreateDateStr)
		if err != nil {
			return nil, errors.New("กรุณาตรวจสอบวันที่ใหม่อีกครั้ง")
		}
		req.CreateAt = createDate
	} else {
		req.CreateAt = time.Now()
	}
	req.IsActive = true
	req.IsReportOut = false
	req.UpdateAt = time.Now()
	return s.repository.Insert(ctx, req)
}

func (s service) Find(ctx context.Context, req FindRoomContractRequest) (*RoomContract, error) {
	if !req.IsRoot {
		if req.BusinessIDStr != "" {
			return nil, errors.New(" [Find Customer] เกิดข้อผิดพลาด")
		}
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = businessID
	}

	objID, err := primitive.ObjectIDFromHex(req.IDStr)
	if err != nil {
		return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.ID = objID
	return s.repository.Find(ctx, req)
}

func (s service) GetRoomContractList(ctx context.Context, req GetRoomContractRequest) ([]RoomContract, *PaginatedData, error) {
	if !req.IsRoot {
		if req.BusinessIDStr != "" {
			return nil, nil, errors.New(" [Find Customer] เกิดข้อผิดพลาด")
		}
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return nil, nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = businessID
	}

	roomID, err := primitive.ObjectIDFromHex(req.RoomIDStr)
	if err != nil {
		return nil, nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
	}
	req.RoomID = roomID
	return s.repository.GetRoomContractList(ctx, req)
}

func (s service) Update(ctx context.Context, req UpdateRoomContractRequest) error {
	if !req.IsRoot {
		if req.BusinessIDStr != "" {
			return errors.New(" [Find Customer] เกิดข้อผิดพลาด")
		}
		businessID, err := primitive.ObjectIDFromHex(req.BusinessIDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.BusinessID = businessID

		idObj, err := primitive.ObjectIDFromHex(req.IDStr)
		if err != nil {
			return errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		req.ID = idObj
	}
	req.UpdateAt = time.Now()
	return s.repository.Update(ctx, req)
}
