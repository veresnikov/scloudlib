package server

import (
	"net"

	"github.com/pkg/errors"
	log "github.com/veresnikov/scloudlib/pkg/app/logger"
	"google.golang.org/grpc"
)

type GrpcServerConfig struct {
	ServeAddress string
}

func NewGrpcServer(server *grpc.Server, config GrpcServerConfig, logger log.Logger) Server {
	return &grpcServer{
		baseServer: server,
		config:     config,
		logger:     logger,
	}
}

type grpcServer struct {
	baseServer *grpc.Server
	config     GrpcServerConfig
	logger     log.Logger
}

func (g *grpcServer) Serve() error {
	grpcListener, grpcErr := net.Listen("tcp", g.config.ServeAddress)
	if grpcErr != nil {
		return errors.Wrapf(grpcErr, "failed to listen port %s", g.config.ServeAddress)
	}

	g.logger.Info("GRPC Server started")
	grpcErr = g.baseServer.Serve(grpcListener)
	return errors.Wrap(grpcErr, "failed to serve GRPC")
}

func (g *grpcServer) Stop() error {
	g.baseServer.GracefulStop()
	return nil
}
