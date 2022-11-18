package main

import "fmt"
import "sync"

// channels
// a basic channel is a 2-way stream

// overview
// channel basics
// restricted data flow
// buffered channels
// closing channels
// for...range loops with channels
// select statements

var wg = sync.WaitGroup{}

func main () {
  // channels are created by passing in the keyword 'chan' and the data type to be transmitted through the channel to the make function
  ch := make(chan int) // this channel can only transmit data of type int

  wg.Add(2)

  // this goroutine receives data from the channel
  go func () {
    i := <- ch
    fmt.Println(i)
    wg.Done()
  }()

  // this goroutine sends data to the channel
  go func () {
    i := 42
    ch <- i // this is a blocking operation. If there is no corresponding data sent (that is expected to be recieved), it would cause a DEADLOCK!
    i = 27
    wg.Done()
  }()

  readersAndWriters()
  readersAndWritersWithoutMutex()
  readersAndWritersWithMutex()
  wg.Wait()
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

func readersAndWritersWithMutex2 () {


}
