package nats

import (
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"github.com/nats-io/nats"
)

//https://github.com/nats-io/nats
//https://github.com/nats-io/nats/tree/master/examples

type ChReceive struct {
	Conn   *nats.Conn
	ChWait chan bool
	ChCMsg chan *nats.Msg
	Error  error
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------
func Connection(host, user, pass string, port int) (*nats.Conn, error) {
	lg.Info("Connection()")
	//nats://derek:pass@localhost:4222
	//nats.DefaultURL
	var err error
	var nc *nats.Conn
	if host == "" {
		//"nats://localhost:4222"
		nc, err = nats.Connect(nats.DefaultURL)
	} else {
		if user == "" {
			nc, err = nats.Connect(fmt.Sprintf("nats://%s:%d", host, port))
		} else {
			nc, err = nats.Connect(fmt.Sprintf("nats://%s:%s@%s:%d", user, pass, host, port))
		}
	}

	if err != nil {
		return nil, err
	}
	return nc, nil
}

//-----------------------------------------------------------------------------
// Subscribe
//-----------------------------------------------------------------------------
func (ch ChReceive) Subscribe(subject string) {
	lg.Info("Subscribe()")

	// Async Subscriber
	var counter int = 0
	ch.Conn.Subscribe(subject, func(msg *nats.Msg) {
		counter += 1
		//lg.Debugf("[%d]Received msg:%s", counter, msg)
		ch.ChCMsg <- msg
	})
	ch.ChWait <- true //notification for being ready.

	//wait for finish
	<-ch.ChWait

	defer func() {
		ch.ChWait <- true //notification for finished
	}()

	ch.Conn.Flush()
	ch.Error = ch.Conn.LastError()

	return
}

func Unsubscribe(sub *nats.Subscription) {
	// Unsubscribe
	sub.Unsubscribe()
}

//-----------------------------------------------------------------------------
// Publish
//-----------------------------------------------------------------------------
//This func is not used because of just example.
func Publish(host, user, pass, subject, msg string, port int) error {
	lg.Info("Publish()")

	// Connection
	nc, err := Connection(host, user, pass, port)
	if err != nil {
		return err
	}
	defer nc.Close()

	// Publish
	message := []byte(msg)
	nc.Publish(subject, message)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		return err
	}

	return nil
}
