package rabbitmq_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/messaging/rabbitmq"
	tu "github.com/hiromaily/golibs/testutil"
)

var (
	queueName = "testQueue99"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[RabitMQ]")
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------
func startReceiver(chWait chan bool) {
	reveiver := New("localhost", "hiromaily", "hiropass", 5672)

	chBody := make(chan []byte)

	//receiver
	go func() {
		for {
			reveiver.CreateReceiver(queueName, chBody)
		}
	}()

	var body []byte
	for {
		body = <-chBody
		lg.Debugf("%s", body)
		if string(body) == "body:100" {
			chWait <- true
		}
	}
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSend(t *testing.T) {

	chWait := make(chan bool)

	//sender
	//first, queue have to be made.
	sender := New("localhost", "hiromaily", "hiropass", 5672)
	q := sender.Declare(queueName)

	//receiver
	go startReceiver(chWait)
	//wait
	time.Sleep(time.Second * 1)

	//send
	body := "aaaaabbbbbcccccdddddeeeeefffffggggg"
	sender.Send([]byte(body), q)

	for i := 0; i < 100; i++ {
		body = fmt.Sprintf("body:%d", i+1)
		sender.Send([]byte(body), q)
		//fmt.Println(i)
	}

	//wait
	<-chWait
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkSend(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
