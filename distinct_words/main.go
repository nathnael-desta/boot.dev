package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	uniqueWords := map[string]struct{}{}
	for _, message := range messages {
		fmt.Println("message: ", message)
		words := strings.Split(message, " ")
		fmt.Println("words: ", words)
		for _, word := range words {
			if _, ok := uniqueWords[strings.ToLower(word)]; !ok {
				uniqueWords[strings.ToLower(word)] = struct{}{}
			}
		}
	}
	fmt.Println(uniqueWords)
	return len(uniqueWords)
}

func main() {
	messages := []string{"Hello world", "hello there", "General Kenobi"}
	countDistinctWords(messages)	
}
