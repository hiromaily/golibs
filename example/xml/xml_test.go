// Package xml_test is just sample
package xml_test

import (
	"encoding/xml"
	"errors"
	"flag"

	//lg "github.com/hiromaily/golibs/log"
	//. "github.com/hiromaily/golibs/xml"
	"io/ioutil"
	"os"
	"testing"

	tu "github.com/hiromaily/golibs/testutil"
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
	xmlFile = flag.String("fp", "", "XML File Path")
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[XML]")

	//check file
	if *xmlFile == "" {
		//default
		*xmlFile = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/example/xml/rssfeeds/"
		//lg.Error("xml file parameter is required to run")
		//os.Exit(1)
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
// LoadXMLFile is to load XML file
func loadXMLFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("Nothing XML File")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestReadXml(t *testing.T) {

	xmlData, err := loadXMLFile(*xmlFile)
	if err != nil {
		t.Fatal("xml load error")
		return
	}

	//defined as struct
	var rss Rss
	err = xml.Unmarshal(xmlData, &rss)
	if err != nil {
		t.Errorf("[01]XML Unmarshal error: %s", err)
	}

	if rss.Root.Title != "TechCrunch" {
		t.Errorf("[02]Result of XML Value is incorrect: %#v", rss)
	}

	//if *tu.LogFlg == 1 {
	//	lg.Info("Testoooooooooooo")
	//}
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
