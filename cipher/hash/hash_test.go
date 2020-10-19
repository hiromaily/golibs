package hash_test

import (
	"os"
	"testing"

	. "github.com/hiromaily/golibs/cipher/hash"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[HASH]")
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
	//tu.SkipLog(t)

	testData := "password"
	result := GetScrypt(testData)

	if result != "DeM69ar6oKwKyRZS0JnI2hI0in1OyVv/NT7U21TrGcU=" {
		t.Errorf("TestGetScrypt result: %s", result)
	}
}
