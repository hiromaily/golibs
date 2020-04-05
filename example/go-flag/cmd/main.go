package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

// Options is command line options
type Options struct {
	Cmd1 struct {
		IsSet bool `short:"i"`
	} `command:"cmd1"`
	Cmd2 struct {
		Key bool `short:"k"`
	} `command:"cmd2"`
	//Configパス
	Path string `short:"p" long:"path" default:"" description:"Path"`
}

var (
	opts Options
)

func init() {
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("cmd1: ", opts.Cmd1)
	fmt.Println("cmd2: ", opts.Cmd2)
	fmt.Println("path: ", opts.Path)
	if opts.Cmd1.IsSet {
		fmt.Println("path: ", opts.Cmd1.IsSet)
	}
}
