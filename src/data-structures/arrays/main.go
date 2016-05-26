package main

import (
  "fmt"
)

func main() {
  var x [58]string
  
  // fmt.Println(x)
  // fmt.Println(len(x))
  // fmt.Println(x[0])
  
  for i := 0; i < len(x); i++ {
    x[i] = string(i + 65)
  }
  
  fmt.Println(x)
  // fmt.Println(len(x))
  fmt.Println(x[42])
}