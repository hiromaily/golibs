package validator_test

import (
	"os"

	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/validator"

	//"reflect"
	"sort"
	"testing"
)

type LoginRequest struct {
	Email string `valid:"nonempty,email,min=5,max=40" field:"email" dispName:"E-Mail"`
	Pass  string `valid:"nonempty,min=8,max=16" field:"pass" dispName:"Password"`
	Code  string `valid:"nonempty,number" field:"code" dispName:"Code"`
	Alpha string `valid:"alphabet" field:"alpha" dispName:"Alpha"`
}

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

func setup() {
	tu.InitializeTest("[Validator]")
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

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCheckValidation(t *testing.T) {
	//tu.SkipLog(t)

	//TODO: TableDrivenTests
	// https://github.com/golang/go/wiki/TableDrivenTests

	posted := &LoginRequest{Email: "abc", Pass: "pass", Code: "aa"}
	mRet := CheckValidation(posted, false)
	lg.Debugf("CheckValidation() %+v", mRet)

	//expected error number
	if len(mRet) != 3 {
		t.Errorf("[01]CheckValidation error: %#v", mRet)
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
		t.Errorf("[02]ConvertErrorMsgs error: %#v", msgs)
	}
	lg.Debugf("ConvertErrorMsgs() %#v", msgs)
}

func TestCheckSkipValidation(t *testing.T) {
	posted := &LoginRequest{Email: "aa", Pass: "", Code: ""}
	mRet := CheckValidation(posted, true)

	//expected error number
	if len(mRet) != 1 {
		t.Errorf("[01]CheckValidation error: %#v", mRet)
	}

}

// Table Test
func TestCheckValidationOnTable(t *testing.T) {
	for i, tt := range validTests {
		mRet := CheckValidation(&tt.input, false)

		//expected error number
		if len(mRet) != len(tt.errMsg) {
			t.Errorf("[01]CheckValidation(index:%d) error: %#v", i, mRet)
		}

		msgs := ConvertErrorMsgs(mRet, ErrFmt)
		//search message
		//expected error message
		if !sliceMatching(tt.errMsg, msgs) {
			t.Errorf("[02]ConvertErrorMsgs(index:%d) error: %#v", i, msgs)
			t.Error("input data:", tt.input)
		}
	}
}
