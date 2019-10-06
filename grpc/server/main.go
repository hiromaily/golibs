package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
	"github.com/pkg/errors"
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
func (s *server) ServerStreamingRespondManytimes(req *samplepb.ManyClients, stream samplepb.SampleService_ServerStreamingRespondManytimesServer) error {
	log.Println("[ServerStreamingRespondManytimes]")

	if len(req.GetClients()) == 0 {
		//TODO: error something
		return errors.New("clients has nothing")
	}

	// create multiple answer
	for i, client := range req.GetClients() {
		//log.Printf("[ServerStreamingAskManytimes] Received: name: %s, question code: %d", client.Name, client.QuestionCode)
		log.Printf("[ServerStreamingAskManytimes] Received: name: %s, question code: %d", client.GetName(), client.GetQuestionCode())

		answer := &samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("[%d]Hi %s", i, client.GetName()),
		}
		// send
		if err := stream.Send(answer); err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func (s *server) ClientStreamingAskManytimes(stream samplepb.SampleService_ClientStreamingAskManytimesServer) error {
	log.Println("[ClientStreamingAskManytimes]")

	answers := make([]*samplepb.Answer, 0)

	var idx uint64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&samplepb.ManyAnswers{Answers: answers})
		} else if err != nil {
			log.Println("err: ", err)
			return err
		}

		answer := &samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("[%d]Hi %s", idx, req.GetName()),
		}
		answers = append(answers, answer)

		idx++
	}
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
	log.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
