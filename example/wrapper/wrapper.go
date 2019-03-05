package wrapper

import (
	"fmt"
)

func Parent(param1 int, param2 string) func() {
	return func() {
		fmt.Println("param1:", param1)
		fmt.Println("param2:", param2)
	}
}
