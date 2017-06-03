package files_test

import (
	"bufio"
	"fmt"
	. "github.com/hiromaily/golibs/files"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"io/ioutil"
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
// Example
//-----------------------------------------------------------------------------
func TestWriteFile1(t *testing.T) {
	//tu.SkipLog(t)

	//1
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./test1.txt", d1, 0644)
	if err != nil {
		lg.Error(err)
	}

	//2 (overwrite)
	d1 = []byte("hello\ngo\n")
	err = ioutil.WriteFile("./test1.txt", d1, 0644)
	if err != nil {
		lg.Error(err)
	}
}

func TestWriteFile2(t *testing.T) {
	//tu.SkipLog(t)
	f, err := os.Create("./test2.txt")
	if err != nil {
		lg.Error(err)
	}
	defer f.Close()

	//1
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		lg.Error(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	//2
	n3, err := f.WriteString("writes\n")
	if err != nil {
		lg.Error(err)
	}
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
}

func TestWriteFile3(t *testing.T) {
	//tu.SkipLog(t)

	f, err := os.Create("./test3.txt")
	if err != nil {
		lg.Error(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	//1
	n4, err := w.WriteString("buffered\n")
	if err != nil {
		lg.Error(err)
	}
	fmt.Printf("wrote %d bytes\n", n4)

	//2
	n4, err = w.WriteString("buffered\n")
	if err != nil {
		lg.Error(err)
	}
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetFileList(t *testing.T) {
	tu.SkipLog(t)

	ext := []string{"tmpl"}

	files := GetFileList(basePath, ext)
	for _, file := range files {
		lg.Debug(file)
	}
}

func TestGetFileList2(t *testing.T) {
	tu.SkipLog(t)

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
