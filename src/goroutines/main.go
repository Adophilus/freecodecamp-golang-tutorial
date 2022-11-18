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

  readersAndWriters()
  readersAndWritersWithoutMutex()
  readersAndWritersWithMutex()
  readersAndWritersWithMutex2()

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


func readersAndWriters () {
  ch := make(chan int)

  wg.Add(2)
  // read only channel
  go func (ch <-chan int) {
    fmt.Println(<-ch)
    // ch <- 28 // would cause the program to panic because ch is a read only channel
    wg.Done()
  }(ch)

  // write only channel ch
  go func (ch chan<- int) {
    ch <- 4
    // fmt.Println(<-ch) // would cause the program to panic because ch is a write only channel
    wg.Done()
  }(ch)
}

func sendData (data int) {
  fmt.Println("Sending data %v", data)
  wg.Done()
}

func readersAndWritersWithoutMutex () {
  counter := 0
  increment := func () {
    counter++
    wg.Done()
  }

  for i := 0; i < 10; i++ {
    wg.Add(2)
    go sendData(counter)
    go increment()
  }

  // Notice the unpredictible behaviour observed when this function is run. In order to control the
  // order of execution (so that it would be just as we expect), we would need to make use of mutex.
  // See the function below
}

func readersAndWritersWithMutex () {
  counter := 0
  m := sync.RWMutex{} // the mutex
  sendDataUsingMutex := func (data int) {
    m.RLock()
    fmt.Println("Sending data (using mutex) %v", data)
    m.RUnlock()
    wg.Done()
  }
  increment := func () {
    m.Lock()
    counter++
    m.Unlock()
    wg.Done()
  }

  for i := 0; i < 10; i++ {
    wg.Add(2)
    go sendDataUsingMutex(counter)
    go increment()
  }

  // Notice that this time, the order of execution goes in the expected direction. The issue now is that
  // we still get some of that unpredictible behaviour. See the function below to discover the solution
}

func readersAndWritersWithMutex2 (){
  counter := 0
  m := sync.RWMutex{} // the mutex
  sendDataUsingMutex := func (data int) {
    fmt.Println("Sending data (using mutex 2) %v", data)
    m.RUnlock()
    wg.Done()
  }
  increment := func () {
    counter++
    m.Unlock()
    wg.Done()
  }

  for i := 0; i < 10; i++ {
    wg.Add(2)
    m.RLock()
    go sendDataUsingMutex(counter)
    m.Lock()
    go increment()
  }

  // Gives expected behaviour. But takes away the key feature of parallelism and thus making this code
  // run worse than it would've if you weren't even using goroutines.
} 
