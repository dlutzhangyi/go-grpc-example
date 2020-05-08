package main

import (
	"context"
	"flag"
	"github.com/go-grpc-example/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

var (
	address string
)

func init() {
	flag.StringVar(&address, "address", ":8080", "address to dial")
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial err:%s", err)
	}

	client := pb.NewEchoServiceClient(conn)
	ctx := context.Background()
	req := &pb.EchoRequest{
		Request: "hello",
	}
	for {
		resp, err := client.Echo(ctx, req)
		if err != nil {
			log.Fatalf("echo err:%s", err)
		}
		log.Infof("resp:%s", resp.Response)
		time.Sleep(2 * time.Second)
	}

}
