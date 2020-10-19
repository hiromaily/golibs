package serial_test

import (
	"encoding/hex"
	"os"
	"testing"

	. "github.com/hiromaily/golibs/serial"
	tu "github.com/hiromaily/golibs/testutil"
)

type User struct {
	ID   int
	Name string
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Serial]")
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
//-----------------------------------------------------------------------------
// Gob
//-----------------------------------------------------------------------------
//1.struct
func TestSerializeStruct(t *testing.T) {
	t.Skip("skipping TestSerializeStruct")

	u := User{ID: 10, Name: "harry dayo"}
	result, err := ToGOB64(u)
	if err != nil {
		t.Errorf("ToGOB64() error: %s", err)
	}
	if result != "Iv+BAwEBBFVzZXIB/4IAAQIBAklkAQQAAQROYW1lAQwAAAAR/4IBFAEKaGFycnkgZGF5bwA=" {
		t.Errorf("ToGOB64 result: %+v", result)
	}
}

func TestDeSerializeStruct(t *testing.T) {
	t.Skip("skipping TestDeSerializeStruct")

	u := User{}
	err := FromGOB64("Iv+BAwEBBFVzZXIB/4IAAQIBAklkAQQAAQROYW1lAQwAAAAR/4IBFAEKaGFycnkgZGF5bwA=", &u)
	if err != nil {
		t.Errorf("FromGOB64 error: %s", err)
	}
	if u.ID != 10 {
		t.Errorf("FromGOB64 result: %+v", u)
	}
}

//2.map (map is not stable.)
//Iteration order is not guaranteed
//https://blog.golang.org/go-maps-in-action
func TestSerializeMap(t *testing.T) {
	t.Skip("skipping TestSerializeMap")

	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	//when using map, result is not stable.
	result, err := ToGOB64(m)
	if err != nil {
		t.Errorf("ToGOB64 error: %s", err)
	}
	if result != "Dv+DBAEC/4QAAQwBBAAAIP+EAAMFYXBwbGX+ASwGYmFuYW5h/gJYBWxlbW9u/gJY" {
		//if result != "Dv+DBAEC/4QAAQwBBAAAIP+EAAMGYmFuYW5h/gJYBWxlbW9u/gJYBWFwcGxl/gEs" {
		t.Errorf("ToGOB64 result: %#v", result)
	}
}

//
func TestDeSerializeMap(t *testing.T) {
	t.Skip("skipping TestDeSerializeMap")

	m := map[string]int{}
	//FromGOB64("Dv+DBAEC/4QAAQwBBAAAIP+EAAMFYXBwbGX+ASwGYmFuYW5h/gJYBWxlbW9u/gJY", &u)
	err := FromGOB64("Dv+DBAEC/4QAAQwBBAAAIP+EAAMGYmFuYW5h/gJYBWxlbW9u/gJYBWFwcGxl/gEs", &m)
	if err != nil {
		t.Errorf("FromGOB64 error: %s", err)
	}
	if m["apple"] != 150 {
		t.Errorf("FromGOB64 result: %#v", m)
	}
}

//Iteration order is not guaranteed
//https://blog.golang.org/go-maps-in-action
//TODO:how can map order be guaranteed.
func TestSerializeMap2(t *testing.T) {
	t.Skip("skipping TestSerializeMap")

	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	//when using map, result is not stable.
	//TODO:convert map data to something others type.
	result, err := ToGOB64(m)
	if err != nil {
		t.Errorf("ToGOB64 error: %s", err)
	}
	if result != "Dv+DBAEC/4QAAQwBBAAAIP+EAAMFYXBwbGX+ASwGYmFuYW5h/gJYBWxlbW9u/gJY" {
		//if result != "Dv+DBAEC/4QAAQwBBAAAIP+EAAMGYmFuYW5h/gJYBWxlbW9u/gJYBWFwcGxl/gEs" {
		t.Errorf("ToGOB64 result: %#v", result)
	}
}

//interface x slice
func TestSerializeInterfaces(t *testing.T) {
	t.Skip("skipping TestSerializeInterfaces")

	inf := []interface{}{1, "abcde", true}
	result, err := ToGOB64(inf)
	if err != nil {
		t.Errorf("ToGOB64 error: %s", err)
	}
	if result != "DP+BAgEC/4IAARAAACX/ggADA2ludAQCAAIGc3RyaW5nDAcABWFiY2RlBGJvb2wCAgAB" {
		t.Errorf("ToGOB64 result: %+v", result)
	}
}

func TestDeSerializeInterfaces(t *testing.T) {
	t.Skip("skipping TestDeSerializeInterfaces")

	inf := []interface{}{}
	err := FromGOB64("DP+BAgEC/4IAARAAACX/ggADA2ludAQCAAIGc3RyaW5nDAcABWFiY2RlBGJvb2wCAgAB", &inf)
	if err != nil {
		t.Errorf("FromGOB64 error: %s", err)
	}
	if inf[0] != 1 || inf[1] != "abcde" || inf[2] != true {
		t.Errorf("FromGOB64 result: %+v", inf)
	}
}

// map[string]interface{} x slice
func TestSerializeMapInterface(t *testing.T) {
	t.Skip("skipping TestSerializeMapInterface")

	//[]map[string]interface{}
	args := make([]map[string]interface{}, 3)
	args[0] = map[string]interface{}{"field1": 10, "field2": "somethings", "field3": true}
	args[1] = map[string]interface{}{"field1": 15, "field2": "testdata", "field3": false}
	args[2] = map[string]interface{}{"field1": 30, "field2": "vvvvvvvv", "field3": true}
	result, err := ToGOB64(args)
	if err != nil {
		t.Errorf("ToGOB64 error: %s", err)
	}
	if result != "Df+FAgEC/4YAAf+EAAAO/4MEAQL/hAABDAEQAAD/tP+GAAMDBmZpZWxkMQNpbnQEAgAUBmZpZWxkMgZzdHJpbmcMDAAKc29tZXRoaW5ncwZmaWVsZDMEYm9vbAICAAEDBmZpZWxkMQNpbnQEAgAeBmZpZWxkMgZzdHJpbmcMCgAIdGVzdGRhdGEGZmllbGQzBGJvb2wCAgAAAwZmaWVsZDEDaW50BAIAPAZmaWVsZDIGc3RyaW5nDAoACHZ2dnZ2dnZ2BmZpZWxkMwRib29sAgIAAQ==" &&
		result != "Df+DAgEC/4QAAf+CAAAO/4EEAQL/ggABDAEQAAD/tP+EAAMDBmZpZWxkMQNpbnQEAgAUBmZpZWxkMgZzdHJpbmcMDAAKc29tZXRoaW5ncwZmaWVsZDMEYm9vbAICAAEDBmZpZWxkMQNpbnQEAgAeBmZpZWxkMgZzdHJpbmcMCgAIdGVzdGRhdGEGZmllbGQzBGJvb2wCAgAAAwZmaWVsZDEDaW50BAIAPAZmaWVsZDIGc3RyaW5nDAoACHZ2dnZ2dnZ2BmZpZWxkMwRib29sAgIAAQ==" &&
		result != "Df+DAgEC/4QAAf+CAAAO/4EEAQL/ggABDAEQAAD/tP+EAAMDBmZpZWxkMwRib29sAgIAAQZmaWVsZDEDaW50BAIAFAZmaWVsZDIGc3RyaW5nDAwACnNvbWV0aGluZ3MDBmZpZWxkMQNpbnQEAgAeBmZpZWxkMgZzdHJpbmcMCgAIdGVzdGRhdGEGZmllbGQzBGJvb2wCAgAAAwZmaWVsZDEDaW50BAIAPAZmaWVsZDIGc3RyaW5nDAoACHZ2dnZ2dnZ2BmZpZWxkMwRib29sAgIAAQ==" {
		t.Errorf("ToGOB64 result: %#v", result)
	}
}

func TestDeSerializeMapInterface(t *testing.T) {
	t.Skip("skipping TestDeSerializeInterfaces")

	//TODO:check make is required or not.
	//args := make([]map[string]interface{}, 3)
	args := []map[string]interface{}{}
	err := FromGOB64("Df+FAgEC/4YAAf+EAAAO/4MEAQL/hAABDAEQAAD/tP+GAAMDBmZpZWxkMQNpbnQEAgAUBmZpZWxkMgZzdHJpbmcMDAAKc29tZXRoaW5ncwZmaWVsZDMEYm9vbAICAAEDBmZpZWxkMQNpbnQEAgAeBmZpZWxkMgZzdHJpbmcMCgAIdGVzdGRhdGEGZmllbGQzBGJvb2wCAgAAAwZmaWVsZDEDaW50BAIAPAZmaWVsZDIGc3RyaW5nDAoACHZ2dnZ2dnZ2BmZpZWxkMwRib29sAgIAAQ==", &args)
	if err != nil {
		t.Errorf("FromGOB64 error: %s", err)
	}
	if args[0]["field1"] != 10 || args[0]["field2"] != "somethings" || args[0]["field3"] != true {
		t.Errorf("FromGOB64 result: %+v", args)
	}
	if args[1]["field1"] != 15 || args[1]["field2"] != "testdata" || args[1]["field3"] != false {
		t.Errorf("FromGOB64 result: %+v", args)
	}
}

//-----------------------------------------------------------------------------
//github.com/ugorji/go/codec
//-----------------------------------------------------------------------------
func TestEncodeStruct(t *testing.T) {
	//t.Skip("skipping TestEncodeStruct")
	//*
	u := User{ID: 10, Name: "harry dayo"}
	byteData := CodecEncode(u)

	//if fmt.Sprintf("%x", byteData) != "82a249640aa44e616d65aa6861727279206461796f" {
	if hex.EncodeToString(byteData) != "82a249640aa44e616d65aa6861727279206461796f" {
		t.Errorf("CodecEncode result: %x", byteData)
	}
}

func TestDcodeStruct(t *testing.T) {
	//t.Skip("skipping TestDcodeStruct")

	u := User{}
	_ = CodecDecode("82a249640aa44e616d65aa6861727279206461796f", &u)
	if u.ID != 10 {
		t.Errorf("CodecDecode result: %+v", u)
	}
}

//Iteration order is not guaranteed
//https://blog.golang.org/go-maps-in-action
func TestEncodeMap(t *testing.T) {
	t.Skip("skipping TestEncodeMap")
	//*
	m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
	byteData := CodecEncode(m)

	if hex.EncodeToString(byteData) != "83a56170706c65cc96a662616e616e61cd012ca56c656d6f6ecd012c" {
		t.Errorf("CodecEncode result: %x", byteData)
	}
}

func TestDcodeMap(t *testing.T) {
	t.Skip("skipping TestDcodeMap")

	m := map[string]int{}
	_ = CodecDecode("83a56170706c65cc96a662616e616e61cd012ca56c656d6f6ecd012c", &m)

	if m["apple"] != 150 {
		t.Errorf("CodecDecode result: %+v", m)
	}
}

//-----------------------------------------------------------------------------
// Bench
//-----------------------------------------------------------------------------
func BenchmarkSerializeStruct(b *testing.B) {
	//b.Skip("skipping BenchmarkSerializeStruct")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := User{ID: 10, Name: "harry dayo"}
		ToGOB64(u)
	}
	b.StopTimer()
	//4774 ns/op
}

func BenchmarkDeSerializeStruct(b *testing.B) {
	//b.Skip("skipping BenchmarkDeSerializeStruct")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := User{}
		FromGOB64("Iv+BAwEBBFVzZXIB/4IAAQIBAklkAQQAAQROYW1lAQwAAAAR/4IBFAEKaGFycnkgZGF5bwA=", &u)
	}
	b.StopTimer()
	//33301 ns/op
}

func BenchmarkSerializeMap(b *testing.B) {
	//b.Skip("skipping BenchmarkSerializeMap")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := map[string]int{"apple": 150, "banana": 300, "lemon": 300}
		ToGOB64(m)
	}
	b.StopTimer()
	//5021 ns/op
}

func BenchmarkDeSerializeMap(b *testing.B) {
	//b.Skip("skipping BenchmarkDeSerializeMap")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := map[string]int{}
		FromGOB64("Dv+DBAEC/4QAAQwBBAAAIP+EAAMGYmFuYW5h/gJYBWxlbW9u/gJYBWFwcGxl/gEs", &m)
	}
	b.StopTimer()
	//33445 ns/op
}

func BenchmarkEncodeStruct(b *testing.B) {
	//b.Skip("skipping BenchmarkEncodeStruct")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := User{ID: 10, Name: "harry dayo"}
		hex.EncodeToString(CodecEncode(u))
	}
	b.StopTimer()
	//2259 ns/op
}

func BenchmarkDcodeStruct(b *testing.B) {
	//b.Skip("skipping BenchmarkDcodeStruct")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := User{}
		CodecDecode("82a249640aa44e616d65aa6861727279206461796f", &u)
	}
	b.StopTimer()
	//2654 ns/op
}
