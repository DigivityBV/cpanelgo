package cpanelgo

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/smallnest/goreq"
	"github.com/tidwall/gjson"
)

func (cp *CP) runQuery(function string, paramaters ...map[string]string) ([]byte, error) {

	params := ""
	if paramaters != nil {
		for k, v := range paramaters[0] {
			params += fmt.Sprintf("%s=%s&", k, v)
		}
	}
	url := fmt.Sprintf("%s://%s:%s/%s/%s?%s", cp.Protocol, cp.WHMHostname, cp.WHMPort, cp.Format, function, params)
	fmt.Println(url)
	request := goreq.New().SetBasicAuth(cp.WHMUsername, cp.WHMAuthKey)
	//Unsecure ssl bypass
	request = request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, _, _ := request.Get(url).End()
	body, _ := ioutil.ReadAll(resp.Body)

	value := gjson.Get(string(body), "status")
	if value.String() == "0" {
		return nil, errors.New("Status Code:0, Please check arguments.")
	}
	return body, nil
}
