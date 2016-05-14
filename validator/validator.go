package validator

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Validator func(str string) bool
type ValidatorCal func(str string, num int) bool

var TagMap = map[string]Validator{
	"nonempty": isNonEmpty,
	"email":    isEmail,
	"url":      isURL,
	"alphanum": isNumber,
}

var TagMapCal = map[string]ValidatorCal{
	"min": isMinOK,
	"max": isMaxOK,
}

//github.com/asaskevich/govalidator/validator.go
func isNonEmpty(str string) bool {
	return str != ""
}

func isEmail(str string) bool {
	regEx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regEx.MatchString(str)
}

func isURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	regEx := regexp.MustCompile(`http(s)?://([\w-]+\.)+[\w-]+(/[\w- ./?%&=]*)?`)
	return regEx.MatchString(str)
}

func isNumber(str string) bool {
	_, ok := strconv.Atoi(str)
	return ok == nil
}

func isMinOK(str string, num int) bool {
	return utf8.RuneCountInString(str) >= num
}

func isMaxOK(str string, num int) bool {
	return utf8.RuneCountInString(str) <= num
}

//str: nonempty,min=3,max=40
//val: test@test.jp
func checkValidation(str string, val string) []string {
	ret := []string{}
	strs := strings.Split(str, ",")
	for _, v := range strs {
		//fmt.Println(v)
		//fmt.Println(TagMap[v])

		//When included「=」on v, divide it.
		equals := strings.Split(v, "=")
		var bRet bool
		if len(equals) > 1 {
			num, _ := strconv.Atoi(equals[1])
			bRet = TagMapCal[equals[0]](val, num)
			if !bRet {
				ret = append(ret, equals[0])
			}
		} else {
			bRet = TagMap[v](val)
			if !bRet {
				ret = append(ret, v)
			}
		}
	}
	//[]string{"min"}
	return ret
}

//Check validation after extracted tag from struct type.
func CheckValidation(s interface{}) map[string][]string {
	//ここではやむを得なく、構造体の値が渡されてくる

	mRet := make(map[string][]string)

	rt, rv := reflect.TypeOf(s), reflect.ValueOf(s)
	//fmt.Printf("%v\n", rt)
	//fmt.Printf("%v\n", rv)
	//fmt.Printf("%v\n", rt.NumField())

	//check each field of struct type
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		valid := field.Tag.Get("valid")
		fld := field.Tag.Get("field")
		disp := field.Tag.Get("dispName")

		fmt.Printf("[Tag] valid:%s, disp:%s, field:%s\n", valid, disp, fld)
		//[Tag] valid:nonempty,min=3,max=40, disp name:Eメール
		fmt.Printf("[Value] %s\n\n", rv.Field(i).Interface())
		//[Value] aa@jj.jp

		if valid != "" {
			//When check is required, check specific field
			val, _ := rv.Field(i).Interface().(string)
			//Returned value is slice, stored name of error
			ret := checkValidation(valid, val)
			if len(ret) != 0 {
				mRet[fld] = ret
			}
		}
	}
	//map[string][]string{"pass":[]string{"min"}, "test":[]string{"nonempty"}}
	return mRet
}
