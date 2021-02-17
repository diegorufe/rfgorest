package beans

import "rfgorest/constants"

// ResponseError for send error when request failt
type ResponseError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"httpStatus"`
}

// NewResponseError : function to create ResponseError
func NewResponseError() *ResponseError {
	var responseError *ResponseError = new(ResponseError)
	responseError.HttpStatus = int(constants.HttpStatusInternalServerError)
	return responseError
}
