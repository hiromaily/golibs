package main

import (
	ud "github.com/hiromaily/golibs/unixdomainsocket"
)

func main() {
	//server / listener
	server := ud.NewServer()
	err := server.Open()
	if err != nil {
		panic(err)
	}

	// wait shutdown by signal
	close := make(chan error)
	server.WaitShutdown(close)

	// handle request from client
	server.Run()

	//wait from shutdown
	err = <-close
	if err != nil {
		panic(err)
	}
}
