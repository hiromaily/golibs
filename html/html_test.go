package html_test

import (
	. "github.com/hiromaily/golibs/html"
	//ht "github.com/hiromaily/golibs/http"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"net/http"
	"os"
	"testing"
)

var (
	benchFlg bool = false
)

var htmlData string = `
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
 <title>sample title</title>
</head>
<body>
 <p>sample html</p>
</body>
</html>
`

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Html_TEST]", "/var/log/go/test.log")
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
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestHTTPResponse(t *testing.T) {

	url := "http://info.finance.yahoo.co.jp/ranking/?kd=1&tm=d&mk=1"

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("http.Get error: %s", err)
	}
	//debug body
	//lg.Debugf("body: \n%s\n", ht.HandleResponse(res))
	//ParseToken(ParseHttpBody(res.Body), "a")

	//<input type="hidden" name="gintoken" value="{{ .gintoken }}">
	ParseToken(ParseHttpBody(res.Body), "input")

}

func TestHTMLString(t *testing.T) {
	nodes, err := ParseHTMLText(htmlData)
	if err != nil {
		t.Errorf("ParseHTMLText() error: %s", err)
	}
	ParseNode(nodes, "a")
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkHtml(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
