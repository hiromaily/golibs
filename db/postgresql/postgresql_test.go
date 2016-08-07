package postgresql_test

import (
	"flag"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/postgresql"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Postgresql_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Postgresql
//-----------------------------------------------------------------------------
func TestPostgresql(t *testing.T) {
	if *benchFlg == 1 {
		t.Skip("skipping TestPostgresql")
	}

	//if err != nil {
	//	t.Errorf("TestPostgresql error: %s", err)
	//}

}

//-----------------------------------------------------------------------------
//Benchmark
//-----------------------------------------------------------------------------
func BenchmarkPostgresql(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
