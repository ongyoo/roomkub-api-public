package custom_error

import (
	"github.com/ongyoo/roomkub-api/pkg/error/code"
	"github.com/pkg/errors"
)

type Conflict struct {
	error
	Code code.Conflict
}

func NewConflict(error error, code code.Conflict) Conflict {
	if error == nil {
		error = errors.New("")
	}
	return Conflict{error: error, Code: code}
}

func (c Conflict) GetType() Type {
	return TypeConflict
}
