package log_test

import (
	"os"
	"testing"

	. "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[LOG]")
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
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestNewLog(t *testing.T) {
	logObj := New(DebugStatus, NoDateNoFile, "[LOG_NEW_TEST]", "/var/log/go/test2.log", "hiromaily")
	logObj.Debug("New->Debug: test debug")
	logObj.Debugf("New->Debugf: %d - %s", 1, "test debugf")

	logObj.Info("New->Info: test info")
	logObj.Infof("New->Infof: %d - %s", 1, "test infof")

	logObj.Warn("New->Warn: test warn")
	logObj.Warnf("New->Warnf: %d - %s", 1, "test warnf")

	logObj.Error("New->Error: test error")
	logObj.Errorf("New->Errorf: %d - %s", 1, "test errorf")

	logObj.Fatal("New->Fatal: test fatal")
	logObj.Fatalf("New->Fatalf: %d - %s", 1, "test fatalf")
}

func TestNewLog2(t *testing.T) {
	logObj := New(DebugStatus, TimeShortFile, "[LOG_NEW_TEST]", "/var/log/go/test2.log", "hiromaily")
	logObj.Debug("New->Debug: test debug")
	logObj.Debugf("New->Debugf: %d - %s", 1, "test debugf")

	logObj.Info("New->Info: test info")
	logObj.Infof("New->Infof: %d - %s", 1, "test infof")

	logObj.Warn("New->Warn: test warn")
	logObj.Warnf("New->Warnf: %d - %s", 1, "test warnf")

	logObj.Error("New->Error: test error")
	logObj.Errorf("New->Errorf: %d - %s", 1, "test errorf")

	logObj.Fatal("New->Fatal: test fatal")
	logObj.Fatalf("New->Fatalf: %d - %s", 1, "test fatalf")
}

func TestNewLog3(t *testing.T) {
	logObj := New(DebugStatus, TimeShortFile, "", "", "hiromaily")
	logObj.Debug("New->Debug: test debug")
	logObj.Debugf("New->Debugf: %d - %s", 1, "test debugf")

	logObj.Info("New->Info: test info")
	logObj.Infof("New->Infof: %d - %s", 1, "test infof")

	logObj.Warn("New->Warn: test warn")
	logObj.Warnf("New->Warnf: %d - %s", 1, "test warnf")

	logObj.Error("New->Error: test error")
	logObj.Errorf("New->Errorf: %d - %s", 1, "test errorf")

	logObj.Fatal("New->Fatal: test fatal")
	logObj.Fatalf("New->Fatalf: %d - %s", 1, "test fatalf")
}

func TestInitializedLog(t *testing.T) {
	//InitializeLog(DebugStatus, TimeShortFile, "[LOG_INIT_TEST]", "/var/log/go/test.log", "hiromaily")
	InitializeLog(DebugStatus, TimeShortFile, "[LOG_INIT_TEST]", "", "hiromaily")

	Debug("New->Debug: test debug")
	Debugf("New->Debugf: %d - %s", 1, "test debugf")

	Info("New->Info: test info")
	Infof("New->Infof: %d - %s", 1, "test infof")

	Warn("New->Warn: test warn")
	Warnf("New->Warnf: %d - %s", 1, "test warnf")

	Error("New->Error: test error")
	Errorf("New->Errorf: %d - %s", 1, "test errorf")

	Fatal("New->Fatal: test fatal")
	Fatalf("New->Fatalf: %d - %s", 1, "test fatalf")
}

func TestInitializedLog2(t *testing.T) {
	InitializeLog(DebugStatus, DateTimeShortFile, "", "", "hiromaily")

	Debug("New->Debug: test debug")
	Debugf("New->Debugf: %d - %s", 1, "test debugf")

	Info("New->Info: test info")
	Infof("New->Infof: %d - %s", 1, "test infof")

	Warn("New->Warn: test warn")
	Warnf("New->Warnf: %d - %s", 1, "test warnf")

	Error("New->Error: test error")
	Errorf("New->Errorf: %d - %s", 1, "test errorf")

	Fatal("New->Fatal: test fatal")
	Fatalf("New->Fatalf: %d - %s", 1, "test fatalf")

	Stack()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
