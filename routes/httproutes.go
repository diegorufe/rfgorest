package routes

import (
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

// HandleCrudListRoute : Method for handle list route
func HandleCrudListRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	http.HandleFunc(pathRoute+"/list", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodPost:

			// Serve the resource.

			var service service.IService = rfHTTP.GetService(keyService).(service.IService)
			var mapParamsService map[string]interface{} = make(map[string]interface{})

			// Start transaction context
			StartTransactionContext(rfHTTP, &mapParamsService, req)

			// Call list service
			var err error
			var data interface{}

			// Catch panic errors
			defer func() {
				if err := recover(); err != nil {
					ForcerFinishRequestRespose(res, data, err.(error), rfHTTP, &mapParamsService)
				}
			}()

			// get request body
			var requestBody beans.RestRequestBody
			requestBody, err = utils.EncodeRequestBody(req)

			if err != nil {
				panic(err)
			}

			// Call list service
			data, err = (service).List(requestBody.Fields, requestBody.Filters, requestBody.Joins, requestBody.Orders, nil, requestBody.Limit, &mapParamsService)

			// If has error finish transaction context
			defer ForcerFinishRequestRespose(res, data, err, rfHTTP, &mapParamsService)

			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
	})
}

// HandleCrudCountRoute : Method for handle count route
func HandleCrudCountRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	http.HandleFunc(pathRoute+"/count", func(res http.ResponseWriter, req *http.Request) {

		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodPost:

			// Serve the resource.
			var service service.IService = rfHTTP.GetService(keyService).(service.IService)
			var mapParamsService map[string]interface{} = make(map[string]interface{})

			// Start transaction context
			StartTransactionContext(rfHTTP, &mapParamsService, req)

			// Call list service
			var err error
			var data interface{}

			// Catch panic errors
			defer func() {
				if err := recover(); err != nil {
					ForcerFinishRequestRespose(res, data, err.(error), rfHTTP, &mapParamsService)
				}
			}()

			// Count
			data, err = (service).Count(nil, nil, nil, &mapParamsService)

			// If has error finish transaction context
			defer ForcerFinishRequestRespose(res, data, err, rfHTTP, &mapParamsService)

			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
	})
}

// HandleCrudLoadNewRoute : Method for handle load new route
func HandleCrudLoadNewRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	http.HandleFunc(pathRoute+"/loadNew", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodPost:

			// Serve the resource.

			var service service.IService = rfHTTP.GetService(keyService).(service.IService)
			var mapParamsService map[string]interface{} = make(map[string]interface{})

			// Start transaction context
			StartTransactionContext(rfHTTP, &mapParamsService, req)

			// Call list service
			var err error
			var data interface{}

			// Catch panic errors
			defer func() {
				if err := recover(); err != nil {
					ForcerFinishRequestRespose(res, data, err.(error), rfHTTP, &mapParamsService)
				}
			}()

			// Call load new service
			data, err = (service).LoadNew(&mapParamsService)

			// If has error finish transaction context
			defer ForcerFinishRequestRespose(res, data, err, rfHTTP, &mapParamsService)

			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
	})
}

// HandleCrudReadRoute : Method for handle load read data
func HandleCrudReadRoute(rfHTTP *rfhttp.RFHttp, pathRoute string, keyService string) {
	http.HandleFunc(pathRoute+"/read", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case http.MethodOptions:
			break

		case http.MethodPost:

			// Serve the resource.

			var service service.IService = rfHTTP.GetService(keyService).(service.IService)
			var mapParamsService map[string]interface{} = make(map[string]interface{})

			// Start transaction context
			StartTransactionContext(rfHTTP, &mapParamsService, req)

			// Call list service
			var err error
			var data interface{}

			// Catch panic errors
			defer func() {
				if err := recover(); err != nil {
					ForcerFinishRequestRespose(res, data, err.(error), rfHTTP, &mapParamsService)
				}
			}()

			// get request body
			var requestBody beans.RestRequestBody
			requestBody, err = utils.EncodeRequestBody(req)

			if err != nil {
				panic(err)
			}

			// Call read service
			data, err = (service).Read(requestBody.Data, &mapParamsService)

			// If has error finish transaction context
			defer ForcerFinishRequestRespose(res, data, err, rfHTTP, &mapParamsService)

			break

		default:
			// Give an error message.
			http.Error(res, utilsstring.IntToString(int(constants.CodeErrorMethodRequest)), http.StatusInternalServerError)
		}
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
		utils.StatusKoInResponseRequest(response)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	} else {
		// send response
		response.Data = data
		utils.StatusOkInResponseRequest(response)
		utils.EncodeJsonDataResponseWriter(res, *response)
	}
}
