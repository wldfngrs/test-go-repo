package main

import "fmt"

func main() {
  sum := Add(7, 8)
  fmt.Printf("%d", sum)
  fmt.Printf("Test statement")
}

func Add(a int, b int) int {
  if a == 0 {
    return b
  } else if b == 0 {
    return a
  } else {
    return a + b
  }
}
