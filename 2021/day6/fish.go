package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

var oldFishDays = 6
var newFishDays = 8

func main() {

	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt", 80))
	fmt.Printf("Test Part 2: %v\n", part1All("test-input.txt", 256))

	fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt", 80))
	fmt.Printf("Real Part 2: %v\n", part1All("real-input.txt", 256))
}

// ReadFileOfCommaSeperatedIntsToSlice - read ints from single line of file
func readFileOfCommaSeperatedIntsToSlice(filename string) (theInts []int) {
	//sliceSize := 26984457539
	sliceSize := 99999999999

	str := utils.ReadFileToString(filename)
	theStrings := strings.Split(str, ",")
	//	theInts = make([]int, 0, len(theStrings))
	theInts = make([]int, 0, sliceSize)
	for _, theString := range theStrings {
		theInt, _ := strconv.Atoi(theString)
		theInts = append(theInts, theInt)
	}
	return
}

func part1All(filename string, numDays int) int {
	theFish := readFileOfCommaSeperatedIntsToSlice(filename) //drb: ensure we do correct thing on day 0 and day 80
	return part1(theFish, numDays)
}

//drb: for now accept we dont know how many fish, so cant pre-alloc whole slice...
func part1(theFish []int, numDays int) (count int) {
	//	fmt.Printf("intFish: %v\n", initFish)

	for d := 0; d < numDays; d++ {
		numFishAtStartOfDay := len(theFish)
		for i := 0; i < numFishAtStartOfDay; i++ {
			switch theFish[i] {
			case 0:
				theFish[i] = oldFishDays
				theFish = append(theFish, newFishDays)
			default:
				theFish[i]--
			}
		}

	}
	return len(theFish)
}

/*
func part2() () {

	return
}
*/
