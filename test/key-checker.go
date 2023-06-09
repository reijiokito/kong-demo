package main

import "github.com/Kong/go-pdk"

// it represents to config parameters into the config.yml
type Config struct {
	Apikey string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(pdk *pdk.PDK) {
	key, err := pdk.Request.GetQueryArg("key")
	apiKey := conf.Apikey

	if err != nil {
		pdk.Log.Err(err.Error())
	}

	//it adjusts the header parameters in this way.
	x := make(map[string][]string)
	x["Content-Type"] = append(x["Content-Type"], "application/json")

	//If the key of the consumer is not equal to the claimed key, test doesn't ensure the proxy
	if apiKey != key {
		pdk.Response.Exit(403, "You have no correct consumer key.", x)
	}
}
