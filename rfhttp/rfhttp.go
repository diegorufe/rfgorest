package rfhttp

import (
	"net/http"
	"rfgocore/utils/utilsstring"
	"rfgorest/constants"
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
	return rfHTTP.MapProperties[constants.ParamAppName].(string)
}

// HandleGetRoute : method for handler get route
func (rfHTTP *RFHttp) HandleGetRoute(route string, handler http.HandlerFunc) {
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
func (rfHTTP *RFHttp) HandlePostRoute(route string, handler http.HandlerFunc) {
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
}
