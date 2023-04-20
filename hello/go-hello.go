package main

/*
	A "hello world" plugin in Go,
	which reads a request header and sets a response header.
*/
import (
	"fmt"
	"github.com/Kong/go-pdk"
	"log"
)

var Version = "0.2"
var Priority = 1

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(pdk *pdk.PDK) {
	host, err := pdk.Request.GetHeader("host")
	if err != nil {
		log.Printf("Error reading 'host' header: %s", err.Error())
	}

	message := conf.Message
	if message == "" {
		message = "hello"
	}
	pdk.Response.SetHeader("x-hello-from-go", fmt.Sprintf("Go says %s to %s", message, host))
}
