package main

import "fmt"

func numbers() {
	nums := []int{2, 3, 5, 7, 11}

	for numbers := range nums {
		fmt.Println(numbers)
	}

}
