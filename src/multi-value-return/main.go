package main
import "fmt"

// SumAndProduct the types of both the outputs will be int
func SumAndProduct(A, B int) (int, int) {
  return A+B, A*B
}

func main() {
  x := 3
  y := 4
  
  xPLUSy, xTIMESy := SumAndProduct(x, y)
  
  fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)

  fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}