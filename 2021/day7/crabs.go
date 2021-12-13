package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", part2All("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n", part2All("real-input.txt"))
}

// obvious plan
// for each position between min and max, move all crabs, summing number of moves...
//  - crazy inefficient, must be a mathematical method for this...
//  - if we have to go that way, can stop each time we exceed current lowest
// - where to start?

func part1All(filename string) int {
	input, min, max := utils.ReadFileOfCommaSeperatedIntsToSliceWithMinMax(filename)
	return part1(input, min, max)
}

func part1(crabs []int, min int, max int) (bestSoFar int) {

	for pos := min; pos < max+1; pos++ {
		fuelUsed := 0
		for _, crab := range crabs {
			fuelUsed += utils.IntAbs(pos - crab)
			if fuelUsed > bestSoFar && pos != min {
				continue
			}
		}
		if pos == min || fuelUsed < bestSoFar {
			bestSoFar = fuelUsed
		}
	}
	return
}

// part 2 - dont want to actually count - must have a mathematical formula...
// - its the nth triangular number...
//  - formula: n(n+1)/2
//test: 5->2 = 3 = 1+2+3 = 6
//test: 5->2 = 3 = (3 * 4) / 2 11 * 12

func part2All(filename string) int {
	input, min, max := utils.ReadFileOfCommaSeperatedIntsToSliceWithMinMax(filename)
	return part2(input, min, max)
}
func part2(crabs []int, min int, max int) (bestSoFar int) {

	for pos := min; pos < max+1; pos++ {
		fuelUsed := 0
		for _, crab := range crabs {
			numMoves := utils.IntAbs(pos - crab)
			fuelUsed += nthTriangularNumber(numMoves)
			if fuelUsed > bestSoFar && pos != min {
				continue
			}
		}
		if pos == min || fuelUsed < bestSoFar {
			bestSoFar = fuelUsed
		}
	}
	return
}

func nthTriangularNumber(n int) (tri int) {
	return (n * (n + 1)) / 2
}
