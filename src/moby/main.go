package main

import (
  "fmt"
  "io"
  "bufio"
  "net/http"
  "os"
  "log"
)

func main() {
  res, err := http.Get("")
  
  if err != nil {
    log.Fatalln(err)
  }
  
  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close()
  
  buckets := make([]int, 200)
  
  for scanner.Scan() {
    n := HashBucket(scanner.Text())   
    buckets[n]++
  }
}