package main

import (
	ud "github.com/hiromaily/golibs/unixdomainsocket"
)

func main() {
	//listner
	listener, _, err := ud.NewListner()
	if err != nil {
		panic(err)
	}

	// wait shutdown by signal
	close := make(chan error)
	ud.WaitShutdown(listener, close)

	// handle request from client
	ud.Server(listener)

	//wait from shutdown
	err = <-close
	if err != nil {
		panic(err)
	}
}
