package rfhttputils

import (
	"encoding/json"
	"net/http"
	"rfgorest/beans/rfhttpresponserequest"
	"rfgorest/constants/rfhttpresponsestatusconstants"
)

// EncodeJsonDataResponseWriter : pass json data to response writer
func EncodeJsonDataResponseWriter(responseWrite http.ResponseWriter, responseRequest rfhttpresponserequest.RestRequestResponse) {
	jsonResult, err := json.Marshal(responseRequest)

	if err != nil {
		http.Error(responseWrite, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWrite.WriteHeader(int(responseRequest.Status))
	responseWrite.Header().Set("Content-Type", "application/json")
	responseWrite.Write(jsonResult)
}

// StatusOkInResponseRequest : Method to ser status ok in response
func StatusOkInResponseRequest(responseRequest *rfhttpresponserequest.RestRequestResponse) {
	responseRequest.Status = rfhttpresponsestatusconstants.HttpStatusOk
}
