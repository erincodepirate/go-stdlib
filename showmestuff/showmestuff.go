package main

import (
  "runtime"
  "fmt"
  "os"
  "bufio"
)

func main() {
  // this is where stuff happens
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("What is your name?")
  text, _ := reader.ReadString('\n')
  fmt.Printf("Hello %v", text)
  fmt.Printf("Our current version of Go is %v running in %v\n", runtime.Version(), runtime.GOOS)
}
