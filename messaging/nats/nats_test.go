package nats_test

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/nats-io/nats.go"

	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/messaging/nats"
	tu "github.com/hiromaily/golibs/testutil"
)

var (
	subjectName  = "msg.hiromaily"
	subjectName2 = "msg2.hiromaily"
	subjectName3 = "msg3.hiromaily"
)

var msgTests = []struct {
	msg string
}{
	{"aaaaa"},
	{"bbbbb"},
	{"ccccc"},
	{"ddddd"},
	{"eeeee"},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[NATS]")
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

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestNats(t *testing.T) {
	t.SkipNow()

	ch := ChReceive{
		Conn:   nil,
		ChWait: make(chan bool),
		ChCMsg: make(chan *nats.Msg),
		Error:  nil,
	}
	var err error

	// 1.Connection
	ch.Conn, err = Connection("", "", "", 0)
	if err != nil {
		t.Errorf("Connection() error: %s", err)
	}
	defer ch.Conn.Close()

	// 2.Subscribe
	go ch.Subscribe(subjectName)
	<-ch.ChWait //After being ready.

	// 3.Connection
	nc, err := Connection("", "", "", 0)
	if err != nil {
		t.Errorf("Connection() error: %s", err)
	}
	defer nc.Close()

	// 4.Publish(send)
	go func() {
		for i, tt := range msgTests {
			// Publish
			message := []byte(fmt.Sprintf("%d:%s", i, tt.msg))
			nc.Publish(subjectName, message)
			nc.Flush()

			if err := nc.LastError(); err != nil {
				t.Errorf("Publish() error: %s", err)
			}
		}
	}()

	// 5.receive
	count := 0
	for {
		m := <-ch.ChCMsg
		lg.Debugf("subject:%s, msg: %s\n", m.Subject, m.Data)

		count++
		if count == len(msgTests) {
			break
		}
	}

	//finish notification to Subscriber
	ch.ChWait <- true

	<-ch.ChWait //After finished
	if ch.Error != nil {
		t.Errorf("Subscribe() error: %s", err)
	}
}

func TestNats2(t *testing.T) {
	t.SkipNow()

	// Connection
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	// Sync Subscriber
	sub, err := nc.SubscribeSync(subjectName2)
	if err != nil {
		t.Errorf("SubscribeSync() error: %s", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Publisher
	go func() {
		count := 0
		for {
			nc.Publish(subjectName2, []byte("Hello World"))
			count++
			if count == 100 {
				break
			}
		}
	}()

	// Subscriber
	go func() {
		count := 0
		for {
			m, err := sub.NextMsg(time.Second * 1)
			if err == nil {
				lg.Debugf("message is %s", m.Data)
				count++
				if count == 100 {
					wg.Done()
					break
				}
			}
		}
	}()

	//wait
	wg.Wait()

}

func TestNats3(t *testing.T) {
	t.SkipNow()

	var msg *nats.Msg
	//msg := &nats.Msg{}

	// Connection
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(subjectName3, ch)
	if err != nil {
		t.Errorf("ChanSubscribe() error: %s", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Publisher
	go func() {
		count := 0
		for {
			nc.Publish(subjectName3, []byte("Hello World"))
			count++
			if count == 100 {
				lg.Debug("finished publish")
				break
			}
		}
	}()

	// Subscriber
	go func() {
		count := 0
		for {
			//invalid operation: msg <- ch (send to non-chan type *"github.com/nats-io/nats".Msg)
			msg = <-ch
			lg.Debugf("message is %s", msg.Data)

			count++
			if count == 100 {
				wg.Done()
				break
			}
		}
	}()

	//wait
	wg.Wait()

	// Unsubscribe
	sub.Unsubscribe()
}

func TestNats4(t *testing.T) {
	//t.SkipNow()

	// Connection
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Requests
	go func() {
		count := 0
		for {
			msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
			if err == nil {
				lg.Debugf("message is %s", msg.Data)
			}

			count++
			if count == 100 {
				wg.Done()
				break
			}
		}
	}()

	//wait
	wg.Wait()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkNats(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
