package beans

// RFHttpProperties properties
type RFHttpProperties struct {
	AppName   string `default:"RFHttp"`
	Host      string `default:"localhost"`
	Port      int    `default:7000`
	MapParams map[string]interface{}
	// Services for RFHttp
	MapServices map[string]interface{}
}
