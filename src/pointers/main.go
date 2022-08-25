package main

import "fmt"

func main () {
  a := 10
  b := &a
  fmt.Println(a, *b)
  a = 11
  fmt.Println(a, *b)
  *b = 13
  fmt.Println(a, *b)

  uchenna := Person {
    age: 18,
    name: "Uchenna Ofoma",
  }

  fmt.Println(uchenna)

  randomPerson := &uchenna
  // or
  // randomPerson = new(uchenna)

  randomPerson.name = "Random Person"
  fmt.Println(randomPerson)
  fmt.Println(uchenna)
}

type Person struct {
  name string
  age int
}
