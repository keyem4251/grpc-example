package main

import (
	"fmt"
	"grpc-example/pb"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":8080"
)

// server is used to implement EchoSVCServer
type server struct{}

// AuthFuncOverride is to handle authentication
func (s *server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if fullMethodName == "/pb.EchoSVC/EchoAuthSkip" {
		return ctx, nil
	}

	ctx, err := authenticate(ctx)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

// Echo is endpoint to demo
func (s *server) Echo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

// Echo is endpoint to demo
func (s *server) EchoAuthSkip(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoAuth not implemented")
}

func authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	fmt.Print(token)
	return ctx, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_auth.UnaryServerInterceptor(authenticate))))
	pb.RegisterEchoSVCServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
