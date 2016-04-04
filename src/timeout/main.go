package main

import (
  "time"
  "fmt"
)

func main() {
  c := make(chan int)
  o := make(chan bool)
  
  go func() {
    for {
      select {
        case v := <- c:
          fmt.Println(v)
          
        case <- time.After(5 * time.Second):
          println("timeout")
          o <- true
          break
      }
    }
  }()
  <- o
}