package compress_test

import (
	"encoding/hex"
	. "github.com/hiromaily/golibs/compress"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"testing"
)

type User struct {
	Id   int
	Name string
}

var benchFlg bool = false

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	//Here is [slower] than included file's init()
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[COMPRESS_TEST]", "/var/log/go/test.log")
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
func TestGZipString(t *testing.T) {
	//t.Skip("skipping TestEncodeStruct")
	//*
	str := `SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=1
	`
	byteData, _ := GZipString(str)

	t.Logf("buf x: %x", byteData) //基数16、10以上の数には小文字(a-f)を使用
	//t.Logf("buf %o: %o", b.Bytes()) //%oは基数8だが、ここではsliceのuint8なのでエラー / uint8の配列
	t.Logf("buf v: %v", byteData) //基数8 = uint8の配列
	//t.Logf("buf s: %s", buf.String()) //string
	t.Logf("buf hex.Encode string: %s", hex.EncodeToString(byteData)) //string

	if hex.EncodeToString(byteData) != "82a249640aa44e616d65aa6861727279206461796f" {
		t.Errorf("TestGZipString result: %x", byteData)
	}
}
