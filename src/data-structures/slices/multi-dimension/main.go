package main

import (
  "fmt"
)

func main() {
  matrix := make([][]int, 0, 3)
  
  for i := 0; i < 3; i++ {
    subArray := make([]int, 0, 4)
    
    for j := 0; j < 4; j++ {
      subArray = append(subArray, j)
    }
    
    matrix = append(matrix, subArray)
  }
  
  fmt.Println(matrix)
}