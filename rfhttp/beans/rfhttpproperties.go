package beans

// RFHttpProperties properties
type RFHttpProperties struct {
	AppName   string
	Host      string
	Port      int
	MapParams map[string]interface{}
	// Services for RFHttp
	MapServices map[string]interface{}
}
