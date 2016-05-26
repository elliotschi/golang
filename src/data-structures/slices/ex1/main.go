package main

import (
  "fmt"
)

func main() {
  mySlice := []string{"hello", "world"}
  myOtherSlice := []string{"1", "2", "3"}
  
  mySlice = append(mySlice, myOtherSlice...)
  fmt.Println(mySlice)
  
  mySlice = append(mySlice[:2], mySlice[3:]...)
  fmt.Println(mySlice)
}