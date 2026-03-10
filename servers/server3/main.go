package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-load-balancing/grpc-load-balancing/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Println("Request handled by Server 3")

	return &pb.HelloResponse{
		Message: "Hello " + req.Name + " from Server 3",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, &server{})

	fmt.Println("Server 3 running on port 50053")

	grpcServer.Serve(lis)
}
