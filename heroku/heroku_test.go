package heroku_test

import (
	"os"
	"testing"

	. "github.com/hiromaily/golibs/heroku"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[HEROKU]")
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
		t.Errorf("GetMySQLInfo error: %s", err)
	}
	lg.Debugf("host: %s", host)
	lg.Debugf("dbname: %s", dbname)
	lg.Debugf("user: %s", user)
	lg.Debugf("pass: %s", pass)
}

func TestGetRedisInfo(t *testing.T) {
	host, pass, port, err := GetRedisInfo("redis://h:pf217irr4gts39d29o0012bghsi@ec2-50-19-83-130.compute-1.amazonaws.com:20819")

	if err != nil {
		t.Errorf("GetRedisInfo error: %s", err)
	}
	lg.Debugf("host: %s", host)
	lg.Debugf("pass: %s", pass)
	lg.Debugf("port: %d", port)
}

func TestGetMongoInfo(t *testing.T) {
	host, dbname, user, pass, port, err := GetMongoInfo("mongodb://heroku_7lbnd77m:7r8f631nv2idt0fhj9ok9714j9@ds161495.mlab.com:61495/heroku_7lbnd77m")

	if err != nil {
		//unexpected EOF
		t.Errorf("GetMongoInfo error: %s", err)
	}
	lg.Debugf("host: %s", host)
	lg.Debugf("dbname: %s", dbname)
	lg.Debugf("user: %s", user)
	lg.Debugf("pass: %s", pass)
	lg.Debugf("port: %d", port)
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
