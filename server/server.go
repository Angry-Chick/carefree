package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/carefree/server/rpc"
)

// Server serves both RPC and HTTP.
type Server struct {
	rpc   *rpc.Server
	http  *http.Server
	errCh chan error // errCh holds error occurred while serving.
}

// New creates an initialized Server instance.
func New() *Server {
	return &Server{
		rpc:   rpc.NewServer(rpc.DefaultServeMux),
		http:  &http.Server{},
		errCh: make(chan error, 2),
	}
}

// Serve starts HTTP server on hport, gRPC server on rport, and awaits
// shutdown signal.
func (s *Server) Serve(hport, rport int) error {
	if err := s.Start(hport, rport); err != nil {
		return err
	}
	return s.Wait()
}

// Start starts HTTP server on localhost:hport, and gRPC server on
// localhost:rport.
func (s *Server) Start(hport, rport int) (err error) {
	hln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", hport))
	if err != nil {
		return err
	}
	rln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", rport))
	if err != nil {
		return err
	}

	log.Printf("HTTP server listen on %v", hln.Addr())
	log.Printf("gRPC server listen on %v", rln.Addr())

	go s.startHTTPServer(hln)
	go s.startRPCServer(rln)

	return nil
}

func (s *Server) startRPCServer(ln net.Listener) {
	s.errCh <- s.rpc.Serve(ln)
}

func (s *Server) startHTTPServer(ln net.Listener) {
	s.errCh <- s.http.Serve(ln)
}

// Wait awaits shutdown signal.
func (s *Server) Wait() error {
	return <-s.errCh
}

// Serve starts HTTP server on hport, gRPC server on rport.
func Serve(hport, rport int) error {
	return New().Serve(hport, rport)
}
