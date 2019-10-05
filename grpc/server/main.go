package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

//TODO: check `SampleServiceServer interface` in pb.go

func (s *server) UnaryAsk(ctx context.Context, in *samplepb.Client) (*samplepb.Answer, error) {
	log.Printf("[UnaryAsk] Received: name: %s, question code: %d", in.Name, in.QuestionCode)

	return &samplepb.Answer{
		Code:   200,
		Answer: fmt.Sprintf("Hi %s, ", in.Name),
	}, nil
}

func (s *server) ServerStreamingAskManytimes(req *samplepb.ManyClients, stream samplepb.SampleService_ServerStreamingAskManytimesServer) error {
	log.Printf("[ServerStreamingAskManytimes] Received: name: %s, question code: %d", req.Client.Name, req.Client.QuestionCode)
	log.Printf("[ServerStreamingAskManytimes] Received: name: %s, question code: %d", req.GetClient().GetName(), req.GetClient().GetQuestionCode())

	// create response
	for i := 0; i < 10; i++ {
		answer := &samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("[%d]Hi %s", i, req.GetClient().GetName()),
		}
		answers := &samplepb.ManyAnswers{
			Result: answer,
		}
		// send
		if err := stream.Send(answers); err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}

func main() {
	//
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//register services
	s := grpc.NewServer()
	samplepb.RegisterSampleServiceServer(s, &server{})

	//serve
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
