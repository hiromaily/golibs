package http_test

import (
	"encoding/json"
	. "github.com/hiromaily/golibs/http"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
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
func TestGetRequest(t *testing.T) {
	t.Skip("skipping TestGetRequest")
	status, _, err := GetRequestSimple("https://www.google.co.jp/")

	if err != nil {
		t.Fatalf("TestGetRequest[1]: %s", err)
	}
	if status != 200 {
		t.Errorf("TestGetRequest[2]: %d", status)
	}
	//if body != "hoge" {
	//	t.Errorf("TestGetRequest[3]: %s", body)
	//}

}

func TestGetRequest2(t *testing.T) {
	status, body, err := GetRequestWithData("http://www.yahoo.co.jpp/")
	if err != nil {
		t.Fatalf("TestGetRequest2[1]: %s", err)
	}
	if status != 200 {
		t.Errorf("TestGetRequest2[2]: %d", status)
	}
	t.Logf("body: %v", body)
}

func TestPostRequest(t *testing.T) {
	t.Skip("skipping TestPostRequest")

	url := "https://www.google.co.jp/"

	message := MessagesJson{
		ContentType: 1,
		Text:        "something code",
	}

	byteBody, _ := json.Marshal(message)

	status, _, err := PostRequest(url, byteBody)
	if err != nil {
		t.Fatalf("TestPostRequest1: %s", err)
	}
	if status != 200 {
		t.Errorf("TestPostRequest2: %d", status)
	}
	t.Logf("byteBody: %v", byteBody)
}
