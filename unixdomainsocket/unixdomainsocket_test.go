package unixdomainsocket_test

import (
	"testing"

	. "github.com/hiromaily/golibs/unixdomainsocket"
)

func TestServer(t *testing.T) {
	//listner
	listener, _, err := NewListner()
	if err != nil {
		t.Fatal(err)
	}

	// wait shutdown by signal
	close := make(chan error)
	WaitShutdown(listener, close)

	// handle request from client
	Server(listener)

	//wait from shutdown
	err = <-close
	if err != nil {
		t.Fatal(err)
	}

	// remove socket file
	// if err := os.Remove(tempDir); err != nil {
	// 	t.Fatal(err)
	// }
}
