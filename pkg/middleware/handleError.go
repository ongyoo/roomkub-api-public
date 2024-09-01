package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	custom_error "github.com/ongyoo/roomkub-api/pkg/error"
	"github.com/ongyoo/roomkub-api/pkg/error/code"
)

type DefaultError struct {
	Message      string `json:"message"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewConflictError(err custom_error.Conflict) DefaultError {
	return DefaultError{
		Message:      "Data conflict",
		ErrorCode:    string(err.Code),
		ErrorMessage: err.Error(),
	}
}

func NewInternalServerError(err error) DefaultError {
	return DefaultError{Message: "Internal Server Error", ErrorCode: string(code.DefaultInternalServerError), ErrorMessage: err.Error()}
}

func NewSpecifiedInternalServerError(err custom_error.Internal) DefaultError {
	return DefaultError{
		Message:      "Internal Server Error",
		ErrorCode:    string(err.Code),
		ErrorMessage: err.Error(),
	}
}

func NewBadRequestError(err custom_error.BadRequest) DefaultError {
	return DefaultError{
		Message:      "Bad Request",
		ErrorCode:    string(err.Code),
		ErrorMessage: err.Error(),
	}
}

type ValidationError struct {
	ErrorCode map[string][]string `json:"errorCode"`
}

func HandleError(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 && c.Writer.Status() == http.StatusOK {
		return
	}

	// handle single, non-validation error
	if len(c.Errors) == 1 &&
		reflect.TypeOf(c.Errors.Last().Err) != reflect.TypeOf(validator.ValidationErrors{}) {
		if c.Errors.Last().Type == gin.ErrorTypePrivate {
			err, ok := c.Errors.Last().Err.(custom_error.Error)
			if ok {
				switch err.GetType() {
				case custom_error.TypeConflict:
					conflictError, ok := err.(custom_error.Conflict)
					if !ok {
						err := fmt.Errorf("unable to cast error to conflict")
						c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
					}
					c.JSON(http.StatusConflict, NewConflictError(conflictError))
					return
				case custom_error.TypeBadRequest:
					badRequestError, ok := err.(custom_error.BadRequest)
					if !ok {
						err := fmt.Errorf("unable to cast error to bad request")
						c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
					}
					c.JSON(http.StatusBadRequest, NewBadRequestError(badRequestError))
					return
				case custom_error.TypeInternal:
					internalError, ok := err.(custom_error.Internal)
					if !ok {
						err := fmt.Errorf("unable to cast error to internal")
						c.JSON(http.StatusInternalServerError, NewInternalServerError(err))
					}
					c.JSON(http.StatusInternalServerError, NewSpecifiedInternalServerError(internalError))
					return
				default:
					break
				}
			}
			c.JSON(http.StatusInternalServerError, NewInternalServerError(c.Errors.Last().Err))
			return
		}
	}

	// handle validation error
	errors := map[string][]string{}
	for _, ginErr := range c.Errors {
		if errs, ok := ginErr.Err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				key := regexp.MustCompile("^[a-z].+").FindString(e.Namespace())
				if key == "" {
					key = regexp.MustCompile("\\.[a-z].+").FindString(e.Namespace())[1:]
				}
				if key == "" {
					continue
				}

				// replace array [0] with empty string
				errCode := e.ActualTag()
				if _, ok := errors[key]; ok {
					errors[key] = append(errors[key], errCode)
				} else {
					errors[key] = []string{errCode}
				}
			}
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, ValidationError{ErrorCode: errors})
		return
	}
}
