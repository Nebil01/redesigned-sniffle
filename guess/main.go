package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input. Exiting.")
		return
	}
	switch choice {
	case 1:
		fmt.Println("Great! You have 10 chances to guess the correct number.")
		fmt.Println("Let's start the game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		var guess int
		random := rand.Intn(100) + 1
		chances := 10
		for chances > 0 {
			fmt.Print("Enter your Guess: ")
			_, err := fmt.Scan(&guess)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 1 and 100.")
				var discard string
				fmt.Scan(&discard)
				continue
			}
			if guess == random {
				fmt.Println("Congratulations! You Guessed the correct number.")
				break
			} else if guess < random {
				fmt.Printf("Incorrect! The number is greater than %d\n", guess)
			} else if guess > random {
				fmt.Printf("Incorrect! The number is less than %d\n", guess)
			}
			chances--
		}
		fmt.Printf("The Correct number was %d", random)

	case 2:
		fmt.Println("Great! You have 5 chances to guess the correct number.")
		fmt.Println("Let's start the game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		var guess int
		random := rand.Intn(100) + 1
		chances := 5
		for chances > 0 {
			fmt.Print("Enter your Guess: ")
			_, err := fmt.Scan(&guess)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 1 and 100.")
				var discard string
				fmt.Scan(&discard)
				continue
			}
			if guess == random {
				fmt.Println("Congratulations! You Guessed the correct number.")
				break
			} else if guess < random {
				fmt.Printf("Incorrect! The number is greater than %d\n", guess)
			} else if guess > random {
				fmt.Printf("Incorrect! The number is less than %d\n", guess)
			}
			chances--
		}
		fmt.Printf("The Correct number was %d", random)
	case 3:
		fmt.Println("Great! You have 3 chances to guess the correct number.")
		fmt.Println("Let's start the game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		var guess int
		random := rand.Intn(100) + 1
		chances := 3
		for chances > 0 {
			fmt.Print("Enter your Guess: ")
			_, err := fmt.Scan(&guess)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number between 1 and 100.")
				var discard string
				fmt.Scan(&discard)
				continue
			}
			if guess == random {
				fmt.Println("Congratulations! You Guessed the correct number.")
				break
			} else if guess < random {
				fmt.Printf("Incorrect! The number is greater than %d\n", guess)
			} else if guess > random {
				fmt.Printf("Incorrect! The number is less than %d\n", guess)
			}
			chances--
		}
		fmt.Printf("The Correct number was %d", random)
	default:
		fmt.Println("Invalid choice. Please select a valid difficulty level.")
	}
}
