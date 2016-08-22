package times_test

import (
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	. "github.com/hiromaily/golibs/times"
	"os"
	"testing"
	"time"
)

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[TIME_TEST]", "/var/log/go/test.log")
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
func TestBasic(t *testing.T) {
	//t.Skip("skipping TestBasic")

	ti := time.Now()
	lg.Debug(ti.Date())
	lg.Debugf("t.Day(): %v", ti.Day())
	lg.Debugf("t.Unix(): %v", ti.Unix())
	lg.Debugf("t.Location(): %v", ti.Location())
}

func TestGetFormatDate(t *testing.T) {
	t.Skip("skipping TestGetFormatDate")

	result := GetFormatDate("2016-06-13 20:20:24", "", false)
	t.Logf("TestGetFormatDate[01] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "", true)
	t.Logf("TestGetFormatDate[02] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "【1月2日】", false)
	t.Logf("TestGetFormatDate[03] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "【1月2日(%s)】", true)
	t.Logf("TestGetFormatDate[04] result: %s", result)
}

func TestGetFormatTime(t *testing.T) {
	t.Skip("skipping TestGetFormatTime")

	result := GetFormatTime("2016-06-13 20:20:24", "")
	t.Logf("TestGetFormatTime[01] result: %s", result)

	result = GetFormatTime("2016-06-13 20:20:24", "15:04:05")
	t.Logf("TestGetFormatTime[02] result: %s", result)

	result = GetFormatTime("2016-06-13 20:20:24", "15時04分")
	t.Logf("TestGetFormatTime[03] result: %s", result)
}

func TestGetCurrentTimeByStr(t *testing.T) {
	strT := GetCurrentTimeByStr()
	t.Log(strT)
}

func TestTrack(t *testing.T) {
	t.Skip("skipping TestTrack")

	defer Track(time.Now(), "TestTrack()")

	//sleep
	time.Sleep(1000 * time.Millisecond)
}
