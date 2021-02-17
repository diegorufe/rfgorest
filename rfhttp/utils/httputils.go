package utils

import (
	"encoding/json"
	"net/http"
	"rfgocore/utils/utilsstring"
	"rfgorest/rfhttp/beans"
	"rfgorest/rfhttp/constants"
)

// EncodeJsonDataResponseWriter : pass json data to response writer
func EncodeJsonDataResponseWriter(responseWrite http.ResponseWriter, responseRequest beans.RestRequestResponse) {
	jsonResult, err := json.Marshal(responseRequest)

	if err != nil {

		http.Error(responseWrite, utilsstring.IntToString(int(constants.CodeErrorMarshalResponseWriter)), http.StatusInternalServerError)

	} else {
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

// EncodeRequestBody : Method for encode request body
func EncodeRequestBody(req *http.Request) (beans.RestRequestBody, error) {
	var requestBody beans.RestRequestBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	return requestBody, err
}

// SetupCorsResponseOriginAll : method for set cors reponse origin all
func SetupCorsResponseOriginAll(res *http.ResponseWriter, req *http.Request) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*res).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
