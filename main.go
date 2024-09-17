package main

import (
	p "removeidentical/pipeline" // can alias packages
)

func main() {
	// makes two channels
	c0 := make(chan string)
	c1 := make(chan string)
	// words have repeated words to be removed
	words := []string{
		"roses", "are", "red", "red", "red",
		"violets", "are", "blue",
		"blue", "sugar", "is", "sweet",
		"and", "so", "are", "you", "you"}
	// go let's other functions run concurrently
	go p.ReadWords(c0, words)
	go p.RemoveIdentical(c0, c1)
	// print does not allow concurrency to block the main function from ending
	p.PrintWords(c1)
}
