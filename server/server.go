package main

import (
	"github.com/go-grpc-example/pb"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type EchoServer struct {
}

func (s *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Infof("get request: %s", req.Request)
	resp := &pb.EchoResponse{
		Response: req.Request,
	}
	return resp, nil
}

func newEchoServer() *EchoServer {
	return &EchoServer{}
}
