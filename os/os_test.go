package os_test

import (
	"os"
	"strings"
	"testing"

	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/os"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[OS]")
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
// Check
//-----------------------------------------------------------------------------
func TestOSHost(t *testing.T) {
	//tu.SkipLog(t)

	host, _ := os.Hostname()
	//centos7
	//hy-MacBook-Pro.local

	lg.Debugf("os.Hostname() host: %s", host)
}

func TestEnv(t *testing.T) {
	//#1 all Environment Variables can be got
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		lg.Debug(pair[0], " : ", pair[1])
	}

	//#2
	os.Setenv("TEST_FLG", "1")
	flg := os.Getenv("TEST_FLG")
	lg.Debugf("TEST_FLG is %s\n", flg)

	// all Environment Variables can be removed.
	os.Clearenv()
	flg = os.Getenv("TEST_FLG")
	lg.Debugf("After Clearenv(), TEST_FLG is %s\n", flg)
}

func TestGetArgs(t *testing.T) {
	lg.Debugf("GetArgs(1): %s", GetArgs(1))
	//-test.v=true
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestAddParam(t *testing.T) {
	key := "abcdef=100"
	if FindParam(key) {
		t.Errorf("[01:FindParam] Result is wrong.")
	}

	AddParam(key)
	if !FindParam(key) {
		t.Errorf("[02:FindParam or AddParam] Result is wrong.")
	}
}
