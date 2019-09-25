package protobuf_test

import (
	"log"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	jsoniter "github.com/json-iterator/go"

	samplepb "github.com/hiromaily/golibs/protobuf/pb/sample"
)

func TestSample(t *testing.T) {
	// For samplepb.Category1Sample
	var cg1Sample = samplepb.Category1Sample{
		Something1: &samplepb.Something1{
			Id:   10,
			UId:  20,
			Type: samplepb.SOMETHING_TYPE_A,
			Data: []string{"abc", "def", "ghi"},
		},
		Something2: &samplepb.Something2{
			Name:  "Ron",
			PTime: 10000,
		},
		Client: &samplepb.Client{
			Name:   "Mike",
			Age:    30,
			Height: 170,
		},
		Md: map[string]int64{
			"a": 123,
			"b": 125,
			"c": 127,
		},
		LastUpdated: &types.Timestamp{Seconds: 1296000000, Nanos: 0},
		//Definition:
		Details: &types.Any{TypeUrl: "", Value: []byte{}},
	}

	// Marshal
	out, err := proto.Marshal(&cg1Sample)
	if err != nil {
		log.Fatalln("Failed to encode address cg1Sample:", err)
	}
	log.Println(out)

	// Unmarshal
	cg1Sample2 := &samplepb.Category1Sample{}
	if err := proto.Unmarshal(out, cg1Sample2); err != nil {
		log.Fatalln("Failed to parse cg1Sample2:", err)
	}
	log.Println(cg1Sample2)

	// For samplepb.Category2Sample
	var cg2Sample = samplepb.Category2Sample{
		Something1: &samplepb.Something1{
			Id:   10,
			UId:  20,
			Type: samplepb.SOMETHING_TYPE_B,
			Data: []string{"abc", "def", "ghi"},
		},
		AId:    999,
		BId:    0.62,
		SName:  "something-name",
		BFlag:  true,
		BtData: []byte{},
	}

	// Marshal
	out, err = proto.Marshal(&cg2Sample)
	if err != nil {
		log.Fatalln("Failed to encode address cg2Sample:", err)
	}
	log.Println(out)

	// SampleBase
	var sampleBase = samplepb.SampleBase{
		SampleName: "sample01",
		Time: time.Now().Unix(),
		SampleData: &samplepb.SampleBase_Category1{Category1: &cg1Sample},
	}

	// Marshal
	out, err = proto.Marshal(&sampleBase)
	if err != nil {
		log.Fatalln("Failed to encode sampleBase:", err)
	}
	log.Println(sampleBase)

	// condition for oneof value
	switch sampleBase.SampleData.(type) {
	case *samplepb.SampleBase_Category1:
		log.Println("type of oneof is SampleBase_Category1")
		if v, ok := sampleBase.SampleData.(*samplepb.SampleBase_Category1); ok {
			log.Println(v.Category1)
		}
	case *samplepb.SampleBase_Category2:
		log.Println("type of oneof is SampleBase_Category2")
		if v, ok := sampleBase.SampleData.(*samplepb.SampleBase_Category2); ok {
			log.Println(v.Category2)
		}
	}

	// Marshal as JSON
	js, err := jsoniter.Marshal(&sampleBase)
	if err != nil {
		log.Fatalln("Failed to encode sampleBase:", err)
	}
	log.Println(string(js))
}
