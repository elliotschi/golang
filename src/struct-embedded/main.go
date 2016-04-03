package main
import "fmt"

// Skills for a student
type Skills []string

//Human is a struct with name, age, and weight
type Human struct {
  name string
  age int
  weight int
}

// Student is a human kind of like a subclass that has a specialty
type Student struct {
  Human
  Skills
  int
  weight int
  specialty string
}

func main() {
  elliot := Student{Human:Human{"elliot", 22, 130}, specialty: "troll", weight: 131}
  
  fmt.Println("his name is: ", elliot.name)
  fmt.Println("his age is: ", elliot.age)
  fmt.Println("his weight is: ", elliot.weight)
  fmt.Println("this specialty is: ", elliot.specialty)
  
  elliot.Skills = []string{"troll"}
  fmt.Println("his skills are: ", elliot.Skills)
  
  elliot.Skills = append(elliot.Skills, "golang")
  
  fmt.Println("he has gained a new skill: ", elliot.Skills)
  
  fmt.Println("this is his human weight", elliot.Human.weight)
}
