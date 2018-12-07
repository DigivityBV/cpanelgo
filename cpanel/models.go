package cpanelgo

import (
	"github.com/asaskevich/govalidator"
	"github.com/soheilhy/args"
)

type CP struct {
	Protocol    string `valid:"required~Protocol not found!"`
	WHMHostname string `valid:"required~Hostname not found!"`
	WHMPort     string `valid:"required~Port not found!,int"`
	WHMUsername string `valid:"required~Username not found!"`
	WHMAuthKey  string `valid:"required~Auth Key not found!"`
	Format      string `valid:"required~Format not found!"`
	APIVersion  string `valid:"-"`
}

var ssl = args.NewBool(args.Default(true))
var apiformat = args.NewString(args.Default("json-api"))

//Constructr for cpanel client
func New(username string, host string, passkey string, opts ...args.V) (CP, error) {
	//Protocol and port assignment
	protocol := "https"
	port := "2087"

	if !ssl.Get(opts) {
		protocol = "http"
		port = "2086"
	}

	//Api format selection
	format := apiformat.Get(opts)

	//Validate
	client := CP{
		Protocol:    protocol,
		WHMHostname: host,
		WHMPort:     port,
		WHMUsername: username,
		WHMAuthKey:  passkey,
		Format:      format,
		APIVersion:  "2",
	}
	if _, err := govalidator.ValidateStruct(client); err != nil {
		return CP{}, err
	}

	return client, nil
}
