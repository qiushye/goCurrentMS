package main

import (
	"flag"
	"log"
	"net"

	string_service "github.com/longjoy/micro-go-book/ch7-rpc/grpc/string-service"
	"github.com/longjoy/micro-go-book/ch7-rpc/pb"
	"google.golang.org/grpc"
)

func main1() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
