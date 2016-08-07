package exec_test

import (
	"flag"
	. "github.com/hiromaily/golibs/exec"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[EXEC_TEST]", "/var/log/go/test.log")
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
// Exec
//-----------------------------------------------------------------------------
func TestExec(t *testing.T) {
	t.Skip("skipping TestExec")

	err := Exec("ls", "-al")
	if err != nil {
		t.Errorf("TestExec[01] error: %s", err)
	}

	err = Exec("ls", "-a -l")
	if err != nil {
		t.Errorf("TestExec[02] error: %s", err)
	}
}

func TestGetExec(t *testing.T) {
	t.Skip("skipping TestGetExec")

	result, err := GetExec("ls", "-al")
	if err != nil {
		t.Errorf("TestExec[01] error: %s", err)
	}
	t.Logf("GetExec ls -al: %v", result)

	result, err = GetExec("ls", "-a -l")
	if err != nil {
		t.Errorf("TestExec[02] error: %s", err)
	}
	t.Logf("GetExec ls -a -l: %v", result)
}

func TestCurl(t *testing.T) {
	t.Skip("skipping TestCurl")

	option := `'http://www.google.co.jp' -H 'Cookie: xxxx=uuuu"`
	result, err := GetExec("curl", option)
	if err != nil {
		t.Errorf("TestCurl: error: %s", err)
	}
	t.Logf("TestCurl xxxx: %v", result)
}

func TestExecSh(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	result, err := GetExec(goPath+"/src/github.com/hiromaily/golibs/exec/sh/test.sh", "")
	if err != nil {
		t.Errorf("TestExecSh: error: %s", err)
	}
	t.Logf("TestExecSh xxxx: %v", result)
}

func TestAsyncExecSh(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	err := AsyncExec(goPath+"/src/github.com/hiromaily/golibs/exec/sh/test.sh", "")
	if err != nil {
		t.Errorf("TestAsyncExecSh: error: %s", err)
	}
}

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
