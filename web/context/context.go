package context

import (
	"context"
	"fmt"
	"time"
)

//http://deeeet.com/writing/2016/07/22/context/
//https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39#.whnyr85ju

func WithCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(time.Second, cancel)
	//go func() {
	//	time.Sleep(time.Second)
	//	cancel()
	//}()

	sleepAndTalk(ctx, 5*time.Second, "hello!")
}

func WithTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "hello!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {

	select {
	case <-time.After(d):
		fmt.Printf("[timeout] msg is %s\n", msg)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(ctx, "[canceled]", err.Error())
	}
	//time.Sleep(d)
	//fmt.Println(msg)
}
