package main

import (
	cp "cpanelgo/cpanel"
	"testing"
)

//Whm username
var username = ""

//Whm password
var password = ""

//Whm hostname address
var host = ""

func Auth(t *testing.T) *cp.CP {
	client, err := cp.New(username, host, password)
	if err != nil {
		t.Error(err)
	}
	return &client
}
func TestListAccount(t *testing.T) {
	client := Auth(t)
	_, err := client.ListAccounts()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateAccount(t *testing.T) {
	client := Auth(t)
	_, err := client.CreateAccount("allstats.in", "allstats", "paperplanes123")
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveAccount(t *testing.T) {
	client := Auth(t)
	_, err := client.RemoveAccount("allstats")
	if err != nil {
		t.Error(err)
	}
}
