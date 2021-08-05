package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/bitmaskit/go-tcp-proxy/echo"
	"log"
	"os"
)

var (
	//go:embed help.txt
	helpMessage string
)

const (
	VERSION = "1.0"
)

func main() {
	// App version and help
	var version, help bool
	flag.BoolVar(&version, "version", false, "display version number")
	flag.BoolVar(&help, "help", false, "display help")
	flag.Parse()
	var firstArg string
	if len(os.Args) > 1 {
		firstArg = os.Args[1]
	}

	if version || firstArg == "version" {
		fmt.Println("Echo Server version:", VERSION)
		os.Exit(0)
	} else if help || firstArg == "help" {
		fmt.Println(helpMessage, "\nVersion:", VERSION)
		os.Exit(0)
	}

	// Initialise logger
	logger := log.New(os.Stdout, "ECHO Server: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(logger); err != nil {
		logger.Println("Error while running the application: ", err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	server, err := echo.New(logger)
	if err != nil {
		return err
	}

	logger.Println("Listening on " + server.Listener.Addr().String())

	// Wait for connection.
	return server.HandleConnections()
}
