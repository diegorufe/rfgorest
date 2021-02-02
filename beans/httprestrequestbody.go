package beans

import (
	"rfgodata/beans/query"
)

// RestRequestBody : for get request body
type RestRequestBody struct {
	Data      interface{}            `json:"data"`
	Limit     query.Limit            `json:"limit"`
	Fields    []query.Field          `json:"fields"`
	Joins     []query.Join           `json:"joins"`
	Filters   []query.Filter         `json:"filters"`
	Orders    []query.Order          `json:"orders"`
	MapParams map[string]interface{} `json:"MapParams"`
}
