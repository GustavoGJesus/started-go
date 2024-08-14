package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("A random number will be generated. Try to guess it. The number is an integer between 0 and 100.")

	// Generate a random number between 0 and 100
	x := rand.Int64N(101)

	// Set up a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)
	
	// Array to store the player's guesses
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("What is your guess?")
		
		// Read and process the player's guess
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		// Convert the guess to an integer
		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Your guess needs to be an integer")
			return
		}

		// Compare the guess to the target number
		switch {
		case guessInt < x:
			fmt.Println("Wrong. The number is higher than", guessInt)
		case guessInt > x:
			fmt.Println("Wrong. The number is lower than", guessInt)
		case guessInt == x:
			// If the guess is correct, congratulate the player and display their attempts
			fmt.Printf(
				"Congratulations, the number was: %d\n"+
					"You guessed it in %d attempts.\n"+
					"These were your attempts: %v\n", x, i+1, guesses[:i],
			)
			return
		}

		// Store the player's guess in the array
		guesses[i] = guessInt
	}

	// If the player doesn't guess the number, show the correct answer and their attempts
	fmt.Printf(
		"Unfortunately, you didn't guess the number, which was: %d\n"+
			"You had 10 attempts.\n"+
			"These were your attempts: %v\n", x, guesses)
}
