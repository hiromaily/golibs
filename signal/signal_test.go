package signal_test

import (
	. "github.com/hiromaily/golibs/signal"
	//lg "github.com/hiromaily/golibs/log"
	"fmt"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"sync"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Signal]")
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
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSignal(t *testing.T) {
	tu.SkipLog(t)
	wg := &sync.WaitGroup{}

	go StartSignal()
	fmt.Println("Input Ctrl + c")

	wg.Add(1)
	go func() {
		count := 0
		for {
			fmt.Println(count)
			count++
		}
	}()

	wg.Wait()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkSignal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
