package goroutine_test

import (
	"fmt"
	. "github.com/hiromaily/golibs/goroutine"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	u "github.com/hiromaily/golibs/utils"
	"os"
	"sync"
	"testing"
)

type User struct {
	Id   int
	Name string
}

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GOROUTINE_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
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
	//fmt.Printf("%+v\n", data)

	//convert interface{} to map[string]int
	result := u.ItoMsi(data)
	fmt.Printf("apple: %d, banana:%d, lemon:%d\n", result["apple"], result["banana"], result["lemon"])
}

//-----------------------------------------------------------------------------
// Test
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
