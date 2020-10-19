package goroutine_test

import (
	"fmt"
	"runtime"
	"time"

	. "github.com/hiromaily/golibs/goroutine"
	tu "github.com/hiromaily/golibs/testutil"

	//lg "github.com/hiromaily/golibs/log"
	"os"
	"sync"
	"testing"

	u "github.com/hiromaily/golibs/utils"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[GOROUTINE]")
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------
func something(idx int) {
	fmt.Println(idx)
}

func something2(idx int, data interface{}) {
	fmt.Println(idx)

	//convert interface{} to map[string]int
	result := u.ItoMsi(data)
	fmt.Printf("apple: %d, banana:%d, lemon:%d\n", result["apple"], result["banana"], result["lemon"])
}

func loop(count int) {
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetGOMAXPROCS(t *testing.T) {
	count := 1000000000

	//1.32s
	// loop(count)
	// loop(count)
	// loop(count)
	// loop(count)

	//0.64s
	wg := &sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			loop(count)
			wg.Done()
		}()
	}
	wg.Wait()

	t.Log(GetGOMAXPROCS())
}

func TestSemaphore1(t *testing.T) {
	//tu.SkipLog(t)

	wg := &sync.WaitGroup{}

	concurrencyCnt := 10
	execCnt := 1000
	Semaphore(something, concurrencyCnt, execCnt, wg)
}

func TestSemaphore2(t *testing.T) {
	//tu.SkipLog(t)

	wg := &sync.WaitGroup{}

	concurrencyCnt := 10
	data := []map[string]int{
		{"apple": 150, "banana": 300, "lemon": 300},
		{"apple": 180, "banana": 400, "lemon": 350},
		{"apple": 220, "banana": 500, "lemon": 380},
	}

	//cannot use data (type []map[string]int) as type []interface {} in argument to goroutine.Semaphore2
	Semaphore2(something2, concurrencyCnt, u.SliceMapToInterface(data), wg)
}

//select block until data coming
func TestSelect(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Printf("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func TestSelect2(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	var c1Count, c2Count int
	go func() {
		for {
			select {
			case <-c1:
				c1Count++
			case <-c2:
				c2Count++
			}
			fmt.Println("is this code running?? Not!")
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

//after closing channel, select doesn't block
func TestSelect3(t *testing.T) {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})

	var c1Count, c2Count int
	//after closing channel, select doesn't block
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

}
