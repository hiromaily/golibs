package os

import (
	reg "github.com/hiromaily/golibs/regexp"
	"os"
)

func GetOS() string {
	hostname, _ := os.Hostname()
	//centos7
	//hy-MacBook-Pro.local
	return hostname
}

func GetEnv(name string) string {
	//os.Getenv("GOPATH")
	return os.Getenv(name)
}

// Get args
func GetArgs(i int) string {
	return os.Args[i]
}

// Add value to args
func AddParam(val string) {
	os.Args = append(os.Args, val)
}

// Find specific args
func FindParam(key string) (bRet bool) {
	//fmt.Println(os.Args)
	bRet = false
	for _, v := range os.Args {
		if reg.CheckRegexp(`^`+key, v) {
			bRet = true
			break
		}
	}
	return
}
