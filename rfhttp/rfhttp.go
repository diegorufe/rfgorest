package rfhttp

import (
	"net/http"
	"rfgocore/utils/utilsstring"
	"rfgorest/constants/rfhttpparamsconstants"
)

// RFHttp : struct for store data for http
type RFHttp struct {
	// Map properties store in http
	MapProperties map[string]interface{}
}

// NewRFHttp : method for create new RFHttp
// mapProperties for store in http method
func NewRFHttp(mapProperties map[string]interface{}) *RFHttp {
	var rfHTTP *RFHttp = new(RFHttp)

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
	return rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamAppName].(string)
}

// HandleRoute : method for handler route function
func (rfHTTP *RFHttp) HandleRoute(route string, handler http.HandlerFunc) {
	http.HandleFunc(route, handler)
}

// Listen : method for start server on host and port
func (rfHTTP *RFHttp) Listen() {

	var hostAndPort string = rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamHost].(string) +
		":" +
		utilsstring.IntToString(rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamPort].(int))

	http.ListenAndServe(hostAndPort, nil)
}

// initDefaultRFHttp: method for initialice default data
func initDefaultRFHttp(rfHTTP *RFHttp) {
	// Init map properties
	rfHTTP.MapProperties = make(map[string]interface{})
	rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamAppName] = "RFHttp"
	rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamHost] = "localhost"
	rfHTTP.MapProperties[rfhttpparamsconstants.RFHttpParamPort] = 7000
}
