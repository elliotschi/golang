package main

import (
  "fmt"
)

func main() {
  myMap := map[string]string{}
  
  myMap["hello"] = "world"
  myMap["word"] = "hello"
  
  fmt.Println(myMap)
  delete(myMap, "word")
  fmt.Println(myMap)
}