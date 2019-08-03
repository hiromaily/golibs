// Package forloop_test is just sample
package forloop_test

import (
	//lg "github.com/hiromaily/golibs/log"
	//. "github.com/hiromaily/golibs/xml"
	"os"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[FORLOOP]")
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
func callSomthing() int {
	for i := 0; i < 5; i++ {
		if i == 3 {
			return 10
		}
	}
	return 5
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestForLoop(t *testing.T) {
	t.Logf("result is %d", callSomthing())
}
