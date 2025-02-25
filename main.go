package main

import (
	"fmt"
	"runtime"

	"github.com/ek-170/go-mod-sample/mystring"
)

func main() {
	const s = "Hello World"
	fmt.Println(s)
	fmt.Println(mystring.Reverse(s))
	fmt.Println(runtime.Version())
}
