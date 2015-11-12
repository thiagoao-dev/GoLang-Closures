package main

import (
	"fmt"
)

type Word struct {
	Word string
	IWord string
}

func (w *Word) SetWord(s string) {
	w.Word = s
}

func (w *Word) GetWord() string {
	return w.Word
}

func (w *Word) SetIWord(s string) {
	w.IWord += s
}

func (w *Word) GetIWord() string {
	return w.IWord
}

func main() {
  word := Word{Word:"github.com/thiagoao"}
  closure := InvertWord(&word)

  for _ = range word.Word {
    closure()
  }
  
  fmt.Println(word.IWord)
}

func InvertWord(w *Word) func() {
  s := w
  return func() {
    s.IWord = s.Word[:1] + s.IWord
    s.Word  = s.Word[1:]
  }
}