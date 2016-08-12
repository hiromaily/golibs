package defaultdata_test

import (
	. "github.com/hiromaily/golibs/defaultdata"
	lg "github.com/hiromaily/golibs/log"
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
	//Here is [slower] than included file's init()
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[DEFAULTDATA_TEST]", "/var/log/go/test.log")
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
func TestDefault(t *testing.T) {

	//cannot use nil as type int in argument to defaultdata.CheckInt
	/*
		CheckInt(nil)
	*/

	//cannot use nil as type string in argument to defaultdata.CheckString
	/*
		CheckString(nil)
	*/

	//cannot use nil as type bool in argument to defaultdata.CheckBool
	/*
		CheckBool(nil)
	*/
	CheckByte(nil)

	CheckError(nil)

	CheckSlice(nil)

	CheckMap(nil)

	CheckInterface(nil)

	CheckMultiInterface(nil)

	CheckMultiInterface(nil, nil, nil)

	//----------------------------------------------------
	//What's happened when sending slice data to interface
	//----------------------------------------------------
	data := []int{1, 2, 3, 4, 5}
	CheckInterfaceWhenSlice(data)

	var intData int = 1
	p := &intData
	CheckInterfaceWhenPointer(p)

	//----------------------------------------------------
	//Check givedvalue after calling func.
	//----------------------------------------------------
	strData := []string{"a", "b", "c", "d", "e"}
	ChangeValOnSlice(strData)
	//t.Logf("ChangeValOnSlice: %v", strData)
	if strData[0] != "a" {
		t.Errorf("ChangeValOnSlice value: %v", strData)
	}
	//changed!

	mapInt := map[string]int{"apple": 100, "lemon": 200, "banana": 300}
	ChangeValOnMap(mapInt)
	//t.Logf("ChangeValOnMap: %v", mapInt)
	if mapInt["apple"] != 100 {
		t.Errorf("ChangeValOnMap value: %v", mapInt)
	}
	//changed!

	strData2 := "before"
	ChangeValOnInterface(strData2)
	if strData2 != "before" {
		t.Errorf("ChangeValOnInterface value: %v", strData2)
	}
	//Not changed!

	//set address as pointer
	ChangeValOnPointer(&strData2)
	if strData2 != "before" {
		t.Errorf("ChangeValOnPointer value: %v", strData2)
	}
	//changed!

}
