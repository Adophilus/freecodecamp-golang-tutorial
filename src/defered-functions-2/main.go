package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main () {
  res, err := http.Get("http://google.com/robots.txt")
  if err != nil {
    log.Fatal(err)
  }
  robots, err := ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", robots)
}
