package main

import (
	"fmt"
	"net/http"
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
	counter += 1
	time.Sleep(1 * time.Second)

	fmt.Println("hello world")
	fmt.Fprintf(w, "Hello, World")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":8080", nil)
}
