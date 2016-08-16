package boom_test

import (
	"encoding/json"
	"flag"
	. "github.com/hiromaily/golibs/http/boom"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	//u "github.com/hiromaily/golibs/utils"
	"os"
	"testing"
)

type MessagesJson struct {
	ContentType uint8  `json:"contentType"`
	Text        string `json:"text"`
}

var (
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[HTTP_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
	if benchFlg {
	}
}

func teardown() {
	if benchFlg {
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
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
