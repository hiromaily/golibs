package regexp_test

import (
	//lg "github.com/hiromaily/golibs/log"
	"fmt"
	. "github.com/hiromaily/golibs/regexp"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
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
// Initialize
func init() {
	tu.InitializeTest("[Regexp]")
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

func TestReplace(t *testing.T) {
	//remove last filter
	fmt.Println("01:", Replace("/nl/amsterdam/area1/filter-aaa", `^.*\/filter-aaa$`, "$1"))             //
	fmt.Println("02:", Replace("/nl/amsterdam/area1/filter-aaa", `\/filter-aaa$`, "$1"))                // /nl/amsterdam/area1
	fmt.Println("03:", Replace("/nl/amsterdam/area1/filter-aaa", `\/filter-aaa$|\/filter-ccc$|`, "$1")) // /nl/amsterdam/area1
	fmt.Println("04:", Replace("/nl/amsterdam/area1/filter-bbb", `\/filter-aaa$|\/filter-ccc$|`, "$1")) // /nl/amsterdam/area1/filter-bbb
	fmt.Println("05:", Replace("/nl/amsterdam/area1/filter-ccc", `\/filter-aaa$|\/filter-ccc$|`, "$1")) // /nl/amsterdam/area1

	//if wanna remove not only last, in the middle of path,
	fmt.Println("10:", Replace("/nl/amsterdam/area1", `\/area1(/|\z)`, "$1"))          // /nl/amsterdam
	fmt.Println("11:", Replace("/nl/amsterdam/area111", `\/area1(/|\z)`, "$1"))        // /nl/amsterdam/area111 => error
	fmt.Println("12:", Replace("/nl/amsterdam/area1/filter", `\/area1(/|\z)`, "$1"))   // /nl/amsterdam/filter
	fmt.Println("13:", Replace("/nl/amsterdam/area111/filter", `\/area1(/|\z)`, "$1")) // /nl/amsterdam/area111/filter => error

	fmt.Println("14:", Replace("/nl/amsterdam/area1", `\/(?:area1|area2)(/|\z)`, "$1"))        // /nl/amsterdam
	fmt.Println("15:", Replace("/nl/amsterdam/area1/filter", `\/(?:area1|area2)(/|\z)`, "$1")) // /nl/amsterdam/filter
	fmt.Println("16:", Replace("/nl/amsterdam/area2", `\/(?:area1|area2)(/|\z)`, "$1"))        // /nl/amsterdam
	fmt.Println("17:", Replace("/nl/amsterdam/area2/filter", `\/(?:area1|area2)(/|\z)`, "$1")) // /nl/amsterdam/filter

	fmt.Println("20:", Replace("/be/luik/spa/bornevenlig-hotel", `\/(?:bornevenlig-hotel|area2)(/|\z)`, "$1")) // /be/luik/spa

	// /([^/]+)/(area1|area2|area3|area4)$
	fmt.Println("21:", Replace("/be/luik/spa/area1", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))          // /be/luik/spa
	fmt.Println("22:", Replace("/be/luik/spa/area1/filter", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))   // /be/luik/spa => error
	fmt.Println("23:", Replace("/be/luik/spa/aarea1", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1"))         // /be/luik/spa => error
	fmt.Println("24:", Replace("/be/luik/spa/area111/filter", `/([^/]+)/(area1|area2|area3|area4)$`, "/$1")) // /be/luik/spa => error

	fmt.Println("30:", Replace("https://www.hotelspecials.dk/at/week", `(https?://)([^:^/]*)(:\\d*)?(.*)?`, "$1, $2, $3, $4")) // https://, www.hotelspecials.dk, , /at/week
}
