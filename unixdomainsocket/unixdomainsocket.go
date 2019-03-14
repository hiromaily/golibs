package unixdomainsocket

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// Server object
type Server struct {
	listener   net.Listener
	pid        int
	socketPath string
}

// NewServer is to return server object
func NewServer() *Server {
	s := new(Server)
	return s
}

// Open is to create net.listener object
func (s *Server) Open() error {

	// create sock file
	s.socketPath = filepath.Join(os.TempDir(), "socketfile")
	fmt.Printf("socket file: %s\n", s.socketPath)

	// pid
	s.pid = os.Getpid()
	fmt.Printf("pid: %d\n", s.pid)

	// listen
	var err error
	s.listener, err = net.Listen("unix", s.socketPath)
	if err != nil {
		return err
	}

	// change owner
	if err := os.Chmod(s.socketPath, 0700); err != nil {
		s.Close()
		return err
	}

	return nil
}

// Close is to close listener
func (s *Server) Close() error {
	return s.listener.Close()
}

// Run is to receive request from client
func (s *Server) Run() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}
		go s.process(conn)
	}
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		nr, err := conn.Read(buf)
		if err != nil {
			break
		}
		data := buf[0:nr]
		fmt.Printf("Received: %v", string(data))
		_, err = conn.Write(data)
		if err != nil {
			log.Printf("error: %v\n", err)
			break
		}
	}
}

// WaitShutdown to wait shutdown by signal
func (s *Server) WaitShutdown(close chan error) {
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
		if err := s.Close(); err != nil {
			close <- err
		}
		close <- nil
	}()
}
