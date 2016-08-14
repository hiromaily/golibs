package log_test

import (
	. "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	InitializeLog(DEBUG_STATUS, LOG_OFF_COUNT, 0, "[Log_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		Debug("This is bench test.")
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
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestNewLog(t *testing.T) {
	logObj := New(DEBUG_STATUS, LOG_OFF_COUNT, 0, "[LOG_NEW_TEST]", "/var/log/go/test2.log")
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
	InitializeLog(DEBUG_STATUS, LOG_OFF_COUNT, 0, "[LOG_INIT_TEST]", "/var/log/go/test.log")

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
