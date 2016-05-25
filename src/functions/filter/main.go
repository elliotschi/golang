package main

import (
  "fmt"
)

func filter(collection []int, callback func(int) bool) []int {
  filtered := []int{}
  for _, n := range collection {
    if callback(n) {
      filtered = append(filtered, n)
    }
  }
  return filtered
}

func main()  {
  evenNums := filter([]int{1, 2, 3, 4, 5}, func(number int) bool{
    return number % 2 == 0
  })
  fmt.Println(evenNums)
}