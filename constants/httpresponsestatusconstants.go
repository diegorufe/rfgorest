package constants

import "net/http"

// HttpStatusType : type for http status
type HttpStatusType int

const (
	// HttpStatusOk : 200 -
	HttpStatusOk HttpStatusType = http.StatusOK

	// HttpStatusInternalServerError : 500
	HttpStatusInternalServerError HttpStatusType = http.StatusInternalServerError
)
