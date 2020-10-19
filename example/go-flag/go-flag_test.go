package go_flag

import (
	"testing"

	"github.com/jessevdk/go-flags"
)

// Options is command line options
type Options struct {
	Cmd1 struct {
		Param bool `short:"g"`
	} `command:"cmd1"`
	Cmd2 struct {
		Key bool `short:"k"`
	} `command:"cmd2"`
	//Configパス
	Path string `short:"d" long:"path" default:"" description:"Path"`
}

var (
	opts Options
)

func init() {
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}
}

func TestGoFlag(t *testing.T) {
	t.Log(opts)
}
