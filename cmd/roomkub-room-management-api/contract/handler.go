package roomContract

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/pkg/api"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

// Create Room Contract godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) CreateRoomContract(c *gin.Context) {
	var req CreateRoomContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	contractID, err := h.service.Create(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[CreateRoomContractResponse]{
		Success: true,
		Message: "ทำรายการเพิ่มข้อมูลสำเร็จ (Your create contract has been completed.)",
		Result: CreateRoomContractResponse{
			ContractID: contractID,
		},
	})
}

func (h Handler) GetRoomContractList(c *gin.Context) {
	var req GetRoomContractRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetRoomContractList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[BaseRoomContractItemsResponse[[]RoomContract]]{
		APIResponse: api.APIResponse[BaseRoomContractItemsResponse[[]RoomContract]]{
			Success: true,
			Result:  BaseRoomContractItemsResponse[[]RoomContract]{ContractList: res},
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}

func (h Handler) UpdateRoomContract(c *gin.Context) {
	var req UpdateRoomContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.Update(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update contract has been completed.)",
	})
}

// ROOT
func (h Handler) GetRootRoomContractList(c *gin.Context) {
	var req GetRoomContractRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetRoomContractList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[BaseRoomContractItemsResponse[[]RoomContract]]{
		APIResponse: api.APIResponse[BaseRoomContractItemsResponse[[]RoomContract]]{
			Success: true,
			Result:  BaseRoomContractItemsResponse[[]RoomContract]{ContractList: res},
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}
