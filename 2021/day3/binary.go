package main

import (
	"fmt"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	testStringSlice := utils.ReadFileToStringSlice("test-input.txt")
	realStringSlice := utils.ReadFileToStringSlice("real-input.txt")

	fmt.Printf("Test Part 1: %v\n", part1String(testStringSlice))
	fmt.Printf("Real Part 1: %v\n\n", part1String(realStringSlice))

	testStringMap := utils.ImportFileToStringMap("test-input.txt")
	fmt.Printf("Test Part2 %v\n", part2(testStringMap))

	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	fmt.Printf("Real Part2 %v\n", part2(realStringMap))

}

func part1Bytes(slice [][]byte) (power int64) {
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

func part1String(slice []string) (power int64) {
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

func part2(oxygenMap map[int]string) (power int64) {

	co2Map := utils.CopyIntStringMap(oxygenMap)
	oxygen := getAttribute(oxygenMap, true)
	co2 := getAttribute(co2Map, false)

	return oxygen * co2
}

func getAttribute(themap map[int]string, most bool) (attribute int64) {
	numChars := len(themap[0])
	for char := 0; char < numChars; char++ { // for every column (each char position)
		zeros := make(map[int]string)
		ones := make(map[int]string)
		count1, count0 := 0, 0
		for key, line := range themap { // for every line at that position
			if line[char] == byte('1') {
				count1++
				ones[key] = line
			} else {
				count0++
				zeros[key] = line
			}
		}
		if most {
			if count1 >= count0 {
				themap = ones
			} else {
				themap = zeros
			}
		} else {
			if count1 >= count0 {
				themap = zeros
			} else {
				themap = ones
			}
		}

		if len(themap) == 1 {
			for _, v := range themap {
				attribute, _ = strconv.ParseInt(string(v), 2, 64)
				return attribute
			}
		}
	}
	return 0
}
