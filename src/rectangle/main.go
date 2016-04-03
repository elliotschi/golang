package main
import (
  "fmt"
  "math"
)

// Rectangle is a struct with a width and a height
type Rectangle struct {
  width, height float64
}

// Circle is a struct with a radius so that we can calculate the area later.
type Circle struct {
  radius float64
}

// func area(r Rectangle) float64 {
//   return r.width*r.height
// }

func (r Rectangle) area() float64 {
  return r.width*r.height
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func main() {
  r1 := Rectangle{5,3}
  r2 := Rectangle{2,5}
  
  c1 := Circle{1}
  c2 := Circle{2}
  
  // fmt.Println("area of r1 is: ", area(r1))
  // fmt.Println("area of r2 is: ", area(r2))
  
  fmt.Println("area of r1 is: ", r1.area())
  fmt.Println("area of r2 is: ", r2.area())
  fmt.Println("area of c1 is: ", c1.area())
  fmt.Println("area of c2 is: ", c2.area())
 
}
