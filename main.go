package main

import (
  "crypto/rand"
  "encoding/hex"
  "fmt"
  "os"
  "strconv"
)

var length = 0

func init() {
  if len(os.Args) != 2 {
    fmt.Println("genkey requires exactly one argument") 
    fmt.Println("usage: genkey <length>") 
    return
  }

  length, _ = strconv.Atoi(os.Args[1])
}

func main() {
  if length == 0 {
    fmt.Println("length is 0")
    return
  }

  length /= 8

  buffer := make([]byte, length)
  _, err := rand.Read(buffer)

  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", hex.EncodeToString(buffer))
}
