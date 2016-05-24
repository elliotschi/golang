package main

import (
	"fmt"
)

func main() {
	name, reverseName := greet("elliot ", "chi ")

	fmt.Println(name, reverseName)
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
