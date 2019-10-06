package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
)

const (
	port = ":50051"
)

var (
	certFile = fmt.Sprintf("%s/src/github.com/hiromaily/golibs/grpc/key/server.crt", os.Getenv("GOPATH"))
	keyFile  = fmt.Sprintf("%s/src/github.com/hiromaily/golibs/grpc/key/server.pem", os.Getenv("GOPATH"))
)

type server struct{}

//TODO: check `SampleServiceServer interface` in pb.go

func (s *server) UnaryAsk(ctx context.Context, in *samplepb.Client) (*samplepb.Answer, error) {
	log.Printf("[UnaryAsk] Received: name: %s, question code: %d", in.Name, in.QuestionCode)

	//validate
	if in.QuestionCode >= 100 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid question_code: %v", in.QuestionCode))
	}
	if in.QuestionCode == 10 {
		//for deadline test
		time.Sleep(2000 * time.Millisecond)
	}

	return &samplepb.Answer{
		Code:   200,
		Answer: fmt.Sprintf("Hi %s, your question_id is %d", in.Name, in.QuestionCode),
	}, nil
}

func (s *server) ServerStreamingRespondManytimes(req *samplepb.ManyClients, stream samplepb.SampleService_ServerStreamingRespondManytimesServer) error {
	log.Println("[ServerStreamingRespondManytimes]")

	//validate
	if len(req.GetClients()) == 0 {
		return status.Errorf(codes.InvalidArgument, "client has nothing")
	}

	// create multiple answer
	for i, client := range req.GetClients() {
		//log.Printf("[ServerStreamingAskManytimes] Received: name: %s, question code: %d", client.Name, client.QuestionCode)
		log.Printf("Received: name: %s, question code: %d", client.GetName(), client.GetQuestionCode())

		answer := &samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("[%d]Hi %s", i, client.GetName()),
		}
		// send
		log.Println("send to client")
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
			log.Println("server is closed after sending")
			return stream.SendAndClose(&samplepb.ManyAnswers{Answers: answers})
		} else if err != nil {
			log.Println("fail to call Recv(): ", err)
			return err
		}

		answer := &samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("Hi %s, your question_id is %d", req.GetName(), req.GetQuestionCode()),
		}
		answers = append(answers, answer)

		idx++
	}
}

func (s *server) BidirectionalStreaming(stream samplepb.SampleService_BidirectionalStreamingServer) error {
	log.Println("[BidirectionalStreaming]")

	var idx uint64
	for {
		//receive
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Println("fail to call Recv(): ", err)
			return err
		}

		//send
		err = stream.Send(&samplepb.Answer{
			Code:   200,
			Answer: fmt.Sprintf("[%d]Hi %s", idx, req.GetName()),
		})
		if err != nil {
			log.Println("fail to call Send(): ", err)
			return err
		}

		idx++
	}
}

func main() {
	//
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("fail to call credentials.NewServerTLSFromFile() %v", err)
	}

	//register services
	//s := grpc.NewServer()
	s := grpc.NewServer(grpc.Creds(creds))
	samplepb.RegisterSampleServiceServer(s, &server{})

	//serve
	log.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
