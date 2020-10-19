package regexp_test

import (
	//lg "github.com/hiromaily/golibs/log"
	"fmt"
	"os"
	"testing"

	. "github.com/hiromaily/golibs/regexp"
	tu "github.com/hiromaily/golibs/testutil"
)

//http://ashitani.jp/golangtips/tips_regexp.html

var regExpData = []struct {
	reg         string
	str         string
	expectation bool
}{
	{`a*c`, "abc", true},
	{`a*c`, "ac", true},
	{`a*c`, "aaaaaac", true},
	{`a*c`, "c", true},
	{`a*c`, "abccccc", true},
	{`a*c`, "abd", false},
	{`a+c`, "ac", true},
	{`a+c`, "aaaaaac", true},
	{`a+c`, "abc", false},
	{`a+c`, "c", false},
	{`a+c`, "abccccc", false},
	{`a+c`, "abd", false},
	{`a?c`, "abc", true},
	{`a?c`, "ac", true},
	{`a?c`, "aaaaaac", true},
	{`a?c`, "c", true},
	{`a?c`, "abccccc", true},
	{`a?c`, "abd", false},
	//
	{`[ABZ]`, "A", true},
	{`[ABZ]`, "Z", true},
	{`[ABZ]`, "Q", false},
	{`[0-9]`, "5", true},
	{`[0-9]`, "A", false},
	{`[A-Z]`, "A", true},
	{`[A-Z]`, "5", false},
	{`[A-Z]`, "a", false},
	{`[^0-9]`, "A", true},
	{`[^0-9]`, "5", false},
	//
	{`^[\\.].*$`, ".git", true},
	{`^[\\.].*$`, ".idea", true},
	{`^[\\.].*$`, "..new", true},
	{`^[\\.].*$`, "folder", false},
	{`^[\\.].*$`, "folder.zip", false},
	//
	{`^.*\.go$|^.*\.php$|^.*\.js$|^.*\.py$|^.*\.txt$`, "abc.go", true},
	{`^.*\.go$|^.*\.php$|^.*\.js$|^.*\.py$|^.*\.txt$`, "abc_xx.go", true},
	{`^.*\.go$|^.*\.php$|^.*\.js$|^.*\.py$|^.*\.txt$`, "ooo_qq.php", true},
	//
	{`^http(s)?:\/\/`, "http://google.com", true},
	{`^http(s)?:\/\/`, "https://google.com", true},
	{`^-test.bench`, "-test.bench=.", true},
	//
	{`^.*\/filter$`, "http://www.test.com/de/berlin/filter", true},
	{`^.*\/filter$`, "http://www.test.com/de/berlin/filter", true},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Regexp]")
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
// Check
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCheckRegexp(t *testing.T) {
	for idx, tt := range regExpData {
		bRet := CheckRegexp(tt.reg, tt.str)
		if bRet != tt.expectation {
			t.Errorf("[%d] Result of [%s] by reg[%s] is %v", idx, tt.str, tt.reg, bRet)
		}
	}
}

func TestIsInvisiblefile(t *testing.T) {
	if !IsInvisiblefile(".git") {
		t.Errorf("[01]IsInvisiblefile() doens't work yet")
	}
	if IsInvisiblefile("git") {
		t.Errorf("[02]IsInvisiblefile() doens't work yet")
	}
}

func TestIsGoFile(t *testing.T) {
	if !IsGoFile("aaaa.go") {
		t.Errorf("[01]IsGoFile() doens't work yet")
	}
	if IsGoFile("bbb.txt") {
		t.Errorf("[02]IsGoFile() doens't work yet")
	}
}

func TestIsTmplFile(t *testing.T) {
	if !IsTmplFile("aaaa.tmpl") {
		t.Errorf("[05]IsTmplFile() doens't work yet")
	}
	if IsTmplFile("bbb.html") {
		t.Errorf("[06]IsTmplFile() doens't work yet")
	}
}

func TestIsExtFile(t *testing.T) {
	if !IsExtFile("abcde.go", "go") {
		t.Errorf("[07]IsExtFile() doens't work yet")
	}
	if IsExtFile("index.thml", "tmpl") {
		t.Errorf("[08]IsExtFile() doens't work yet")
	}
}

func TestIsHeaderURL(t *testing.T) {
	if !IsHeaderURL("http://google.com/") {
		t.Errorf("[09]IsHeaderURL() doens't work yet")
	}
	if !IsHeaderURL("https://google.com/") {
		t.Errorf("[10]IsHeaderURL() doens't work yet")
	}
	if IsHeaderURL("httpps://google.com/") {
		t.Errorf("[11]IsHeaderURL() doens't work yet")
	}
	if IsHeaderURL("https:://google.com/") {
		t.Errorf("[12]IsHeaderURL() doens't work yet")
	}
}

func TestIsBenchTest(t *testing.T) {
	if !IsBenchTest("-test.bench=.") {
		t.Errorf("[13]IsBenchTest doens't work yet")
	}
}

func TestIsStaticFile(t *testing.T) {
	testOKData := []string{"aaa.jpg", "bbb.png", "cccc.js", "abd.woff"}
	testNGData := []string{"/", "bbb/png", "cccc/js/", "/abd/woff/", "/abd/woff/gggg"}
	for idx, tt := range testOKData {
		if !IsStaticFile(tt) {
			t.Errorf("[%d]IsStaticFile()OK data doens't work yet.", idx)
		}
	}

	for idx, tt := range testNGData {
		if IsStaticFile(tt) {
			t.Errorf("[%d]IsStaticFile()NG data doens't work yet.", idx)
		}
	}
}

func TestReplaceResult(t *testing.T) {
	// for removing last
	fmt.Println(Replace("/last", `\/last$|\/first$|`, "$1"))
	// Output:

	// for removing any positions
	fmt.Println(Replace("/nl/amsterdam/area2/filter", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam/filter
}

func TestReplace01(t *testing.T) {
	// Remove only last path can be target.

	fmt.Println(Replace("/nl/amsterdam/area1/filter-aaa", `\/filter-aaa$`, "$1"))
	// Output: /nl/amsterdam/area1

	fmt.Println(Replace("/nl/amsterdam/area1/filter-aaa", `\/filter-aaa$|\/filter-ccc$|`, "$1"))
	// Output: /nl/amsterdam/area1

	fmt.Println(Replace("/nl/amsterdam/area1/filter-bbb", `\/filter-aaa$|\/filter-ccc$|`, "$1"))
	// Output: /nl/amsterdam/area1/filter-bbb

	fmt.Println(Replace("/nl/amsterdam/area1/filter-ccc", `\/filter-aaa$|\/filter-ccc$|`, "$1"))
	// Output: /nl/amsterdam/area1

	fmt.Println(Replace("/nl/amsterdam/area1/filter-ccc/last", `\/filter-aaa$|\/filter-ccc$|`, "$1"))
	// Output: /nl/amsterdam/area1/filter-ccc/last

	fmt.Println(Replace("/last", `\/last$|\/first$|`, "$1"))
	// Output:
}

func TestReplace02(t *testing.T) {
	// This pattern seems that only last path can be target.
	// And multiple targets can be set

	fmt.Println(Replace("/be/luik/spa/area1", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))
	// Output: /be/luik/spa

	fmt.Println(Replace("/be/luik/spa/area1/filter", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))
	// Output: /be/luik/spa => error

	fmt.Println(Replace("/be/luik/spa/Xarea1", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))
	// Output: /be/luik/spa => error

	fmt.Println(Replace("/be/luik/spa/area123/filter", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))
	// Output: /be/luik/spa => error

	fmt.Println(Replace("/last", `/([^/]+)/(last|first)$`, "/$1"))
	// Output: error => this result indicates this regexp is risky.
}

func TestReplace03(t *testing.T) {
	// Both last and in the middle of path can be target

	fmt.Println(Replace("/nl/amsterdam/area1", `\/area1(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	fmt.Println(Replace("/nl/amsterdam/area111", `\/area1(/|\z)`, "$1"))
	// Output: /nl/amsterdam/area111 => error

	fmt.Println(Replace("/nl/amsterdam/area1/filter", `\/area1(/|\z)`, "$1"))
	// Output: /nl/amsterdam/filter

	fmt.Println(Replace("/nl/amsterdam/area111/filter", `\/area1(/|\z)`, "$1"))
	// Output: /nl/amsterdam/area111/filter => error
}

func TestReplace04(t *testing.T) {
	// Both last and in the middle of path can be target
	// And multiple targets can be set

	fmt.Println(Replace("/nl/amsterdam/area1", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	fmt.Println(Replace("/nl/amsterdam/area1/filter", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam/filter

	fmt.Println(Replace("/nl/amsterdam/area2", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	fmt.Println(Replace("/nl/amsterdam/area2/filter", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam/filter

	fmt.Println(Replace("/nl/amsterdam/area2/area1", `\/(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam/area1

}

func TestReplace05(t *testing.T) {
	// wanna apply recursively.
	fmt.Println(Replace("/nl/amsterdam/area2/area1", `\/+(?:area1|area2).+(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	//TODO: what is wrong??
	fmt.Println(Replace("/nl/amsterdam/area2", `\/+(?:area1|area2).+(/|\z)`, "$1"))
	// Output: error

	//it's came from TestReplace04
	fmt.Println(Replace("/nl/amsterdam/area2", `\/+(?:area1|area2)(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	//TODO:ongoing
	fmt.Println(Replace2("/nl/amsterdam/area2", `\/+(?:area1|area2)+(/|\z)`, "$1"))
	// Output: /nl/amsterdam

	fmt.Println(Replace2("/nl/amsterdam/area2/area1", `(?:/(?:area1|area2))+(/|\z)`, "$1"))
	//

	fmt.Println(Replace2("/nl/amsterdam/area2", `(?:/(?:area1|area2))+(/|\z)`, "$1"))
	// Output: /nl/amsterdam
	//(?:/(?:area1|area2))+(/|$)
}

func TestReplace06(t *testing.T) {
	fmt.Println(Replace("https://www.hotelspecials.dk/at/week", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$1"))
	// Output: https://

	fmt.Println(Replace("https://www.hotelspecials.dk/at/week", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$2"))
	// Output: www.hotelspecials.dk

	fmt.Println(Replace("https://www.hotelspecials.dk/at/week", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$3"))
	// Output:

	fmt.Println(Replace("https://www.hotelspecials.dk/at/week", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$4"))
	// Output: /at/week

	fmt.Println(Replace("https://www.hotelspecials.dk/at/week?param=1", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$4"))
	// Output: /at/week?param=1

	//remove path
	fmt.Println(Replace("/at/week?param=1&param2=2", `([^\?]+)(\?.*)?`, "$1"))
	// Output: /at/week

	fmt.Println(Replace("/at/week?param=1&param2=2", `([^\?]+)(\?.*)?`, "$2"))
	// Output: ?param=1&param2=2
}

func TestReplace07(t *testing.T) {
	fmt.Println(Replace("/dk/syddanmark/soenderborg/weekendophold", `^/dk/syddanmark/soenderborg(/|\z)`, "/dk/syddanmark/sonderborg$1"))
	// Output: /dk/syddanmark/sonderborg/weekendophold

	fmt.Println(Replace("/bohuslaen/romance", `^/bohuslaen(/|\z)`, "/se/bohusla$1"))
	// Output: /se/bohusla/romance

	fmt.Println(Replace("/se/bohuslaen/romance", `^/(?:bohuslaen|se/bohuslaen)(/|\z)`, "/se/bohusla$1"))
	// Output: /se/bohusla/romance
}
