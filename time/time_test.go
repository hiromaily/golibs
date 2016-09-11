package time_test

import (
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/time"
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
	tu.InitializeTest("[Time]")
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
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	tu.SkipLog(t)

	ti := time.Now()
	lg.Debug(ti.Date())                          //2016 September 11
	lg.Debugf("t.Day(): %v", ti.Day())           //11
	lg.Debugf("t.Unix(): %v", ti.Unix())         //1473565765
	lg.Debugf("t.Location(): %v", ti.Location()) // Local
}

func TestCheckParseTime(t *testing.T) {
	//tu.SkipLog(t)

	//LastModified
	strTime := "Tue, 16 Aug 2016 01:31:09 GMT"
	retI := CheckParseTime(strTime)
	lg.Debugf("LastModified data format: %s", strTime)
	for _, v := range retI {
		lg.Debugf("[index:%d] %s", v, TimeLayouts[v])
	}

	//RSS
	strTime = "Mon, 15 Aug 2016 08:16:28 +0000"
	retI = CheckParseTime(strTime)
	lg.Debugf("RSS data format: %s", strTime)
	for _, v := range retI {
		lg.Debugf("[index:%d] %s", v, TimeLayouts[v])
	}
}

func TestParseTime(t *testing.T) {
	//tu.SkipLog(t)

	//LastModified
	strTime := "Tue, 16 Aug 2016 01:31:09 GMT"
	ti, err := ParseTime(strTime)
	if err != nil {
		t.Errorf("[01]ParseTime error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestParseTimeForLastModified(t *testing.T) {
	//tu.SkipLog(t)

	strTime := "Tue, 16 Aug 2016 01:31:09 GMT"
	ti, err := ParseTimeForLastModified(strTime)
	if err != nil {
		t.Errorf("[01]ParseTimeForLastModified error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestParseTimeForRss(t *testing.T) {
	//tu.SkipLog(t)

	strTime := "Mon, 15 Aug 2016 08:16:28 +0000"
	ti, err := ParseTimeForRss(strTime)
	if err != nil {
		t.Errorf("[01]ParseTimeForRss error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestTrack(t *testing.T) {
	//tu.SkipLog(t)
	defer Track(time.Now(), "TestTrack()")

	//sleep
	time.Sleep(1000 * time.Millisecond)
}

func TestGetCurrentTimeByStr(t *testing.T) {
	//tu.SkipLog(t)

	strT := GetCurrentDateTimeByStr("")
	lg.Debug(strT)
}

func TestGetFormatDate(t *testing.T) {
	//tu.SkipLog(t)

	result := GetFormatDate("2016-06-13 20:20:24", "", false)
	lg.Debugf("TestGetFormatDate[01] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "", true)
	lg.Debugf("TestGetFormatDate[02] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "[1月2日]", false)
	lg.Debugf("TestGetFormatDate[03] result: %s", result)

	result = GetFormatDate("2016-06-13 20:20:24", "[1月2日(%s)]", true)
	lg.Debugf("TestGetFormatDate[04] result: %s", result)
}

func TestGetFormatTime(t *testing.T) {
	tu.SkipLog(t)

	result := GetFormatTime("2016-06-13 20:20:24", "")
	lg.Debugf("TestGetFormatTime[01] result: %s", result)

	result = GetFormatTime("2016-06-13 20:20:24", "15:04:05")
	lg.Debugf("TestGetFormatTime[02] result: %s", result)

	result = GetFormatTime("2016-06-13 20:20:24", "15時04分")
	lg.Debugf("TestGetFormatTime[03] result: %s", result)
}
