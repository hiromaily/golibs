package cast_test

import (
	"os"
	"testing"

	. "github.com/hiromaily/golibs/cast"
	tu "github.com/hiromaily/golibs/testutil"
)

var (
	testStr  = "aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeee"
	testByte = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Cast]")
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
// Bench
//-----------------------------------------------------------------------------
//String to []Byte
func BenchmarkStoB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StoB(testStr)
	}
	//56.7 ns/op
}

func BenchmarkBufferStoB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BufferStoB(testStr)
	}
	//153 ns/op
}

// []Byte to String
func BenchmarkBtoS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BtoS(testByte)
	}
	//19.6 ns/op
}

func BenchmarkBufferBtoS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BufferBtoS(testByte)
	}
	//331 ns/op
}
