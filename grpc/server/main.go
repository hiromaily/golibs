package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) UnaryEcho(ctx context.Context, in *samplepb.SampleBase) (*samplepb.SampleResponse, error) {
	log.Printf("Received: %v", in.SampleName)
	return &samplepb.SampleResponse{
		Code:   200,
		Answer: fmt.Sprintf("accept %s", in.SampleName),
	}, nil
}

//type SrvSampleServer interface {
//	// UnaryEcho is unary echo.
//	UnaryEcho(context.Context, *SampleBase) (*SampleResponse, error)
//	// ServerStreamingEcho is server side streaming.
//	ServerStreamingEcho(*SampleBase, SrvSample_ServerStreamingEchoServer) error
//	// ClientStreamingEcho is client side streaming.
//	ClientStreamingEcho(SrvSample_ClientStreamingEchoServer) error
//	// BidirectionalStreamingEcho is bidi streaming.
//	BidirectionalStreamingEcho(SrvSample_BidirectionalStreamingEchoServer) error
//}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	samplepb.RegisterSrvSampleServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
