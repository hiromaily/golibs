package compress_test

import (
	"encoding/hex"
	"flag"
	. "github.com/hiromaily/golibs/compress"
	lg "github.com/hiromaily/golibs/log"
	"golang.org/x/tools/cmd/vet/testdata/divergent"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

type User struct {
	Id   int
	Name string
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[COMPRESS_TEST]", "/var/log/go/test.log")
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
// Compress
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
