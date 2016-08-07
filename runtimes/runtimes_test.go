package runtimes_test

import (
	"flag"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/runtimes"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

type User struct {
	Id   int
	Name string
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[RUNTIMES_TEST]", "/var/log/go/test.log")
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
// Runtimes
//-----------------------------------------------------------------------------
func TestCallerDebug(t *testing.T) {
	t.Skip("skipping TestCallerDebug")

	CallerDebug(0)
	//CallerDebug(1)
	//CallerDebug(2)
}

func TestArchEnv(t *testing.T) {
	t.Skip("skipping TestArchEnv")
	ArchEnv()
}

func TestCurrentFunc(t *testing.T) {
	//t.Skip("skipping TestCurrentFunc")
	s := CurrentFunc(1)
	t.Logf("[CurrentFunc1] func is %s", s)

	b := CurrentFunc2()
	t.Logf("[CurrentFunc2] func is %s", b)
	//func is command-line-arguments_test.TestCurrentFunc
}
