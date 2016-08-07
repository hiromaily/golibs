package defaultdata

import (
	lg "github.com/hiromaily/golibs/log"
	"reflect"
)

//-----------------------------------------------------------------------------
// check nil
//-----------------------------------------------------------------------------
//when calling these func setting nil of parameter, what is value of parameter.
func CheckInt(val int) {
	//cannot use nil as type int in argument
	lg.Debugf("CheckInt: %v", val)
}

func CheckString(val string) {
	//cannot use nil as type string in argument
	lg.Debugf("CheckString: %v", val)
}

func CheckBool(val bool) {
	//cannot use nil as type bool in argument
	lg.Debugf("CheckBool: %v", val)
}

func CheckByte(val []byte) {
	//cannot use nil as type bool in argument
	lg.Debugf("CheckByte: %v", val)

	var defaultData []byte
	lg.Debugf("CheckByte2: %v", defaultData)
}

func CheckError(val error) {
	//<nil>
	lg.Debugf("CheckError: %v", val)
}

func CheckSlice(val []string) {
	//[]
	lg.Debugf("CheckSlice: %v", val)
}

func CheckMap(val map[string]int) {
	//map[]
	lg.Debugf("CheckMap: %v", val)
}

func CheckInterface(val interface{}) {
	//<nil>
	lg.Debugf("CheckInterface: %v", val)
}

func CheckMultiInterface(args ...interface{}) {
	//[<nil>]
	//[<nil> <nil> <nil>]
	lg.Debugf("CheckMultiInterface: %v", args)
}

//-----------------------------------------------------------------------------
// check Interface
//-----------------------------------------------------------------------------
//If pass slice to interface{} what's happend?
func CheckInterfaceWhenSlice(val interface{}) {
	//[1 2 3 4 5]
	lg.Debugf("CheckInterfaceWhenSlice: %v", val)
	v := reflect.ValueOf(val)
	lg.Debugf("CheckInterfaceWhenSlice v.Type(): %v", v.Type()) //[]int
	lg.Debugf("CheckInterfaceWhenSlice v.Kind(): %v", v.Kind()) //slice

	//invalid operation: val[0] (type interface {} does not support indexing)
	//lg.Debugf("CheckInterfaceWhenSlice: %d,%d,%d", val[0], val[1], val[2])

	//How to convert interface{} to slice...
	vv, ok := val.([]int)
	lg.Debugf("CheckInterfaceWhenSliceval.([]int): %v, %t, %d", vv, ok, vv[0])

}

//pointer
func CheckInterfaceWhenPointer(val interface{}) {
	//0xc82000e640
	lg.Debugf("CheckInterfaceWhenPointer: %v", val)
	v := reflect.ValueOf(val)
	lg.Debugf("CheckInterfaceWhenPointer v.Type(): %v", v.Type()) //*int
	lg.Debugf("CheckInterfaceWhenPointer v.Kind(): %v", v.Kind()) //ptr

}

//-----------------------------------------------------------------------------
// If change parameter, what is value of variable on caller
//-----------------------------------------------------------------------------
func ChangeValOnSlice(val []string) {
	val[0] = "changed"
	lg.Debugf("CheckSlice: %v", val)
}

func ChangeValOnMap(val map[string]int) {
	val["apple"] = 555
	lg.Debugf("CheckMap: %v", val)
}

func ChangeValOnInterface(val interface{}) {
	val = "changed"
	lg.Debugf("CheckInterface: %v", val)
}

func ChangeValOnPointer(val *string) {
	*val = "changed"
	lg.Debugf("CheckInterface: %s", *val)
}
