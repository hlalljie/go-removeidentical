package pipeline

import "fmt"

// Reads words sending them downstream to start the pipeline
func ReadWords(downstream chan<- string, words []string) {
	for _, word := range words {
		downstream <- word
	}
	close(downstream)
}

// Removes repeated words and sends them downstream
func RemoveIdentical(upstream <-chan string, downstream chan<- string) {

	previousWord := ""
	for word := range upstream {
		if word != previousWord {
			downstream <- word
			previousWord = word
		}
	}
	close(downstream)
}

// Prints words ending the pipeline
func PrintWords(upstream <-chan string) {
	for word := range upstream {
		fmt.Print(word, " ")
	}
	fmt.Println()
}
