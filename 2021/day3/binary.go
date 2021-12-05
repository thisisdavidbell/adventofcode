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
	fmt.Printf("Test Part2 %v\n", part2OnePass(testStringMap))
	fmt.Printf("Test Part2 Loops: %v\n\n", part2Loops(testStringMap))

	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	fmt.Printf("Real Part2 OnePass: %v\n", part2OnePass(realStringMap))
	fmt.Printf("Real Part2 Loops: %v\n", part2Loops(realStringMap))
	fmt.Printf("Real Part2 Delete: %v\n", part2Delete(realStringMap))

	realStringSlice2 := utils.ReadFileToByteSliceSlice("real-input.txt")
	fmt.Printf("Real Part2 Slice: %v\n", part2Slice(realStringSlice2))

	realStringSliceReuse := utils.ReadFileToByteSliceSlice("real-input.txt")
	fmt.Printf("Real Part2 Slice Reuse: %v\n", part2SliceReuse(realStringSliceReuse))

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

func part2OnePass(oxygenMap map[int]string) (power int64) {

	co2Map := utils.CopyIntStringMap(oxygenMap)
	oxygen := getAttributeOnePass(oxygenMap, true)
	co2 := getAttributeOnePass(co2Map, false)

	return oxygen * co2
}

func getAttributeOnePass(themap map[int]string, most bool) (attribute int64) {
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

///////////////////////////////////

func part2Loops(oxygenMap map[int]string) (power int64) {

	co2Map := utils.CopyIntStringMap(oxygenMap)
	oxygen := getAttributeLoops(oxygenMap, true)
	co2 := getAttributeLoops(co2Map, false)

	return oxygen * co2
}

func getAttributeLoops(themap map[int]string, most bool) (attribute int64) {
	numChars := len(themap[0])
	for char := 0; char < numChars; char++ { // for every column (each char position)
		numLines := len(themap)
		count1 := 0
		for _, line := range themap { // for every line at that position
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
		keep := make(map[int]string) //ToDo: borrow idea to rewrite and shrink existing themap
		for i, line := range themap {
			if line[char] == match {
				keep[i] = line // wrong. but dont need map[int]string, just map[stringstruct] or slice
			}
		}
		themap = keep
		if len(themap) == 1 {
			for _, v := range themap {
				attribute, _ = strconv.ParseInt(string(v), 2, 64)
				return attribute
			}
		}
	}
	return 0
}

/////////////////

func part2Delete(oxygenMap map[int]string) (power int64) {

	co2Map := utils.CopyIntStringMap(oxygenMap)
	oxygen := getAttributeDelete(oxygenMap, true)
	co2 := getAttributeDelete(co2Map, false)
	return oxygen * co2
}

func getAttributeDelete(themap map[int]string, most bool) (attribute int64) {
	numChars := len(themap[0])
	for char := 0; char < numChars; char++ { // for every column (each char position)
		numLines := len(themap)
		count1 := 0
		for _, line := range themap { // for every line at that position
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
		for i, line := range themap {
			if line[char] != match {
				delete(themap, i)
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

/////////////////

func part2Slice(oxygenSlice [][]byte) (power int64) {

	co2Slice := make([][]byte, len(oxygenSlice))
	copy(co2Slice, oxygenSlice)
	oxygen := getAttributeSlice(oxygenSlice, true)
	co2 := getAttributeSlice(co2Slice, false)

	return oxygen * co2
}

func getAttributeSlice(theSlice [][]byte, most bool) (attribute int64) {
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
		lenMatch := numLines - count1
		if most {
			if count1*2 >= numLines {
				match = byte('1')
				lenMatch = count1
			}
		} else {
			if count1*2 < numLines {
				match = byte('1')
				lenMatch = count1
			}
		}
		keep := make([][]byte, 0, lenMatch)
		for _, line := range theSlice {
			if line[char] == match {
				keep = append(keep, line)
			}
		}
		theSlice = keep
		if len(theSlice) == 1 {
			for _, v := range keep {
				attribute, _ = strconv.ParseInt(string(v), 2, 64)
				return attribute
			}
		}
	}
	return 0
}

/////////////////

func part2SliceReuse(oxygenSlice [][]byte) (power int64) {

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
