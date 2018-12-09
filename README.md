## Cpanel & WHM API for GOLANG

## Contents
- [Installation Guide](#installation-guide)
- [Usage](#usage)
- [Feedback & Contribution](#feedback-&-contribution)

### Installation Guide

To install this package, you can run this code 
```
   go get https://github.com/DigivityBV/cpanelgo
```


### Usage
For example, if you would like to get list accounts of your whm server, 


Firstly auth the whm panel

```
   client, err := cp.New(username, host, password)
	if err != nil {
		panic(err)
	}
```

Then list accounts

```
   accounts, err := client.ListAccounts()
	if err != nil {
		panic(err)
	}

```

### Cpanel API Usage

Visit [CPANEL API](https://documentation.cpanel.net/display/DD/Guide+to+cPanel+API+2) To all modules and functions 

```
   client, err := cp.New(username, host, password)
	if err != nil {
		panic(err)
	}
	message, err := client.Cpanel(module, function, username)
	if err != nil {
		panic(err)
	}
```
#### Feedback and contribution
This package is free and open source, feel free to fork and report some issue to this package. :)


