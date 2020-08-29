package rfhttpresponserequest

import (
	"rfgorest/constants/rfhttpresponsecodeerrors"
	"rfgorest/constants/rfhttpresponsestatusconstants"
)

// RestRequestResponse : struct for store response data
type RestRequestResponse struct {
	Data            interface{}
	Status          rfhttpresponsestatusconstants.HttpStatusType
	MessageResponse string
	CodeError       rfhttpresponsecodeerrors.CodeErrorResponseType
}

// NewRestRequestResponse : function to create RestRequestResponse
func NewRestRequestResponse() *RestRequestResponse {
	var requestResponse *RestRequestResponse = new(RestRequestResponse)
	requestResponse.Status = rfhttpresponsestatusconstants.HttpStatusInternalServerError
	return requestResponse
}
