package encryption_test

import (
	. "github.com/hiromaily/golibs/cipher/encryption"
	//lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Encryption]")
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
func TestEncryption(t *testing.T) {
	//tu.SkipLog(t)

	key := "8#75F%R+&a5ZvM_<"
	iv := "@~wp-7hPs<WEx@R4"

	str := "abcdefg@gmail.com"

	NewCrypt(key, iv)
	crypt := GetCrypt()

	result1 := crypt.EncryptBase64(str)
	if result1 != "GY+hCmXh+xJekHSnpuy6fe7s7adFBqWqfgeuMnBv9GQ=" {
		t.Errorf("[01]EncryptBase64 result: %s", result1)
	}

	result2, _ := crypt.DecryptBase64(result1)
	if result2 != str {
		t.Errorf("[02]DecryptBase64 result: %s", result2)
	}
}
