package time_test

import (
	"math"
	"os"
	"strings"
	"testing"
	"time"

	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/time"
	u "github.com/hiromaily/golibs/utils"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Time]")
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
// Check
//-----------------------------------------------------------------------------
func TestBasic(t *testing.T) {
	tu.SkipLog(t)

	ti := time.Now()
	lg.Debug(ti.Date())                          //2016 September 11
	lg.Debugf("t.Day(): %v", ti.Day())           //11
	lg.Debugf("t.Unix(): %v", ti.Unix())         //1473565765
	lg.Debugf("t.Location(): %v", ti.Location()) // Local
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestTrack(t *testing.T) {
	tu.SkipLog(t)
	defer Track(time.Now(), "TestTrack()")

	//sleep
	time.Sleep(1000 * time.Millisecond)
}

func TestTimeout(t *testing.T) {
	Timeout()
}

func TestCheckParseTime(t *testing.T) {
	tu.SkipLog(t)

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
	tu.SkipLog(t)

	//LastModified
	strTime := "Tue, 16 Aug 2016 01:31:09 GMT"
	ti, err := ParseTime(strTime)
	if err != nil {
		t.Errorf("[01]ParseTime error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestParseTimeForLastModified(t *testing.T) {
	tu.SkipLog(t)

	strTime := "Tue, 16 Aug 2016 01:31:09 GMT"
	ti, err := ParseTimeForLastModified(strTime)
	if err != nil {
		t.Errorf("[01]ParseTimeForLastModified error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestParseTimeForRss(t *testing.T) {
	tu.SkipLog(t)

	strTime := "Mon, 15 Aug 2016 08:16:28 +0000"
	ti, err := ParseTimeForRss(strTime)
	if err != nil {
		t.Errorf("[01]ParseTimeForRss error: %s", err)
	}
	lg.Debugf("time is %v", ti)
}

func TestGetCurrentTimeByStr(t *testing.T) {
	tu.SkipLog(t)

	strT := GetCurrentDateTimeByStr("")
	lg.Debug(strT)
}

func TestGetFormatDate(t *testing.T) {
	tu.SkipLog(t)

	result := GetFormatDate("2016-06-13 20:20:24", "", false)
	lg.Debugf("TestGetFormatDate[01] result: %s", result)
	//【6/13】

	result = GetFormatDate("2016-06-13 20:20:24", "", true)
	lg.Debugf("TestGetFormatDate[02] result: %s", result)
	//6/13(月)】

	result = GetFormatDate("2016-06-13 20:20:24", "[1月2日]", false)
	lg.Debugf("TestGetFormatDate[03] result: %s", result)
	//[6月13日]

	result = GetFormatDate("2016-06-13 20:20:24", "[1月2日(%s)]", true)
	lg.Debugf("TestGetFormatDate[04] result: %s", result)
	//[6月13日(月)]
}

func TestGetFormatTime(t *testing.T) {
	tu.SkipLog(t)

	result := GetFormatTime("2016-06-13 20:20:24", "")
	lg.Debugf("TestGetFormatTime[01] result: %s", result)
	//20:20

	result = GetFormatTime("2016-06-13 20:20:24", "15:04:05")
	lg.Debugf("TestGetFormatTime[02] result: %s", result)
	//20:20:24

	result = GetFormatTime("2016-06-13 20:20:24:555", "15:04:05:999")
	lg.Debugf("TestGetFormatTime[03] result: %s", result)
	//=>it doesn't work

	result = GetFormatTime("2016-06-13 20:20:24", "15時04分")
	lg.Debugf("TestGetFormatTime[04] result: %s", result)
	//20時20分
}

func TestGetFormatTime2(t *testing.T) {
	//tu.SkipLog(t)

	ti := GetFormatTime2(3, 10, 5, 400e6)
	lg.Debug("TestGetFormatTime2[01] result:", ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond)
	//3 10 5 0x111f9f0
	lg.Debug(ti.Format("2006-01-02T15:04:05.000Z07:00"))
	//2017-06-03T03:10:05.400+02:00
	lg.Debug(ti.Format("15:04:05.000Z07:00"))
	//03:10:05.400+02:00
	lg.Debug(ti.Format("15:04:05.000"))
	//03:10:05.400
}

//This logic is used for gotools/gosubstr/main.go
func TestCalcTime(t *testing.T) {
	//tu.SkipLog(t)
	//00:00:10,950
	timeStr := "00:00:10,950"
	//addedTime := -6.2
	//addedTime := 6.2
	addedTime := 0.6

	tims := strings.Split(strings.Replace(timeStr, ",", ":", -1), ":")
	timI := u.ConvertToInt(tims)

	//lg.Debug(950e6, timI[3], timI[3] * (10^6), timI[3] * int(math.Pow10(6))) //9.5e+08 950 -9498
	//1.2Ｅ＋08 = 1.2×10^8（120,000,000)
	//ti := GetFormatTime2(timI[0], timI[1], timI[2], 950e6)
	ti := GetFormatTime2(timI[0], timI[1], timI[2], timI[3]*int(math.Pow10(6)))
	lg.Debug(ti.Format("15:04:05.000"))
	//00:00:10.950

	integerVal := math.Trunc(addedTime)
	decimalVal := math.Trunc((addedTime - math.Trunc(addedTime)) * 1000)
	lg.Debug(integerVal) //-6
	lg.Debug(decimalVal) //-200
	//ti2 := ti.Add(1 * time.Minute)
	if integerVal != 0 {
		ti = ti.Add(time.Duration(integerVal) * time.Second)
		lg.Debug(ti.Format("15:04:05.000"))
		//00:00:04.950
	}
	if decimalVal != 0 {
		ti = ti.Add(time.Duration(int(decimalVal)*int(math.Pow10(6))) * time.Nanosecond)
		lg.Debug(ti.Format("15:04:05.000"))
		//00:00:04.750
	}

	result := strings.Replace(ti.Format("15:04:05.000"), ".", ",", -1)
	lg.Debug(result)

	//p := fmt.Println
	//p(then.Year())
	//p(then.Month())
	//p(then.Day())
	//p(then.Hour())
	//p(then.Minute())
	//p(then.Second())
	//p(then.Nanosecond())
	//p(then.Location()
}
