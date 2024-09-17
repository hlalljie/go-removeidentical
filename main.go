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
	go p.ReadWords(c0, words)
	go p.RemoveIdentical(c0, c1)
	// c1 being sent downstream activates printWords
	p.PrintWords(c1)
}
