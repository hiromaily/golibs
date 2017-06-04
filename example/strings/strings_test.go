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

//TODO:It's useful!!
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

//TODO:It's useful!!
func TestHasSuffix(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.HasSuffix("GopherMIX", "MIX")
	lg.Debugf("[1]strings.HasSuffix():%v", ret) //true

	ret = strings.HasSuffix("GopherMIX", "C")
	lg.Debugf("[2]strings.HasSuffix():%v", ret) //false

	ret = strings.HasSuffix("GopherMIX", "ix")
	lg.Debugf("[3]strings.HasSuffix():%v", ret) //false

	ret = strings.HasSuffix("GopherMIX", "")
	lg.Debugf("[4]strings.HasSuffix():%v", ret) //true
}

//TODO:It's useful!!
func TestIndex(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.Index("chicken", "ken")
	lg.Debugf("[1]strings.Index():%v", ret) // 4

	ret = strings.Index("chicken", "den")
	lg.Debugf("[2]strings.Index():%v", ret) // -1
}

func TestIndexAny(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.IndexAny("chicken", "aeiouy")
	lg.Debugf("[1]strings.IndexAny():%v", ret) // 2

	ret = strings.IndexAny("crwth", "aeiouy")
	lg.Debugf("[2]strings.IndexAny():%v", ret) // -1
}

func TestIndexByte(t *testing.T) {
	tu.SkipLog(t)
	//func IndexByte(s string, c byte) int
}

//IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.
//日本語が混じっているのかどうか、調べることができる
func TestIndexFunc(t *testing.T) {
	tu.SkipLog(t)
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	ret := strings.IndexFunc("Hello, 世界", f)
	lg.Debugf("[1]strings.IndexFunc():%v", ret) // 7

	ret = strings.IndexFunc("Hello, world", f)
	lg.Debugf("[1]strings.IndexFunc():%v", ret) // -1
}

func TestIndexRune(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.IndexRune("chicken", 'k')
	lg.Debugf("[1]strings.IndexRune():%v", ret) // 4

	ret = strings.IndexRune("chicken", 'd')
	lg.Debugf("[2]strings.IndexRune():%v", ret) // -1
}

//複数検索文字が存在する時に、その最後に発見されたものの、indexを返す。
func TestLastIndex(t *testing.T) {
	tu.SkipLog(t)
	ret := strings.Index("go gopher", "go")
	lg.Debugf("[0]strings.Index():%v", ret) // 0

	ret = strings.LastIndex("go gopher", "go")
	lg.Debugf("[1]strings.LastIndex():%v", ret) // 3

	ret = strings.LastIndex("go gopher", "rodent")
	lg.Debugf("[2]strings.LastIndex():%v", ret) // -1
}

func TestLastIndexAny(t *testing.T) {
	tu.SkipLog(t)
	//func LastIndexAny(s, chars string) int
}

func TestLastIndexByte(t *testing.T) {
	tu.SkipLog(t)
	//func LastIndexByte(s string, c byte) int
}

func TestLastIndexFunc(t *testing.T) {
	tu.SkipLog(t)
	//func LastIndexFunc(s string, f func(rune) bool) int
}

//TODO:It's useful!!
func TestMap(t *testing.T) {
	tu.SkipLog(t)
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	reverse := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			//return 'A' + (r-'A'+13)%26
			return r ^ 32 //A => a
		case r >= 'a' && r <= 'z':
			//return 'a' + (r-'a'+13)%26
			return r ^ 32 //a => A
		}
		return r
	}

	ret := strings.Map(rot13, "'Twas brillig and the slithy gopher...")
	lg.Debugf("[1]strings.Map():%v", ret) //

	ret = strings.Map(reverse, "'i like golang and devops...")
	lg.Debugf("[2]strings.Map():%v", ret) //
}

//TODO:It's useful!!
func TestRepeat(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Repeat("na", 2)
	lg.Debugf("[1]strings.Repeat():%v", ret) //nana
}

//TODO:It's useful!!
func TestReplace(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Replace("oink oink oink", "k", "ky", 2)
	lg.Debugf("[1]strings.Replace():%v", ret) //oinky oinky oink

	ret = strings.Replace("oink oink oink", "k", "ky", -1)
	lg.Debugf("[2]strings.Replace():%v", ret) //oinky oinky oinky
}

//TODO:It's useful!!
func TestSplit(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Split("a,b,c", ",")
	lg.Debugf("[1]strings.Split():%q", ret) //["a" "b" "c"]

	ret = strings.Split("a man a plan a canal panama", "a ")
	lg.Debugf("[2]strings.Split():%q", ret) //["" "man " "plan " "canal panama"]

	ret = strings.Split(" xyz ", "")
	lg.Debugf("[3]strings.Split():%q", ret) //[" " "x" "y" "z" " "]

	ret = strings.Split("", "Bernardo O'Higgins")
	lg.Debugf("[4]strings.Split():%q", ret) //[""]
}

//SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
func TestSplitAfter(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.SplitAfter("a,b,c", ",")
	lg.Debugf("[1]strings.SplitAfter():%q", ret) //["a," "b," "c"]
}

//with separated stirng
func TestSplitAfterN(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.SplitAfterN("a,b,c,d,e", ",", 1)
	lg.Debugf("[1]strings.SplitAfterN():%q", ret) //["a,b,c,d,e"]

	ret = strings.SplitAfterN("a,b,c,d,e", ",", 2)
	lg.Debugf("[2]strings.SplitAfterN():%q", ret) //["a," "b,c,d,e"]

	ret = strings.SplitAfterN("a,b,c,d,e", ",", 3)
	lg.Debugf("[3]strings.SplitAfterN():%q", ret) //["a," "b," "c,d,e"]

	ret = strings.SplitAfterN("a,b,c,d,e", ",", 4)
	lg.Debugf("[4]strings.SplitAfterN():%q", ret) //["a," "b," "c," "d,e"]

	ret = strings.SplitAfterN("a,b,c,d,e", ",", 5)
	lg.Debugf("[5]strings.SplitAfterN():%q", ret) //["a," "b," "c," "d," "e"]
}

//without separated stirng
func TestSplitN(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.SplitN("a,b,c,d,e", ",", 1)
	lg.Debugf("[1]strings.SplitN():%q", ret) //["a,b,c,d,e"]

	ret = strings.SplitN("a,b,c,d,e", ",", 2)
	lg.Debugf("[2]strings.SplitN():%q", ret) //["a" "b,c,d,e"]

	ret = strings.SplitN("a,b,c,d,e", ",", 3)
	lg.Debugf("[3]strings.SplitN():%q", ret) //["a" "b" "c,d,e"]

	ret = strings.SplitN("a,b,c,d,e", ",", 4)
	lg.Debugf("[4]strings.SplitN():%q", ret) //["a" "b" "c" "d,e"]

	ret = strings.SplitN("a,b,c,d,e", ",", 5)
	lg.Debugf("[5]strings.SplitN():%q", ret) //["a" "b" "c" "d" "e"]
}

//Title returns a copy of the string s with all Unicode letters that begin words mapped to their title case.
func TestTitle(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Title("her royal highness")
	lg.Debugf("[1]strings.Title():%s", ret) //Her Royal Highness

	ret = strings.Title("This is a pen")
	lg.Debugf("[2]strings.Title():%s", ret) //This Is A Pen
}

func TestToLower(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.ToLower("Gopher Go Go")
	lg.Debugf("[1]strings.ToLower():%s", ret) //gopher go go
}

func TestToTitle(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.ToTitle("loud noises")
	lg.Debugf("[1]strings.ToTitle():%s", ret) //LOUD NOISES
}

func TestToUpper(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.ToUpper("Gopher Go Go")
	lg.Debugf("[1]strings.ToUpper():%s", ret) //GOPHER GO GO
}

//TODO:It's useful!!
//it can remove only both ends of string by cutset letters.
func TestTrim(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.Trim(" !!! Achtung! Achtung! !!! ", "! ")
	lg.Debugf("[1]strings.Trim():%s", ret) //Achtung! Achtung

	ret = strings.Trim("aabbccddee", "ab")
	lg.Debugf("[2]strings.Trim():%s", ret) //ccddee

	ret = strings.Trim("aabbccddee", "bd")
	lg.Debugf("[3]strings.Trim():%s", ret) //aabbccddee

	ret = strings.Trim("aa bb cc dd ee", "be")
	lg.Debugf("[4]strings.Trim():%s", ret) //aa bb cc dd
}

func TestTrimLeft(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! ")
	lg.Debugf("[1]strings.TrimLeft():%s", ret) //Achtung! Achtung! !!!

	ret = strings.TrimLeft("aabbccddee", "ab")
	lg.Debugf("[2]strings.TrimLeft():%s", ret) //ccddee

	ret = strings.TrimLeft("aabbccddee", "bd")
	lg.Debugf("[3]strings.TrimLeft():%s", ret) //aabbccddee

	ret = strings.TrimLeft("aa bb cc dd ee", "be")
	lg.Debugf("[4]strings.TrimLeft():%s", ret) //aa bb cc dd ee
}

func TestTrimRight(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.TrimRight(" !!! Achtung! Achtung! !!! ", "! ")
	lg.Debugf("[1]strings.TrimRight():%s", ret) // !!! Achtung! Achtung

	ret = strings.TrimRight("aabbccddee", "ab")
	lg.Debugf("[2]strings.TrimRight():%s", ret) //aabbccddee

	ret = strings.TrimRight("aabbccddee", "bd")
	lg.Debugf("[3]strings.TrimRight():%s", ret) //aabbccddee

	ret = strings.TrimRight("aa bb cc dd ee", "be")
	lg.Debugf("[4]strings.TrimRight():%s", ret) //aa bb cc dd
}

func TestTrimSpace(t *testing.T) {
	tu.SkipLog(t)

	ret := strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")
	lg.Debugf("[1]strings.TrimSpace():%s", ret) //a lone gopher
}

func TestTrimPrefix(t *testing.T) {
	tu.SkipLog(t)
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "Goodbye,")
	lg.Debugf("[1]strings.TrimPrefix():%s", s) //, world!
	s = strings.TrimPrefix(s, "Howdy,")
	lg.Debugf("[2]strings.TrimPrefix():%s", s) //, world!
	s = strings.TrimPrefix(s, ", ")
	lg.Debugf("[2]strings.TrimPrefix():%s", s) //world!
}

func TestTrimSuffix(t *testing.T) {
	//tu.SkipLog(t)

	var s = "Hello, goodbye, etc!"
	s = strings.TrimSuffix(s, "goodbye, etc!")
	lg.Debugf("[1]strings.TrimSuffix():%s", s) //Hello,
	s = strings.TrimSuffix(s, "planet")
	lg.Debugf("[2]strings.TrimSuffix():%s", s) //Hello,
	s = strings.TrimSuffix(s, ", ")
	lg.Debugf("[2]strings.TrimSuffix():%s", s) //Hello
}

//func TestTrimFunc(t *testing.T) {
//	//TrimFunc(s string, f func(rune) bool)
//}
