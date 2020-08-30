package rfhttpresponserequest

import (
	"rfgorest/constants/rfhttpresponsecodeerrors"
	"rfgorest/constants/rfhttpresponsestatusconstants"
)

// RestRequestResponse : struct for store response data
type RestRequestResponse struct {
	Data            interface{}                                    `json:"data"`
	Status          rfhttpresponsestatusconstants.HttpStatusType   `json:"status"`
	MessageResponse string                                         `json:"messageResponse"`
	CodeError       rfhttpresponsecodeerrors.CodeErrorResponseType `json:"codeError"`
}

// NewRestRequestResponse : function to create RestRequestResponse
func NewRestRequestResponse() *RestRequestResponse {
	var requestResponse *RestRequestResponse = new(RestRequestResponse)
	requestResponse.Status = rfhttpresponsestatusconstants.HttpStatusInternalServerError
	return requestResponse
}
