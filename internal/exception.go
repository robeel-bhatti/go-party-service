package internal

import (
	"errors"
	"net/http"
	"time"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")

	ErrMap = map[error]int{
		ErrNotFound:            http.StatusNotFound,
		ErrInternalServerError: http.StatusInternalServerError,
		ErrBadRequest:          http.StatusBadRequest,
		ErrUnprocessableEntity: http.StatusUnprocessableEntity,
	}
)

// PartyError represents the response payload during an exception
// Fields should be capitalized (exported) so json.Marshal can serialize the struct
type PartyError struct {
	Timestamp time.Time
	Path      string
	Status    string
	Code      int
	Message   string
}
