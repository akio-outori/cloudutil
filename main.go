package main

import (
  "fmt"
  "flag"
  "os"
  "strings"
)

func printUsage() {
  flag.PrintDefaults()
  os.Exit(1)
}

func main() {
  stack   := flag.String("stack", "", "stack to query")
  command := flag.String("command", "", "action to take")
  flag.Parse()

  switch {

    case *command == "getCurrentStacks":
      stacks := getCurrentStacks()
      fmt.Println(strings.Join(stacks[:], "\n"))

    case *command == "getStackStatus": 
      if *stack != "" {
        fmt.Println(getStackStatus(*stack))
      } else {
        printUsage()
      }

    default:
      printUsage()

  }
}

