package cpanelgo

import (
	"github.com/tidwall/gjson"
)

type Maps map[string]string

func MapsJoin(m1, m2 Maps) Maps {
	for ia, va := range m1 {
		if it, ok := m2[ia]; ok {
			va += it
		}
		m2[ia] = va

	}
	return m2
}

func (cp *CP) Cpanel(module string, function string, username string, params ...map[string]string) ([]gjson.Result, error) {
	values := map[string]string{
		"cpanel_jsonapi_version": "2",
		"cpanel_jsonapi_module":  module,
		"cpanel_jsonapi_func":    function,
		"cpanel_jsonapi_user":    username,
	}
	if params != nil {
		values = MapsJoin(params[0], values)
	}

	data, err := cp.runQuery("cpanel", values)
	if err != nil {
		return nil, err
	}
	value := gjson.Get(string(data), "cpanelresult.data")
	return value.Array(), nil
}
