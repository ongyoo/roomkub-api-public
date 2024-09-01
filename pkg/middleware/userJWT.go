package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	//"github.com/ongyoo/roomkub-api/cmd/user/role"
	"github.com/ongyoo/roomkub-api/pkg/api"
)

const userClaimsKey = "userClaims"

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type UserClaims struct {
	Payload UserPayload `json:"payload"`
	//jwt.RegisteredClaims
	jwt.RegisteredClaims
}

type UserPayload struct {
	ID           string `json:"id" `
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	NickName     string `json:"nick_name"`
	ThumbnailURl string `json:"thumbnail_url"`
	//Role         *role.UserRole `json:"role"`
	RoleID     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	BusinessID string `json:"business_id"`
}

func GenerateJWT(payload UserPayload) (tokenString string, err error) {
	expirationTime := time.Now().Add(168 * time.Hour)
	claims := &UserClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			//ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (userClaims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func UserJWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//sellerJWT := c.GetHeader("authentication")
		bearerAuthen := c.Request.Header.Get("authentication")
		userJWT := strings.Split(bearerAuthen, "Bearer ")
		if len(userJWT) == 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"debug_token": userJWT,
				// "debug_header": c.Request.Header,
				"message": "Token fail",
			})
			return
		}
		claims := &UserClaims{}

		token, err := jwt.ParseWithClaims(userJWT[len(userJWT)-1], claims, func(*jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"debug_token": userJWT,
				// "debug_header": c.Request.Header,
				"message": "UserJWT " + err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			err = errors.New("couldn't parse claims")
			return
		}
		if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
			err = errors.New("token expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.APIErrorMessage{
				ErrorCode: http.StatusBadRequest,
				Message:   err.Error(),
			})
			return
		}

		c.Set(userClaimsKey, claims)
	}
}

func GetUserClaims(c *gin.Context) (UserClaims, bool, error) {
	claims, ok := c.Get(userClaimsKey)
	if !ok {
		return UserClaims{}, false, errors.New("user claims not found")
	}

	result, ok := claims.(*UserClaims)
	if !ok {
		return UserClaims{}, false, errors.New("unable to parse user claims from context")
	}

	if result.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return UserClaims{}, true, errors.New("token expired")
	}

	return *result, false, nil
}

func GetContextUserClaims(ctx context.Context) (UserClaims, bool, error) {
	claims := ctx.Value(userClaimsKey)
	if claims == nil {
		return UserClaims{}, false, errors.New("user claims not found")
	}

	result, ok := claims.(*UserClaims)
	if !ok {
		return UserClaims{}, false, errors.New("unable to parse user claims from context")
	}

	if result.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return UserClaims{}, true, errors.New("token expired")
	}

	return *result, false, nil
}
