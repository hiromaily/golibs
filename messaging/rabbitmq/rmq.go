package rabbitmq

import (
	//"errors"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"github.com/streadway/amqp"
	"time"
)

// RM is struct of RabitMQ object
type RM struct {
	DB *amqp.Connection
	Ch *amqp.Channel
}

// ContentTypes is content type of messages
var ContentTypes = []string{
	"text/plain",
	"application/json",
	"text/html",
	"text/xml",
}

var defConType uint8

func failOnError(err error, msg string) {
	if err != nil {
		lg.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// SetContentType is to set default ContentType
func SetContentType(val uint8) {
	defConType = val
}

// New is to connect to RabbitMQ server
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

// Close is to close connection
func (r *RM) Close() {
	r.DB.Close()
}

// CreateChannel is to create channel
func (r *RM) CreateChannel() {
	var err error
	r.Ch, err = r.DB.Channel()
	failOnError(err, "Failed to open a channel")
	//defer r.Ch.Close()
}

// CloseChannel is to close connection to channel
func (r *RM) CloseChannel() {
	r.Ch.Close()
}

// Declare is to declare of queue name
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

// CreateConsume is to create Consume
func (r *RM) CreateConsume(name string, chBody chan []byte) (<-chan amqp.Delivery, error) {
	return r.Ch.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

// CreateReceiver is for consumer
func (r *RM) CreateReceiver(name string, chBody chan []byte) {
	retry := 5
	var err error
	var msgs <-chan amqp.Delivery

	for i := 0; i < retry; i++ {
		msgs, err = r.CreateConsume(name, chBody)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}

	for d := range msgs {
		chBody <- d.Body
	}
}

// Send is for producer
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
