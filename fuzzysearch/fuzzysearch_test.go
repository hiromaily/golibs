package fuzzysearch_test

import (
	"fmt"
	. "github.com/hiromaily/golibs/fuzzysearch"
	tu "github.com/hiromaily/golibs/testutil"
	"testing"
)

var base01 = "/dk/sjaelland/koege"
var testData01 = []struct {
	base     string
	compared string
	distance int
}{
	{base01, "/dk/sjaelland",6},
	{base01, "/dk/sjaelland/koge",1},
	{base01, "/dk/sjaelland/korsor",4},
	{base01, "/dk/sjaelland/kge",2},
	{base01, "/dk/sjaelland/faxe",4},
	{base01, "/dk/sjaelland/rdby",5},
	{base01, "/dk/sjaelland/koge",1},
}

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
func TestGetRank(t *testing.T) {
	for _, val := range testData01 {
		fmt.Println(GetDistance(val.base, val.compared))
	}
}
