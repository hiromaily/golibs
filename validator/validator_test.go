package validator_test

import (
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	//r "github.com/hiromaily/golibs/runtimes"
	. "github.com/hiromaily/golibs/validator"
	"os"
	"reflect"
	"sort"
	"testing"
)

type LoginRequest struct {
	Email string `valid:"nonempty,email,min=5,max=40" field:"email" dispName:"E-Mail"`
	Pass  string `valid:"nonempty,min=8,max=16" field:"pass" dispName:"Password"`
	Code  string `valid:"nonempty,number" field:"code" dispName:"Code"`
	Alpha string `valid:"alphabet" field:"alpha" dispName:"Alpha"`
}

var (
	benchFlg bool = false
)

var ErrFmt = map[string]string{
	"nonempty": "Empty is not allowed on %s",
	"email":    "Format of %s is invalid",
	"alphabet": "Only alphabet is allowd on %s",
	"number":   "Only number is allowd on %s",
	"min":      "At least %s of characters is required on %s",
	"max":      "At a maximum %s of characters is allowed on %s",
}

var validTests = []struct {
	input  LoginRequest
	errMsg []string
}{
	{LoginRequest{Email: "abc", Pass: "pass", Code: "aa", Alpha: ""},
		[]string{"Only number is allowd on Code", "At least 8 of characters is required on Password", "Format of E-Mail is invalid"}},
	{LoginRequest{Email: "abc@test.jp", Pass: "password", Code: "", Alpha: "abcdef"},
		[]string{"Empty is not allowed on Code"}},
	{LoginRequest{Email: "abc@test.jp", Pass: "", Code: "12345", Alpha: "abcdef123"},
		[]string{"Empty is not allowed on Password", "Only alphabet is allowd on Alpha"}},
	{LoginRequest{Email: "abc@test.jp", Pass: "12345678901234567890", Code: "999aaa", Alpha: ""},
		[]string{"Only number is allowd on Code", "At a maximum 16 of characters is allowed on Password"}},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[FLAG_TEST]", "/var/log/go/test.log")
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
// functions
//-----------------------------------------------------------------------------
func sliceMatching(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Strings(s1)
	sort.Strings(s2)

	bRet := true
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			bRet = false
			break
		}
	}

	return bRet
}

func checkStruct(data *LoginRequest) {

	val := reflect.ValueOf(data).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n",
			typeField.Name, valueField.Interface(), tag.Get("valid"))
	}
	fmt.Println("--------------------------")
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCheckValidation(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

	//TODO: TableDrivenTests
	// https://github.com/golang/go/wiki/TableDrivenTests

	posted := &LoginRequest{Email: "abc", Pass: "pass", Code: "aa"}
	//mRet := CheckValidation(*posted, false)
	mRet := CheckValidation(posted, false)
	t.Logf("%+v", mRet)

	//expected error number
	if len(mRet) != 3 {
		t.Errorf("TestCheckValidation[01] error: %#v", mRet)
	}

	msgs := ConvertErrorMsgs(mRet, ErrFmt)
	//search message
	bRet := false
	for _, v := range msgs {
		if v == "Format of E-Mail is invalid" {
			bRet = true
			break
		}
	}

	//expected error message
	if !bRet {
		t.Errorf("TestCheckValidation[02] error: %#v", msgs)
	}
	t.Logf("%#v", msgs)
}

func TestCheckSkipValidation(t *testing.T) {
	posted := &LoginRequest{Email: "aa", Pass: "", Code: ""}
	mRet := CheckValidation(posted, true)
	t.Logf("%+v", mRet)

	//expected error number
	if len(mRet) != 1 {
		t.Errorf("TestCheckSkipValidation[01] error: %#v", mRet)
	}

}

func TestCheckValidationOnTable(t *testing.T) {
	for i, tt := range validTests {
		mRet := CheckValidation(&tt.input, false)
		//t.Logf("%+v", mRet)

		//expected error number
		if len(mRet) != len(tt.errMsg) {
			t.Errorf("TestCheckValidation[01](index:%d) error: %#v", i, mRet)
		}

		msgs := ConvertErrorMsgs(mRet, ErrFmt)
		//search message
		//expected error message
		if !sliceMatching(tt.errMsg, msgs) {
			t.Errorf("TestCheckValidation[02](index:%d) error: %#v", i, msgs)
			t.Error(tt.input)
		}
		//t.Logf("%#v", msgs)
	}
}

//change validation
// it's inpossible bacause of nothing method for update
func TestCheckValidationEg(t *testing.T) {
	//t.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))

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
