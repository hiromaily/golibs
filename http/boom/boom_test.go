package boom_test

import (
	"encoding/json"
	"flag"
	. "github.com/hiromaily/golibs/http/boom"
	lg "github.com/hiromaily/golibs/log"
	//u "github.com/hiromaily/golibs/utils"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

type MessagesJson struct {
	ContentType uint8  `json:"contentType"`
	Text        string `json:"text"`
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[HTTP_TEST]", "/var/log/goweb/ginserver.log")
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

func TestExecBoom(t *testing.T) {
	url := "https://www.google.co.jp/"

	message := MessagesJson{
		ContentType: 1,
		Text:        "something code",
	}

	byteBody, _ := json.Marshal(message)

	//status, _, err := PostRequest(url, byteBody)
	ExecBoom(url, byteBody)
}
