package reflects

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Just Check
func CheckReflect(value interface{}) {
	//reflect.ValueOf()
	v := reflect.ValueOf(value)
	fmt.Printf("reflect.ValueOf(value): %v\n", v)
	fmt.Printf("v.Kind(): %v\n\n", v.Kind())
	//v.String()

	//reflect.TypeOf()
	t := reflect.TypeOf(value)
	fmt.Printf("Type: %T\n", value)
	fmt.Printf("reflect.TypeOf(value): %v\n\n\n", t)
}

// Any formats any value as a string.
func GetValueAsString(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
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

// display
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

func SetDataToStruct(columns []string, values [][]interface{}, x interface{}) error {
	//As precondition parameter is slice of struct by pointer.
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("parameter should not be nil")
		} else {
			//2-1.Slice
			if v.Elem().Kind() == reflect.Slice || v.Elem().Kind() == reflect.Array {
				elemType := v.Elem().Type().Elem() //reflects_test.TeacherInfo
				newElem := reflect.New(elemType).Elem()
				for _, value := range values {
					scan(columns, value, newElem)
					v.Elem().Set(reflect.Append(v.Elem(), newElem))
				}
			} else if v.Elem().Kind() == reflect.Struct {
				scan(columns, values[0], v.Elem())
			} else {
				return errors.New("parameter should be pointer of struct slice or struct")
			}
		}
	} else {
		return errors.New("parameter should be pointer of struct slice or struct")
	}
	return nil
}

//This is for just check
func SetDataToStructForCheck(columns []string, values [][]interface{}, x interface{}) error {
	//As precondition parameter is slice of struct by pointer.
	v := reflect.ValueOf(x) //Value  (v.Elem() also return Value)
	//fmt.Println(v.Kind()) //ptr
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("parameter should not be nil")
		} else {
			fmt.Printf("v.Elem(): %v\n", v.Elem())               //[]
			fmt.Printf("v.Elem().Type(): %v\n", v.Elem().Type()) //[]reflects_test.TeacherInfo
			fmt.Printf("v.Elem().Kind(): %v\n", v.Elem().Kind()) //slice
			//2-1.Slice
			if v.Elem().Kind() == reflect.Slice || v.Elem().Kind() == reflect.Array {
				fmt.Println(1, v.Elem().Kind()) //slice
				//Test01 Basic
				if v.Elem().Len() == 0 {
					fmt.Println(v.Elem().CanAddr())     //true
					fmt.Println(v.Elem().Type())        //[]reflects_test.TeacherInfo
					fmt.Println(v.Elem().Type().Elem()) //[]reflects_test.TeacherInfo
					//fmt.Println(v.Elem().NumField()) //error
					//fmt.Println(v.Elem().Index(0))   //error
				} else {
					for i := 0; i < v.Elem().Len(); i++ {
						//struct
						//v.Index(i)
					}
				}
				//Test01 End

				//Test02: Indirect(), MakeSlice()
				/*
					elem := reflect.Indirect(v.Elem()) //if parameter is ptr, return v.Elem(). it's safe.
					//[]

					elem.Set(reflect.MakeSlice(elem.Type(), 0, 0))
					//[]
				*/
				//Test02 End

				//Test03: It can show content of slice.
				elemType := v.Elem().Type().Elem() //reflects_test.TeacherInfo
				/*
					for i := 0; i < elemType.NumField(); i++ {
						fmt.Println(elemType.Field(i))
						fmt.Println(elemType.Field(i).Index)
						fmt.Println(elemType.Field(i).Name)
						fmt.Println(elemType.Field(i).Type)
						fmt.Println(elemType.Field(i).Tag)
					}
				*/
				//Test03 End

				//Test04: create new element
				newElem := reflect.New(elemType).Elem()
				//newElem:{0  }

				//newType := reflect.ValueOf(newElem.Addr().Interface()).Type()
				//newElem.Addr().Interface(): *reflects_test.TeacherInfo
				//newType: *reflects_test.TeacherInfo

				//if newType.Kind() == reflect.Ptr {
				//	newType = newType.Elem() //struct
				//}

				// Get all fields
				//for i := 0; i < newType.Elem().NumField(); i++ {
				//	newElem.Field(i).Set(reflect.ValueOf(values[0][i])) //OK
				//}
				for _, value := range values {
					scan(columns, value, newElem)
					//append
					v.Elem().Set(reflect.Append(v.Elem(), newElem))
				}
				//Test04 End

			} else if v.Elem().Kind() == reflect.Struct {
				//2-2. struct
				fmt.Println(4, v.Elem().Kind())

				//1.To check type, use v.Type()
				//structType := v.Elem().Type()
				//for i := 0; i < structType.NumField(); i++ {
				//	//fmt.Println(structType.Field(i)) //This is type
				//	fmt.Println(structType.Field(i).Index)
				//	fmt.Println(structType.Field(i).Type)
				//	fmt.Println(structType.Field(i).Name)
				//	fmt.Println(structType.Field(i).Tag)
				//	fmt.Println(structType.Field(i).PkgPath)
				//	fmt.Println(structType.Field(i).Offset)
				//	fmt.Println(structType.Field(i).Anonymous)
				//}

				//2.To check value, use v.Elem()
				//for i := 0; i < v.Elem().NumField(); i++ {
				//	fmt.Println(v.Elem().Field(i)) //This is value
				//	fmt.Println(v.Elem().Field(i).Kind())
				//	fmt.Println(v.Elem().Field(i).CanAddr())
				//}

				//3.Set data (for test)
				//v.Elem().Field(0).SetInt(1)
				//v.Elem().Field(1).SetString("harry")
				//v.Elem().Field(2).SetString("Japan")

				//v.Elem().Field(0).Set(reflect.ValueOf(values[0]))
				//v.Elem().Field(1).Set(reflect.ValueOf(values[1]))
				//v.Elem().Field(2).Set(reflect.ValueOf(values[2]))

				scan(columns, values[0], v.Elem())
			} else {
				return errors.New("parameter should be pointer of struct slice or slice")
			}
		}
	} else {
		return errors.New("parameter should be pointer of struct slice or slice")
	}

	return nil
}

//Set data
func scan(columns []string, values []interface{}, v reflect.Value) error {
	structType := v.Type()
	for i := 0; i < structType.NumField(); i++ {
		v.Field(i).Set(reflect.ValueOf(values[i]))

		//reflect: call of reflect.Value.Elem on struct Value
		/*
			switch v.Kind() {
			case reflect.Invalid:
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v.Elem().Field(i).SetInt(values[i])
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				v.Elem().Field(i).SetUint(values[i])
			case reflect.Bool:
				v.Elem().Field(i).SetBool(values[i])
			case reflect.String:
				v.Elem().Field(i).SetString(values[i])
			case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
			default: // reflect.Array, reflect.Struct, reflect.Interface
			}
		*/
	}
	return nil
}
