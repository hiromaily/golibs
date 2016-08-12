package boltdb_test

import (
	"flag"
	. "github.com/hiromaily/golibs/db/boltdb"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var (
	path          = flag.String("fp", "", "BoltDB File Path")
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	flag.Parse()

	if *path == "" {
		os.Exit(1)
		return
	}

	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[BoltDB_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
	if !benchFlg {
		New(*path)
	}
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
func TestSetAndGetData01(t *testing.T) {

	data := "testdayo"
	err := Set("table01", "key01", []byte(data))
	if err != nil {
		t.Errorf("TestSetAndGetData01 Set error: %s", err)
	}

	b, err := Get("table01", "key01")
	if err != nil {
		t.Errorf("TestSetAndGetData01 Get error: %s", err)
	}
	if string(b) != "testdayo" {
		t.Errorf("TestSetAndGetData01 data: %s", string(b))
	}
	//t.Logf("data:%s", string(b))
	Close()
	// ns/op
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkSetData01(b *testing.B) {
	New(*path)

	data := "testdayo"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Set("table01", "key01", []byte(data))
	}
	b.StopTimer()

	Close()
	//123347 ns/op
}

func BenchmarkGetData01(b *testing.B) {
	New(*path)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Get("table01", "key01")
	}
	b.StopTimer()

	Close()
	//1073 ns/op
}
