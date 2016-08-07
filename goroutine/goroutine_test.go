package goroutine_test

import (
	"flag"
	"fmt"
	. "github.com/hiromaily/golibs/goroutine"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
	"os"
	"sync"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

type User struct {
	Id   int
	Name string
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GOROUTINE_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

func something(idx int) {
	fmt.Println(idx)
}

func something2(idx int, data interface{}) {
	fmt.Println(idx)
	//fmt.Printf("%+v\n", data)

	//convert interface{} to map[string]int
	result := u.ItoMsi(data)
	fmt.Printf("apple: %d, banana:%d, lemon:%d\n", result["apple"], result["banana"], result["lemon"])
}

//-----------------------------------------------------------------------------
// Semaphore
//-----------------------------------------------------------------------------
func TestSemaphore1(t *testing.T) {
	//t.Skip("skipping TestSemaphore1")
	//*
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}

	concurrencyCnt := 10
	execCnt := 1000
	Semaphore(something, concurrencyCnt, execCnt, wg)
}

func TestSemaphore2(t *testing.T) {
	//t.Skip("skipping TestSemaphore2")
	//*
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}

	concurrencyCnt := 10
	data := []map[string]int{
		map[string]int{"apple": 150, "banana": 300, "lemon": 300},
		map[string]int{"apple": 180, "banana": 400, "lemon": 350},
		map[string]int{"apple": 220, "banana": 500, "lemon": 380},
	}

	//cannot use data (type []map[string]int) as type []interface {} in argument to goroutine.Semaphore2
	Semaphore2(something2, concurrencyCnt, u.InterfaceSliceMap(data), wg)
}
