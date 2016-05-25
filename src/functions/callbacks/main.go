package main

import (
  "fmt"
)

func main() {
  ints := []int{1, 2, 3, 4, 5}
  
  each(ints, func(n int) {
    fmt.Println(n)
  })
}

func each(collection []int, callback func(int)) {
  for _, n := range collection {
    callback(n)
  }
}