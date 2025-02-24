package main

import (
	"fmt"
	"go-mod-sample/mystring"
	"runtime"
)

func main() {
	const s = "Hello World"
	fmt.Println(s)
	fmt.Println(mystring.Reverse(s))
	fmt.Println(runtime.Version())
}
