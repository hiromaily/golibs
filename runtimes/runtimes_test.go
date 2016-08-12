package runtimes_test

import (
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	. "github.com/hiromaily/golibs/runtimes"
	"os"
	"testing"
)

type User struct {
	Id   int
	Name string
}

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[RUNTIMES_TEST]", "/var/log/go/test.log")
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
