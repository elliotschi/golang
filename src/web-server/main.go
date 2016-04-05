package main

import (
  "fmt"
  "html/template"
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

func login(w http.ResponseWriter, r *http.Request) {
  // catches method
  fmt.Println("method: ", r.Method)
  
  if r.Method == "GET" {
    t, _ := template.ParseFiles("login.gtpl")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    
    fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
    
    fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
    
    template.HTMLEscape(w, []byte(r.Form.Get("username")))
  }
}

func main() {
  // set route
  http.HandleFunc("/", sayhelloName) 
  http.HandleFunc("/login", login)
  // listen to port number
  err := http.ListenAndServe(":9090", nil)
  
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
