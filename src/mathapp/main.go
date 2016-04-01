//$GOPATH/src/mathapp/main.go source code.
package main

import (
    "mymath"
    "fmt"
)

// func main() {
//     fmt.Printf("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))
// }

func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func main() {
  x := 3
  y := 4
  z := 5
  
  maxXy := max(x, y)
  maxXz := max(x, z)
  
  fmt.Printf("max(%d, %d) = %d\n", x, y, maxXy)
  
  fmt.Printf("max(%d, %d) = %d\n", x, z, maxXz)
  
  fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
}