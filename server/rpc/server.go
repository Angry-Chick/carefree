package rpc

import (
	"net"

	"google.golang.org/grpc"
)

// DefaultServeMux is the default ServeMux.
var DefaultServeMux = NewServeMux()

// Handle registers a service to grpc server.
func Handle(s Service) {
	DefaultServeMux.services[s] = struct{}{}
}

// ServeMux is a RPC service registry. It allows one to register a gRPC service
// without creating a gRPC.Server first. This design matches the http package's
// approach on registering handler.
type ServeMux struct {
	services map[Service]struct{}
}

// NewServeMux returns an initialized ServeMux.
func NewServeMux() *ServeMux {
	return &ServeMux{services: make(map[Service]struct{})}
}

func (mux *ServeMux) registerAll(s *Server) {
	for svc := range mux.services {
		svc.Register(s)
	}
}

// Service defines the methods a rpc server implementation should have.
type Service interface {
	// Register registers the service implementation to a GRPC server.
	Register(*Server)
}

// Server is a RPC server.
type Server struct {
	GRPC *grpc.Server
}

// NewServer returns a new carefree rpc server.
func NewServer(mux *ServeMux, opts ...grpc.ServerOption) *Server {
	svr := &Server{GRPC: grpc.NewServer(opts...)}
	mux.registerAll(svr)
	return svr
}

// Serve serves RPC services on ln.
func (s *Server) Serve(lis net.Listener) error {
	return s.GRPC.Serve(lis)
}

// Stop stops all RPC connections immediately.
func (s *Server) Stop() {
	s.GRPC.Stop()
}

// GracefulStop stops RPC server gracefully. It blocks new incoming connections,
// and allows pending requests to be processed.
func (s *Server) GracefulStop() {
	s.GRPC.GracefulStop()
}
