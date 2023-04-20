package main

import (
	"github.com/ugorji/go/codec"
	"log"
	"net/rpc"
)

func main() {

	socketFile := "/home/cong/Documents/bnm/demo/go-pluginserver/go_pluginserver.sock"

	conn, err := rpc.Dial("unix", socketFile)
	if err != nil {
		log.Fatal("Unable to connect to socket file:", err)
		return
	}
	defer conn.Close()

	// Send data to the socket
	var mh codec.MsgpackHandle
	mh.WriteExt = true

	// Create an encoder
	var data []byte
	enc := codec.NewEncoderBytes(&data, &mh)

	// Encode the value
	_ = enc.Encode("/path/dir")

	var reply error
	var val string
	err = conn.Call("plugin.SetPluginDir", []interface{}{"/path", &val}, reply)
	if err != nil {
		log.Fatal("Failed to call SetPluginDir:", err)
	}

	//log.Printf("Received %d bytes: %s", n, buf)

}

//// Send data to the socket
//var error error
//err = conn.Call("plugin.StartInstanc", server.PluginConfig{Config: []byte("config"), Name: "key-checker"}, &error)
//if err != nil {
//log.Fatal("Error writing to socket:", err)
//return
//}
