package pipeline

import "fmt"

func ReadWords(downstream chan<- string, words []string) {
	for _, word := range words {
		downstream <- word
	}
	close(downstream)
}

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

func PrintWords(upstream <-chan string) {
	for word := range upstream {
		fmt.Print(word, " ")
	}
	fmt.Println()
}
