package main

import (
	"flag"
	"fmt"
	"github.com/go-grpc-example/pb"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	ggp "github.com/grpc-ecosystem/go-grpc-prometheus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8080, "port")
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("tcp listen address:%s error:%s", address, err)
	}

	echoServer := newEchoServer()
	s := grpc.NewServer(grpc.StatsHandler(NewStateHandler()),
		grpc.StreamInterceptor(middleware.ChainStreamServer(ggp.StreamServerInterceptor)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(ggp.UnaryServerInterceptor)))

	pb.RegisterEchoServiceServer(s, echoServer)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("server err:%s", err)
	}
}
