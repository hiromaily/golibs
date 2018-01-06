package color

import (
	"fmt"
)

type Col string

//https://misc.flogisoft.com/bash/tip_colors_and_formatting
//http://jafrog.com/2013/11/23/colors-in-terminal.html
var (
	Reset       Col = "\x1B[0m"
	Red         Col = "\x1B[38;5;1m"
	Orange      Col = "\x1B[38;5;9m"
	DeepPink    Col = "\x1B[38;5;5m"
	Pink        Col = "\x1B[38;5;13m"
	Yellow      Col = "\x1B[38;5;3m"
	Green       Col = "\x1B[38;5;2m"
	SpringGreen Col = "\x1B[38;5;10m"
	Blue        Col = "\x1B[38;5;4m"
	DeepSkyBlue Col = "\x1B[38;5;6m"
	SkyBlue     Col = "\x1B[38;5;14m"
	Grey        Col = "\x1B[38;5;8m"
	Black       Col = "\x1B[38;5;232m"
	White       Col = "\x1B[38;5;15m"
)

func Addf(format string, c Col, a ...interface{}) string {
	return Add(fmt.Sprintf(format, a...), c)
}

func Add(str string, c Col) string {
	return string(c) + str + string(Reset)
}

func Check() {
	for i := 1; i < 256; i++ {
		fmt.Println(Add(fmt.Sprintf("Number %d:", i), Col(fmt.Sprintf("\x1B[38;5;%dm", i))))
	}
}
