package main

import (
  "fmt"
)

func main() {
  avg := average(123,5,32,67,8,35,2)
  
  fmt.Println(avg)
}

func average(sf ...float64) float64 {
	var total float64

	for _, v := range sf {
    total += v
	}
  return total / float64(len(sf))
}
