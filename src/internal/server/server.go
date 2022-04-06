package server

import (
	"context"
	"log"
	"time"

	api "github.com/jasonzou/ezgRPC/src/api/v1"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// var _ api1.Entry = (*grpcServer)(nil)

type grpcServer struct {
	api.UnimplementedEntryServicesServer
	Entries *Entries
}

func (s *grpcServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.Entry, error) {
	resp, err := s.Entries.Retrieve(int(req.Id))
	if err == ErrIDNotFound {
		return nil, status.Error(codes.NotFound, "id was not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *grpcServer) Insert(ctx context.Context, entry *api.Entry) (*api.InsertResponse, error) {
	id, err := s.Entries.Insert(entry)
	if err != nil {
		log.Printf("Error:%s", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	res := api.InsertResponse{Id: int32(id)}
	return &res, nil
}

func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Entries, error) {
	entires, err := s.Entries.List(int(req.Offset))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.Entries{Entries: entires}, nil
}

func NewGRPCServer() *grpc.Server {
	var acc *Entries
	var err error
	if acc, err = NewEntries(); err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	srv := grpcServer{
		Entries: acc,
	}
	api.RegisterEntryServicesServer(gsrv, &srv)
	return gsrv
}

type Entry struct {
	Time        time.Time
	Description string
	ID          int
}
