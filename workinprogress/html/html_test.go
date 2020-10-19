package html_test

import (
	. "github.com/hiromaily/golibs/workinprogress/html"
	//ht "github.com/hiromaily/golibs/http"
	//lg "github.com/hiromaily/golibs/log"
	"net/http"
	"os"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
)

var htmlData = `
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

func setup() {
	tu.InitializeTest("[HTML]")
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
	ParseToken(ParseHTTPBody(res.Body), "input")

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
func BenchmarkHTML(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
