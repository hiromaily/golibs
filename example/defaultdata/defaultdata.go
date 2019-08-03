// Package defaultdata is just sample for check about what is default data
package defaultdata

import (
	"reflect"

	lg "github.com/hiromaily/golibs/log"
)

//-----------------------------------------------------------------------------
// check nil
// when calling these func setting nil of parameter, what is value of parameter.
//-----------------------------------------------------------------------------

// CheckInt is to check nil argument at int type
func CheckInt(val int) {
	//cannot use nil as type int in argument
	lg.Debugf("CheckInt: %v", val)
}

// CheckString is to check nil argument at string type
func CheckString(val string) {
	//cannot use nil as type string in argument
	lg.Debugf("CheckString: %v", val)
}

// CheckBool is to check nil argument at bool type
func CheckBool(val bool) {
	//cannot use nil as type bool in argument
	lg.Debugf("CheckBool: %v", val)
}

// CheckByte is to check nil argument at []byte type
func CheckByte(val []byte) {
	//cannot use nil as type bool in argument
	lg.Debugf("CheckByte: %v", val)

	var defaultData []byte
	lg.Debugf("CheckByte2: %v", defaultData)
}

// CheckError is to check nil argument at error type
func CheckError(val error) {
	//<nil>
	lg.Debugf("CheckError: %v", val)
}

// CheckSlice is to check nil argument at slice type
func CheckSlice(val []string) {
	//[]
	lg.Debugf("CheckSlice: %v", val)
}

// CheckMap is to check nil argument at map type
func CheckMap(val map[string]int) {
	//map[]
	lg.Debugf("CheckMap: %v", val)
}

// CheckInterface is to check nil argument at interface{} type
func CheckInterface(val interface{}) {
	//<nil>
	lg.Debugf("CheckInterface: %v", val)
}

// CheckMultiInterface is to check nil argument at multiple interface{} type
func CheckMultiInterface(args ...interface{}) {
	//[<nil>]
	//[<nil> <nil> <nil>]
	lg.Debugf("CheckMultiInterface: %v", args)
}

//-----------------------------------------------------------------------------
// check Interface
// If passing slice or pointer to interface{} what's happened?
//-----------------------------------------------------------------------------

// CheckInterfaceWhenSlice is to check slice argument in interface{} type
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

// CheckInterfaceWhenPointer is to check pointer argument in interface{} type
func CheckInterfaceWhenPointer(val interface{}) {
	//0xc82000e640
	lg.Debugf("CheckInterfaceWhenPointer: %v", val)
	v := reflect.ValueOf(val)
	lg.Debugf("CheckInterfaceWhenPointer v.Type(): %v", v.Type()) //*int
	lg.Debugf("CheckInterfaceWhenPointer v.Kind(): %v", v.Kind()) //ptr

}

//-----------------------------------------------------------------------------
// If changing value of parameter in func, what is value of variable on caller
//-----------------------------------------------------------------------------

// ChangeValOnSlice is to check when passing slice
func ChangeValOnSlice(val []string) {
	val[0] = "changed"
	lg.Debugf("CheckSlice: %v", val)
}

// ChangeValOnMap is to check when passing map
func ChangeValOnMap(val map[string]int) {
	val["apple"] = 555
	lg.Debugf("CheckMap: %v", val)
}

// ChangeValOnInterface is to check when passing interface{}
func ChangeValOnInterface(val interface{}) {
	val = "changed"
	lg.Debugf("CheckInterface: %v", val)
}

// ChangeValOnPointer is to check when passing pointer
func ChangeValOnPointer(val *string) {
	*val = "changed"
	lg.Debugf("CheckInterface: %s", *val)
}
