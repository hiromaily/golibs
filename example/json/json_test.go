package json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	. "github.com/hiromaily/golibs/example/json"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//https://scene-si.org/2016/06/13/advanced-go-tips-and-tricks/
//Nested structures when parsing JSON data
type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}

type TeacherInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type SiteInfo struct {
	URL      string        `json:"url"`
	Teachers []TeacherInfo `json:"teachers"`
}

var (
	fileData []byte
)

var jsonData = `
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
	tu.InitializeTest("[JSON]")
}

func setup() {
	var err error
	fileData, err = LoadJSONFile(*tu.JSONFile)
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

	//defined as struct
	var siteInfo SiteInfo

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("[01] Unmarshal error: %s", err)
	}

	if siteInfo.URL != "http://eikaiwa.dmm.com/" {
		t.Errorf("[02] siteInfo: %#v", siteInfo)
	}
}

func TestJsonAsMap(t *testing.T) {

	//defined as map
	var siteInfo map[string]interface{}

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("[01] Unmarshal error: %s", err)
	}

	if siteInfo["url"] != "http://eikaiwa.dmm.com/" {
		t.Errorf("[02] siteInfo: %#v", siteInfo)
	}
}

func TestJsonAsInterface(t *testing.T) {

	//defined as interface{}
	var siteInfo interface{}

	err := json.Unmarshal(fileData, &siteInfo)
	if err != nil {
		t.Fatalf("[01] Unmarshal error: %s", err)
	}

	if siteInfo.(map[string]interface{})["url"].(string) != "http://eikaiwa.dmm.com/" {
		t.Errorf("[02] siteInfo: %#v", siteInfo)
	}
}

func TestEmbededJsonData(t *testing.T) {

	//defined as interface{}
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		t.Fatalf("[01] Unmarshal error: %s", err)
	}

	if person.Name != "Tit Petric" {
		t.Errorf("[02] siteInfo: %#v", person)
	}
}

func TestMarshalJson(t *testing.T) {

	//it can't work
	//missing type in composite literal
	//http://stackoverflow.com/questions/17912893/missing-type-in-composite-literal
	//person := Person{Id: 1, Name: "Harry", Address: {City: "Tokyo", Country: "Japan"}}
	//b0, err := json.Marshal(person)

	siteInfo := SiteInfo{URL: "http://google.com",
		Teachers: []TeacherInfo{{ID: 123, Name: "Harry", Country: "Japan"}, {ID: 456, Name: "Taro", Country: "America"}}}
	b, err := json.Marshal(siteInfo)
	if err != nil {
		t.Fatalf("[01] error: %s", err)
	}
	lg.Debugf("siteInfo: %s", string(b))

	b2, err := json.MarshalIndent(siteInfo, "", "\t")
	if err != nil {
		t.Fatalf("[02] error: %s", err)
	}
	lg.Debugf("siteInfo: %s", string(b2))
}

func TestLoadWithDecode(t *testing.T) {
	var siteInfo SiteInfo
	err := LoadWithDecode(*tu.JSONFile, &siteInfo)
	if err != nil {
		t.Fatalf("[01] error: %s", err)
	}
	lg.Debugf("siteInfo: %#v", siteInfo)
}
