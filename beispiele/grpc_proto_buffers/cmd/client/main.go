package main

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	pb "protobuff-example/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoServiceClient(conn)

	request := &pb.HelloRequest{Name: "client", Id: 0x_ff00_00}

	protoRequest, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("failed to marshal request: %v", err)
	}
	fmt.Println(protoRequest)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("failed to marshal request: %v", err)
	}
	fmt.Println(jsonRequest)

	reply, err := c.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", reply.Message)
}
