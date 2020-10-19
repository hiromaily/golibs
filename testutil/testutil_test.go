package testutil_test

import (
	. "github.com/hiromaily/golibs/testutil"
	//lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	InitializeTest("[TestUtil]")
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
func TestLogf(t *testing.T) {
	Logf(t, "Logf test: %s", "12345")

	Log(t, "Log test.")
}

func TestSkipLog(t *testing.T) {
	SkipLog(t)

	Log(t, "This code would be not shown")
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkTestutil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
