package main
import "fmt"

type Human struct {
  name string
  age int
  phone string
}

type Student struct {
  Human
  school string
}

type Employee struct {
  Human
  company string
}

func (h *Human) sayHi() {
  fmt.Printf("name: %s, call me at: %s\n", h.name, h.phone)
}

func (e *Employee) sayHi() {
  fmt.Printf("%s %s %s\n", e.name, e.company, e.phone)
}

func main() {
  elliot := Student{Human{"elliot", 22, "123-4567"}, "MKS"}
  
  elliot2 := Employee{Human{"elliot", 22, "123-4567"}, "MKS2"}
  
  elliot.sayHi()
  elliot2.sayHi()
}