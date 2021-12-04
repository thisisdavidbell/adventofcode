package main

import (
	"fmt"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	testByteSlice := utils.ReadFileToByteSliceSlice("test-input.txt")
	realByteSlice := utils.ReadFileToByteSliceSlice("real-input.txt")

	testStringSlice := utils.ReadFileToStringSlice("test-input.txt")
	realStringSlice := utils.ReadFileToStringSlice("real-input.txt")

	fmt.Printf("Test Part 1 Answer Byte: %v\n", part1Bytes(testByteSlice))
	fmt.Printf("Real Part 1 Answer Byte: %v\n\n", part1Bytes(realByteSlice))

	fmt.Printf("Test Part 1 Answer String: %v\n", part1String(testStringSlice))
	fmt.Printf("Real Part 1 Answer String: %v\n", part1String(realStringSlice))

	//	fmt.Printf("Test Part 2 Answer: %v\n\n", part2(testSlice))

	//	fmt.Printf("Real Part 1 Answer: %v\n", part1(realSlice))
	//	fmt.Printf("PERF Real Part 2 Answer: %v\n", part2(realslice))
}

func part1Bytes(slice [][]byte) (power int64) {
	numChars := len(slice[0])
	numLines := len(slice)
	gamma := make([]byte, 0, numChars)
	epsilon := make([]byte, 0, numChars)
	for char := 0; char < numChars; char++ { // for every collumn (each char position)
		count1, count0 := 0, 0
		for line := 0; line < numLines; line++ { // for every line at that position
			if slice[line][char] == byte('1') {
				count1++
			} else {
				count0++
			}
		}
		if count1 > count0 {
			gamma = append(gamma, byte('1'))
			epsilon = append(epsilon, byte('0'))
		} else {
			gamma = append(gamma, byte('0'))
			epsilon = append(epsilon, byte('1'))
		}
	}
	gammaInt, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonInt, _ := strconv.ParseInt(string(epsilon), 2, 64)

	return gammaInt * epsilonInt
}

func part1String(slice []string) (power int64) {
	numChars := len(slice[0])
	numLines := len(slice)
	gamma := make([]byte, 0, numChars)
	epsilon := make([]byte, 0, numChars)
	for char := 0; char < numChars; char++ { // for every collumn (each char position)
		count1, count0 := 0, 0
		for line := 0; line < numLines; line++ { // for every line at that position
			if slice[line][char] == byte('1') {
				count1++
			} else {
				count0++
			}
		}
		if count1 > count0 {
			gamma = append(gamma, byte('1'))
			epsilon = append(epsilon, byte('0'))
		} else {
			gamma = append(gamma, byte('0'))
			epsilon = append(epsilon, byte('1'))
		}
	}
	gammaInt, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonInt, _ := strconv.ParseInt(string(epsilon), 2, 64)
	return gammaInt * epsilonInt
}

/*
func part2() () {

	return
}
*/
