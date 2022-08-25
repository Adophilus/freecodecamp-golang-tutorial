package main

import "fmt"

func main () {
  var message string = "'Ello"
  displayMessage(message)
  displayMessageTo("uchenna", message)
  displayMessageToMany(message, "uchenna", "unnamedNoob1", "unnamedNoob2")
  fmt.Println(displayMessageWithReturn(message))
  fmt.Println(displayMessageWithReturnPtr(message))
  fmt.Println(displayMessageWithReturnVar(message))
  fmt.Println(displayMessageWithMultipleReturn(message))
}

// regular function
func displayMessage (message string) {
  fmt.Println(message)
}

// grouping function types
func displayMessageTo (name, message string) {
  fmt.Println(message, name)
}

// variadic function
func displayMessageToMany (message string, names ...string) {
  for _,name := range names {
    fmt.Println(message, name)
  }
}

// function with return type
func displayMessageWithReturn (message string) bool {
  fmt.Println(message)
  return true
}

// function returning a pointer to a local variable
func displayMessageWithReturnPtr (message string) *bool {
  fmt.Println(message)
  status := true
  return &status
}

// function with named return value
func displayMessageWithReturnVar (message string) (status bool) {
  status = true
  fmt.Println(message)
  return
}

// function with multiple return values
func displayMessageWithMultipleReturn (message string) (bool, bool) {
  fmt.Println(message)
  return true, false
}
