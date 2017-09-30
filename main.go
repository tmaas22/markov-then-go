package main

import "fmt"

type Meh struct {
	TheList map[string]WordsOnWords
}

type WordsOnWords struct {
	Followers map[string]int
	Total     int
}

func main() {
	fmt.Println("Yay! We back in go!")
}
