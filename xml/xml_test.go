package xml_test

import (
	"encoding/xml"
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	. "github.com/hiromaily/golibs/xml"
	"os"
	"testing"
)

type Rss struct {
	Root Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	//Date  time.Time `xml:"lastBuildDate"`
	Date  string `xml:"lastBuildDate"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"pubDate"`
}

var (
	xmlFile       = flag.String("fp", "", "XML File Path")
	benchFlg bool = false
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	flag.Parse()

	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[Xml_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}

	if *xmlFile == "" {
		fmt.Println("xml file parameter is required to run.")
		os.Exit(1)
		return
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
func TestReadXml(t *testing.T) {

	xmlData, err := LoadXmlFile(*xmlFile)
	if err != nil {
		t.Fatal("xml load error.")
		return
	}

	//defined as struct
	var rss Rss
	err = xml.Unmarshal(xmlData, &rss)
	if err != nil {
		t.Errorf("XML Unmarshal error: %s", err)
	}

	if rss.Root.Title != "TechCrunch" {
		t.Errorf("%#v", rss)
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkXml(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
