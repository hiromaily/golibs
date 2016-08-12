package hash_test

import (
	. "github.com/hiromaily/golibs/cipher/hash"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

var benchFlg bool = false

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	//Here is [slower] than included file's init()
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[HASH_TEST]", "/var/log/go/test.log")
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
func TestGetMD5(t *testing.T) {
	testData := "password"
	result := GetMD5(testData)

	if result != "5f4dcc3b5aa765d61d8327deb882cf99" {
		t.Errorf("TestGetMD5 result: %s", result)
	}
}

func TestGetSHA1(t *testing.T) {
	testData := "password"
	result := GetSHA1(testData)

	if result != "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8" {
		t.Errorf("TestGetSHA1 result: %s", result)
	}
}

func TestGetSHA256(t *testing.T) {
	testData := "password"
	result := GetSHA256(testData)

	if result != "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8" {
		t.Errorf("TestGetSHA256 result: %s", result)
	}
}

func TestGetMD5Plus(t *testing.T) {
	testData := "password"
	result := GetMD5Plus(testData, "")

	if result != "02aaa55939a894316cfc3427234bf201" {
		t.Errorf("TestGetMD5Plus result: %s", result)
	}
}

func TestGetScrypt(t *testing.T) {
	//t.Skip("skipping TestGetScrypt")

	testData := "password"
	result := GetScrypt(testData)

	if result != "DeM69ar6oKwKyRZS0JnI2hI0in1OyVv/NT7U21TrGcU=" {
		t.Errorf("TestGetScrypt result: %s", result)
	}
}
