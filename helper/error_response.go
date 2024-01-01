package helper

import "errors"

var (
	ErrorNotFound            = errors.New("data not found")
	ErrorInternalServerError = errors.New("internal server error")
	ErrorBadRequest          = errors.New("bad request")
	ErrorConflict            = errors.New("data already exist")
	ErrorUnauthorized        = errors.New("unauthorized")
	ErrorForbidden           = errors.New("forbidden")
)
