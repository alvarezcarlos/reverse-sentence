package main

import (
	"flag"
	"fmt"
	"strings"
)

const defaultSentence = "today is the last day of the rest of my life"

func main() {
	sentencePrt := flag.String("s", defaultSentence, "sentence to reverse")
	flag.Parse()
	s := reverseSentence(*sentencePrt)
	fmt.Println(s)
}

func reverseSentence(sentence string) string {
	words := strings.Fields(sentence)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}