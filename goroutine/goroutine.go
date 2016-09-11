package goroutine

import (
	"os"
	"strconv"
	"sync"
)

//TODO: How many goroutine is possible

// GetGOMAXPROCS is to get number of core
func GetGOMAXPROCS() int {
	if os.Getenv("GOMAXPROCS") != "" {
		coreNum, _ := strconv.Atoi(os.Getenv("GOMAXPROCS"))
		return coreNum
	}
	//fmt.Println(runtime.NumCPU())        //number of logical CPUs
	//fmt.Println(runtime.GOMAXPROCS(0))   //GOMAXPROCS sets the maximum number of CPUs(Don't need to use)
	//fmt.Println(runtime.NumGoroutine())  //number of goroutines
	return 0
}

// Semaphore is management of concurrency
// It's to execute simultaneously specific func at designated number
func Semaphore(paramFunc func(int), pool int, cnt int, wg *sync.WaitGroup) {
	chanSemaphore := make(chan bool, pool)

	for i := 0; i < cnt; i++ {
		wg.Add(1)
		chanSemaphore <- true

		//chanSemaphore <- true
		go func(cnt int) {
			defer func() {
				<-chanSemaphore
				wg.Done()
			}()
			//concurrent func
			paramFunc(cnt)
		}(i)
	}
	wg.Wait()
}

// Semaphore2 is management of concurrency
// It's to execute simultaneously specific func at number of slice data
func Semaphore2(paramFunc func(int, interface{}), pool int, data []interface{}, wg *sync.WaitGroup) {
	chanSemaphore := make(chan bool, pool)

	for i, d := range data {
		wg.Add(1)
		chanSemaphore <- true

		//chanSemaphore <- true
		go func(cnt int, one interface{}) {
			defer func() {
				<-chanSemaphore
				wg.Done()
			}()
			//concurrent func
			paramFunc(cnt, one)
		}(i, d)
	}
	//close(chanSemaphore)
	wg.Wait()
}
