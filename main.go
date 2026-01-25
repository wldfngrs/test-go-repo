package main

import "fmt"

func main() {
  sum := Add(3, 8)
  sub := Sub(8, 7)
  fmt.Printf("%d", sum)
  fmt.Printf("%d", sub)
  fmt.Printf("Test statement 121")
}

func Sub(a int, b int) int {
  return a - b
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
