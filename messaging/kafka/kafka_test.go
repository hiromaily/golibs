package kafka_test

import (
	"github.com/Shopify/sarama"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/messaging/kafka"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

var (
	topicName string = "topic"
	host      string = "127.0.0.1"
	//port      int    = 32768 //TODO:this port number may change.
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
	tu.InitializeTest("[KAFKA]")
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
// Check
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
	c, err := CreateConsumer(host, *tu.KafkaIP)
	if err != nil {
		t.Errorf("CreateConsumer() error: %s", err)
	}
	go Consumer(c, topicName, ch)

	lg.Debug("wait Reveiver()")
	<-ch.ChWait //After being ready.
	lg.Debug("go after Reveiver()")

	//
	//2.Producer(Sender)
	p, err := CreateProducer(host, *tu.KafkaIP)
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
		lg.Debugf("Key: %v, Value: %v", string(m.Key), string(m.Value))

		count++
		if count == len(msgTests) {
			break
		}
	}

	//finish notification to Consumer
	ch.ChWait <- true

	<-ch.ChWait //After finish
	//time.Sleep(1 * time.Second)

	//consumer
	//c.Close()
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
