package cast_test

import (
	"flag"
	. "github.com/hiromaily/golibs/cast"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"testing"
)

var (
	benchFlg = flag.Int("bc", 0, "Normal Test or Bench Test")
)

var testStr = "aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeee"
var testByte = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func setup() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[CAST_TEST]", "/var/log/go/test.log")
	if *benchFlg == 0 {
	}
}

func teardown() {
	if *benchFlg == 0 {
	}
}

// Initialize
func TestMain(m *testing.M) {
	flag.Parse()

	//TODO: According to argument, it switch to user or not.
	//TODO: For bench or not bench
	setup()

	code := m.Run()

	teardown()

	// 終了
	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Operator
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
