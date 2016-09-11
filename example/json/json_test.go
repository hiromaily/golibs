package json_test

import (
	"encoding/json"
	"flag"
	"fmt"
	. "github.com/hiromaily/golibs/example/json"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

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

type TeacherInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type SiteInfo struct {
	Url      string        `json:"url"`
	Teachers []TeacherInfo `json:"teachers"`
}

var (
	jsonFile = flag.String("fp", "", "Json File Path")
	fileData []byte
)

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

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[EXEC]")

	if *jsonFile == "" {
		fmt.Println("json parameter is required to run.")
		os.Exit(1)
		return
	}
}

func setup() {
	var err error
	fileData, err = LoadJSONFile(*jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
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