package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	fmt.Println("Hello World")
// 	fmt.Println(stringutils.MyName)
// }

// func wrapper() func() int {
// 	x := 0

// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func main() {
// 	increment := wrapper()
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

func main() {
	res, _ := http.Get("https://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
