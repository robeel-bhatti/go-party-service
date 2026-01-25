package service

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

func NewPartyError(path string, err error) *PartyError {
	p := &PartyError{
		Timestamp: time.Now(),
		Path:      path,
		Message:   err.Error(),
	}

	for e, c := range ErrMap {
		if errors.Is(err, e) {
			p.Code = c
			p.Status = http.StatusText(c)
			break
		}
	}

	return p
}
