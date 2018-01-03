package runtimes_test

import (
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/runtimes"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"runtime"
	"runtime/debug"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Runtimes]")
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
// functions
//-----------------------------------------------------------------------------
// CallerDebug is just sample of runtime.Caller
func callerDebug(skip int) {
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(skip)
	lg.Debugf("ok: %t", ok)
	lg.Debugf("programCounter: %v", programCounter)
	lg.Debugf("sourceFileName: %s", sourceFileName)
	lg.Debugf("sourceFileLineNum: %d", sourceFileLineNum)
	lg.Debug("- - - - - - - - - - - - - - - - - - - - - - - - - - - - -")

	//0.1.2...と増える毎に呼び出し元を辿っていく
	//_, file, line, ok = runtime.Caller(calldepth)
	//pc, src, line, ok := runtime.Caller(0)
	//fmt.Println(pc, src, line, ok)
	//runtime.Caller(0)->582751 {GOPATH}/src/github.com/hiromaily/golibs/log/log.go 138 true

	//pc, src, line, ok = runtime.Caller(1)
	//fmt.Println(pc, src, line, ok)
	//8525 {GOPATH}/src/github.com/hiromaily/goweb/ginserver.go 20 true

	//pc, src, line, ok = runtime.Caller(2)
	//fmt.Println(pc, src, line, ok)
	//11667 {GOPATH}/src/github.com/hiromaily/goweb/ginserver.go 100 true

	//PrintStack prints to standard error the stack trace returned by runtime.Stack.
	debug.PrintStack()
}

//-----------------------------------------------------------------------------
// Check
//-----------------------------------------------------------------------------
func TestCallerDebug(t *testing.T) {
	tu.SkipLog(t)

	callerDebug(0)
	callerDebug(1)
	callerDebug(2)
}

func TestArchEnv(t *testing.T) {
	tu.SkipLog(t)

	lg.Debugf("GOOS: %s", runtime.GOOS)     //[mac]darwin
	lg.Debugf("GOARCH: %s", runtime.GOARCH) //[mac]amd64
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCurrentFunc(t *testing.T) {
	//tu.SkipLog(t)
	s := CurrentFunc(1)
	if s != "TestCurrentFunc" {
		t.Errorf("result of CurrentFunc(1) is wrong: %s", s)
	}
	lg.Debugf("CurrentFunc(1) :%s", s)

	b := CurrentFuncV2()
	lg.Debugf("CurrentFunc2() :%s", b)
}

func TestGetStackTrace(t *testing.T) {
	info := GetStackTrace("hiromaily")
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Printf("%02d: [Function]%s [File]%s:%d\n", i, v.FunctionName, v.FileName, v.FileLine)
	}
	// Output
	// 03: [Function]goexit [File]/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
	// 02: [Function]tRunner [File]/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
	// 01: [Function]TestGetStackTrace [File]./golibs/runtimes/runtimes_test.go:102
	// 00: [Function]GetStackTrace [File]./golibs/runtimes/runtimes.go:58
}

func TestTraceAllHistory(t *testing.T) {
	TraceAllHistory(os.Stdout, "hiromaily")
	// Output
	//03: [Function]goexit [File]/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
	//02: [Function]tRunner [File]/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
	//01: [Function]TestTraceAllHistory [File]./golibs/runtimes/runtimes_test.go:115
	//00: [Function]TraceAllHistory [File]./golibs/runtimes/runtimes.go:63}
}
