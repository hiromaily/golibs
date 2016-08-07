package http_test

import (
	"encoding/json"
	"flag"
	. "github.com/hiromaily/golibs/http"
	lg "github.com/hiromaily/golibs/log"
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
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[HTTP_TEST]", "/var/log/go/test.log")
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
