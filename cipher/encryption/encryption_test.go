package encryption_test

import (
	"flag"
	. "github.com/hiromaily/golibs/cipher/encryption"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[ENCRYPTION_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Encryption
//-----------------------------------------------------------------------------
func TestEncryption(t *testing.T) {
	//t.Skip("skipping TestExec")

	size := 16
	key := "8#75F%R+&a5ZvM_<"
	iv := "@~wp-7hPs<WEx@R4"

	str := "abcdefg@gmail.com"

	NewCrypt(size, key, iv)
	crypt := GetCryptInstance()

	/*
		tmp := crypt.Encrypt([]byte(str))
		//when assert byte type to string, be careful.
		result := string(tmp[:])
		if result != "xxxx" {
			//result is not readable
			t.Errorf("TestEncryption[01] result: %s", result)
		}
	*/

	result2 := crypt.EncryptBase64(str)
	if result2 != "GY+hCmXh+xJekHSnpuy6fe7s7adFBqWqfgeuMnBv9GQ=" {
		t.Errorf("TestEncryption[02] result: %s", result2)
	}

	result3, _ := crypt.DecryptBase64(result2)
	if result3 != str {
		t.Errorf("TestEncryption[03] result: %s", result3)
	}

	/*
		key := "bardzotrudnykluczszyfrujący"
		aes, err := simpleaes.New(16, key)
		if err != nil {
			panic(err)
		}
		phrase := "czy nie mają koty na nietoperze ochoty?"
		buf := aes.Encrypt([]byte(phrase))
		fmt.Println(buf)
		buf = aes.Decrypt(buf)
		fmt.Println(string(buf))
	*/

}
