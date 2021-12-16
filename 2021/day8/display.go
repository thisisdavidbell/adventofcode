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

func part1All(filename string) (count int) {
	displays := readInput(filename)
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

func readInput(filename string) (displays []display) {
	lines := utils.ReadFileToStringSlice(filename)
	displays = make([]display, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "|")
		d := display{example: strings.Fields(parts[1])}

		//digitsMap := makeDigitsMap()
		for _, digit := range strings.Fields(parts[0]) {
			switch len(digit) {
			case 2:
				d.lenTwo = append(d.lenTwo, digit)
			case 3:
				d.lenThree = append(d.lenThree, digit)
			case 4:
				d.lenFour = append(d.lenFour, digit)
			case 5:
				d.lenFive = append(d.lenFive, digit)
			case 6:
				d.lenSix = append(d.lenSix, digit)
			case 7:
				d.lenSeven = append(d.lenSeven, digit)
			}
		}
		//d := display{digits: digitsMap, example: strings.Fields(parts[1])}
		displays = append(displays, d)
	}
	return
}

func part2All(filename string) (count int) {
	displays := readInput(filename)
	count = part2(displays)
	return
}

func part2(displays []display) (count int) {
	for _, d := range displays {
		correctDigitLetters := make(map[int]string, 10)
		//find top:
		top := removeChars(d.lenThree[0], (d.lenTwo[0]))
		applyLetters(correctDigitLetters, top, []int{0, 2, 3, 5, 6, 7, 8, 9})

		otherCharsinFour := removeChars(d.lenFour[0], d.lenTwo[0])

		// find three
		three := ""
		for _, f := range d.lenFive {
			if strings.Contains(f, string(d.lenTwo[0][0])) && strings.Contains(f, string(d.lenTwo[0][1])) {
				three = f
				break
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
		bottom := removeChars(three, d.lenThree[0])
		bottom = removeChars(bottom, middle)

		applyLetters(correctDigitLetters, bottom, []int{0, 2, 3, 5, 6, 8, 9})

		// bottom left must be 8 less 3, less topleft
		bottomleft := removeChars(d.lenSeven[0], three)
		bottomleft = removeChars(bottomleft, topleft)
		applyLetters(correctDigitLetters, bottomleft, []int{0, 2, 6, 8})

		//find top right and bottom right -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
		topright := ""
		for _, f := range d.lenSix {
			if !(strings.Contains(f, string(d.lenTwo[0][0])) && strings.Contains(f, string(d.lenTwo[0][1]))) {
				if strings.Contains(f, string(d.lenTwo[0][0])) {
					topright = string(d.lenTwo[0][1])
				} else {
					topright = string(d.lenTwo[0][0])
				}
				break
			}
		}
		applyLetters(correctDigitLetters, topright, []int{0, 1, 2, 3, 4, 7, 8, 9})

		//bottom right remaining segment of 1
		bottomright := removeChars(d.lenTwo[0], topright)
		applyLetters(correctDigitLetters, bottomright, []int{0, 1, 3, 4, 5, 6, 7, 8, 9})

		// now match the examples
		count += findActualDigits(correctDigitLetters, d.example)
	}

	return
}

func findActualDigits(correctDigitLetters map[int]string, examples []string) (value int) {
	actualDigits := make([]int, 0, 4)
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

func applyLetters(correctDigitLetters map[int]string, ch string, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = correctDigitLetters[s] + ch
		//correctDigitLetters[s] = strings.Join([]string{correctDigitLetters[s], ch}, "")
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

type display struct {
	lenTwo   []string
	lenThree []string
	lenFour  []string
	lenFive  []string
	lenSix   []string
	lenSeven []string

	example []string
}
