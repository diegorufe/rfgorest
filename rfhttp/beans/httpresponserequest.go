package beans

import (
	"rfgorest/constants"
)

// RestRequestResponse : struct for store response data
type RestRequestResponse struct {
	Data            interface{}                     `json:"data"`
	Status          constants.HttpStatusType        `json:"status"`
	MessageResponse string                          `json:"messageResponse"`
	CodeError       constants.CodeErrorResponseType `json:"codeError"`
	Token           string                          `json:"token"`
	MapParams       map[string]interface{}          `json:"MapParams"`
}

// NewRestRequestResponse : function to create RestRequestResponse
func NewRestRequestResponse() *RestRequestResponse {
	var requestResponse *RestRequestResponse = new(RestRequestResponse)
	requestResponse.Status = constants.HttpStatusInternalServerError
	return requestResponse
}
