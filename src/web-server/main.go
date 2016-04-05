package main

import (
  "fmt"
  "html/template"
  "net/http"
  "strings"
  "log"
  "time"
  "crypto/md5"
  "io"
  "strconv"
  "os"
  "goredis"
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
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    
    t, _ := template.ParseFiles("login.gtpl")
    t.Execute(w, token)
  } else {
    r.ParseForm()
    
    token := r.Form.Get("token")
    
    if token != "" {
      
    } else {
      
    }
    
    fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
    
    fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
    
    template.HTMLEscape(w, []byte(r.Form.Get("username")))
  }
}

func upload(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method)
  if r.Method == "GET" {
    crutime := time.Now().Unix()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))

    t, _ := template.ParseFiles("upload.gtpl")
    t.Execute(w, token)
  } else {
    r.ParseMultipartForm(32 << 20)
    file, handler, err := r.FormFile("uploadfile")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Fprintf(w, "%v", handler.Header)
    f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)
  }
}

func main() {
  // set route
  http.HandleFunc("/", sayhelloName) 
  http.HandleFunc("/login", login)
  http.HandleFunc("/upload", upload)
  // listen to port number
  err := http.ListenAndServe(":9090", nil)
  
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
