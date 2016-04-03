package main
import "fmt"

type testInt func(int) bool
// function that returns a bool and takes in an int

func isOdd(integer int) bool {
  if integer % 2 == 0 {
    return false
  } 
  return true
}

func isEven(integer int) bool {
  if integer % 2 == 0 {
    return true
  }
  return false
}

// pass function f into filter

func filter(slice []int, f testInt) []int {
  var result []int
  
  for _, value := range slice {
    if f(value) {
      result = append(result, value)
    }
  }
  return result
}

func main() {
  slice := []int {1,2,3,4,5,6,7}
  
  fmt.Println("slice = ", slice)
  
  odd := filter(slice, isOdd)
  fmt.Println("odd = ", odd)
  
  even := filter(slice, isEven)
  fmt.Println("even = ", even)
}