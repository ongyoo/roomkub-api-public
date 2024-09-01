package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/ongyoo/roomkub-api/cmd/roomkub-user-customer-api/user"
	"github.com/ongyoo/roomkub-api/cmd/roomkub-user-and-customer-api/user/userRolePermission"
	"github.com/ongyoo/roomkub-api/pkg/api"
	"github.com/ongyoo/roomkub-api/pkg/model/rolePermission"
)

type Service interface {
	ValidatePermission() gin.HandlerFunc
}

type service struct {
	rootSlug              string
	rolePermissionService userRolePermission.Service
}

func NewPermissionService(rolePermissionService userRolePermission.Service, rootSlug string) *service {
	return &service{rootSlug, rolePermissionService}
}

func (s service) ValidatePermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, _, err := GetUserClaims(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
				ErrorCode: http.StatusUnauthorized,
				Message:   "something wrong (เกิดข้อผิดพลาดกรุณาติดต่อผู้ดูแลระบบ) " + err.Error(),
			})
			return
		}
		if userClaims.Payload.RoleID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
				ErrorCode: http.StatusUnauthorized,
				Message:   "something wrong [1][role id not found] (เกิดข้อผิดพลาดกรุณาติดต่อผู้ดูแลระบบ)",
			})
			return
		}

		objRoleID, objRoleIDErr := primitive.ObjectIDFromHex(userClaims.Payload.RoleID)
		if objRoleIDErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
				ErrorCode: http.StatusUnauthorized,
				Message:   "something wrong [2][role id not found] (เกิดข้อผิดพลาดกรุณาติดต่อผู้ดูแลระบบ)",
			})
			return
		}
		res, err := s.rolePermissionService.GetRootUserRoleByID(c, rolePermission.GetUserRoleByIdRequest{IDStr: objRoleID.Hex()})
		if err != nil || len(res.Permissions) == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, api.APIErrorMessage{
				ErrorCode: http.StatusNotFound,
				Message:   "roleService.GetById " + err.Error(),
			})
			return
		}

		if res.Slug == s.rootSlug {
			// for root role only
			return
		}

		filter := userRolePermission.UserPermissionResponse{}
		path := c.FullPath()
		permissionItems := res.Permissions
		for i := range permissionItems {
			if strings.Contains(path, permissionItems[i].Path) {
				filter = permissionItems[i]
			}
		}

		errDeniedMsg := errors.New("Permission denied (ไม่มีสิทธิ์เข้าถึงข้อมูล)")
		if filter.Path == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, api.APIErrorMessage{
				ErrorCode: http.StatusBadRequest,
				Message:   errDeniedMsg.Error(),
			})
			return
		}

		switch c.Request.Method {
		case "GET":
			if !filter.AllowGetMethod {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
					ErrorCode: http.StatusUnauthorized,
					Message:   errDeniedMsg.Error(),
				})
				return

			}
		case "POST":
			if !filter.AllowPostMethod {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
					ErrorCode: http.StatusUnauthorized,
					Message:   errDeniedMsg.Error(),
				})
				return

			}
		case "PUT":
			if !filter.AllowPutMethod {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
					ErrorCode: http.StatusUnauthorized,
					Message:   errDeniedMsg.Error(),
				})
				return

			}

		case "DELETE":
			if !filter.AllowDeleteMethod {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
					ErrorCode: http.StatusUnauthorized,
					Message:   errDeniedMsg.Error(),
				})
				return

			}
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
				ErrorCode: http.StatusUnauthorized,
				Message:   errDeniedMsg.Error(),
			})
			return
		}
	}
}
