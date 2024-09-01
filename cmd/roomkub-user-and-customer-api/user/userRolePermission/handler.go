package userRolePermission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/pkg/api"
	rolePermission "github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
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
func (h Handler) CreateUserRole(c *gin.Context) {
	var req rolePermission.UserRoleCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.CreateUserRole(c, req)
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
func (h Handler) CreateUserPermission(c *gin.Context) {
	var req rolePermission.UserPermissionCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.CreateUserPermission(c, req)
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
func (h Handler) UpdateUserRoleInfo(c *gin.Context) {
	var req rolePermission.UserRoleCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateUserRoleInfo(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update transaction has been completed.)",
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
func (h Handler) UpdateUserRolePermission(c *gin.Context) {
	var req rolePermission.UserRoleCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateUserRolePermission(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update transaction has been completed.)",
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
func (h Handler) UpdateUserPermissionInfo(c *gin.Context) {
	var req rolePermission.UserPermissionCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.UpdateUserPermissionInfo(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการอัพเดตข้อมูลสำเร็จ (Your update transaction has been completed.)",
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
func (h Handler) DeleteUserRolePermission(c *gin.Context) {
	var req rolePermission.UserRoleCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.DeleteUserRolePermission(c, req)
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
func (h Handler) DeleteUserPermission(c *gin.Context) {
	var req rolePermission.UserPermissionCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	err := h.service.DeleteUserPermission(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการลบข้อมูลสำเร็จ (Your update transaction has been completed.)",
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
func (h Handler) GetUserRoleAllList(c *gin.Context) {
	var req rolePermission.GetUserRoleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	req.Active = 1

	res, paginatedData, err := h.service.GetUserRoleAllList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserRole]{
		APIResponse: api.APIResponse[[]rolePermission.UserRole]{
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
func (h Handler) GetRootUserRoleAllList(c *gin.Context) {
	var req rolePermission.GetUserRoleListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.IsRoot = true
	res, paginatedData, err := h.service.GetUserRoleAllList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserRole]{
		APIResponse: api.APIResponse[[]rolePermission.UserRole]{
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
func (h Handler) GetUserPermissionAllList(c *gin.Context) {
	var req rolePermission.GetUserPermissionListRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserPermissionAllList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserPermission]{
		APIResponse: api.APIResponse[[]rolePermission.UserPermission]{
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
func (h Handler) GetRootUserPermissionAllList(c *gin.Context) {
	var req rolePermission.GetUserPermissionListRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.IsRoot = true
	res, paginatedData, err := h.service.GetUserPermissionAllList(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserPermission]{
		APIResponse: api.APIResponse[[]rolePermission.UserPermission]{
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
func (h Handler) GetUserRoleByID(c *gin.Context) {
	var req rolePermission.GetUserRoleByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, err := h.service.GetUserRoleByID(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[*UserRoleResponse]{
		Success: true,
		Message: "Success",
		Result:  res,
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
func (h Handler) GetUserRoleByIDs(c *gin.Context) {
	var req rolePermission.GetUserRoleByIdsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserRoleByIDs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserRole]{
		APIResponse: api.APIResponse[[]rolePermission.UserRole]{
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
func (h Handler) GetUserPermissionByID(c *gin.Context) {
	var req rolePermission.GetRolePermissionByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, err := h.service.GetUserPermissionByID(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[*rolePermission.UserPermission]{
		Success: true,
		Message: "Success",
		Result:  res,
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
func (h Handler) GetUserPermissionByIDs(c *gin.Context) {
	var req rolePermission.GetRolePermissionByIdsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserPermissionByIDs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]rolePermission.UserPermission]{
		APIResponse: api.APIResponse[[]rolePermission.UserPermission]{
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
