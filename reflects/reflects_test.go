package reflects_test

import (
	"flag"
	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/reflects"
	"os"
	"testing"
	"time"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Reflects_TEST]", "/var/log/go/test.log")
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

var dInt int = 10
var dInt64 int64 = 99999
var dStr string = "testdata"
var dBool bool = false
var dSlice []int = []int{1, 2, 3, 4, 5}
var dTime time.Duration = 1 * time.Nanosecond
var dMap = map[string]int{"apple": 150, "banana": 300, "lemon": 300}

type TeacherInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type SiteInfo struct {
	Url      string        `json:"url"`
	Teachers []TeacherInfo `json:"teachers"`
}

var siteInfo SiteInfo = SiteInfo{Url: "http://google.com",
	Teachers: []TeacherInfo{{Id: 123, Name: "Harry", Country: "Japan"}, {Id: 456, Name: "Taro", Country: "America"}}}
var tInfo []TeacherInfo = []TeacherInfo{{Id: 123, Name: "Harry", Country: "Japan"}, {Id: 456, Name: "Taro", Country: "America"}}

//-----------------------------------------------------------------------------
// Reflects
//-----------------------------------------------------------------------------
func TestCheckReflect(t *testing.T) {
	t.Skip("skipping TestCheckReflect")

	CheckReflect(dInt)     //int/int
	CheckReflect(dInt64)   //int64/int64
	CheckReflect(dStr)     //v.Kind(): string  Type: string
	CheckReflect(dSlice)   //v.Kind(): slice   Type: []int
	CheckReflect(dMap)     //v.Kind(): map     Type: map[string]int
	CheckReflect(siteInfo) //v.Kind(): struct  Type: reflects_test.SiteInfo
	CheckReflect(tInfo)    //v.Kind(): struct  Type: reflects_test.SiteInfo

	var techerInfo []TeacherInfo
	CheckReflect(&techerInfo) //v.Kind(): ptr  Type: *[]reflects_test.TeacherInfo
}

func TestGetValueAsString(t *testing.T) {
	t.Skip("skipping TestReflects")

	t.Log(GetValueAsString(dInt))
	t.Log(GetValueAsString(dInt64))
	t.Log(GetValueAsString(dStr))
	t.Log(GetValueAsString(dBool))
	t.Log(GetValueAsString(dSlice))
	t.Log(GetValueAsString(dTime))
	t.Log(GetValueAsString(dMap))
	t.Log(GetValueAsString([]time.Duration{dTime}))
	t.Log(GetValueAsString(siteInfo))
}

func TestDisplay(t *testing.T) {
	t.Skip("skipping TestDisplay")

	Display("int", dInt)
	//Display int (int):
	//int = 10

	Display("int64", dInt64)
	//Display int64 (int64):
	//int64 = 99999

	Display("string", dStr)
	//Display string (string):
	//string = "testdata"

	Display("bool", dBool)
	//Display bool (bool):
	//bool = false

	Display("slice", dSlice)
	//Display slice ([]int):
	//slice[0] = 1
	//slice[1] = 2
	//slice[2] = 3
	//slice[3] = 4
	//slice[4] = 5

	Display("time", dTime)
	//Display time (time.Duration):
	//time = 1

	Display("map", dMap)
	//Display map (map[string]int):
	//map["apple"] = 150
	//map["banana"] = 300
	//map["lemon"] = 300

	Display("struct", siteInfo)
	//Display struct (reflects_test.SiteInfo):
	//struct.Url = "http://google.com"
	//struct.Teachers[0].Id = 123
	//struct.Teachers[0].Name = "Harry"
	//struct.Teachers[0].Country = "Japan"
	//struct.Teachers[1].Id = 456
	//struct.Teachers[1].Name = "Taro"
	//struct.Teachers[1].Country = "America"

	Display("struct", tInfo)
	//Display struct ([]reflects_test.TeacherInfo):
	//struct[0].Id = 123
	//struct[0].Name = "Harry"
	//struct[0].Country = "Japan"
	//struct[1].Id = 456
	//struct[1].Name = "Taro"
	//struct[1].Country = "America"

	Display("struct", &tInfo)
	//Display struct (*[]reflects_test.TeacherInfo):
	//(*struct)[0].Id = 123
	//(*struct)[0].Name = "Harry"
	//(*struct)[0].Country = "Japan"
	//(*struct)[1].Id = 456
	//(*struct)[1].Name = "Taro"
	//(*struct)[1].Country = "America"
}

func TestSetDataToStruct(t *testing.T) {
	//t.Skip("skipping TestCheckStruct")
	var columns []string = []string{"field1", "field2", "field3"}
	//values := []interface{}{10, "Harry", "Japan"}
	values := make([][]interface{}, 2)
	values[0] = []interface{}{10, "Harry", "UK"}
	values[1] = []interface{}{15, "Hiroki", "Japan"}

	//1.struct
	//*
	var techerInfo TeacherInfo
	err := SetDataToStruct(columns, values, &techerInfo)
	if err != nil {
		t.Errorf("TestSetDataToStruct: error: %s", err)
	}
	t.Logf("techerInfo:%#v", techerInfo)
	//*/

	//2.slice struct
	//*
	var techerInfos []TeacherInfo
	err = SetDataToStruct(columns, values, &techerInfos)
	if err != nil {
		t.Errorf("TestSetDataToStruct: error: %s", err)
	}
	t.Logf("techerInfo:%#v", techerInfos)
	//*/
}

//-----------------------------------------------------------------------------
//Benchmark
//-----------------------------------------------------------------------------
func BenchmarkReflects(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
