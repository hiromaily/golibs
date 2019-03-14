package main

import (
	"fmt"
	"time"

	"github.com/hiromaily/golibs/files"
)

func main() {
	l, err := files.NewFileLock("cmd/lock/main.go")
	if err != nil {
		panic(err)
	}

	fmt.Println("lock")
	l.Lock()

	time.Sleep(10 * time.Second)

	l.Unlock()
	fmt.Println("unlock")
}
