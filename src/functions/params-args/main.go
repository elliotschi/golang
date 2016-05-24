package main

import (
	"fmt"
)

func main() {
	greet("elliot", "chi")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
