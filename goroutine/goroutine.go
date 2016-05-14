package goroutine

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Get number of core
func GetGOMAXPROCS() int{
	if os.Getenv("GOMAXPROCS") != "" {
		coreNum, _ := strconv.Atoi(os.Getenv("GOMAXPROCS"))
		return coreNum
	}
	//runtime.NumCPU()
	return 0
}

// Manage Concurrency
// 指定した数を並列で実行し、終了したらchannelで通知する
func Semaphore(paramFunc func(), pool int){
	chanSemaphore := make(chan bool, pool)

	//TODO:終了の判断ロジックは各々で異なるので、それをどうするか
	for {
		chanSemaphore <- true

		//chanSemaphore <- true
		go func() {
			defer func() {
				<-chanSemaphore
			}()
			//concurrent func
			paramFunc()
		}()
	}
}
