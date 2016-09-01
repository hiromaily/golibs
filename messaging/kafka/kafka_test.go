package kafka_test

import (
	"fmt"
	"github.com/Shopify/sarama"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/messaging/kafka"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
	"time"
)

var (
	benchFlg  bool   = false
	topicName string = "topic"
	host      string = "127.0.0.1"
	port      int    = 32769
)

var msgTests = []struct {
	key   string
	value string
}{
	{"key1", "aaaaa"},
	{"key2", "bbbbb"},
	{"key3", "ccccc"},
	{"key4", "ddddd"},
	{"key5", "eeeee"},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Kafka_TEST]", "/var/log/go/test.log")
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

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestKafka(t *testing.T) {
	ch := ChReceive{
		ChWait: make(chan bool),
		ChCMsg: make(chan *sarama.ConsumerMessage),
	}

	//
	//1.Consumer(Receiver)
	c, err := CreateConsumer(host, port)
	if err != nil {
		t.Errorf("CreateConsumer() error: %s", err)
	}
	go Consumer(c, topicName, ch)

	fmt.Println("wait Reveiver()")
	<-ch.ChWait //After being ready.
	fmt.Println("go after Reveiver()")

	//
	//2.Producer(Sender)
	p, err := CreateProducer(host, port)
	if err != nil {
		t.Errorf("CreateProducer() error: %s", err)
	}
	defer p.Close()

	go func() {
		for i, tt := range msgTests {
			msg := CreateMsg(topicName, tt.key, tt.value)
			err = Producer(p, msg)
			if err != nil {
				t.Errorf("[%d]Sender() error: %s", i, err)
			}
		}
	}()

	count := 0
	for {
		m := <-ch.ChCMsg
		fmt.Printf("Key: %v, Value: %v\n", string(m.Key), string(m.Value))

		count++
		if count == len(msgTests) {
			break
		}
	}

	//finish notification to Consumer
	ch.ChWait <- true
	time.Sleep(1 * time.Second)

}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkKafka(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
