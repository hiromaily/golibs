package deferpkg

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Println("call first()")
	fmt.Println(first()) //4
	fmt.Println("call second()")
	fmt.Println(second())
}

func first() (code int) {
	code = 1
	defer func() {
		fmt.Println(code) //3
		code = 4
	}()
	code = 2
	return 3 //this return means code = 3; return
}

func second() (code int) {
	code = 1
	defer func() {
		fmt.Println(code)
		//return 4 this is invalid
		//code can be changed in defer func
	}()
	code = 2
	return 3
}
