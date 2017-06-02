package strings_test

import (
	//. "github.com/hiromaily/golibs/example/flag"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"strings"
	"testing"
	"unicode"
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
	tu.InitializeTest("[STRINGS]")
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
func TestCompare(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.Compare(data1, data1)
	lg.Debugf("[1]strings.Compare():%v", ret) //0

	ret = strings.Compare(data1, "abc")
	lg.Debugf("[2]strings.Compare():%v", ret) //1

	ret = strings.Compare("abc", data1)
	lg.Debugf("[3]strings.Compare():%v", ret) //-1
}

//TODO:It's useful!!
func TestContains(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.Contains(data1, data1)
	lg.Debugf("[1]strings.Contains():%v", ret) //true

	ret = strings.Contains(data1, "abc")
	lg.Debugf("[2]strings.Contains():%v", ret) //true

	ret = strings.Contains(data1, "")
	lg.Debugf("[3]strings.Contains():%v", ret) //true

	ret = strings.Contains(data1, "ABC")
	lg.Debugf("[4]strings.Contains():%v", ret) //false

	ret = strings.Contains(data1, "what")
	lg.Debugf("[5]strings.Contains():%v", ret) //false
}

func TestContainsAny(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.ContainsAny(data1, data1)
	lg.Debugf("[1]strings.ContainsAny():%v", ret) //true

	ret = strings.ContainsAny(data1, "abc")
	lg.Debugf("[2]strings.ContainsAny():%v", ret) //true

	ret = strings.ContainsAny(data1, "c")
	lg.Debugf("[3]strings.ContainsAny():%v", ret) //true

	ret = strings.ContainsAny(data1, "a & z")
	lg.Debugf("[4]strings.ContainsAny():%v", ret) //true

	ret = strings.ContainsAny(data1, "z")
	lg.Debugf("[5]strings.ContainsAny():%v", ret) //false

	ret = strings.ContainsAny(data1, "")
	lg.Debugf("[6]strings.ContainsAny():%v", ret) //false
}

func TestContainsRune(t *testing.T) {
	tu.SkipLog(t)

	var data3 rune
	data3 = 'a'
	ret := strings.ContainsRune(data1, data3)
	lg.Debugf("[1]strings.ContainsRune():%v", ret) //true

	data3 = 'あ'
	ret = strings.ContainsRune(data1, data3)
	lg.Debugf("[2]strings.ContainsRune():%v", ret) //false

	data3 = ' '
	ret = strings.ContainsRune(data1, data3)
	lg.Debugf("[3]strings.ContainsRune():%v", ret) //false

	//data3 = ''
	//ret = strings.ContainsRune(data1, data3)
	//lg.Debugf("[4]strings.ContainsRune():%v", ret) //
}

func TestCount(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Count(data1, "a")
	lg.Debugf("[1]strings.Count():%v", ret) //1

	ret = strings.Count(data1, "bc")
	lg.Debugf("[2]strings.Count():%v", ret) //1

	ret = strings.Count(data1, "ac")
	lg.Debugf("[3]strings.Count():%v", ret) //0

	ret = strings.Count(data1, "")
	lg.Debugf("[4]strings.Count():%v", ret) //12
}

func TestEqualFold(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.EqualFold("Go", "go")
	lg.Debugf("[1]strings.EqualFold():%v", ret) //true

	ret = strings.EqualFold("go", "go")
	lg.Debugf("[2]strings.EqualFold():%v", ret) //true

	ret = strings.EqualFold("ABC", "abc")
	lg.Debugf("[3]strings.EqualFold():%v", ret) //true

	ret = strings.EqualFold("ABCd", "abc")
	lg.Debugf("[4]strings.EqualFold():%v", ret) //false

	ret = strings.EqualFold("ABC", "bc")
	lg.Debugf("[5]strings.EqualFold():%v", ret) //false
}

//TODO:It's useful!!
func TestFields(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Fields("  foo bar  baz   ")
	lg.Debugf("[1]strings.Fields():%q", ret)   //["foo" "bar" "baz"]
	lg.Debugf("[1.1]strings.Fields():%v", ret) //[foo bar baz]
}

func TestFieldsFunc(t *testing.T) {
	tu.SkipLog(t)
	f := func(c rune) bool {
		//fmt.Println(c)
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	lg.Debugf("[1]strings.FieldsFunc():%q", strings.FieldsFunc("  foo1;bar2,baz3...", f)) //["foo1" "bar2" "baz3"]
}

func TestHasPrefix(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.HasPrefix("Gopher", "Go")
	lg.Debugf("[1]strings.HasPrefix():%v", ret) //true

	ret = strings.HasPrefix("Gopher", "C")
	lg.Debugf("[2]strings.HasPrefix():%v", ret) //false

	ret = strings.HasPrefix("Gopher", "go")
	lg.Debugf("[3]strings.HasPrefix():%v", ret) //false

	ret = strings.HasPrefix("Gopher", "")
	lg.Debugf("[4]strings.HasPrefix():%v", ret) //true
}

func TestHasSuffix(t *testing.T) {
	//tu.SkipLog(t)
	ret := strings.HasSuffix("GopherMIX", "MIX")
	lg.Debugf("[1]strings.HasSuffix():%v", ret) //true

	ret = strings.HasSuffix("GopherMIX", "C")
	lg.Debugf("[2]strings.HasSuffix():%v", ret) //false

	ret = strings.HasSuffix("GopherMIX", "ix")
	lg.Debugf("[3]strings.HasSuffix():%v", ret) //false

	ret = strings.HasSuffix("GopherMIX", "")
	lg.Debugf("[4]strings.HasSuffix():%v", ret) //true
}


//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}

//func TestContains(t *testing.T) {
//	tu.SkipLog(t)
//}
