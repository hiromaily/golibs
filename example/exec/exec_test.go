package exec_test

import (
	. "github.com/hiromaily/golibs/example/exec"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
	"os/exec"
	"strings"
	"log"
	"fmt"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[EXEC]")
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
func TestExec(t *testing.T) {
	tu.SkipLog(t)

	err := Exec("ls", "-al")
	if err != nil {
		t.Errorf("TestExec[01] error: %s", err)
	}

	err = Exec("ls", "-a -l")
	if err != nil {
		t.Errorf("TestExec[02] error: %s", err)
	}
}

func TestExecParams(t *testing.T) {
	//tu.SkipLog(t)
	//cmd := "./example/exec/tool/cmdtool"
	var cmd *exec.Cmd
	cmdName := "./tool/cmdtool"
	strParam := "-s bbbb -n 999 -b true"

	params := strings.Split(strParam, " ")

	cmd = exec.Command(cmdName, params...)

	//stderr, err := cmd.StderrPipe()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err := cmd.Start()
	//if err != nil {
	//	log.Fatal(err)
	//}

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func TestGetExec(t *testing.T) {
	tu.SkipLog(t)

	result, err := GetExec("ls", "-al")
	if err != nil {
		t.Errorf("TestExec[01] error: %s", err)
	}
	lg.Debugf("GetExec ls -al: %v", result)

	result, err = GetExec("ls", "-a -l")
	if err != nil {
		t.Errorf("TestExec[02] error: %s", err)
	}
	lg.Debugf("GetExec ls -a -l: %v", result)
}

func TestCurl(t *testing.T) {
	//TODO:work in progress
	tu.SkipLog(t)

	//option := `'http://www.google.co.jp' -H 'Cookie: xxxx=uuuu'`
	option := `'http://www.yahoo.co.jp/'`
	result, err := GetExec("curl", option)
	if err != nil {
		t.Errorf("TestCurl: error: %s", err)
	}
	lg.Debugf("TestCurl xxxx: %v", result)
}

func TestExecSh(t *testing.T) {
	tu.SkipLog(t)

	goPath := os.Getenv("GOPATH")
	result, err := GetExec(goPath+"/src/github.com/hiromaily/golibs/example/exec/sh/test.sh", "")
	if err != nil {
		t.Errorf("TestExecSh: error: %s", err)
	}
	t.Logf("TestExecSh xxxx: %v", result)
}

func TestAsyncExecSh(t *testing.T) {
	tu.SkipLog(t)

	goPath := os.Getenv("GOPATH")
	err := AsyncExec(goPath+"/src/github.com/hiromaily/golibs/example/exec/sh/test.sh", "")
	if err != nil {
		t.Errorf("TestAsyncExecSh: error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Bench
//-----------------------------------------------------------------------------
func BenchmarkExec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		_ = Exec("ls", "-al")
		//
	}
	b.StopTimer()
}

func BenchmarkGetExec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		_, _ = GetExec("ls", "-al")
		//
	}
	b.StopTimer()
}

func BenchmarkAsyncExec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		_ = AsyncExec("ls", "-al")
		//
	}
	b.StopTimer()
}
