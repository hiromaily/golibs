package runtimes

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

var (
	re = regexp.MustCompile(`^(\S.+)\.(\S.+)$`)
)

type CallerInfo struct {
	PackageName  string
	FunctionName string
	FileName     string
	FileLine     int
}

func dump() (callerInfo []*CallerInfo) {
	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i) // https://golang.org/pkg/runtime/#Caller
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)

		_fn := re.FindStringSubmatch(fn.Name())
		callerInfo = append(callerInfo, &CallerInfo{
			PackageName:  _fn[1],
			FunctionName: _fn[2],
			FileName:     fileName,
			FileLine:     fileLine,
		})
	}
	return
}

func TraceAllHistory() {
	info := dump()
	for i := len(info) - 1; i > -1; i-- {
		v := info[i]
		fmt.Printf("%02d: %s.%s@%s:%d\n", i, v.PackageName, v.FunctionName, v.FileName, v.FileLine)
	}
}

// CurrentFunc is to get current func name
func CurrentFunc(skip int) string {
	programCounter, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	sl := strings.Split(runtime.FuncForPC(programCounter).Name(), ".")
	return sl[len(sl)-1]
}

// CurrentFuncV2 is to get current func name
func CurrentFuncV2() []byte {
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
