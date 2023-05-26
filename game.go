package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the Math Game!")
	fmt.Println("You will be given random multiplication problems to solve.")
	fmt.Println("Enter the correct answer to score a point.")

	score := 0
	numQuestions := 5

	for i := 0; i < numQuestions; i++ {
		num1 := rand.Intn(10)
		num2 := rand.Intn(10)
		expectedAnswer := num1 * num2

		fmt.Printf("\nQuestion %d:\n", i+1)
		fmt.Printf("%d x %d = ", num1, num2)

		var userAnswer int
		fmt.Scanln(&userAnswer)

		if userAnswer == expectedAnswer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Wrong!")
			fmt.Printf("The correct answer is %d\n", expectedAnswer)
		}
	}

	fmt.Printf("\nGame over! You scored %d out of %d.\n", score, numQuestions)
}
