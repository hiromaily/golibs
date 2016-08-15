package flag

import (
	"flag"
	"fmt"
	"os"
)

var (
	intVal = flag.Int("iv", 0, "this is just check val for int")
	strVal = flag.String("sv", "", "this is just check val for string")
)

var usage = `Usage: %s [options...] <url>

Options:
  -iv  Number of something.
  -sv  String of something
       bra bra bra.
`

// Set Usage
func SetUsage(msg string) {
	// -h option
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(msg, os.Args[0]))
	}
}

// Show usage before exit
func ShowUsageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()

	os.Exit(1)
}

//just example
func GetIntVal() int {
	return *intVal
}

//just example
func GetStrVal() string {
	return *strVal
}