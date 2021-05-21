package rfhttp

import (
	"reflect"
	"rfgorest/rfhttp"
	"testing"
)

func TestRFHttpContextName(t *testing.T) {

	var data *rfhttp.RFHttp = rfhttp.NewRFHttp()
	data.Properties.AppName = "TEST"
	var desireResult string = "TEST"

	if !reflect.DeepEqual(data.AppName(), desireResult) {
		t.Errorf("TestRFHttpContextName %s == %s", data.AppName(), desireResult)
	}

}
