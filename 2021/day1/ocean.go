package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	testIntSlice := utils.ImportFileToIntSlice("test-input.txt")
	realIntSlice := utils.ImportFileToIntSlice("real-input.txt")

	fmt.Printf("Test Part 1 Answer: %v\n", part1(testIntSlice))
	fmt.Printf("Test Part 2 Answer: %v\n\n", part2(testIntSlice))

	fmt.Printf("Real Part 1 Answer: %v\n", part1(realIntSlice))
	fmt.Printf("PERF Real Part 2 Answer: %v\n", part2(realIntSlice))
}

func part1(intSlice []int) (count int) {
	for i := 1; i < len(intSlice); i++ {
		if intSlice[i-1] < intSlice[i] {
			count++
		}
		//fmt.Printf("Compare: %v < %v: %v\n", intSlice[i-1], intSlice[i], count)
	}
	return count
}

func part2(intSlice []int) (count int) {
	for i := 3; i < len(intSlice); i++ {
		if intSlice[i-3] < intSlice[i] {
			count++
		}
	}
	return
}
