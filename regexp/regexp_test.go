package regexp_test

import (
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/regexp"
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
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Regexp_TEST]", "/var/log/go/test.log")
}

func teardown() {
}

// Initialize
func TestMain(m *testing.M) {

	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Regexp
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
}
