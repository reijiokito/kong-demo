package main

import (
	"log"
	"net"
)

func main() {
	socketFile := "go_pluginserver.sock"

	conn, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Fatal("Unable to connect to socket file:", err)
		return
	}
	defer conn.Close()

	// Send data to the socket
	_, err = conn.Write([]byte("Hello, socket"))
	if err != nil {
		log.Fatal("Error writing to socket:", err)
		return
	}

	// Receive data from the socket
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("vfahhvhza  Error reading from socket:", err)
		return
	}

	log.Printf("Received %d bytes: %s", n, buf)
	log.Printf("")

}
