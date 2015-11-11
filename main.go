package main

import (
	"fmt"
)

func main() {
  word := "thiagoao"
  closure := InvertWord(word)

  for _ = range word {
    fmt.Printf("%s\n",closure())
  }	
}

func InvertWord(s string) func() string {
  word, iword := s, ""
  return func() string {
    iword = word[:1] + iword
    word  =  word[1:]
    return iword
  }
}