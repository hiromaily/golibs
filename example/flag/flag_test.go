package flag_test

import (
	"flag"
	. "github.com/hiromaily/golibs/example/flag"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

var (
	intVal = flag.Int("iv", 0, "this is just check val for int")
	strVal = flag.String("sv", "", "this is just check val for string")

	usage = `Usage: %s [options...] <url>

Options:
  -iv  Number of something.
  -sv  String of something
       bra bra bra.
`
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	SetUsage(usage)

	tu.InitializeTest("[FLAG]")
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
// Check
//-----------------------------------------------------------------------------
func TestInitFlag(t *testing.T) {
	tu.SkipLog(t)

	//SetUsage(usage)
	//flag.Usage = func() {
	//	fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	//}

	//command-line
	//flag.Parse()

	lg.Debugf("flag.NArg():%v", flag.NArg())
	lg.Debugf("flag.Args():%v", flag.Args())

	//show error
	ShowUsageAndExit("something error")
}

func TestInitFlag2(t *testing.T) {
	//tu.SkipLog(t)

	//./argtest FILE1 -opt1 aaa -opt2 bbb
	//fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//fs.Parse(os.Args[2:])

	//go test -v flag/flag_test.go -iv 1 -sv abcde
	for i, v := range os.Args {
		lg.Debugf("os.Args[%d]: %v", i, v)
	}

	//lg.Debugf("os.Args[]: %v", os.Args)     //flag.test
	//lg.Debugf("os.Args[1]: %v", os.Args[1]) //os.Args[1]: -test.v=true
	lg.Debugf("flag.NArg(): %v", flag.NArg())
	lg.Debugf("flag.Args(): %v", flag.Args())
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Bench
//-----------------------------------------------------------------------------
func BenchmarkFlag(b *testing.B) {
	tu.SkipBLog(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, v := range os.Args {
			lg.Debugf("os.Args[%d]: %v", i, v)
			//os.Args[0]: /var/folders/zw/fjz6w8n17c7_m47wjbzdqw1c0000gn/T/go-build666805607/github.com/hiromaily/golibs/flag/_test/flag.test
			//os.Args[1]: -test.bench=.
			//os.Args[2]: -test.benchmem=true
		}
	}
	b.StopTimer()
	//2654 ns/op
}
