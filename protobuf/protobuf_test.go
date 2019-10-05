package protobuf_test

import (
	//"bytes"
	"log"
	"testing"
	"time"

	//"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	jsoniter "github.com/json-iterator/go"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
)

func TestSample(t *testing.T) {
	// For NormalType
	normal := &samplepb.NormalType{
		I32: 1,
		I64: time.Now().Unix(),
		U32: 3,
		U64: 4,
		Fl:  1.2,
		Db:  123.123,
		Bl:  true,
		St:  "foobar",
		Bt:  []byte{0x00, 0x01, 0x02, 0x03},
	}

	client := &samplepb.Client{
		QuestionCode: 1,
		Name:         "Mike",
	}

	extension := &samplepb.ExtensionType{
		Data: []string{"hello", "friend"},
		Mp: map[string]int64{
			"a": 123,
			"b": 125,
			"c": 127,
		},
		Ts:     &types.Timestamp{Seconds: 1296000000, Nanos: 0}, //*types.Timestamp
		Struct: &types.Struct{},                                 //*types.Struct
		Any:    &types.Any{TypeUrl: "", Value: []byte{}},        //*types.Any
		//OneofData: &samplepb.ExtensionType_I32{I32: 12345},
		OneofData: &samplepb.ExtensionType_St{St: "12345"},
		Type:      samplepb.SOMETHING_TYPE_A,
		Client:    client,
	}

	// Marshal
	out, err := proto.Marshal(normal)
	if err != nil {
		log.Fatalln("Failed to encode normal:", err)
	}
	log.Println(out)

	// Unmarshal
	normal2 := &samplepb.NormalType{}
	if err := proto.Unmarshal(out, normal2); err != nil {
		log.Fatalln("Failed to parse normal2:", err)
	}
	log.Println(normal2)

	// Marshal
	out, err = proto.Marshal(extension)
	if err != nil {
		log.Fatalln("Failed to encode extension:", err)
	}
	log.Println(out)

	// condition for oneof value
	switch extension.OneofData.(type) {
	case *samplepb.ExtensionType_I32:
		log.Println("type of oneof is ExtensionType_I32")
		if v, ok := extension.OneofData.(*samplepb.ExtensionType_I32); ok {
			log.Println(v.I32)
		}
	case *samplepb.ExtensionType_St:
		log.Println("type of oneof is ExtensionType_St")
		if v, ok := extension.OneofData.(*samplepb.ExtensionType_St); ok {
			log.Println(v.St)
		}
	}

	// Marshal as JSON
	js, err := jsoniter.Marshal(&extension)
	if err != nil {
		log.Fatalln("Failed to encode sampleBase:", err)
	}
	log.Println(string(js))

	// Unmarshal
	// https://qiita.com/yugui/items/238dcdb75cd40d0f1ece
	//var sampleBase2 = samplepb.SampleBase{}
	////func Unmarshal(r io.Reader, pb proto.Message) error {
	//if err := jsonpb.Unmarshal(bytes.NewReader(js), &sampleBase2); err != nil {
	//	log.Fatalln("Failed to parse sampleBase2:", err)
	//	//unknown field "SampleData" in samplepb.SampleBase
	//}
	//log.Println(extension)
}
