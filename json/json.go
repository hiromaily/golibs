package json

import (
	"encoding/json"
	"github.com/hiromaily/golibs/utils"
)

// Json Encode
func JsonEncode(jsonData string) []byte {
	//return of Marshal : byte array encoded by UTF-8
	//when json of struct is exsiting, create json string from struct

	//func Marshal(v interface{}) ([]byte, os.Error)
	bJson, err := json.Marshal(jsonData)

	//error check
	utils.GoPanicWhenError(err)

	return bJson
}

// Json Encode
func JsonEncodeIndent(jsonData string) []byte {

	//func MarshalIndent(v interface{}, prefix, indent string) ([]byte, os.Error)
	b_json, err := json.MarshalIndent(jsonData, "", "\t")

	//error check
	utils.GoPanicWhenError(err)

	return b_json
}

// Json Analyze
func JsonAnalyze(jsonData string, retObj *map[string]interface{}) {
	//func Unmarshal(data []byte, v interface{}) os.Error
	json.Unmarshal([]byte(jsonData), retObj)
}
