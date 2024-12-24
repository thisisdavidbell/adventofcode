package main

import "fmt"

func main() {
	answer1test, answer2test := calculateAnswers("test-input.txt")
	answer1real, answer2real := calculateAnswers("real-input.txt")

	fmt.Printf("Test answer 1: %v\n", answer1test)
	fmt.Printf("Test answer 2: %v\n", answer2test)

	fmt.Printf("\nReal answer 1: %v\n", answer1real)
	fmt.Printf("Real answer 2: %v\n", answer2real)
}

func calculateAnswers(filename string) (answer1, answer2 int) {

	return answer1, answer2
}
