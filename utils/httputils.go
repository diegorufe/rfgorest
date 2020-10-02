package utils

import (
	"encoding/json"
	"net/http"
	"rfgocore/utils/utilsstring"
	"rfgorest/beans"
	"rfgorest/constants"
)

// EncodeJsonDataResponseWriter : pass json data to response writer
func EncodeJsonDataResponseWriter(responseWrite http.ResponseWriter, responseRequest beans.RestRequestResponse) {
	jsonResult, err := json.Marshal(responseRequest)

	if err != nil {

		http.Error(responseWrite, utilsstring.IntToString(int(constants.CodeErrorMarshalResponseWriter)), http.StatusInternalServerError)

	} else {
		// if logger.IsDebugEnabled() {
		// 	logger.Debug("$$EncodeJsonDataResponseWriter response: " + string(jsonResult))
		// }
		responseWrite.Header().Set("Content-Type", "application/json; charset=utf-8")
		responseWrite.WriteHeader(int(responseRequest.Status))
		responseWrite.Write(jsonResult)

	}

}

// StatusOkInResponseRequest : Method to ser status ok in response
func StatusOkInResponseRequest(responseRequest *beans.RestRequestResponse) {
	responseRequest.Status = constants.HttpStatusOk
}

// StatusKoInResponseRequest : Method to ser status ok in response
func StatusKoInResponseRequest(responseRequest *beans.RestRequestResponse) {
	responseRequest.Status = constants.HttpStatusInternalServerError
}
