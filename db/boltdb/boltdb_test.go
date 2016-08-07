package boltdb_test

import (
	"flag"
	. "github.com/hiromaily/golibs/db/boltdb"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
	path     = flag.String("fp", "", "BoltDB File Path")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[BoltDB_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
		New(*path)
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	if *path == "" {
		os.Exit(1)
		return
	}

	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

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
//Benchmark
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
