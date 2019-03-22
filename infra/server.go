package infra

import (
	"os"
	"strconv"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// AppPort gets the port from certain env variable
// if it is set, or use the provided port otherwise.
func AppPort(port int) int {
	if v, ok := os.LookupEnv(AppPortEnv); ok {
		if p, err := strconv.Atoi(v); err != nil {
			port = p
		}
	}
	return port
}

type serverOption func(*grpc.Server)

// WithHealth registers a health server.
func WithHealth(srv *grpc.Server) {
	healthpb.RegisterHealthServer(srv, health.NewServer())
}

// WithReflection turns on the server's
// reflection protocol.
func WithReflection(srv *grpc.Server) {
	reflection.Register(srv)
}

// NewServer returns a grpc server with
// opencensus statHandler.
func NewServer(opts ...serverOption) *grpc.Server {
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}
