package main

import "fmt"
import "runtime"
import "sync"
// import "time"

// NOTE: don't create goroutines in libraries

// waitgroup 
var wg = sync.WaitGroup {}
var counter = 0
var m = sync.RWMutex {}

// the main function is actually executing in a goroutine
func main () {
  runtime.GOMAXPROCS(100) // sets the number of threads
  // var numberOfCores int = runtime.GOMAXPROCS(-1) // gets the number of threads available

  msg := "This is a message"

  wg.Add(2) // 2 go routines

  // go sayHello() // spins off a 'green' thread and runs this function inside it

  go func () {
    // making reference to a variable in another go routine
    // this could be dangerous because it could result in a race condition
    fmt.Println(msg) // the go runtine understands where to get the 'msg' variable from
    wg.Done()
  }()

  go func (message string) {
    // this is a much better way of handling data between go routines than the one above
    // and just like that, you've decoupled the data in the main go routine from this go
    // go routine because you passed in the dependencies of this go routine as function
    // arguments
    fmt.Println(message)
    wg.Done()
  }(msg)

  msg = "Goodbye"

  for i := 0; i < 10; i++ {
    wg.Add(2)
    m.RLock()
    go sayHello()
    m.Lock()
    go increment()
  }

  // not best practice to use sleep so that the program can complete all its go routines
  // time.Sleep(100 * time.Millisecond)

  // waitgroup (a better way of waiting for go routines to finish)
  // waitgroup is designed to sync multiple go routines together
  wg.Wait()
}

func sayHello () {
  fmt.Println("Hello", counter)
  m.RUnlock()
  wg.Done()
}

func increment () {
  counter++
  m.Unlock()
  wg.Done()
}
