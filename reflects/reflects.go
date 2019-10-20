package reflects

import (
	"errors"
	"fmt"
	"reflect"
)

// SetDataToStruct is to set value to struct passed as interface{}
func SetDataToStruct(values [][]interface{}, x interface{}) error {
	//As precondition parameter is slice of struct by pointer.
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("parameter should not be nil")
		}
		//2-1.Slice
		if v.Elem().Kind() == reflect.Slice || v.Elem().Kind() == reflect.Array {
			elemType := v.Elem().Type().Elem() //reflects_test.TeacherInfo
			newElem := reflect.New(elemType).Elem()
			for _, value := range values {
				scan(value, newElem)
				v.Elem().Set(reflect.Append(v.Elem(), newElem))
			}
		} else if v.Elem().Kind() == reflect.Struct {
			scan(values[0], v.Elem())
		} else {
			return errors.New("parameter should be pointer of struct slice or struct")
		}
	} else {
		return errors.New("parameter should be pointer of struct slice or struct")
	}
	return nil
}

// SetDataToStructForDev is just check (for development of SetDataToStruct)
func SetDataToStructForDev(values [][]interface{}, x interface{}) error {
	//As precondition parameter is slice of struct by pointer.
	v := reflect.ValueOf(x) //Value  (v.Elem() also return Value)
	//fmt.Println(v.Kind()) //ptr
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("parameter should not be nil")
		}
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
				scan(value, newElem)
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

			scan(values[0], v.Elem())
		} else {
			return errors.New("parameter should be pointer of struct slice or slice")
		}
	} else {
		return errors.New("parameter should be pointer of struct slice or slice")
	}

	return nil
}

// scan is to set data to container
func scan(values []interface{}, v reflect.Value) {
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
}
