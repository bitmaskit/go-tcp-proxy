package echo

import (
	"io"
	"log"
	"net"
)

const ADDR = "0.0.0.0:7" // per specification

type Server struct {
	Listener net.Listener
	Log      *log.Logger
}

func New(log *log.Logger) (*Server, error) {
	listener, err := net.Listen("tcp", ADDR)
	if err != nil {
		return nil, err
	}
	echoServer := Server{Listener: listener, Log: log}
	return &echoServer, nil
}

func (s *Server) Echo(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			s.Log.Println("error while closing connection", err)
		}
		s.Log.Println("disconnecting: ", conn.RemoteAddr().String())
	}(conn)

	if _, err := io.Copy(conn, conn); err != nil {
		s.Log.Println("Unable to read/write data")
	}
}

func (s *Server) HandleConnections() error {
	for {
		// Create net.Conn on connection established.
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Println("Unable to accept connection: ")
			return err
		}
		// Log at echo server level
		s.Log.Println("Someone has connected to echo server.", conn.RemoteAddr().String())

		// Handle the connection
		go s.Echo(conn)
	}
}
