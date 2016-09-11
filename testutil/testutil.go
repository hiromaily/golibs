// Package testutil is utility for test
package testutil

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	r "github.com/hiromaily/golibs/runtimes"
	"testing"
)

var (
	//LogFlg is for switch to output log or not
	LogFlg = flag.Int("log", 0, "Log Flg: 0:OFF, 1:ON")
	//BenchFlg is when benchmark test, value is true
	BenchFlg = false
)

// InitializeTest is to run common initial code for test
func InitializeTest(prefix string) {
	flag.Parse()

	//log
	logLevel := lg.InfoStatus
	if *LogFlg == 1 {
		logLevel = lg.DebugStatus
	}
	lg.InitializeLog(logLevel, lg.LogOff, 0, prefix, "/var/log/go/test.log")

	//bench
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		BenchFlg = true
	}
}

// Logf is t.Logf()
//	tu.Logf(t, "%+v", mRet)
func Logf(t *testing.T, format string, args ...interface{}) {
	if *LogFlg == 1 {
		t.Logf(format, args...)

	}
}

// Log is t.Log()
//	tu.Log(t, mRet)
func Log(t *testing.T, args ...interface{}) {
	if *LogFlg == 1 {
		t.Log(args...)
	}
}

// SkipLog is to skip test with func name
func SkipLog(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(2)))
}
