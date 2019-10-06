package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/codes"
	"io"
	"log"
	"net"
	"time"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	address = "0.0.0.0:50051"
)

var (
	mode     = flag.Int("mode", 1, "mode")
	name     = flag.String("name", "", "name")
	question = flag.Int64("q", 0, "question code")
)

func init() {
	flag.Parse()
}

func validate() {
	if *name == "" {
		log.Fatal("-name parameter is required e.g. `-name Mike`")
	} else if *question == 0 {
		log.Fatal("-q parameter is required e.g. `-q 1`")
	}
}

func main() {
	validate()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to connect: %v", err)
	}
	defer conn.Close()
	cli := samplepb.NewSampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	//TODO: check `SampleServiceClient interface` in pb.go

	switch *mode {
	case 1:
		//Unary
		doUnary(ctx, cli)
	case 2:
		//ServerStreaming
		doServerStreaming(ctx, cli)
	case 3:
		//ClientStreaming
		doClientStreaming(ctx, cli)
	default:
		log.Fatal("-mode is out of range")
	}
}

func doUnary(ctx context.Context, cli samplepb.SampleServiceClient) {
	res, err := cli.UnaryAsk(ctx, &samplepb.Client{
		QuestionCode: *question,
		Name:         *name,
	})
	if err != nil {
		log.Fatalf("fail to call cli.UnaryAsk(): %v", err)
	}
	log.Printf("Response: [code] %d, [answer] %s", res.Code, res.Answer)
}

func doServerStreaming(ctx context.Context, cli samplepb.SampleServiceClient) {
	req := &samplepb.ManyClients{
		Clients: []*samplepb.Client{
			{
				QuestionCode: *question,
				Name:         *name,
			},
			{
				QuestionCode: *question + 1,
				Name:         *name + "2",
			},
			{
				QuestionCode: *question + 2,
				Name:         *name + "3",
			},
		},
	}
	resStream, err := cli.ServerStreamingRespondManytimes(ctx, req)
	if err != nil {
		log.Fatalf("fail to call cli.ServerStreamingRespondManytimes(): %v", err)
	}
	for {
		res, err := resStream.Recv()
		if isError(err) {
			//rpc error: code = DeadlineExceeded desc = context deadline exceeded
			log.Fatalf("fail to receive response from Recv(): %v", err)
			break
		}
		log.Printf("Response: [code] %d, [answer] %s", res.Code, res.Answer)
	}
}

func doClientStreaming(ctx context.Context, cli samplepb.SampleServiceClient) {
	stream, err := cli.ClientStreamingAskManytimes(ctx)
	if err != nil {
		log.Fatalf("fail to call cli.ClientStreamingAskManytimes(): %v", err)
	}

	//req := &samplepb.Client{
	//	QuestionCode: *question,
	//	Name:         *name,
	//}

	reqs := &samplepb.ManyClients{
		Clients: []*samplepb.Client{
			{
				QuestionCode: *question,
				Name:         *name,
			},
			{
				QuestionCode: *question + 1,
				Name:         *name + "2",
			},
			{
				QuestionCode: *question + 2,
				Name:         *name + "3",
			},
		},
	}

	for _, req := range reqs.Clients {
		stream.Send(req)
		time.Sleep(200 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("fail to receive response from CloseAndRecv(): %v", err)
	}

	for idx, answer := range res.Answers {
		log.Printf("Response: [%d][code] %d, [answer] %s", idx, answer.Code, answer.Answer)
	}
}

func isError(err error) bool {
	switch err {
	case io.EOF:
		log.Println("server was closed")
		return true
	case context.DeadlineExceeded:
		//FIXME: this would not return by grpc error
		log.Println("timeout by context.DeadlineExceeded")
		return true
	default:
		e, ok := err.(net.Error)
		if ok && e.Timeout() {
			log.Println("timeout by net.Error")
			return true
		}
	}

	switch status.Code(err) {
	case codes.DeadlineExceeded:
		log.Println("timeout by grpc.DeadlineExceeded")
		return true
	case
		codes.Canceled,
		codes.Unknown,
		codes.InvalidArgument,
		codes.NotFound,
		codes.AlreadyExists,
		codes.PermissionDenied,
		codes.ResourceExhausted,
		codes.FailedPrecondition,
		codes.Aborted,
		codes.OutOfRange,
		codes.Unimplemented,
		codes.Internal,
		codes.Unavailable,
		codes.DataLoss,
		codes.Unauthenticated:

		log.Println("grpc error")
	}

	return false
}
