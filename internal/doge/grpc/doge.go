package grpc

import (
	"context"
	"my-go/internal/doge/store"
	"my-go/pkg/protobuf/doge"
	"my-go/pkg/protobuf/uuid"

	pb "my-go/pkg/protobuf/doge/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateDoge(ctx context.Context, req *pb.CreateDogeRequest) (*doge.Doge, error) {
	dogeType, err := doge.FromDogeTypeProto(&req.Type)
	if err != nil {
		return nil, err
	}

	result, err := s.Api.CreateDoge(store.CreateDogeParams{
		Name: req.Name, MagicNumber: uint32(req.MagicNumber), Type: *dogeType,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create Doge: %v", err)
	}

	return doge.ToDogeProto(result), nil
}

func (s *Server) GetDoge(ctx context.Context, id *uuid.Uuid) (*doge.Doge, error) {
	result, err := s.Api.GetDoge(uuid.FromUUIDProto(id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to get Doge: %v", err)
	}

	return doge.ToDogeProto(result), nil
}
