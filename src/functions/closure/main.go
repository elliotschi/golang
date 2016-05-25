package main

import (
  "fmt"
)

func main() {
  increment := closureFn()
  
  fmt.Println(increment())
  fmt.Println(increment())
}

func closureFn() func() int {
  var counter int
  
  return func() int {
    counter++
    return counter
  }
}