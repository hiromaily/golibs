package color_test

import (
	"fmt"
	c "github.com/hiromaily/golibs/color"
	tu "github.com/hiromaily/golibs/testutil"
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
	fmt.Println(c.Add("This is Blue", c.Blue))
	// Output
	// This is Blue
}

func ExampleAddf() {
	str := c.Addf("This is %s", c.Yellow, "Yellow")
	fmt.Println(str)
	// Output
	// This is Yellow
}
