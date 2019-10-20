package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var counter int

func subtask(i int) {
	fmt.Println(" start subtask:", i)
	time.Sleep(2 * time.Second)
	fmt.Println(" end subtask", i)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----- get request -----")
	go subtask(counter)
	counter++
	time.Sleep(1 * time.Second)

	fmt.Println("hello world")
	fmt.Fprintf(w, "Hello, World")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----- get request 2 -----")

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	chResult := make(chan []int, 1)
	sum := make([]int, 0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		chResult <- []int{1, 2, 3}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		chResult <- nil
	}()

	go func() {
		wg.Wait()
		close(chResult)
	}()

	var isDone bool
	for {
		select {
		case ret, isOpen := <-chResult:
			if !isOpen {
				isDone = true
			}
			if ret == nil {
				fmt.Println("ret == nil")
			}
			sum = append(sum, ret...)
		case <-ctx.Done():
			fmt.Println("timeout")
			isDone = true
		}

		if isDone {
			break
		}
	}

	fmt.Printf("hello world: %d\n", len(sum))
	fmt.Fprintf(w, "Hello, World")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/sample1", handler)
	http.HandleFunc("/sample2", handler2)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":8080", nil)
}
