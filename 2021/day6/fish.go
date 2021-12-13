package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {

	fmt.Printf("Test Part 1: %v\n", solveAll("test-input.txt", 80))
	fmt.Printf("Test Part 2: %v\n", solveAll("test-input.txt", 256))

	fmt.Printf("Real Part 1: %v\n", solveAll("real-input.txt", 80))
	fmt.Printf("Real Part 2: %v\n", solveAll("real-input.txt", 256))
}

func readInput(filename string) (theCount []int) {
	str := utils.ReadFileToString(filename)
	theStrings := strings.Split(str, ",")
	theCount = make([]int, 9)
	for _, theString := range theStrings {
		fishDay, _ := strconv.Atoi(theString)
		theCount[fishDay]++
	}
	return
}

func solveAll(filename string, numDays int) int {
	theCount := readInput(filename)
	return solve(theCount, numDays)
}

func solve(theCount []int, numDays int) (count int) {

	for d := 0; d < numDays; d++ {
		day0 := theCount[0]
		for i := 1; i < 9; i++ {
			theCount[i-1] = theCount[i]
		}
		theCount[6] = theCount[6] + day0
		theCount[8] = day0
	}
	for c := 0; c < 9; c++ {
		count = count + theCount[c]
	}
	return
}
