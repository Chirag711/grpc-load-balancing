package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "grpc-load-balancing/grpc-load-balancing/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	servers := []string{
		"localhost:50051",
		"localhost:50052",
		"localhost:50053",
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {

		server := servers[rand.Intn(len(servers))]

		conn, err := grpc.Dial(
			server,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		if err != nil {
			log.Fatal(err)
		}

		client := pb.NewGreetingServiceClient(conn)

		res, err := client.SayHello(context.Background(), &pb.HelloRequest{
			Name: "Chirag",
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Message)

		conn.Close()

		time.Sleep(1 * time.Second)
	}
}
