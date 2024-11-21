package main

import "fmt"

func words() {
	words := []string{"lol", "a", "b", "c", "d"}

	for wordle := range words {
		fmt.Println(wordle)
	}

}
