package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Phone string
}

func main() {
  session, err := mgo.Dial("localhost:27017")
  if err != nil {
    panic(err)
  }
  
  defer session.Close()
  
  session.SetMode(mgo.Monotonic, true)
  
  c := session.DB("test").C("people")
  err = c.Insert(&Person{"Elliot", "123-456-7891"})
  if err != nil {
    panic(err)
  }
  
  result := Person{}
  err = c.Find(bson.M{"name": "Elliot"}).One(&result)
  if err != nil {
    panic(err)
  }
  
  fmt.Println("Phone: ", result.Phone)
}