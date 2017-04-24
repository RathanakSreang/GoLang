package main

import "fmt"

type Person struct {
  name string
  age int
}

type testing func(sample Person) (rs Person, isMyturn bool)

func main() {
  fmt.Println("Rathanak")
  p := Person {"Rathanak_123", 23}
  fmt.Println(p)
  // testing(p)
}
