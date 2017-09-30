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
	test = NewMeh()
	fmt.Println(test)
}

func NewMeh() *Meh {
	var m Meh
	var w WordsOnWords
	w.Followers = make(map[string]int)
	m.TheList = make(map[string]w)
}
