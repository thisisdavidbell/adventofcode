package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	testIntSlice := utils.ImportFileToIntSlice("test-input.txt")
	realIntSlice := utils.ImportFileToIntSlice("real-input.txt")

	fmt.Printf("Test Part 1 Answer: %v\n", part1CountIncreases(testIntSlice))
	fmt.Printf("Test Part 2 Answer: %v\n\n", part2countSumIncrease(testIntSlice))

	fmt.Printf("Real Part 1 Answer: %v\n", part1CountIncreases(realIntSlice))
	fmt.Printf("Real Part 2 Answer: %v\n", part2countSumIncrease(realIntSlice))
}

func part1CountIncreases(intSlice []int) (count int) {
	for i := 1; i < len(intSlice); i++ {
		if intSlice[i-1] < intSlice[i] {
			count++
		}
		//fmt.Printf("Compare: %v < %v: %v\n", intSlice[i-1], intSlice[i], count)
	}
	return count
}

// quick and dirty first pass - process twice
func produceSumSlice(intSlice []int) (sumSlice []int) {
	sumSlice = make([]int, 0)
	for i := 2; i < len(intSlice); i++ {
		sumSlice = append(sumSlice, intSlice[i-2]+intSlice[i-1]+intSlice[i])
	}
	return
}

func part2countSumIncrease(intSlice []int) (count int) {
	sumSlice := produceSumSlice(intSlice)
	return part1CountIncreases(sumSlice)
}
