package main

import (
  "fmt"
)

func main()  {
  fmt.Println(ex1(2))
  fmt.Println(ex2(4))
  
  nums := []int{1,2,3,4,5}
  fmt.Println(ex3(nums...))
  
  ex5(1, 2)
  ex5(1, 2, 3)
  aSlice := []int{1, 2, 3, 4}
  ex5(aSlice...)
  ex5()
}

func ex1(n int) (int, bool) {
  return n/2, n % 2 == 0
}

func ex1float(n int) (float64, bool) {
  return float64(n)/2, n % 2 == 0
}

var ex2 = func(n int) (int, bool) {
  return n/2, n % 2 == 0
}

func ex3(nums ...int) int {
  var max int
  for _, n := range nums {
    if n > max {
      max = n
    }
  }
  return max
}

func ex5(nums ...int) {
  for _, n := range nums {
    fmt.Println(n)
  }
}