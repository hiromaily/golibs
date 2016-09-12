package files_test

import (
	. "github.com/hiromaily/golibs/files"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

var basePath = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/testdata/tmpl/pages"

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[FILES]")
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
	ext := []string{"tmpl"}

	files := GetFileList(basePath, ext)
	for _, file := range files {
		lg.Debug(file)
	}
}

func TestGetFileList2(t *testing.T) {
	ext := []string{"tmpl"}

	files := GetFileListSingle(basePath, ext)
	for _, file := range files {
		lg.Debug(file)
	}
}

//-----------------------------------------------------------------------------
// Bench
//-----------------------------------------------------------------------------
func BenchmarkGetFileList(b *testing.B) {
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileList(basePath, ext)
	}
	b.StopTimer()
	//456996 ns/op	  405576 B/op	    1203 allocs/op
}

func BenchmarkGetFileList2(b *testing.B) {
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileListSingle(basePath, ext)
	}
	b.StopTimer()
	//408523 ns/op	  406034 B/op	    1206 allocs/op
}

func BenchmarkGetFileListJIC(b *testing.B) {
	ext := []string{"tmpl"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetFileListJIC(basePath, ext)
	}
	b.StopTimer()
	//388937 ns/op	  405478 B/op	    1202 allocs/op
}
