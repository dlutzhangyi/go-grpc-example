package main

import (
	"github.com/go-grpc-example/pb"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/stats"
	"sync"
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

type StateHandler struct {
	connCount int
	sync.RWMutex
}

func NewStateHandler() *StateHandler {
	h := &StateHandler{}
	return h
}

func (s *StateHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return context.WithValue(ctx, "RPCTagInfo", info)
}

func (s *StateHandler) HandleRPC(context.Context, stats.RPCStats) {
}

func (s *StateHandler) TagConn(ctx context.Context, cti *stats.ConnTagInfo) context.Context {
	return context.WithValue(ctx, "ConnTagInfo", cti)
}

func (s *StateHandler) HandleConn(ctx context.Context, cs stats.ConnStats) {
	s.Lock()
	defer s.Unlock()

	switch cs.(type) {
	case *stats.ConnBegin:
		s.connCount++
		log.Infof("create a new conn,current conn count:%d", s.connCount)
	case *stats.ConnEnd:
		s.connCount--
		log.Infof("end a conn,current conn count:%d", s.connCount)
	default:
		log.Infof("can't not find the conn,current conn count:%d", s.connCount)
	}
}
