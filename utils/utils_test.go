package utils_test

import (
	"fmt"
	"os"
	"testing"

	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/utils"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Utils]")
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
func TestCheckInterface(t *testing.T) {
	tu.SkipLog(t)

	lg.Debug(CheckInterface(10)) //int

	lg.Debug(CheckInterface(5.4)) //default

	lg.Debug(CheckInterface("aaaaa")) //string

	lg.Debug(CheckInterface('A')) //int

	lg.Debug(CheckInterface(true)) //bool

	lg.Debug(CheckInterface([]byte("hello"))) //[]uint8

	lg.Debug(CheckInterface([]int{1, 2, 3})) //default

	lg.Debug(CheckInterface(func() {})) //default

}

func TestCheckInterfaceByIf(t *testing.T) {
	tu.SkipLog(t)

	lg.Debug(CheckInterfaceByIf(10)) //int

	lg.Debug(CheckInterfaceByIf(5.4)) //float64

	lg.Debug(CheckInterfaceByIf("aaaaa")) //string

	lg.Debug(CheckInterfaceByIf('A')) //int32

	lg.Debug(CheckInterfaceByIf(true)) //bool

	lg.Debug(CheckInterfaceByIf([]byte("hello"))) //slice

	lg.Debug(CheckInterfaceByIf([]int{1, 2, 3})) //slice

	lg.Debug(CheckInterfaceByIf(func() {})) //func
}

func TestStoType(t *testing.T) {
	tu.SkipLog(t)

	//return reflect.Kind
	lg.Debug(StoType("int")) //int
}

func TestSearchString(t *testing.T) {
	tu.SkipLog(t)

	data := []string{"ABC", "Abc", "abc"}
	lg.Debug(SearchString(data, "abc")) //2
}

func TestSearchStringLower(t *testing.T) {
	tu.SkipLog(t)

	data := []string{"ABC", "Abc", "abc"}
	lg.Debug(SearchStringLower(data, "abc")) //0
}

func TestSlice(t *testing.T) {
	tu.SkipLog(t)

	data := "0123456789"
	lg.Debug(Slice(data, 3))    //3456789
	lg.Debug(Slice(data, -2))   //89
	lg.Debug(Slice(data, 1, 5)) //12345
}

func TestSubstr(t *testing.T) {
	tu.SkipLog(t)

	data := "0123456789"
	lg.Debug(Substr(data, 3))     //3456789
	lg.Debug(Substr(data, 1, 5))  //12345
	lg.Debug(Substr(data, -5, 3)) //567
}

func TestOperateSlice(t *testing.T) {
	tu.SkipLog(t)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lg.Debug(PopInt(data))           //[1 2 3 4 5 6 7 8] 9 is gone
	lg.Debug(PushInt(data, 10))      //[1 2 3 4 5 6 7 8 9 10] 10 is added
	lg.Debug(ShiftInt(data))         //[2 3 4 5 6 7 8 9] 1 is gone
	lg.Debug(UnshiftInt(data, 0))    //[0 1 2 3 4 5 6 7 8 9] 0 is added
	lg.Debug(SpliceInt(data, 1, 99)) //[1 99 2 3 4 5 6 7 8 9]
	lg.Debug(DeleteInt(data, 1, 3))  //[1 4 5 6 7 8 9]
}

func TestOperateSlice2(t *testing.T) {
	tu.SkipLog(t)

	data2 := "abcdefghijk"
	lg.Debug(PopStr(data2))              //abcdefghij => k is gone
	lg.Debug(PushStr(data2, "l"))        //abcdefghijkl => l is added
	lg.Debug(ShiftStr(data2))            //bcdefghijk => a is gone
	lg.Debug(UnshiftStr(data2, "x"))     //xabcdefghijk
	lg.Debug(SpliceStr(data2, "xyz", 1)) //axyzbcdefghijk => xyz is added
	lg.Debug(DeleteStr(data2, 1, 3))     //aefghijk
}

func TestRandom(t *testing.T) {
	//tu.SkipLog(t)
	CheckRandom()
}

func TestGenerateIntData(t *testing.T) {
	tu.SkipLog(t)

	lg.Debug(GenerateIntData(4, 10))         //[9,1,9,4]
	lg.Debug(GenerateUniquieArray(3, 5, 10)) //[7,6,10]
	lg.Debug(GenerateRandom(6, 9))           //7

	arr := []int{1, 2, 3, 4, 5}
	lg.Debug(DeleteElement(arr, 3)) //1,2,4,5
}

func TestPickOneFromEnum(t *testing.T) {
	tu.SkipLog(t)

	enumData := []string{"apple", "grape", "strawberry"}
	lg.Debug(PickOneFromEnum(enumData))
	lg.Debug(PickOneFromEnum(enumData))

	enumData2 := []string{"apple"}
	lg.Debug(PickOneFromEnum(enumData2))
	lg.Debug(PickOneFromEnum(enumData2))
}

func TestUniqueStringSlice(t *testing.T) {
	tu.SkipLog(t)
	fmt.Println(UniqueStringSlice([]string{"aaa", "bbb", "ccc", "bbb"}))
}

func TestSortStructSlice(t *testing.T) {
	SortStructSlice()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkUtils(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
