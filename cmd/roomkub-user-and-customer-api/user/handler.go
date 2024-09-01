package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ongyoo/roomkub-api/pkg/api"
	crypto "github.com/ongyoo/roomkub-api/pkg/crypto"

	//. "github.com/gobeam/mongo-go-pagination"
	userModel "github.com/ongyoo/roomkub-api/pkg/model/user"
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
func (h Handler) CreateUser(c *gin.Context) {
	var req userModel.UserCreateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	if len(req.Password) > 8 {
		hash, err := crypto.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, api.APIErrorMessage{
				ErrorCode: http.StatusBadRequest,
				Message:   err.Error(),
			})
			return
		}
		req.Password = hash
	}

	userId, err := h.service.Create(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[*string]{
		Success: true,
		Message: "ลงทะเบียนสำเร็จ",
		Result:  userId,
	})
}

// Login User godoc
// @Security     ApiKeyAuth
// @Summary      Post LoginUser
// @Description  Post LoginUser
// @Tags         LoginUser
// @Success      200       {object}  object{data=String}
// @Failure      400
// @Failure      500
// @Router       /user/create [post]
func (h Handler) LoginUser(c *gin.Context) {
	var req userModel.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	jwtToken, err := h.service.Login(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[LoginUserResponse]{
		Success: true,
		Result: LoginUserResponse{
			Token: jwtToken,
		},
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
func (h Handler) UpdateProfileInfo(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdateProfileInfo(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UpdateMyProfileInfo(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdateMyProfileInfo(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UpdatePassword(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdatePassword(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UpdateMyPassword(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdateMyPassword(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UpdateProfileImage(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdateProfileImage(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UpdateMyProfileImage(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}
	err := h.service.UpdateMyProfileImage(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) BanUser(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.User.IsBaned = true
	err := h.service.UpdateBan(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) UnBanUser(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.User.IsBaned = false
	err := h.service.UpdateBan(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) VerifyUser(c *gin.Context) {
	var req userModel.UpdateUserFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	req.User.IsVerify = true
	err := h.service.UpdateVerify(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIMessage{
		Success: true,
		Message: "ทำรายการสำเร็จ",
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
func (h Handler) GetUserDetail(c *gin.Context) {
	var req userModel.GetUserByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, err := h.service.GetById(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.APIResponse[*userModel.User]{
		Success: true,
		Message: "",
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
func (h Handler) GetUserAllListByUserIDs(c *gin.Context) {
	var req userModel.GetUserListByIDsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserAllListByIDs(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]GetUserItemList]{
		APIResponse: api.APIResponse[[]GetUserItemList]{
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
func (h Handler) GetUserAllListByBusinessChannel(c *gin.Context) {
	var req userModel.GetUserListByChannelIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserAllListByBusinessChannel(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]GetUserItemList]{
		APIResponse: api.APIResponse[[]GetUserItemList]{
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
func (h Handler) GetUserAllList(c *gin.Context) {
	var req userModel.GetUserListRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	res, paginatedData, err := h.service.GetUserAllList(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, api.APIErrorMessage{
			ErrorCode: http.StatusBadRequest,
			Message:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PaginatedContent[[]GetUserItemList]{
		APIResponse: api.APIResponse[[]GetUserItemList]{
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
