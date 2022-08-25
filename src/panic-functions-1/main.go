package main

import "fmt"

func main () {
  fmt.Println("start")
  panic("Something bad happened")
  fmt.Println("end")
}
