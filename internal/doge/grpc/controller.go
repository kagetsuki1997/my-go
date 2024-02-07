package grpc

import (
	"context"
	"my-go/internal/doge/api"
	"my-go/internal/doge/config"
	pb "my-go/pkg/protobuf/doge/v1"
	"net"

	"github.com/chelpis/member-center/pkg/telemetry/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	healthPb "google.golang.org/grpc/health/grpc_health_v1"
)

const logLabel = "gRPCServer"

func Run(ctx context.Context, config *config.Config, api api.Api) <-chan error {
	ctx, logger := log.WithNamedLoggerContext(ctx, logLabel)

	server := newServer(api)

	grpcLogger := zapgrpc.NewLogger(logger)
	grpclog.SetLoggerV2(grpcLogger)

	s := grpc.NewServer(grpc.UnaryInterceptor(GrpcLogger(ctx)))
	pb.RegisterDogeServiceServer(s, server)
	healthPb.RegisterHealthServer(s, server)

	addr := config.GrpcServer.Addr.String()
	logger.Info("gRPC server started", zap.String("addr", addr))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil
	}

	errChan := make(chan error, 1)

	// start server and handle server error
	go func() {
		err := s.Serve(lis)
		if err != nil {
			errChan <- err
		}
	}()

	// handle stop server
	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return errChan
}
