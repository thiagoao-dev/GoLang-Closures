package main

import (
	"fmt"
)

func main() {
  word := "github.com/thiagoao"
  closure := InvertWord(word)

  for _ = range word {
    fmt.Printf("%s\n",closure())
  }	
}

func InvertWord(s string) func() string {
  w, iw := s, ""
  return func() string {
    iw = w[:1] + iw
    w  = w[1:]
    return iw
  }
}