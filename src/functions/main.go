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

  // anonymous function
  fmt.Println(func () bool {
    return true
  })

  // using methods
  greeter := Greeter {
    name: "Uchenna",
    greeting: "Hello",
  }
  greeter.greet()
  fmt.Println(greeter)
  greeter.greetAndUpdateName()
  fmt.Println(greeter)
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


// method
// a method is a function that exists in a defined context
type Greeter struct {
  greeting string
  name string
}


// this method will actually get a copy of the Greeter 'instance'
func (g Greeter) greet () {
  fmt.Println(g.greeting, g.name)
  g.name = "<" + g.name + ">" // this will only modify the copy struct that this method receives
}

// this method will get a reference to the original Greeter 'instance'
func (g *Greeter) greetAndUpdateName () {
  fmt.Println(g.greeting, g.name)
  g.name = "<" + g.name + ">"
}
