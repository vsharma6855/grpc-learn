package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/vsharma6855/grpc-learn/proto/greeting"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("gRPC server is running...")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
