package unixdomainsocket_test

import (
	"testing"

	. "github.com/hiromaily/golibs/unixdomainsocket"
)

func TestServer(t *testing.T) {
	//server / listener
	server := NewServer()
	err := server.Open()
	if err != nil {
		t.Fatal(err)
	}

	// wait shutdown by signal
	close := make(chan error)
	server.WaitShutdown(close)

	// handle request from client
	server.Run()

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
