package join_test

import (
	//. "github.com/hiromaily/golibs/join"
	"bytes"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
	"os"
	"strings"
	"testing"
)

//http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37

var (
	benchFlg bool = false
	m             = []string{
		"AAAAAAAAAAAAAAAAAAAAAAAAAA",
		"BBBBBBBBBBBBBBBBBBBBBBBBBB",
		"CCCCCCCCCCCCCCCCCCCCCCCCCC",
		"DDDDDDDDDDDDDDDDDDDDDDDDDD",
		"EEEEEEEEEEEEEEEEEEEEEEEEEE",
		"FFFFFFFFFFFFFFFFFFFFFFFFFF",
		"GGGGGGGGGGGGGGGGGGGGGGGGGG",
	}
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[JOIN_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
}

func setup() {
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
func BenchmarkCapByteArray_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = make([]byte, 0, 100)
		for _, v := range m {
			m2 = append(m2, v...)
			m2 = append(m2, ',')
		}
		_ = string(m2)
	}
	//271 ns/op
}

func BenchmarkHardCoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m[0] + "," + m[1] + "," + m[2] + "," + m[3] + "," + m[4] + "," + m[5] + "," + m[6]
	}
	//272 ns/op
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_ = strings.Join(m, ",")
	}
	//274 ns/op
}

func BenchmarkStringsJoin2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = []string{
			"AAAAAAAAAAAAAAAAAAAAAAAAAA",
			"BBBBBBBBBBBBBBBBBBBBBBBBBB",
			"CCCCCCCCCCCCCCCCCCCCCCCCCC",
			"DDDDDDDDDDDDDDDDDDDDDDDDDD",
			"EEEEEEEEEEEEEEEEEEEEEEEEEE",
			"FFFFFFFFFFFFFFFFFFFFFFFFFF",
			"GGGGGGGGGGGGGGGGGGGGGGGGGG",
		}
		_ = strings.Join(m2, ",")
	}
	//282 ns/op
}

func BenchmarkByteArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 []byte
		for _, v := range m {
			m2 = append(m2, v...)
			m2 = append(m2, ',')
		}
		_ = string(m2)
	}
	//515 ns/op
}

func BenchmarkCapBytesBuffer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		for _, v := range m {
			m2.WriteString(v)
			m2.WriteString(",")
		}
		_ = m2.String()
	}
	//661 ns/op
}

func BenchmarkCapBytesBuffer_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		for _, v := range m {
			m2.Write([]byte(v))
			m2.Write([]byte{','})
		}
		_ = m2.String()
	}
	//752 ns/op
}

func BenchmarkBytesBuffer____(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 bytes.Buffer
		for _, v := range m {
			m2.Write([]byte(v))
			m2.Write([]byte{','})
		}
		_ = m2.String()
	}
	//900 ns/op
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", m[0], m[1], m[2], m[3], m[4], m[5], m[6])
	}
	//956 ns/op
}

func BenchmarkAppendOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m2 string
		for _, v := range m {
			m2 += m2 + "," + v
		}
	}
	//2482 ns/op
}
