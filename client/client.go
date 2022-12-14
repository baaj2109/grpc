package main

import (
	"flag"

	_ "github.com/baaj2109/grpc/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊阜編號")
	flag.Parse()
}
