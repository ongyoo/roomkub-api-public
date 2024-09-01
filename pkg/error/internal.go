package custom_error

import (
	"github.com/ongyoo/roomkub-api/pkg/error/code"
	"github.com/pkg/errors"
)

type Internal struct {
	error
	Code code.Internal
}

func NewInternal(error error, code code.Internal) Internal {
	if error == nil {
		error = errors.New("")
	}
	return Internal{error: error, Code: code}
}

func (i Internal) GetType() Type {
	return TypeInternal
}
