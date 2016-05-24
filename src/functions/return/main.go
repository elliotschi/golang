package main

import (
	"fmt"
)

func main() {
	name := greet("elliot ", "chi")
	fmt.Println(name)
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
