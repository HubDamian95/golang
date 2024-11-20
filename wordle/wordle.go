package wordle

import (
	"fmt"
	"strings"

	"github.com/Hubdamian95/wordle/words"
)

func userGuess() string {
	var guess string
	fmt.Println("Please enter guess:")
	fmt.Scanln(&guess)
	return guess
}

func validateGuess(g string) bool {
	if len(g) != 5 {
		fmt.Println("Input must be exactly 5 characters")
		return false
	}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range g { //
		if !strings.Contains(alphabet, string(v)) {
			fmt.Println("The guess must only contain the 26 characters in the English alpahbet")
			return false
		}
	}
	if !words.IsWord(g) {
		fmt.Println("Not a valid word")
		return false
	}

	return true
	// Func - check user input is only 5 char long
	// No numbers allowed

}

func main() {
	secret := words.GetWord()
	secret = strings.ToLower(secret)

	attempts := 6
	for attempts > 0 {
		guess := userGuess()
		if !validateGuess(guess) {
			fmt.Printf("Invalid guess.")
			continue
		}
		if guess == secret {
			fmt.Printf("The secret has been found")
			return
		}

		attempts--
		fmt.Printf("%d attempts left", attempts)
		// guesses (5 letters - each letter - one of 26 english characters and in one of the three states - Absent, Present, correct)
	}

}
