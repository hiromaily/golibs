package main

import (
	"flag"
	"fmt"
	sig "github.com/hiromaily/golibs/signal"
	"sync"
	"time"
)

var timeOut = flag.Int("time", 1, "TimeOut(s)")

func init() {
	flag.Parse()
}

func main() {
	setup()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			fmt.Println("child running ...")
			time.Sleep(time.Duration(*timeOut) * time.Second)
		}
	}()

	wg.Wait()
}

func setup() {
	sig.StartSignal()
}
