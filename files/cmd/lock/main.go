package main

import (
	"fmt"
	"time"

	"github.com/hiromaily/golibs/files"
	t "github.com/hiromaily/golibs/time"
)

func main() {
	defer t.Track(time.Now(), "lock/main()")

	l, err := files.NewFileLock("cmd/lock/main.go")
	if err != nil {
		panic(err)
	}

	fmt.Println("lock")
	if l.Lock() != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	if l.Unlock() != nil {
		panic(err)
	}
	fmt.Println("unlock")
}
