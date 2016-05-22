package main

import (
	"fmt"
)

const meterToYards float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)

	yards := meters * meterToYards
	fmt.Println(meters, "meters is ", yards, " yards")

	a := "hello"

	fmt.Println(a)

	fmt.Println(&a)
}
