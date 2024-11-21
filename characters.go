package main

import "fmt"

func main() {
	str := "hi there!"

	for _, value := range str {
		fmt.Printf("%c\n", value)
	}
}
