package main

import (
  "fmt"
  "flag"
  "os"
  "strings"
)

func main() {

  stack := flag.String("stack", "", "stack to query")
  flag.Parse()

  if *stack == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  stacks := getCurrentStacks()
  fmt.Println(strings.Join(stacks[:], "\n"))

  fmt.Println(getStackStatus(*stack))

}

