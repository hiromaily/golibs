package goroutine_test

import (
	"fmt"
	"runtime"

	. "github.com/hiromaily/golibs/goroutine"

	//lg "github.com/hiromaily/golibs/log"
	"os"
	"sync"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
	u "github.com/hiromaily/golibs/utils"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[GOROUTINE]")
}

func setup() {
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
		map[string]int{"apple": 150, "banana": 300, "lemon": 300},
		map[string]int{"apple": 180, "banana": 400, "lemon": 350},
		map[string]int{"apple": 220, "banana": 500, "lemon": 380},
	}

	//cannot use data (type []map[string]int) as type []interface {} in argument to goroutine.Semaphore2
	Semaphore2(something2, concurrencyCnt, u.SliceMapToInterface(data), wg)
}
