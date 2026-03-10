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

	fmt.Println("Request handled by Server 2")

	return &pb.HelloResponse{
		Message: "Hello " + req.Name + " from Server 2",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetingServiceServer(grpcServer, &server{})

	fmt.Println("Server 2 running on port 50052")

	grpcServer.Serve(lis)
}
