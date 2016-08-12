package files_test

import (
	. "github.com/hiromaily/golibs/files"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var benchFlg bool = false

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	//Here is [slower] than included file's init()
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
func TestGetFileList(t *testing.T) {
	basePath := "/Users/hy/work/go/src/github.com/hiromaily/go-gin-wrapper/templates/pages"
	ext := []string{"tmpl"}

	files := GetFileList(basePath, ext)
	for _, file := range files {
		t.Log(file)
	}
}

func TestGetFileList2(t *testing.T) {
	basePath := "/Users/hy/work/go/src/github.com/hiromaily/go-gin-wrapper/templates/pages"
	ext := []string{"tmpl"}

	files := GetFileListSingle(basePath, ext)
	for _, file := range files {
		t.Log(file)
	}
}

//-----------------------------------------------------------------------------
// Bench
//-----------------------------------------------------------------------------
func BenchmarkGetFileList(b *testing.B) {
	basePath := "/Users/hy/work/go/src/github.com/hiromaily/go-gin-wrapper/templates/pages"
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileList(basePath, ext)
	}
	b.StopTimer()
	//456996 ns/op	  405576 B/op	    1203 allocs/op
}

func BenchmarkGetFileList2(b *testing.B) {
	basePath := "/Users/hy/work/go/src/github.com/hiromaily/go-gin-wrapper/templates/pages"
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileListSingle(basePath, ext)
	}
	b.StopTimer()
	//408523 ns/op	  406034 B/op	    1206 allocs/op
}

func BenchmarkGetFileListJIC(b *testing.B) {
	basePath := "/Users/hy/work/go/src/github.com/hiromaily/go-gin-wrapper/templates/pages"
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileListJIC(basePath, ext)
	}
	b.StopTimer()
	//388937 ns/op	  405478 B/op	    1202 allocs/op
}
