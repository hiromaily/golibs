package compress_test

import (
	"encoding/hex"
	"os"
	"testing"

	. "github.com/hiromaily/golibs/compress"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Compress]")
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

	lg.Debugf("buf x: %x", byteData) //基数16、10以上の数には小文字(a-f)を使用
	//lg.Debugf("buf %o: %o", b.Bytes()) //%oは基数8だが、ここではsliceのuint8なのでエラー / uint8の配列
	lg.Debugf("buf v: %v", byteData) //基数8 = uint8の配列
	//lg.Debugf("buf s: %s", buf.String()) //string
	lg.Debugf("buf hex.Encode string: %s", hex.EncodeToString(byteData)) //string

	//if hex.EncodeToString(byteData) != "82a249640aa44e616d65aa6861727279206461796f" {
	//	t.Errorf("TestGZipString result: %x", byteData)
	//}
}
