package cpanelgo

import (
	"github.com/tidwall/gjson"
)

func (cp *CP) ListAccounts() ([]gjson.Result, error) {

	data, err := cp.runQuery("listaccts", nil)
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(data), "acct")
	return value.Array(), nil
}

func (cp *CP) ListUsers() ([]gjson.Result, error) {

	data, err := cp.runQuery("list_users", nil)
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(data), "data.users")
	return value.Array(), nil
}

func (cp *CP) GetDomainInfo() ([]gjson.Result, error) {

	data, err := cp.runQuery("get_domain_info", nil)
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(data), "data.domains")
	return value.Array(), nil
}

func (cp *CP) CreateAccount(domain string, username string, password string) ([]gjson.Result, error) {

	data, err := cp.runQuery("createacct", map[string]string{
		"username": username,
		"domain":   domain,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(data), "data")
	return value.Array(), nil
}

func (cp *CP) RemoveAccount(username string) ([]gjson.Result, error) {
	data, err := cp.runQuery("removeacct", map[string]string{
		"username": username,
	})
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(data), "data")
	return value.Array(), nil
}
