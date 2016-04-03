package main

import (
 . "fmt"
)

type person struct {
  name string
  age int
}

func Older(p1, p2 person) (person, int) {
  if p1.age > p2.age {
    return p1, p1.age-p2.age
  }
  return p2, p2.age-p1.age
}

func main() {
  var me person
  
  me.name = "elliot"
  me.age = 22
  
  oldDude := person{age: 90, name: "old"}
  oldestDude := person{"oldest", 100}
  
  meOldOlder, meOldDiff := Older(me, oldDude)
  
  oldOldestOlder, oldOldestDiff := Older(oldDude, oldestDude)
  
  meOldestOlder, meOldestDiff := Older(me, oldestDude)
  
  Printf("Of %s and %s, %s is older by %d years\n", me.name, oldDude.name, meOldOlder.name, meOldDiff)
  
  Printf("Of %s and %s, %s is older by %d years\n", oldDude.name, oldestDude.name, oldOldestOlder.name, oldOldestDiff)
  
  Printf("Of %s and %s, %s is older by %d years\n", me.name, oldestDude.name, meOldestOlder.name, meOldestDiff)
}