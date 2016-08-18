package sqs

//TODO:work in progress
import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	awslib "github.com/hiromaily/golibs/aws"
	conf "github.com/hiromaily/golibs/config"
	"strconv"
)

var svc *sqs.SQS

// send message to queue
func SendMessageToQueue(params *sqs.SendMessageInput) {
	req, resp := svc.SendMessageRequest(params)

	err := req.Send()

	if err == nil {
		// resp is now filled
		fmt.Println(resp)
	}
}

// send multiple messages to queue
func SendMultipleMessagesToQueue(params *sqs.SendMessageBatchInput) {
	req, resp := svc.SendMessageBatchRequest(params)

	err := req.Send()

	if err == nil {
		// resp is now filled
		fmt.Println(resp)
	}
}

// purge queue
// 削除するタイミングに制限あり(一定間隔内における連続削除はできない)
// http://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_PurgeQueue.html
func PurgeQueue(url *string) {
	fmt.Println(*url)
	params := &sqs.PurgeQueueInput{
		QueueUrl: aws.String(*url), // Required
	}

	req, resp := svc.PurgeQueueRequest(params)

	err := req.Send()

	if err != nil { // resp is now filled
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

// create input parameter
// TODO:このAttributesのパラメータは動的に変更せねばならない。
// TODO:contentTypeも動的に変更できるように
func CreateSendMessageInput(url *string, body *string, acid *string, ot string, ct string) (params *sqs.SendMessageInput) {
	//fmt.Println(*url)
	//fmt.Println(*body)

	var acidData string = "736afbdc388752eb"
	if *acid != "" {
		acidData = *acid
	}

	if ot != "2" {
		params = &sqs.SendMessageInput{
			MessageBody:  aws.String(*body),
			QueueUrl:     aws.String(*url),
			DelaySeconds: aws.Int64(0),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"operationType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String(ot),
				},
				"clientId": {
					DataType:    aws.String("String"),
					StringValue: aws.String(acidData),
				},
				"contentType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String(ct),
				},
			},
		}
	} else {
		params = &sqs.SendMessageInput{
			MessageBody:  aws.String(*body),
			QueueUrl:     aws.String(*url),
			DelaySeconds: aws.Int64(0),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"operationType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String(ot),
				},
				"clientId": {
					DataType:    aws.String("String"),
					StringValue: aws.String(acidData),
				},
				"messageScheduledId": {
					DataType:    aws.String("Number"),
					StringValue: aws.String("1122"),
				},
				"trackingFlag": {
					DataType:    aws.String("Number"),
					StringValue: aws.String("0"),
				},
				"contentType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String("99"),
				},
			},
		}
	}
	return
}

// create input parameter for batch
// up to ten messages!!
func CreateSendMessageBatchInput(url *string, body *string, acid *string, ot string, ct string, num int) (params *sqs.SendMessageBatchInput) {
	var acidData string = "736afbdc388752eb"
	if *acid != "" {
		acidData = *acid
	}
	conf := conf.GetConf()
	if ot == "0" {
		ot = conf.Aws.Sqs.MsgAttr.OpType
	}
	if ct == "0" {
		ct = conf.Aws.Sqs.MsgAttr.OpType
	}

	var entries []*sqs.SendMessageBatchRequestEntry
	for i := 0; i < num; i++ {
		entry := &sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(strconv.Itoa(i + 1)),
			MessageBody: aws.String(*body),
			//QueueUrl:     aws.String(*url),
			DelaySeconds: aws.Int64(0),
			MessageAttributes: map[string]*sqs.MessageAttributeValue{
				"operationType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String(ot),
				},
				"clientId": {
					DataType:    aws.String("String"),
					StringValue: aws.String(acidData),
				},
				"contentType": {
					DataType:    aws.String("Number"),
					StringValue: aws.String(ct),
				},
			},
		}
		entries = append(entries, entry)
	}

	params = &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(*url), // Required
	}
	return
}

// Queueを作成. 既に存在していれば、urlを返してくれる。
func CreateNewQueue(params *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	//create
	resp, err := svc.CreateQueue(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		return nil, err
	}

	// Pretty-print the response data.
	return resp, nil
}

// create input parameter
func CreateInputParam(queueName string) *sqs.CreateQueueInput {
	//parameter
	params := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			//"VisibilityTimeout": aws.String("30"),
			//"MessageRetentionPeriod": aws.String("345600"),
			//"MaximumMessageSize": aws.String("262144"),
			"DelaySeconds":                  aws.String("0"),
			"ReceiveMessageWaitTimeSeconds": aws.String("20"),
		},
	}
	return params
}

// CreateAttributesParams
func CreateAttributesParams(url *string) *sqs.GetQueueAttributesInput {
	params := &sqs.GetQueueAttributesInput{
		QueueUrl: aws.String(*url), // Required
		AttributeNames: []*string{
			aws.String("All"), // Required
			// More values...
		},
	}
	return params
}

// Get Queue Attributes
func GetQueueAttributes(params *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	resp, err := svc.GetQueueAttributes(params)
	return resp, err
}

// Create sqs client
func New() {
	//set environment variable
	awslib.InitAwsEnv("", "")

	//get config
	conf := conf.GetConf()

	//create client for sqs
	svc = sqs.New(session.New(), aws.NewConfig().WithRegion(conf.Aws.Region))
}
