package businessChannel

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

// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) CreateBusinessChannel(c *gin.Context) {
	var req BusinessChannelCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.CreateBusinessChannel(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการเพิ่มข้อมูลสำเร็จ (Your create transaction has been completed.)",
	})
}

// Get
// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) GetBusinessChannelListByUserId(c *gin.Context) {
	var req GetBusinessChannelListByUserIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetBusinessChannelListByUserId(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]GetBusinessChannelItem]{
		APIResponse: api.APIResponse[[]GetBusinessChannelItem]{
			Success: true,
			Result:  res,
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}

// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) UpdateDelete(c *gin.Context) {
	var req UpdateDeleteBusinessFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateDelete(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการลบข้อมูลสำเร็จ (Your delete transaction has been completed.)",
	})
}

// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) UpdateSetting(c *gin.Context) {
	var req UpdateBusinessChannelSettingFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateSetting(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update data transaction has been completed.)",
	})
}

// Root
// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) GetBusinessChannelAllList(c *gin.Context) {
	var req GetBusinessChannelListRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetBusinessChannelAllList(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]GetRootBusinessChannelItem]{
		APIResponse: api.APIResponse[[]GetRootBusinessChannelItem]{
			Success: true,
			Result:  res,
		},
		Total:     paginatedData.Pagination.Total,
		Page:      paginatedData.Pagination.Page,
		PerPage:   paginatedData.Pagination.PerPage,
		Prev:      paginatedData.Pagination.Prev,
		Next:      paginatedData.Pagination.Next,
		TotalPage: paginatedData.Pagination.TotalPage,
	})
}

// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) UpdateBan(c *gin.Context) {
	var req UpdateBanBusinessFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateBan(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการ Ban ข้อมูลสำเร็จ (Your ban transaction has been completed.)",
	})
}

// Create User godoc
// @Security     ApiKeyAuth
// @Summary      Post CreateUser
// @Description  Post CreateUser
// @Tags         Create User
// @Param        X-BP-JWT  header    string  false  "Insert X-BP-JWT value"
// @Success      200       {object}  object{data=[]String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) UpdateUnBan(c *gin.Context) {
	var req UpdateBanBusinessFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateUnBan(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการ Ban ข้อมูลสำเร็จ (Your ban transaction has been completed.)",
	})
}
