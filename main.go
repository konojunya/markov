package main

import (
	"fmt"

	"github.com/konojunya/markov/salad"
)

func main() {
	text, count := salad.GetTextAndCount()
	h := salad.NewSalad(text)
	for i := 0; i < count; i++ {
		fmt.Println(h.MakeWord())
	}
}
