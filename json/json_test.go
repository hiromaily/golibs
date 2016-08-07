package json_test

import (
	"encoding/json"
	"errors"
	"flag"
	//. "github.com/hiromaily/golibs/json"
	lg "github.com/hiromaily/golibs/log"
	"io/ioutil"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
	jsonFile = flag.String("fp", "", "Json File Path")
)

var fileData []byte

//https://scene-si.org/2016/06/13/advanced-go-tips-and-tricks/
//Nested structures when parsing JSON data
type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}

var jsonData string = `
{
    "id": 1,
    "name": "Tit Petric",
    "address": {
        "city": "Ljubljana",
        "country": "Slovenia"
    }
}
`

type TeacherInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type SiteInfo struct {
	Url      string        `json:"url"`
	Teachers []TeacherInfo `json:"teachers"`
}

func LoadJsonFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("Nothing Json File")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[JSON_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}

	var err error
	fileData, err = LoadJsonFile(*jsonFile)
	if err != nil {
		os.Exit(1)
		return
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	if *jsonFile == "" {
		os.Exit(1)
		return
	}

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Json
//-----------------------------------------------------------------------------
func TestJsonAsStruct(t *testing.T) {
	funcName := "TestJsonAsStruct"

	//defined as struct
	var siteInfo SiteInfo

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}

	if siteInfo.Url != "http://eikaiwa.dmm.com/" {
		t.Errorf("%s[02] siteInfo: %#v", funcName, siteInfo)
	}
}

func TestJsonAsMap(t *testing.T) {
	funcName := "TestJsonAsMap"

	//defined as map
	var siteInfo map[string]interface{}

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}

	if siteInfo["url"] != "http://eikaiwa.dmm.com/" {
		t.Errorf("%s[02] siteInfo: %#v", funcName, siteInfo)
	}
}

func TestJsonAsInterface(t *testing.T) {
	funcName := "TestJsonAsInterface"

	//defined as interface{}
	var siteInfo interface{}

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}

	if siteInfo.(map[string]interface{})["url"].(string) != "http://eikaiwa.dmm.com/" {
		t.Errorf("%s[02] siteInfo: %#v", funcName, siteInfo)
	}
}

func TestEmbededJsonData(t *testing.T) {
	funcName := "TestEmbededJsonData"

	//defined as interface{}
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}

	if person.Name != "Tit Petric" {
		t.Errorf("%s[02] siteInfo: %#v", funcName, person)
	}
}

func TestMarshalJson(t *testing.T) {
	funcName := "TestMarshalJson"

	//it can't work
	//missing type in composite literal
	//http://stackoverflow.com/questions/17912893/missing-type-in-composite-literal
	//person := Person{Id: 1, Name: "Harry", Address: {City: "Tokyo", Country: "Japan"}}
	//b0, err := json.Marshal(person)

	siteInfo := SiteInfo{Url: "http://google.com",
		Teachers: []TeacherInfo{{Id: 123, Name: "Harry", Country: "Japan"}, {Id: 456, Name: "Taro", Country: "America"}}}
	b, err := json.Marshal(siteInfo)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}
	if string(b) != "test" {
		t.Errorf("%s[02] siteInfo: %s", funcName, string(b))
	}

	b2, err := json.MarshalIndent(siteInfo, "", "\t")
	if err != nil {
		t.Fatalf("%s[03] error: %s", funcName, err)
	}
	if string(b2) != "test" {
		t.Errorf("%s[04] siteInfo: %s", funcName, string(b2))
	}
}
