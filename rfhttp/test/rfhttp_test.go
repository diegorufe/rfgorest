package rfhttp

import (
	"reflect"
	"rfgorest/constants"
	"rfgorest/rfhttp"
	"testing"
)

func TestRFHttpContextName(t *testing.T) {

	var mapProperties map[string]interface{} = make(map[string]interface{})
	mapProperties[constants.ParamAppName] = "TEST"
	var data *rfhttp.RFHttp = rfhttp.NewRFHttp(mapProperties)
	var desireResult string = "TEST"

	if !reflect.DeepEqual(data.AppName(), desireResult) {
		t.Errorf("TestRFHttpContextName %s == %s", data, desireResult)
	}

}
