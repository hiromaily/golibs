package json

import (
	"encoding/json"
	//u "github.com/hiromaily/golibs/utils"
	"errors"
	"io/ioutil"
)

//https://ukiahsmith.com/blog/go-marshal-and-unmarshal-json-with-time-and-url-data/

//-----------------------------------------------------------------------------
// This is just about how to use basic package
//-----------------------------------------------------------------------------
// Json Encode
func JsonEncode(jsonData string) ([]byte, error) {
	//return of Marshal : byte array encoded by UTF-8
	//when json of struct is exsiting, create json string from struct

	//func Marshal(v interface{}) ([]byte, os.Error)
	//bJson, err := json.Marshal(jsonData)
	return json.Marshal(jsonData)
}

// Json Encode
func JsonEncodeIndent(jsonData string) ([]byte, error) {
	//func MarshalIndent(v interface{}, prefix, indent string) ([]byte, os.Error)

	//bJson, err := json.MarshalIndent(jsonData, "", "\t")
	return json.MarshalIndent(jsonData, "", "\t")
}

// Json Analyze
func JsonAnalyze(jsonData string, retObj *map[string]interface{}) {
	//func Unmarshal(data []byte, v interface{}) os.Error
	json.Unmarshal([]byte(jsonData), retObj)
}

//-----------------------------------------------------------------------------
// Util
//-----------------------------------------------------------------------------
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
