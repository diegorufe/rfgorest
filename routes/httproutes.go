package routes

import (
	"net/http"
	"rfgocore/utils/utilsstring"
	"rfgodata/beans/query"
	rfgodataconst "rfgodata/constants"
	"rfgodata/service"
	"rfgodata/transactions"
	transactiongorm "rfgodata/transactions/gorm"
	"rfgorest/beans"
	"rfgorest/constants"
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
			data, err := (service).List(nil, nil, nil, nil, nil, query.Limit{0, 1}, &mapParamsService)

			if err != nil {
				// Send error
				http.Error(res, err.Error(), http.StatusInternalServerError)
			} else {
				// send response
				var response *beans.RestRequestResponse = beans.NewRestRequestResponse()
				response.Data = data
				utils.StatusOkInResponseRequest(response)
				utils.EncodeJsonDataResponseWriter(res, *response)
			}

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
		var transaction *gorm.DB = db.WithContext(req.Context())
		var transactionGorm interface{} = transactiongorm.TransactionGorm{Transaction: transaction}
		var iTransaction transactions.ITransaction = transactionGorm.(transactions.ITransaction)
		(*mapParamsService)[rfgodataconst.ParamTransaction] = iTransaction
	}
}
