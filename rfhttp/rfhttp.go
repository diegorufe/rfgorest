package rfhttp

import (
	"fmt"
	"net/http"
	"rfgocore/logger"
	"rfgocore/utils/utilsstring"
	rfgodataconst "rfgodata/constants"
	"rfgorest/rfhttp/beans"
)

// RFHttp : struct for store data for http
type RFHttp struct {
	// Properties RFHttp
	Properties beans.RFHttpProperties
	// Transaction type for database
	TransactionTypeContext rfgodataconst.TransactionType
	// DbConnection
	DbConnection interface{}
}

// NewRFHttp : method for create new RFHttp
func NewRFHttp() *RFHttp {
	var rfHTTP *RFHttp = new(RFHttp)

	rfHTTP.Properties.AppName = "RFHttp"
	rfHTTP.Properties.Host = "localhost"
	rfHTTP.Properties.Port = 7000

	// Init default properties
	initDefaultRFHttp(rfHTTP)

	// For default transaction type is Gorm
	rfHTTP.TransactionTypeContext = rfgodataconst.TransactionGorm

	return rfHTTP
}

// AppName : Method for get appName
func (rfHTTP *RFHttp) AppName() string {
	return rfHTTP.Properties.AppName
}

// AddService : method for add service by key
func (rfHTTP *RFHttp) AddService(keyService string, service interface{}) {
	if service != nil {
		rfHTTP.Properties.MapServices[keyService] = service
	}
}

// GetService : method for get service by key
func (rfHTTP *RFHttp) GetService(keyService string) interface{} {
	return rfHTTP.Properties.MapServices[keyService]
}

// Listen : method for start server on host and port
func (rfHTTP *RFHttp) Listen() error {

	var hostAndPort string = rfHTTP.Properties.Host +
		":" +
		utilsstring.IntToString(rfHTTP.Properties.Port)

	if logger.IsInfoEnabled() {
		logger.Info(fmt.Sprintf("Serving %s\n", hostAndPort))
	}

	return http.ListenAndServe(hostAndPort, nil)
}

// initDefaultRFHttp: method for initialice default data
func initDefaultRFHttp(rfHTTP *RFHttp) {
	// Init map properties
	rfHTTP.Properties.MapServices = make(map[string]interface{})
	rfHTTP.Properties.MapParams = make(map[string]interface{})
}
