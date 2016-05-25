package main

import (
  "fmt"
)

func main() {
  nums := []float64{15, 57, 26, 57}
  
  avg := average(nums...)
  
  fmt.Println(avg)
}

func average(sf ...float64) float64 {
  var total float64
  
  for _, v := range sf {
    total += v
  }
  
  return total / float64(len(sf))
}