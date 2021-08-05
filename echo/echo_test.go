package echo

import (
	"bytes"
	"log"
	"net"
	"os"
	"testing"
)

var logger = log.New(os.Stdout, "ECHO Server: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
var server *Server

func init() {
	var err error
	server, err = New(logger)
	if err != nil {
		logger.Println("Could not start TCP Server")
		return
	}

	// Wait for connection in go routine, so we can perform test.
	go func() {
		err := server.HandleConnections()
		if err != nil {
			logger.Println("could not handle connections")
			return
		}
	}()
}
func TestEchoServer_OneConnection(t *testing.T) {
	tt := []struct {
		test string
		send []byte
		want []byte
	}{{
		"Sending a simple request: Hello",
		[]byte("hello world\n"),
		[]byte("hello world\n"),
	}, {
		"Sending a simple request: Goodbye",
		[]byte("goodbye world\n"),
		[]byte("goodbye world\n"),
	}}

	// Test 1 connection, 2 payloads
	conn, err := net.Dial("tcp", ":7")
	// create connection to our tcp server
	if err != nil {
		t.Error("could not connect to TCP server: ", err)
	}
	defer conn.Close()

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			if _, err := conn.Write(tc.send); err != nil {
				t.Error("could not write send to TCP server:", err)
			}
			got := make([]byte, len(tc.send)) //

			if _, err := conn.Read(got); err != nil {
				t.Error("could not read from connection")
			}

			if bytes.Compare(got, tc.want) != 0 {
				t.Error("response did not match expected output")
			}
		})
	}
}
func TestEchoServer_TwoConnections(t *testing.T) {
	tt := []struct {
		test string
		send []byte
		want []byte
	}{{
		"Sending a simple request returns result",
		[]byte("hello world\n"),
		[]byte("hello world\n"),
	}, {
		"Sending another simple request works",
		[]byte("goodbye world\n"),
		[]byte("goodbye world\n"),
	}}

	// Test 2 connection, 1 payload
	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			// create connection to our tcp server
			conn, err := net.Dial("tcp", ":7")
			if err != nil {
				t.Error("could not connect to TCP server: ", err)
			}

			if _, err := conn.Write(tc.send); err != nil {
				t.Error("could not write send to TCP server:", err)
			}

			got := make([]byte, len(tc.send)) //

			if _, err := conn.Read(got); err != nil {
				t.Error("could not read from connection")
			}

			if bytes.Compare(got, tc.want) != 0 {
				t.Error("response did not match expected output")
			}
			// Close connection before next test
			conn.Close()
		})
	}
}
