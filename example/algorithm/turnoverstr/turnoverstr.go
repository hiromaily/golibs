package main

import (
	"fmt"
)

var strData = "test data da yo"

func main() {
	//strを配列に変換
	//runes := []rune(strData)

	ret := reverse(strData)
	fmt.Println(ret)
}

func reverse(s string) string {
	//1.配列に変換
	runes := []rune(s)

	//2. 最前の文字と最後の文字を入れ替えるが、文字数分のループは行わない
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
