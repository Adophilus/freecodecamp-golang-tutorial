package main

import "fmt"

func main () {

  // your typical for loop
  for i := 0; i < 100; i++ {
    if i % 3 == 0 {
      fmt.Println("fizz")
    }

    if i % 7 == 0 {
      fmt.Println("buzz")
    }

    if i % 3 == 0 && i % 7 == 0 {
      fmt.Println("fizzbuzz")
    }
  }

  // go's version of a while loop
  i := 0
  for i < 100 {
    if i % 3 == 0 {
      fmt.Println("fizz")
    }

    if i % 7 == 0 {
      fmt.Println("buzz")
    }

    if i % 3 == 0 && i % 7 == 0 {
      fmt.Println("fizzbuzz")
    }

    i++
  }

  // go's version of an infinite loop
  for {
    if i == 2000 {
      break
    }

    if i % 3 == 0 {
      fmt.Println("fizz")
    }

    if i % 7 == 0 {
      fmt.Println("buzz")
    }

    if i % 3 == 0 && i % 7 == 0 {
      fmt.Println("fizzbuzz")
    }

    i++
  }

  // exploring the continue statement
  for i := 0; i < 20; i++ {
    if i % 2 == 0 {
      continue
    }

    if i == 11 {
      break
    }

    fmt.Println(i)
  }

  // exploring the label feature
  // you can use this to break out of multiple levels of loops
  Loop: // this is a label
    for i := 0; i < 3; i++ {
      for j := 0; j < 3; j++ {
        fmt.Printf("%v ", i + j + 1)

        if j == 2 {
          break Loop
        }
      }

      fmt.Println()
    }

  // exploring the for-range feature
  // this also works for maps
  var nums = [3]int { 1, 2, 3 }
  for k, v := range nums {
    fmt.Println(k, v)
  }
}
