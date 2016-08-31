package rabbitmq

import (
	//"errors"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"github.com/streadway/amqp"
)

type RM struct {
	DB *amqp.Connection
	Ch *amqp.Channel
}

var ContentTypes []string = []string{
	"text/plain",
	"application/json",
	"text/html",
	"text/xml",
}

var (
	//rm RM
	defConType uint8 = 0
)

func failOnError(err error, msg string) {
	if err != nil {
		lg.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func SetContentType(val uint8) {
	defConType = val
}

//connect to RabbitMQ server
func New(host, user, pass string, port int) *RM {
	var err error
	var rm RM
	//rm.DB, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	rm.DB, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	rm.CreateChannel()

	return &rm
}

//func GetConn() *RM {
//	if rm.DB == nil {
//		failOnError(errors.New("RM object is nil."), "Failed to get object")
//	}
//	return &rm
//}

func (r *RM) Close() {
	r.DB.Close()
}

//
func (r *RM) CreateChannel() {
	var err error
	r.Ch, err = r.DB.Channel()
	failOnError(err, "Failed to open a channel")
	//defer r.Ch.Close()
}

func (r *RM) CloseChannel() {
	r.Ch.Close()
}

func (r *RM) Declare(name string) *amqp.Queue {
	q, err := r.Ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &q
}

func (r *RM) CreateReceiver(name string, chBody chan []byte) {
	msgs, err := r.Ch.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		chBody <- d.Body
	}
}

func (r *RM) Send(body []byte, q *amqp.Queue) {
	// It is common to use serialisation formats like JSON, Thrift, Protocol Buffers
	// and MessagePack to serialize structured data in order to publish it as the message payload.
	err := r.Ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: ContentTypes[defConType], //"text/plain"
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}
