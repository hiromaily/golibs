package regexp_test

import (
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	. "github.com/hiromaily/golibs/regexp"
	"os"
	"testing"
)

//http://ashitani.jp/golangtips/tips_regexp.html

var (
	benchFlg bool = false
)

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
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Regexp_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
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
func TestRegexp(t *testing.T) {
	for idx, tt := range regExpData {
		bRet := CheckRegexp(tt.reg, tt.str)
		if bRet != tt.expectation {
			t.Errorf("[%d] Result of [%s] by reg[%s] is %v", idx, tt.str, tt.reg, bRet)
		}
	}
}

func TestRegexp2(t *testing.T) {
	if !IsInvisiblefile(".git") {
		t.Errorf("[01]IsInvisiblefile() doens't work yet")
	}
	if IsInvisiblefile("git") {
		t.Errorf("[02]IsInvisiblefile() doens't work yet")
	}
	if !IsGoFile("aaaa.go") {
		t.Errorf("[03]IsGoFile() doens't work yet")
	}
	if IsGoFile("bbb.txt") {
		t.Errorf("[04]IsGoFile() doens't work yet")
	}
	if !IsTmplFile("aaaa.tmpl") {
		t.Errorf("[05]IsTmplFile() doens't work yet")
	}
	if IsTmplFile("bbb.html") {
		t.Errorf("[06]IsTmplFile() doens't work yet")
	}
	if !IsExtFile("abcde.go", "go") {
		t.Errorf("[07]IsExtFile() doens't work yet")
	}
	if IsExtFile("index.thml", "tmpl") {
		t.Errorf("[08]IsExtFile() doens't work yet")
	}
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
	if !IsBenchTest("-test.bench=.") {
		t.Errorf("[13]IsBenchTest doens't work yet")
	}
}

func TestRegexp3(t *testing.T) {
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
