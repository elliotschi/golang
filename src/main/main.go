package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Hello World")
// 	fmt.Println(stringutils.MyName)
// }

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
