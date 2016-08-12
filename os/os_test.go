package os_test

import (
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/os"
	"os"
	"strings"
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
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[OS_TEST]", "/var/log/go/test.log")
	if FindParam("-test.bench") {
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
func TestOSHost(t *testing.T) {

	host := GetOS()
	t.Logf("TestOSHost[01] host: %s", host)
}

func TestEnv(t *testing.T) {
	//#1 all Environment Variables can be got
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		t.Logf(pair[0], pair[1])
	}

	//#2
	os.Setenv("TEST_FLG", "1")
	flg := os.Getenv("TEST_FLG")
	t.Logf("TEST_FLG is %s\n", flg)

	// all Environment Variables can be removed.
	os.Clearenv()
	flg = os.Getenv("TEST_FLG")
	t.Logf("After Clearenv(), TEST_FLG is %s\n", flg)
}

func TestGetArgs(t *testing.T) {
	t.Log(GetArgs(1))
}

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
