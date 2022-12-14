package main

import (
	"context"
	"flag"

	pb "github.com/baaj2109/grpc/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊阜編號")
	flag.Parse()
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) {

}
