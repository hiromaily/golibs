package wrapper_test

import (
	"os"
	"testing"

	. "github.com/hiromaily/golibs/example/wrapper"
	tu "github.com/hiromaily/golibs/testutil"
)

//https://golang.org/pkg/strings/

var (
	data1 = "abcdefghijk"
	data2 = "ABCDEFGHIJK"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Wrapper]")
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
func TestParent(t *testing.T) {
	fn := Parent(10, "test")
	fn()
}
