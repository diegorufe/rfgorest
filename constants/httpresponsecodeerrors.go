package constants

// CodeErrorResponseType : code error for response
type CodeErrorResponseType int

const (
	// CodeErrorMethodRequest : error method request
	CodeErrorMethodRequest CodeErrorResponseType = 0xE000000001

	// CodeErrorMarshalResponseWriter : error unable convert response request to json
	CodeErrorMarshalResponseWriter CodeErrorResponseType = 0xE000000002
)
