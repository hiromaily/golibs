package main

import (
	"fmt"

	"github.com/hiromaily/golibs/execution_order/sub"
)

func init() {
	fmt.Println("main.go: init()")
}

func main() {
	fmt.Println("main.go: main()")

	sub.Something()

	//1. sub/sub.go: init()
	//2. aaa.go: init()
	//3. main.go: init()
	//4. sub1.go: init()
	//5. sub2.go: init()
	//6. main.go: main()
	//7. sub/sub.go: something()

	//First,  except main, imported package's init() is called.
	//Second, if main package includes several go files, these are executed by alphabetical order.
}
