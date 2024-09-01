package contractTemplate

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
func (h Handler) CreateContractTemplate(c *gin.Context) {
	var req CreateContractTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.Insert(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการเพิ่มข้อมูลสำเร็จ (Your create contract template has been completed.)",
	})
}

func (h Handler) GetMyContractTemplateList(c *gin.Context) {
	var req GetContractTemplateRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.MyChannel = true
	res, paginatedData, err := h.service.GetContractTemplateList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[BaseContractTemplateItemsResponse[[]ContractTemplateResponse]]{
		APIResponse: api.APIResponse[BaseContractTemplateItemsResponse[[]ContractTemplateResponse]]{
			Success: true,
			Result:  BaseContractTemplateItemsResponse[[]ContractTemplateResponse]{ContractList: res},
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}

func (h Handler) GetAllPublisherContractTemplateList(c *gin.Context) {
	var req GetContractTemplateRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.MyChannel = false
	res, paginatedData, err := h.service.GetContractTemplateList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[BaseContractTemplateItemsResponse[[]ContractTemplateResponse]]{
		APIResponse: api.APIResponse[BaseContractTemplateItemsResponse[[]ContractTemplateResponse]]{
			Success: true,
			Result:  BaseContractTemplateItemsResponse[[]ContractTemplateResponse]{ContractList: res},
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}

func (h Handler) GetContractTemplateDetail(c *gin.Context) {
	var req GetContractTemplateDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, err := h.service.Find(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[BaseContractTemplateItemResponse[*ContractTemplate]]{
		Success: true,
		Message: "",
		Result:  BaseContractTemplateItemResponse[*ContractTemplate]{Item: res},
	})
}

func (h Handler) UpdateContract(c *gin.Context) {
	var req UpdateContractTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateContract(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update contract template has been completed.)",
	})
}

func (h Handler) UpdateActive(c *gin.Context) {
	var req UpdateContractTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateActive(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update contract template has been completed.)",
	})
}

func (h Handler) UpdatePublisher(c *gin.Context) {
	var req UpdateContractTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdatePublisher(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update contract template has been completed.)",
	})
}
