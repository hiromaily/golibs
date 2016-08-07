package os_test

import (
	"flag"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/os"
	"os"
	"strings"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[OS_TEST]", "/var/log/go/test.log")
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
// OS
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
