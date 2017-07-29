package main

import (
	"flag"
	"fmt"
)

var (
	str  = flag.String("s", "", "String")
	num  = flag.Int("n", 0, "Int")
	bool = flag.Bool("b", false, "Bool") //it's true when there is any value
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("[Debug]str is %s\n", *str)
	fmt.Printf("[Debug]num is %d\n", *num)

	fmt.Printf("[Debug]bool is %t\n", *bool)
	fmt.Printf("[Debug]bool is %v\n", *bool)
}
