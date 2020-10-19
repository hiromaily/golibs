package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
)

const (
	address = "localhost:50051"
)

var (
	mode     = flag.Int("mode", 1, "mode")
	name     = flag.String("name", "", "name")
	question = flag.Int64("q", 0, "question code")
	isTLS    = flag.Bool("tls", false, "tls mode")
	certFile = fmt.Sprintf("%s/src/github.com/hiromaily/golibs/grpc/key/ca.crt", os.Getenv("GOPATH"))
)

var clients = []*samplepb.Client{
	{
		QuestionCode: 1,
		Name:         "Mike",
	},
	{
		QuestionCode: 5,
		Name:         "Salon",
	},
	{
		QuestionCode: 10,
		Name:         "Quen",
	},
	{
		QuestionCode: 18,
		Name:         "Lin",
	},
	{
		QuestionCode: 29,
		Name:         "Michael",
	},
	{
		QuestionCode: 50,
		Name:         "Rocky",
	},
}

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

	// Set up a connection to the server.
	var conn = new(grpc.ClientConn)

	if *isTLS {
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("fail to call credentials.NewClientTLSFromFile(): %v", err)
		}
		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatalf("fail to connect: %v", err)
		}
	} else {
		var err error
		conn, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to connect: %v", err)
		}
	}

	defer conn.Close()
	cli := samplepb.NewSampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	//TODO: check `SampleServiceClient interface` in pb.go

	switch *mode {
	case 1:
		validate()
		//Unary
		doUnary(ctx, cli)
	case 2:
		//ServerStreaming
		doServerStreaming(ctx, cli)
	case 3:
		//ClientStreaming
		doClientStreaming(ctx, cli)
	case 4:
		doBidirectionalStreaming(ctx, cli)
	default:
		log.Fatal("-mode is out of range")
	}
}

func doUnary(ctx context.Context, cli samplepb.SampleServiceClient) {
	//grpc metadata
	//https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md
	md := metadata.New(map[string]string{"key1": "val1", "key2": "val2"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := cli.UnaryAsk(ctx, &samplepb.Client{
		QuestionCode: *question,
		Name:         *name,
	})
	if err != nil {
		if respErr, ok := status.FromError(err); ok {
			switch respErr.Code() {
			case codes.InvalidArgument:
				log.Println("parameter is invalid")
			case codes.DeadlineExceeded:
				log.Println("timeout")
			case codes.Unavailable:
				log.Println("authentication handshake failed")
			default:
				log.Printf("something grpc error: %d, %s", respErr.Code(), respErr.Message())
			}
		} else {
			log.Fatalf("fail to call cli.UnaryAsk(): %v", err)
		}
		return
	}
	log.Printf("Response: [code] %d, [answer] %s", res.Code, res.Answer)
}

func doServerStreaming(ctx context.Context, cli samplepb.SampleServiceClient) {
	req := &samplepb.ManyClients{
		Clients: clients,
	}
	log.Println("[doServerStreaming] send from client")
	resStream, err := cli.ServerStreamingRespondManytimes(ctx, req)
	if err != nil {
		log.Fatalf("fail to call cli.ServerStreamingRespondManytimes(): %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			log.Println("server is closed")
			break
		} else if err != nil {
			if respErr, ok := status.FromError(err); ok {
				switch respErr.Code() {
				case codes.InvalidArgument:
					log.Println("parameter is invalid")
				case codes.DeadlineExceeded:
					log.Println("timeout")
				default:
					log.Printf("something grpc error: %d, %s", respErr.Code(), respErr.Message())
				}
			} else {
				log.Fatalf("fail to receive response from Recv(): %v", err)
			}
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

	reqs := &samplepb.ManyClients{
		Clients: clients,
	}

	for _, req := range reqs.Clients {
		log.Println("send from client")
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("receive from server")
	res, err := stream.CloseAndRecv()
	if err != nil {
		if respErr, ok := status.FromError(err); ok {
			switch respErr.Code() {
			case codes.InvalidArgument:
				log.Println("parameter is invalid")
			case codes.DeadlineExceeded:
				log.Println("timeout")
			default:
				log.Printf("something grpc error: %d, %s", respErr.Code(), respErr.Message())
			}
		} else {
			log.Fatalf("fail to receive response from Recv(): %v", err)
		}
		return
	}

	for _, answer := range res.Answers {
		log.Printf("Response: [code] %d, [answer] %s", answer.Code, answer.Answer)
	}
}

func doBidirectionalStreaming(ctx context.Context, cli samplepb.SampleServiceClient) {
	stream, err := cli.BidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("fail to call cli.BidirectionalStreaming(): %v", err)
	}

	waitc := make(chan struct{})

	reqs := &samplepb.ManyClients{
		Clients: clients,
	}

	//send
	go func() {
		for _, client := range reqs.Clients {
			log.Println("send from client")
			stream.Send(client)
			time.Sleep(200 * time.Millisecond)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("fail to call CloseSend(): %v", err)
		}
	}()

	//receive
	go func() {
		for {
			log.Println("receive from server")
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				if respErr, ok := status.FromError(err); ok {
					switch respErr.Code() {
					case codes.InvalidArgument:
						log.Println("parameter is invalid")
					case codes.DeadlineExceeded:
						log.Println("timeout")
					default:
						log.Printf("something grpc error: %d, %s", respErr.Code(), respErr.Message())
					}
				} else {
					log.Fatalf("fail to receive response from Recv(): %v", err)
				}
				break
			}
			log.Printf("Response: [code] %d, [answer] %s", res.Code, res.Answer)
		}
		close(waitc)
	}()

	<-waitc
}

//func isError(err error) bool {
//	switch err {
//	case io.EOF:
//		log.Println("server was closed")
//		return true
//	case context.DeadlineExceeded:
//		//FIXME: this would not return by grpc error
//		log.Println("timeout by context.DeadlineExceeded")
//		return true
//	default:
//		e, ok := err.(net.Error)
//		if ok && e.Timeout() {
//			log.Println("timeout by net.Error")
//			return true
//		}
//	}
//
//	switch status.Code(err) {
//	case codes.DeadlineExceeded:
//		log.Println("timeout by grpc.DeadlineExceeded")
//		return true
//	case codes.InvalidArgument:
//		log.Println("invalid error")
//		return true
//	case
//		codes.Canceled,
//		codes.Unknown,
//		codes.NotFound,
//		codes.AlreadyExists,
//		codes.PermissionDenied,
//		codes.ResourceExhausted,
//		codes.FailedPrecondition,
//		codes.Aborted,
//		codes.OutOfRange,
//		codes.Unimplemented,
//		codes.Internal,
//		codes.Unavailable,
//		codes.DataLoss,
//		codes.Unauthenticated:
//
//		log.Println("grpc error")
//	}
//
//	return false
//}
