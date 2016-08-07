package heroku_test

import (
	"flag"
	. "github.com/hiromaily/golibs/heroku"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Heroku_TEST]", "/var/log/go/test.log")
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
// Heroku
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
//Benchmark
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
