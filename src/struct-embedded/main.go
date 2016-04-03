package main
import "fmt"

//Human is a struct with name, age, and weight
type Human struct {
  name string
  age int
  weight int
}

// Student is a human kind of like a subclass that has a specialty
type Student struct {
  Human
  specialty string
}

func main() {
  elliot := Student{Human{"elliot", 22, 130}, "troll"}
  
  fmt.Println("his name is: ", elliot.name)
  fmt.Println("his age is: ", elliot.age)
  fmt.Println("his weight is: ", elliot.weight)
  fmt.Println("this specialty is: ", elliot.specialty)
  
}
