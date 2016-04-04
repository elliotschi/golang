package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  // parse arguments ? body parser?
  r.ParseForm()
  
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("val: ", strings.Join(v, ""))
  }
  // send data to the client
  fmt.Fprintf(w, "hello elli") 
}

func main() {
  // set route
  http.HandleFunc("/", sayhelloName) 
  // listen to port number
  err := http.ListenAndServe(":9090", nil)
  
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}