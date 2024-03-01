package server

import (
	"bytes"
	"context"
	"fmt"
	"io"

	pb "gonfoot/proto"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (*server) ExampleGet(ctx context.Context, r *pb.ExampleRequest) (*pb.ExampleReply, error) {
	reply := fmt.Sprintf("Hello %s", r.GetName())
	return &pb.ExampleReply{Message: reply}, nil
}

func (*server) ExamplePost(ctx context.Context, r *pb.ExampleRequest) (*pb.ExampleReply, error) {
	reply := fmt.Sprintf("Hello %s", r.GetName())
	return &pb.ExampleReply{Message: reply}, nil
}

func (s *server) Download(r *pb.ExampleRequest, stream pb.API_DownloadServer) error {
	file := bytes.NewReader([]byte("abcd"))
	filename := r.GetName()
	if filename == "" {
		filename = "default.csv"
	}

	if err := stream.SetHeader(metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))); err != nil {
		s.log.Error("error setting download header", "err", err)
		return status.Error(codes.Unavailable, codes.Unavailable.String())
	}
	buff := make([]byte, 1024)
	for {
		n, err := file.Read(buff)
		if err != nil && err != io.EOF {
			s.log.Error("error reading file", "err", err, "request", r)
			return status.Error(codes.Unavailable, codes.Unavailable.String())
		} else if err == io.EOF {
			break
		}
		msg := &httpbody.HttpBody{
			ContentType: "text/csv",
			Data:        buff[:n],
		}
		if err := stream.Send(msg); err != nil {
			s.log.Error("error sending chunk", "err", err, "request", r)
			return status.Error(codes.Unavailable, codes.Unavailable.String())
		}
	}
	return nil
}

func (s *server) GetMigrations(ctx context.Context, _ *emptypb.Empty) (*pb.MigrationList, error) {
	rows, err := s.db.Queries.ListAllMigrations(ctx)
	if err != nil {
		s.log.Error("error getting migrations", "err", err)
		return nil, status.Error(codes.Unavailable, codes.Unavailable.String())
	}
	result := &pb.MigrationList{
		Migrations: make([]*pb.Migration, 0),
	}
	for _, row := range rows {
		result.Migrations = append(result.Migrations, &pb.Migration{
			Id:        row.ID,
			Name:      row.Name,
			Query:     row.Query,
			CreatedAt: timestamppb.New(row.CreatedAt.Time),
		})
	}
	return result, nil
}
