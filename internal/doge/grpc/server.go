package grpc

import (
	"my-go/internal/doge/api"
	pb "my-go/pkg/protobuf/doge/v1"

	"google.golang.org/grpc/health"
)

type Server struct {
	*health.Server
	pb.UnimplementedDogeServiceServer
	api.Api
}

func newServer(api api.Api) *Server {
	return &Server{
		Server: health.NewServer(),
		Api:    api,
	}
}
