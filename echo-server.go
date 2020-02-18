package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const ADDR = "0.0.0.0:9000"

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	// Bind to TCP port 8887 on all interfaces.
	listener, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on " + ADDR)
	// Wait for connection.
	for {
		// Create net.Conn on connection established.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Log at echo server level
		fmt.Println("Someone has connected to echo server.")

		// Handle the connection
		go echo(conn)
	}
}
