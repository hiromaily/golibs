package heroku_test

import (
	. "github.com/hiromaily/golibs/heroku"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Heroku_TEST]", "/var/log/go/test.log")
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
func TestGetMySQLInfo(t *testing.T) {
	host, dbname, user, pass, err := GetMySQLInfo("mysql://be2ebea7cda583:49eef93c@us-cdbr-iron-east-04.cleardb.net/heroku_aa95a7f43af0311?reconnect=true")

	if err != nil {
		//unexpected EOF
		t.Errorf("TestGetMySQLInfo error: %s", err)
	}
	t.Logf("host: %s", host)
	t.Logf("dbname: %s", dbname)
	t.Logf("user: %s", user)
	t.Logf("pass: %s", pass)
}

func TestGetRedisInfo(t *testing.T) {
	host, pass, port, err := GetRedisInfo("redis://h:pf217irr4gts39d29o0012bghsi@ec2-50-19-83-130.compute-1.amazonaws.com:20819")

	if err != nil {
		t.Errorf("TestGetRedisInfo error: %s", err)
	}
	t.Logf("host: %s", host)
	t.Logf("pass: %s", pass)
	t.Logf("port: %d", port)
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkHeroku(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
