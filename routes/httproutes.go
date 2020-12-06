package routes

import (
	"fmt"
	"net/http"
	"rfgocore/utils/utilsstring"
	rfgodataconst "rfgodata/constants"
	"rfgodata/service"
	"rfgodata/transactions"
	transactiongorm "rfgodata/transactions/gorm"
	datautils "rfgodata/utils"
	"rfgorest/beans"
	"rfgorest/constants"
	"rfgorest/logger"
	"rfgorest/rfhttp"
	"rfgorest/utils"

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

// HandleCrudListRoute : Method for handle list route
func HandleCrudListRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/list", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		return (service).List(requestBody.Fields, requestBody.Filters, requestBody.Joins, requestBody.Orders, nil, requestBody.Limit, mapParamsService)
	})
}

// HandleCrudCountRoute : Method for handle count route
func HandleCrudCountRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/count", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		return (service).Count(requestBody.Filters, requestBody.Joins, nil, mapParamsService)
	})
}

// HandleCrudLoadNewRoute : Method for handle load new route
func HandleCrudLoadNewRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/loadNew", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		return (service).LoadNew(mapParamsService)
	})
}

// HandleCrudReadRoute : Method for handle load read data
func HandleCrudReadRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	HandlePostRouteWithTransaction(rfHTTP, pathRoute+"/read", keyService, true, func(service service.IService, mapParamsService *map[string]interface{}, requestBody beans.RestRequestBody) (interface{}, error) {
		return (service).Read(requestBody.Data, mapParamsService)
	})
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

	var response *beans.RestRequestResponse = beans.NewRestRequestResponse()

	if err != nil {
		// Send error to logger
		logger.Error(err.Error())

		// Send error
		fmt.Println(err.Error())
		utils.StatusKoInResponseRequest(response)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	} else {
		// send response
		response.Data = data
		utils.StatusOkInResponseRequest(response)
		utils.EncodeJsonDataResponseWriter(res, *response)
	}
}
