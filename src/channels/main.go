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

  wg.Wait()
}
