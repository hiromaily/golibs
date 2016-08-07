package os

import (
	"fmt"
	"os"
)

func GetOS() string {
	hostname, _ := os.Hostname()
	fmt.Println(hostname)
	//centos7
	//hy-MacBook-Pro.local
	return hostname
}

func GetEnv(name string) string {
	//os.Getenv("GOPATH")
	return os.Getenv(name)
}
