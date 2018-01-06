package color

import (
	"fmt"
)

type Colour string

var (
	Reset  Colour = "\x1B[0m"
	Red    Colour = "\x1B[38;5;124m"
	Yellow Colour = "\x1B[38;5;208m"
	Blue   Colour = "\x1B[38;5;33m"
	Grey   Colour = "\x1B[38;5;144m"
	Green  Colour = "\x1B[38;5;34m"
	Gold   Colour = "\x1B[38;5;3m"
)

func Addf(format string, c Colour, a ...interface{}) string {
	return Add(fmt.Sprintf(format, a...), c)
}

func Add(str string, c Colour) string {
	return string(c) + str + string(Reset)
}
