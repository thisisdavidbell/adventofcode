package main

import (
	"fmt"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", part2All("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n", part2All("real-input.txt"))
}

func readInputPart1(filename string) (displays []display) {
	lines := utils.ReadFileToStringSlice(filename)
	displays = make([]display, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "|")
		d := display{digits: strings.Fields(parts[0]), example: strings.Fields(parts[1])}
		displays = append(displays, d)
	}
	return
}

func part1All(filename string) (count int) {
	displays := readInputPart1(filename)
	count = part1(displays)
	return
}

func part1(displays []display) (count int) {
	for _, d := range displays {
		for _, digit := range d.example {
			switch len(digit) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return
}

func readInputPart2(filename string) (displays []displayPart2) {
	lines := utils.ReadFileToStringSlice(filename)
	displays = make([]displayPart2, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "|")
		digitsMap := makeDigitsMap()
		for _, digit := range strings.Fields(parts[0]) {
			digitsMap[len(digit)] = append(digitsMap[len(digit)], digit)
		}
		d := displayPart2{digits: digitsMap, example: strings.Fields(parts[1])}
		displays = append(displays, d)
	}
	return
}

func part2All(filename string) (count int) {
	displays := readInputPart2(filename)
	//fmt.Printf("displays: %+v\n", displays)
	count = part2(displays)
	return
}

func part2(displays []displayPart2) (count int) {
	for _, d := range displays { // each line in input
		correctDigitLetters := make(map[int]string, 0)
		//find top:
		top := removeChars(d.digits[3][0], (d.digits[2][0]))
		applyLetters(correctDigitLetters, top, []int{0, 2, 3, 5, 6, 7, 8, 9})

		otherCharsinFour := removeChars(d.digits[4][0], d.digits[2][0])

		// find three
		three := ""
		for _, f := range d.digits[5] {
			if strings.Contains(f, string(d.digits[2][0][0])) && strings.Contains(f, string(d.digits[2][0][1])) {
				three = f
				break
				//fmt.Printf("three: %v\n", three)
			}
		}
		// find middle
		middle := ""
		if strings.Contains(three, string(otherCharsinFour[0])) {
			middle = string(otherCharsinFour[0])
		} else {
			middle = string(otherCharsinFour[1])
		}
		applyLetters(correctDigitLetters, middle, []int{2, 3, 4, 5, 6, 8, 9})

		// top-left must be remaining char in 4:
		topleft := removeChars(otherCharsinFour, middle)
		applyLetters(correctDigitLetters, topleft, []int{0, 4, 5, 6, 8, 9})

		// find char left after removing 7 and middle from three, must be bottom.
		bottom := removeChars(three, d.digits[3][0])
		bottom = removeChars(bottom, middle)

		applyLetters(correctDigitLetters, bottom, []int{0, 2, 3, 5, 6, 8, 9})

		// bottom left must be 8 less 3, less topleft

		bottomleft := removeChars(d.digits[7][0], three)
		bottomleft = removeChars(bottomleft, topleft)
		applyLetters(correctDigitLetters, bottomleft, []int{0, 2, 6, 8})

		//find top right and bottom right -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
		topright := ""
		for _, f := range d.digits[6] {
			if !(strings.Contains(f, string(d.digits[2][0][0])) && strings.Contains(f, string(d.digits[2][0][1]))) {
				if strings.Contains(f, string(d.digits[2][0][0])) {
					topright = string(d.digits[2][0][1])
				} else {
					topright = string(d.digits[2][0][0])
				}
				break
			}
		}
		applyLetters(correctDigitLetters, topright, []int{0, 1, 2, 3, 4, 7, 8, 9})

		//bottom right remaining segment of 1
		bottomright := removeChars(d.digits[2][0], topright)
		applyLetters(correctDigitLetters, bottomright, []int{0, 1, 3, 4, 5, 6, 7, 8, 9})

		// now match the examples
		count += findActualDigits(correctDigitLetters, d.example)
	}

	return
}

func findActualDigits(correctDigitLetters map[int]string, examples []string) (value int) {
	actualDigits := make([]int, 0)
	for _, ex := range examples {
		matchedInt := 0
		for k, s := range correctDigitLetters {
			if len(ex) == len(s) {
				match := true
				for _, c := range ex {
					if !strings.Contains(s, string(c)) {
						match = false
						break
					}
				}
				if match {
					matchedInt = k
					break
				}
			}
		}
		actualDigits = append(actualDigits, matchedInt)
	}
	value = actualDigits[0]*1000 + actualDigits[1]*100 + actualDigits[2]*10 + actualDigits[3]
	return
}

type display struct {
	digits  []string
	example []string
}

func applyLetters(correctDigitLetters map[int]string, ch string, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = correctDigitLetters[s] + ch
	}
}

func removeChars(s string, remove string) (changed string) {
	filter := func(r rune) rune {
		if strings.ContainsRune(remove, r) {
			return -1
		}
		return r
	}
	return strings.Map(filter, s)
}

func makeDigitsMap() (digitsMap map[int][]string) {
	digitsMap = make(map[int][]string)
	digitsMap[2] = make([]string, 0, 1)
	digitsMap[3] = make([]string, 0, 1)
	digitsMap[4] = make([]string, 0, 1)
	digitsMap[5] = make([]string, 0, 3)
	digitsMap[6] = make([]string, 0, 3)
	digitsMap[7] = make([]string, 0, 1)
	return
}

type displayPart2 struct {
	digits  map[int][]string //map of length of digit to lsice containing all digits of that length
	example []string
}
