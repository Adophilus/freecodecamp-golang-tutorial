package main

import "fmt"

func main () {
  // defered functions are executed in LIFO order
  // LIFO: Last In First Out
  fmt.Println("start")
  defer fmt.Println("executed")
  fmt.Println("end")
}
