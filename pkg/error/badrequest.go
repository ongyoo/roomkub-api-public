package custom_error

import (
	"github.com/ongyoo/roomkub-api/pkg/error/code"
	"github.com/pkg/errors"
)

type BadRequest struct {
	error
	Code code.BadRequest
}

func NewBadRequest(error error, code code.BadRequest) BadRequest {
	if error == nil {
		error = errors.New("")
	}
	return BadRequest{error: error, Code: code}
}

func (b BadRequest) GetType() Type {
	return TypeBadRequest
}
