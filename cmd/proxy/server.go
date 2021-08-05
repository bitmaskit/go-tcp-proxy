package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const (
	TARGET    = "0.0.0.0:9000"
	THIS_PORT = ":80"
)

func handle(src net.Conn) {
	defer src.Close()
	dst, err := net.DialTimeout("tcp", TARGET, 10*time.Second)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	fmt.Println("We connected to the echo server")
	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("We have disconnected from the echo server.")
}

func main() {
	// Listen on local 80
	listener, err := net.Listen("tcp", THIS_PORT)
	fmt.Printf("Listening on %s\n", THIS_PORT)
	if err != nil {
		log.Fatalln("Unable to bind to port.")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Log info at Proxy level
		fmt.Println("Someone has connected to the proxy.")

		// TODO: Disconnect from echo server when client disconnects
		// Send message to the client
		conn.Write([]byte("You have connected to the proxy.\n"))
		go handle(conn)
	}
}
