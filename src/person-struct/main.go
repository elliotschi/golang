package main
import "fmt"

// type person struct {
//   name string
//   age int
// }

func main() {
  // var P person
  
  // P.name = "elliot"
  // P.age = 22
  
  // P := person{name: "elliot", age: 22}
  
  // P := person{"elliot", 22}
  
  // anonymous struct without declaring it above
  
  P := struct{name string; age int}{"elliot", 22}
  
  fmt.Printf("name %s, age %d\n", P.name, P.age)
}