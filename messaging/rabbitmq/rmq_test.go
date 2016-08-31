package rabbitmq_test

import (
	"fmt"
	. "github.com/hiromaily/golibs/messaging/rabbitmq"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var (
	benchFlg  bool   = false
	queueName string = "testQueur"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[RabitMQ_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
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
	//if err != nil {
	//	t.Errorf("TestSend error: %s", err)
	//}

	chWait := make(chan bool)

	//receiver
	go startReceiver(chWait)

	//sender
	sender := New("localhost", "hiromaily", "hiropass", 5672)

	q := sender.Declare(queueName)
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
