package main

import (
	"fmt"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1(utils.ReadFileToStringSlice("test-input.txt")))
	fmt.Printf("Test Part 2: %v\n\n", part2(utils.ReadFileToByteSliceSlice("test-input.txt")))

	fmt.Printf("Real Part 1: %v\n", part1(utils.ReadFileToStringSlice("real-input.txt")))
	fmt.Printf("Real Part 2: %v\n", part2(utils.ReadFileToByteSliceSlice("real-input.txt")))
}

func part1(slice []string) (power int64) {
	numChars := len(slice[0])
	numLines := len(slice)
	gamma := make([]byte, 0, numChars)
	epsilon := make([]byte, 0, numChars)
	for char := 0; char < numChars; char++ { // for every column (each char position)
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

func part2(oxygenSlice [][]byte) (power int64) {

	co2Slice := make([][]byte, len(oxygenSlice))
	copy(co2Slice, oxygenSlice)
	oxygen := getAttributeSliceReuse(oxygenSlice, true)
	co2 := getAttributeSliceReuse(co2Slice, false)

	return oxygen * co2
}

func getAttributeSliceReuse(theSlice [][]byte, most bool) (attribute int64) {
	numChars := len(theSlice[0])
	for char := 0; char < numChars; char++ { // for every column (each char position)
		numLines := len(theSlice)
		count1 := 0
		for _, line := range theSlice { // for every line at that position
			if line[char] == byte('1') {
				count1++
			}
		}
		match := byte('0')
		if most {
			if count1*2 >= numLines {
				match = byte('1')
			}
		} else {
			if count1*2 < numLines {
				match = byte('1')
			}
		}
		j := 0
		for _, line := range theSlice {
			if line[char] == match {
				theSlice[j] = line
				j++
			}
		}
		theSlice = theSlice[:j]
		if len(theSlice) == 1 {
			for _, v := range theSlice {
				attribute, _ = strconv.ParseInt(string(v), 2, 64)
				return attribute
			}
		}
	}
	return 0
}
