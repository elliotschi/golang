package main

import (
  "fmt"
)

func main() {
  greeting := map[int]string{
    0: "good morning",
    1: "bonjour",
  }
  
  for key, val := range greeting {
    fmt.Println(key, " - ", val)
  }
}