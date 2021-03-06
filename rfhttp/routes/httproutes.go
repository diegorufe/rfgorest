package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"rfgocore/logger"
	"rfgocore/utils/utilsstring"
	databcore "rfgodata/beans/core"
	rfgodataconst "rfgodata/constants"
	"rfgodata/service"
	"rfgodata/transactions"
	transactiongorm "rfgodata/transactions/gorm"
	datautils "rfgodata/utils"
	"rfgorest/rfhttp"
	"rfgorest/rfhttp/beans"
	"rfgorest/rfhttp/constants"
	"rfgorest/rfhttp/utils"

	"gorm.io/gorm"
)

// HandleGetRoute : method for handler get route
func HandleGetRoute(rfHTTP *rfhttp.RFHttp, route string, handler http.HandlerFunc) {
	http.HandleFunc(route, func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodGet:
			// Serve the resource.
			handler(res, req)
			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
	})
}

// HandlePostRoute : method for handler get route
func HandlePostRoute(rfHTTP *rfhttp.RFHttp, route string, handler http.HandlerFunc) {
	http.HandleFunc(route, func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodPost:
			// Serve the resource.
			handler(res, req)
			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
	})
}

// HandlePostRouteWithTransaction : method for handle route with transaction
func HandlePostRouteWithTransaction(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string, passRequestBodyToFnAction bool,
	fnActionRoute func(service.IService, *map[string]interface{}, beans.RestRequestBody) (interface{}, error)) {

	http.HandleFunc(pathRoute, func(res http.ResponseWriter, req *http.Request) {

		// Setup cors
		utils.SetupCorsResponseOriginAll(&res, req)

		switch req.Method {

		case http.MethodOptions:

			break

		case http.MethodPost:

			// Serve the resource.
			var service service.IService = rfHTTP.GetService(keyService).(service.IService)
			var mapParamsService map[string]interface{} = make(map[string]interface{})

			// Start transaction context
			StartTransactionContext(rfHTTP, &mapParamsService, req)

			var err error
			var data interface{}
			var requestBody beans.RestRequestBody

			// If passRequestBodyToFnAction get it
			if passRequestBodyToFnAction {
				// Get request body
				requestBody, err = utils.EncodeRequestBody(req)

				if err != nil {
					panic(err)
				}
			}

			// Catch panic errors
			defer func() {
				if err := recover(); err != nil {
					// Finish request response on error
					ForcerFinishRequestRespose(res, data, err.(error), rfHTTP, &mapParamsService)
				}
			}()

			// execute function
			data, err = fnActionRoute(service, &mapParamsService, requestBody)

			// If has error finish transaction context
			defer ForcerFinishRequestRespose(res, data, err, rfHTTP, &mapParamsService)

			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
			break
		}
	})
}

// HandleCrudBrowserRoute : Method for handle browse route
func HandleCrudBrowserRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/browser", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		var requestBrowser databcore.RequestBrowser
		jsonbody, err := json.Marshal(requestBody.Data.(map[string]interface{}))

		if err == nil {
			json.Unmarshal(jsonbody, &requestBrowser)
		}

		responseService := (service).Browser(requestBrowser, mapParamsService)
		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudListRoute : Method for handle list route
func HandleCrudListRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/list", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		responseService := (service).List(requestBody.Fields, requestBody.Filters, requestBody.Joins, requestBody.Orders, nil, requestBody.Limit, mapParamsService)
		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudCountRoute : Method for handle count route
func HandleCrudCountRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/count", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		responseService := (service).Count(requestBody.Filters, requestBody.Joins, nil, mapParamsService)
		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudLoadNewRoute : Method for handle load new route
func HandleCrudLoadNewRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/loadNew", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		responseService := (service).LoadNew(mapParamsService)
		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudReadRoute : Method for handle load read data
func HandleCrudReadRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/read", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		responseService := (service).Read(requestBody.Data, mapParamsService)
		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudEditRoute : Method for handle load read data
func HandleCrudEditRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/edit", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {

		jsonBody, err := json.Marshal(requestBody.Data.(map[string]interface{}))
		var responseService databcore.ResponseService

		if err == nil {

			modelStruct := reflect.New(service.GetTypeModel()).Interface()
			err = json.Unmarshal(jsonBody, &modelStruct)

			if err == nil {
				responseService = (service).Edit(modelStruct, mapParamsService)
			} else {
				responseService.ResponseError = err
			}
		} else {
			responseService.ResponseError = err
		}

		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudAddRoute : Method for handle load add data
func HandleCrudAddRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/add", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {

		jsonBody, err := json.Marshal(requestBody.Data.(map[string]interface{}))
		var responseService databcore.ResponseService

		if err == nil {

			modelStruct := reflect.New(service.GetTypeModel()).Interface()
			err = json.Unmarshal(jsonBody, &modelStruct)

			if err == nil {
				responseService = (service).Add(modelStruct, mapParamsService)
			} else {
				responseService.ResponseError = err
			}
		} else {
			responseService.ResponseError = err
		}

		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudDeleteRoute : Method for handle load delete data
func HandleCrudDeleteRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/delete", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {

		jsonBody, err := json.Marshal(requestBody.Data.(map[string]interface{}))
		var responseService databcore.ResponseService

		if err == nil {

			modelStruct := reflect.New(service.GetTypeModel()).Interface()
			err = json.Unmarshal(jsonBody, &modelStruct)

			if err == nil {
				responseService = (service).Delete(modelStruct, mapParamsService)
			} else {
				responseService.ResponseError = err
			}
		} else {
			responseService.ResponseError = err
		}

		return responseService.Data, responseService.ResponseError
	})
}

// HandleCrudRoutes : Method for handle all crud routes. Count, List, Browser, Read, Edit, Add, Delete, LoadNew
func HandleCrudRoutes(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandleCrudBrowserRoute(rfHTTP, pathRoute, keyService)
	HandleCrudListRoute(rfHTTP, pathRoute, keyService)
	HandleCrudCountRoute(rfHTTP, pathRoute, keyService)
	HandleCrudAddRoute(rfHTTP, pathRoute, keyService)
	HandleCrudDeleteRoute(rfHTTP, pathRoute, keyService)
	HandleCrudEditRoute(rfHTTP, pathRoute, keyService)
	HandleCrudLoadNewRoute(rfHTTP, pathRoute, keyService)
	HandleCrudReadRoute(rfHTTP, pathRoute, keyService)
}

// StartTransactionContext : method for start transaction context
func StartTransactionContext(rfHTTP *rfhttp.RFHttp, mapParamsService *map[string]interface{}, req *http.Request) {
	// Transaction type gorm
	if rfHTTP.TransactionTypeContext == rfgodataconst.TransactionGorm {
		var db *gorm.DB = rfHTTP.DbConnection.(*gorm.DB)
		var transaction *gorm.DB = db.WithContext(req.Context()).Begin()
		var transactionGorm interface{} = transactiongorm.TransactionGorm{Transaction: transaction}
		var iTransaction transactions.ITransaction = transactionGorm.(transactions.ITransaction)
		(*mapParamsService)[rfgodataconst.ParamTransaction] = iTransaction
	}
}

// FinishTransactionContext : method for finish transaction context
func FinishTransactionContext(err error, rfHTTP *rfhttp.RFHttp, mapParamsService *map[string]interface{}) error {
	var returnError error
	var transactionError error
	// Transaction type gorm
	if rfHTTP.TransactionTypeContext == rfgodataconst.TransactionGorm {
		transaction, transactionError := datautils.GetTransactionInParams(mapParamsService)

		if transactionError == nil {
			transactionError = transaction.FinishTransaction(err)
		}

	}
	if err != nil {
		returnError = err
	} else {
		returnError = transactionError
	}
	return returnError
}

// ForcerFinishRequestRespose : method for force response error request
func ForcerFinishRequestRespose(res http.ResponseWriter, data interface{}, err error, rfHTTP *rfhttp.RFHttp, mapParamsService *map[string]interface{}) {

	err = FinishTransactionContext(err, rfHTTP, mapParamsService)

	if err != nil {
		var responseError *beans.ResponseError = beans.NewResponseError()
		// Send error to logger
		logger.Error(err.Error())

		// Send error
		fmt.Println(err.Error())
		responseError.Message = err.Error()
		jsonResult, errorJSON := json.Marshal(responseError)

		if errorJSON != nil {
			http.Error(res, errorJSON.Error(), http.StatusInternalServerError)
		} else {
			http.Error(res, string(jsonResult), http.StatusInternalServerError)
		}

	} else {
		var response *beans.RestRequestResponse = beans.NewRestRequestResponse()
		// send response
		response.Data = data
		utils.StatusOkInResponseRequest(response)
		utils.EncodeJsonDataResponseWriter(res, *response)
	}
}
