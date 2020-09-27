package rfhttp

import (
	"net/http"
	"rfgocore/utils/utilsstring"
	rfgodataconst "rfgodata/constants"
	"rfgorest/constants"
)

// RFHttp : struct for store data for http
type RFHttp struct {
	// Map properties store in http
	MapProperties map[string]interface{}
	// Transaction type for database
	TransactionTypeContext rfgodataconst.TransactionType
	// DbConnection
	DbConnection interface{}
}

// NewRFHttp : method for create new RFHttp
// mapProperties for store in http method
func NewRFHttp(mapProperties map[string]interface{}) *RFHttp {
	var rfHTTP *RFHttp = new(RFHttp)

	// For default transaction type is Gorm
	rfHTTP.TransactionTypeContext = rfgodataconst.TransactionGorm

	// Init default data
	initDefaultRFHttp(rfHTTP)

	// Add map properties pass from user
	if mapProperties != nil {
		for key := range mapProperties {
			rfHTTP.MapProperties[key] = mapProperties[key]
		}
	}

	return rfHTTP
}

// AppName : Method for get appName
func (rfHTTP *RFHttp) AppName() string {
	return rfHTTP.MapProperties[constants.ParamAppName].(string)
}

// AddService : method for add service by key
func (rfHTTP *RFHttp) AddService(keyService string, service interface{}) {
	if service != nil {
		var mapServices map[string]interface{} = rfHTTP.MapProperties[constants.ParamMapServices].(map[string]interface{})
		mapServices[keyService] = service
	}
}

// GetService : method for get service by key
func (rfHTTP *RFHttp) GetService(keyService string) interface{} {
	var mapServices map[string]interface{} = rfHTTP.MapProperties[constants.ParamMapServices].(map[string]interface{})
	return mapServices[keyService]
}

// Listen : method for start server on host and port
func (rfHTTP *RFHttp) Listen() {

	var hostAndPort string = rfHTTP.MapProperties[constants.ParamHost].(string) +
		":" +
		utilsstring.IntToString(rfHTTP.MapProperties[constants.ParamPort].(int))

	http.ListenAndServe(hostAndPort, nil)
}

// initDefaultRFHttp: method for initialice default data
func initDefaultRFHttp(rfHTTP *RFHttp) {
	// Init map properties
	rfHTTP.MapProperties = make(map[string]interface{})
	rfHTTP.MapProperties[constants.ParamAppName] = "RFHttp"
	rfHTTP.MapProperties[constants.ParamHost] = "localhost"
	rfHTTP.MapProperties[constants.ParamPort] = 7000
	rfHTTP.MapProperties[constants.ParamMapServices] = make(map[string]interface{})
}
