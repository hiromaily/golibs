package runtimes

import (
	"bytes"
	lg "github.com/hiromaily/golibs/log"
	"runtime"
	"runtime/debug"
	"strings"
)

//Caller
func CallerDebug(skip int) {
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

func ArchEnv() {
	lg.Debugf("GOOS: %s", runtime.GOOS)     //[mac]darwin
	lg.Debugf("GOARCH: %s", runtime.GOARCH) //[mac]amd64
}

func CurrentFunc(skip int) string {
	programCounter, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	sl := strings.Split(runtime.FuncForPC(programCounter).Name(), ".")
	return sl[len(sl)-1]
}

func CurrentFunc2() []byte {
	b := make([]byte, 250)
	b = b[:runtime.Stack(b, false)]
	for i := 0; i < 3; i++ {
		j := bytes.IndexByte(b, '\n')
		if j < 0 {
			return nil
		}

		b = b[j+1:]
	}
	i := bytes.IndexByte(b, '(')
	if i < 0 {
		return nil
	}

	return b[:i]
}
