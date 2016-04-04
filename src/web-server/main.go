// package main

// import (
//   "fmt"
//   "html/template"
//   "net/http"
//   "strings"
//   "log"
// )

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
//   // parse arguments ? body parser?
//   r.ParseForm()
  
//   fmt.Println(r.Form)
//   fmt.Println("path", r.URL.Path)
//   fmt.Println("scheme", r.URL.Scheme)
//   fmt.Println(r.Form["url_long"])
  
//   for k, v := range r.Form {
//     fmt.Println("key: ", k)
//     fmt.Println("val: ", strings.Join(v, ""))
//   }
//   // send data to the client
//   fmt.Fprintf(w, "hello elli") 
// }

// func login(w http.ResponseWriter, r *http.Request) {
//   // catches method
//   fmt.Println("method: ", r.Method)
  
//   if r.Method == "GET" {
//     t, _ := template.ParseFiles("login.gtpl")
//     t.Execute(w, nil)
//   } else {
//     r.ParseForm()
    
//     fmt.Println("username: ", r.Form["username"])
    
//     fmt.Println("[password]: ", r.Form["password"])
//   }
// }

// func main() {
//   // set route
//   http.HandleFunc("/", sayhelloName) 
//   http.HandleFunc("/login", login)
//   // listen to port number
//   err := http.ListenAndServe(":9090", nil)
  
//   if err != nil {
//     log.Fatal("ListenAndServe: ", err)
//   }
// }

package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}