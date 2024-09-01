package custom_error

type Type int

const (
	TypeConflict Type = iota
	TypeInternal
	TypeBadRequest
)

type Error interface {
	error
	GetType() Type
}
