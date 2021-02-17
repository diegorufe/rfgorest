package constants

import "net/http"

// HttpStatusType : type for http status
type HttpStatusType int

const (
	// HttpStatusOk : 200 -
	HttpStatusOk HttpStatusType = http.StatusOK

	// HttpStatusCreated : 201 -
	HttpStatusCreated HttpStatusType = http.StatusCreated

	// HttpStatusInternalServerError : 500
	HttpStatusInternalServerError HttpStatusType = http.StatusInternalServerError
)
