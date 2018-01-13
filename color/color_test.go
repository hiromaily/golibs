package color_test

import (
	"fmt"
	c "github.com/hiromaily/golibs/color"
	tu "github.com/hiromaily/golibs/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Color]")
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
