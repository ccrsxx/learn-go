package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(rawWords string) map[string]int {
	wordCount := make(map[string]int)

	fmt.Println("Input string:", rawWords)

	words := strings.Fields(rawWords)

	fmt.Printf("Split words: %#v\n", words)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
