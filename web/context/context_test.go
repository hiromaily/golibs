package context_test

import (
	. "github.com/hiromaily/golibs/web/context"
	//lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Context]")
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
// Test
//-----------------------------------------------------------------------------
func TestContextCancel(t *testing.T) {
	ContextWithCancel()
	//if err != nil {
	//	t.Errorf("TestContext error: %s", err)
	//}
}

func TestContextWithTimeout(t *testing.T) {
	ContextWithTimeout()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkContext(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
