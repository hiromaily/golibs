package color_test

import (
	"fmt"
	"os"
	"testing"

	c "github.com/hiromaily/golibs/color"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Color]")
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
func ExampleAdd() {
	fmt.Println(c.Add(c.Blue, "This is Blue"))
	// Output
	// This is Blue
}

func ExampleAddf() {
	str := c.Addf(c.Yellow, "This is %s", "Yellow")
	fmt.Println(str)
	// Output
	// This is Yellow
}

func TestCheck(t *testing.T) {
	c.Check()
}
