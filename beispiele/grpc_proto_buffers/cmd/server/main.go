package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "protobuff-example/proto"
)

const port = ":50051"

// handler is used to implement pb.EchoServiceServer.
type handler struct {
	pb.UnimplementedEchoServiceServer
}

func (s *handler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	return &pb.HelloReply{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	//
	// Network connection.

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//
	// Create a server and register our instance.

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, &handler{})

	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("server listening on %s", port)
}
