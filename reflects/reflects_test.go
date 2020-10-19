package reflects_test

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	lg "github.com/hiromaily/golibs/log"
	. "github.com/hiromaily/golibs/reflects"
	tu "github.com/hiromaily/golibs/testutil"
)

type TeacherInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type SiteInfo struct {
	URL      string        `json:"url"`
	Teachers []TeacherInfo `json:"teachers"`
}

type LoginRequest struct {
	Email string `valid:"nonempty,email,min=5,max=40" field:"email" dispName:"E-Mail"`
	Pass  string `valid:"nonempty,min=8,max=16" field:"pass" dispName:"Password"`
	Code  string `valid:"nonempty,number" field:"code" dispName:"Code"`
	Alpha string `valid:"alphabet" field:"alpha" dispName:"Alpha"`
}

var (
	//test data
	dInt         = 10
	dInt64 int64 = 99999
	dStr         = "testdata"
	dBool  bool
	dSlice = []int{1, 2, 3, 4, 5}
	dTime  = 1 * time.Nanosecond
	dMap   = map[string]int{"apple": 150, "banana": 300, "lemon": 300}

	siteInfo = SiteInfo{URL: "http://google.com",
		Teachers: []TeacherInfo{{ID: 123, Name: "Harry", Country: "Japan"}, {ID: 456, Name: "Taro", Country: "America"}}}
	tInfo = []TeacherInfo{{ID: 123, Name: "Harry", Country: "Japan"}, {ID: 456, Name: "Taro", Country: "America"}}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Reflects]")
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
// functions
//-----------------------------------------------------------------------------

// CheckReflect is just to check ValueOf and TypeOf of value
func checkReflect(value interface{}) {
	//reflect.ValueOf()

	t := reflect.TypeOf(value)
	fmt.Printf("[]Type of parameter: %T\n", value)
	fmt.Printf(" reflect.TypeOf(value): %v\n", t)

	v := reflect.ValueOf(value)
	fmt.Printf(" reflect.ValueOf(value): %v\n", v)
	fmt.Printf(" v.Kind(): %v\n\n", v.Kind())

}

// get tag name
func checkStruct(data *LoginRequest) {

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		lg.Debugf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s",
			typeField.Name, valueField.Interface(), tag.Get("valid"))
	}
	lg.Debug("-------------------------------------")
}

// GetValueAsString is to get any formats any value as a string.
func GetValueAsString(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom is to format a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		//return strconv.FormatBool(v.Bool())
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

// Display is to display kind() and type() and value
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(name string, v reflect.Value) {
	switch v.Kind() { //uint
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", name)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", name, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", name, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", name,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", name)
		} else {
			display(fmt.Sprintf("(*%s)", name), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", name)
		} else {
			fmt.Printf("%s.type = %s\n", name, v.Elem().Type())
			display(name+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", name, formatAtom(v))
	}
}

//-----------------------------------------------------------------------------
// Check
//-----------------------------------------------------------------------------

// TestCheckValidationEg is just check of struct type
func TestCheckValidationEg(t *testing.T) {
	tu.SkipLog(t)

	//1:Normal
	data := &LoginRequest{Email: "abc", Pass: "pass", Code: "aa", Alpha: "abcde"}
	checkStruct(data)

	//2:Normal and blank field
	data = &LoginRequest{Email: "abc", Pass: "pass", Code: "aa", Alpha: ""}
	checkStruct(data)

	//3: there is lack of field
	data = &LoginRequest{Email: "abc", Pass: "pass", Code: "aa"}
	checkStruct(data)

}

func TestCheckReflect(t *testing.T) {
	tu.SkipLog(t)

	checkReflect(dInt)     //int/int
	checkReflect(dInt64)   //int64/int64
	checkReflect(dStr)     //v.Kind(): string  Type: string
	checkReflect(dSlice)   //v.Kind(): slice   Type: []int
	checkReflect(dMap)     //v.Kind(): map     Type: map[string]int
	checkReflect(siteInfo) //v.Kind(): struct  Type: reflects_test.SiteInfo
	checkReflect(tInfo)    //v.Kind(): struct  Type: reflects_test.SiteInfo

	var techerInfo []TeacherInfo
	checkReflect(&techerInfo) //v.Kind(): ptr  Type: *[]reflects_test.TeacherInfo
}

func TestGetValueAsString(t *testing.T) {
	tu.SkipLog(t)

	lg.Debug(GetValueAsString(dInt))
	lg.Debug(GetValueAsString(dInt64))
	lg.Debug(GetValueAsString(dStr))
	lg.Debug(GetValueAsString(dBool))
	lg.Debug(GetValueAsString(dSlice))
	lg.Debug(GetValueAsString(dTime))
	lg.Debug(GetValueAsString(dMap))
	lg.Debug(GetValueAsString([]time.Duration{dTime}))
	lg.Debug(GetValueAsString(siteInfo))
}

func TestDisplay(t *testing.T) {
	tu.SkipLog(t)

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

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSetDataToStruct(t *testing.T) {
	//tu.SkipLog(t)

	//var columns = []string{"field1", "field2", "field3"}
	//values := []interface{}{10, "Harry", "Japan"}
	values := make([][]interface{}, 2)
	values[0] = []interface{}{10, "Harry", "UK"}
	values[1] = []interface{}{15, "Hiroki", "Japan"}

	//1.struct
	//*
	var techerInfo TeacherInfo
	err := SetDataToStruct(values, &techerInfo)
	if err != nil {
		t.Errorf("[01]SetDataToStruct: error: %s", err)
	}
	lg.Debugf("techerInfo:%#v", techerInfo)
	//*/

	//2.slice struct
	//*
	var techerInfos []TeacherInfo
	err = SetDataToStruct(values, &techerInfos)
	if err != nil {
		t.Errorf("[02]SetDataToStruct: error: %s", err)
	}
	lg.Debugf("techerInfo:%#v", techerInfos)
	//*/
}

//-----------------------------------------------------------------------------
// Benchmark
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
