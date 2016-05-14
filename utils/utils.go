package utils

import (
	"fmt"
	"runtime"
	"reflect"
	"errors"
)

//check type of interface
func CheckInterface(v interface{}) string {
	//switch
	switch v.(type) {
	case int, int64, int32, int16, int8:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case []uint8:
		return "[]uint8"
	//case []byte:
	//	return "[]byte"
	default:
		return "default"
	}
}

//check type of interface
func CheckInterfaceByIf(val interface{}) string {
	// ValueOfでreflect.Value型のオブジェクトを取得
	//v := reflect.ValueOf(val).Type()
	v := reflect.ValueOf(val).Kind()

	if v == reflect.Int64 {
		return "int64"
	}
	if v == reflect.Int32 {
		return "int32"
	}
	if v == reflect.Int16 {
		return "int16"
	}
	if v == reflect.Int {
		return "int"
	}

	return ""
}

// search string
func SearchString(ary []string, str string) int {

	var retIdx int = -1
	if len(ary) == 0 {
		return retIdx
	}
	for i, val := range ary {
		if val == str {
			retIdx = i
			break
		}
	}

	return retIdx
}

// Interface型のString型への変更
func Itos(val interface{}) string {
	str, ok := val.(string)
	if !ok {
		return ""
	}
	return str
}

// Interface型のint型への変更
func Itoi(val interface{}) int {

	//v := reflect.ValueOf(val).Type()
	//v := reflect.ValueOf(val).Kind()
	//switch v.(type) {
	//case int, int64, int32, int16, int8:

	//TODO:型を判別して自動でその型にキャストしたい

	num64, ok := val.(int64)
	if ok {
		return int(num64)
	}

	num32, ok := val.(int32)
	if ok {
		return int(num32)
	}

	num, ok := val.(int)
	if ok {
		return int(num)
	}else{
		return 0
	}
}

// String型のError型への変更
func Stoe(val string) error {
	return errors.New(val)
}

// check error and if so execute panic
func GoPanicWhenError(err error) {
	if err != nil {
		fmt.Println(runtime.Caller(1))
		panic(err)
	}
}

// check error and if so print error
func ShowErrorWhenError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
