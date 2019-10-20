// Package flag is just sample
package flag

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// SetUsage is to set usage
func SetUsage(msg string) {
	//msg
	///var/folders/0b/v5tdhbhj58v9x_r2nxtfd_w00000gn/T/go-build296316815/command-line-arguments/_test/flag.test
	strs := strings.Split(os.Args[0], "/")

	// -h option
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(msg, strs[len(strs)-1]))
	}
}

// ShowUsageAndExit is show usage and exit program.
func ShowUsageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Usage()

	os.Exit(1)
}
