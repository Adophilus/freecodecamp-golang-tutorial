package main

import "fmt"
import "time"

// the main function is actually executing in a goroutine
func main () {
  sayHello()
  go sayHello() // spins off a 'green' thread and runs this function inside it
  time.Sleep(100 * time.Millisecond)
}

func sayHello () {
  fmt.Println("Hello")
}
