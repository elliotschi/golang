package main

import (
  "fmt"
)

func main() {
  myMap := make(map[string]string)
  
  myMap["hello"] = "world"
  myMap["world"] = "jello"
  
  fmt.Println(myMap)
  
  delete(myMap, "hello")
  fmt.Println(myMap)
}