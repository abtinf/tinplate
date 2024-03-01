package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) StartupProbe(ctx context.Context, r *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) LivenessProbe(ctx context.Context, r *emptypb.Empty) (*emptypb.Empty, error) {
	if !s.live.Load() {
		return nil, status.Error(codes.Unavailable, codes.Unavailable.String())
	}
	return &emptypb.Empty{}, nil
}

func (s *server) ReadinessProbe(ctx context.Context, r *emptypb.Empty) (*emptypb.Empty, error) {
	if !s.ready.Load() {
		return nil, status.Error(codes.Unavailable, codes.Unavailable.String())
	}
	return &emptypb.Empty{}, nil
}
