package unixdomainsocket

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

// NewListner is to return net.listener object
func NewListner() (net.Listener, string, error) {

	// create sock file
	sockFile := filepath.Join(os.TempDir(), "socketfile")
	fmt.Printf("socket file: %s\n", sockFile)

	// pid
	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("pid: %s\n", pid)

	// listen
	listener, err := net.Listen("unix", sockFile)
	if err != nil {
		return nil, "", err
	}

	// change owner
	if err := os.Chmod(sockFile, 0700); err != nil {
		return nil, "", err
	}

	return listener, sockFile, nil
}

// Server is to receive request from client
func Server(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		nr, err := conn.Read(buf)
		if err != nil {
			break
		}
		data := buf[0:nr]
		fmt.Printf("Recieved: %v", string(data))
		_, err = conn.Write(data)
		if err != nil {
			log.Printf("error: %v\n", err)
			break
		}
	}
}

// WaitShutdown to wait shutdown by signal
func WaitShutdown(listener net.Listener, close chan error) {
	c := make(chan os.Signal, 1)
	// wait Ctrl+C, kill -SIGTERM xxx
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			s := <-c
			switch s {
			case syscall.SIGINT:
				fmt.Println("[Signal] interrupt by control + C")
			case syscall.SIGTERM:
				fmt.Println("[Signal] force stop by kill command")
			}
			break
		}
		if err := listener.Close(); err != nil {
			close <- err
		}
		close <- nil
	}()
}
