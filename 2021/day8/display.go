package main

import (
	"bytes"
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", part2All("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n", part2All("real-input.txt"))
}

func part1All(filename string) (count int) {
	line := readInput(filename)
	count = part1(line)
	return
}

func part1(line []byte) (count int) {
	displays := processInput(line)
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

func readInput(filename string) (line []byte) {
	line = utils.ReadFileToByteSlice(filename)
	return
}

func processInput(line []byte) (displays []display) {
	lines := bytes.Split(line, []byte("\n"))
	displays = make([]display, 0, len(lines))
	for _, line := range lines {
		parts := bytes.Split(line, []byte{'|'})
		d := display{example: bytes.Fields(parts[1])}

		//digitsMap := makeDigitsMap()
		for _, digit := range bytes.Fields(parts[0]) {
			switch len(digit) {
			case 2:
				d.lenTwo = digit
			case 3:
				d.lenThree = digit
			case 4:
				d.lenFour = digit
			case 5:
				d.lenFive = append(d.lenFive, digit)
			case 6:
				d.lenSix = append(d.lenSix, digit)
			case 7:
				d.lenSeven = digit
			}
		}
		//d := display{digits: digitsMap, example: strings.Fields(parts[1])}
		displays = append(displays, d)
	}
	return
}

func part2All(filename string) (count int) {
	line := readInput(filename)
	count = part2(line)
	return
}

// todo: update to remove string cassts - use byte everywhere we can... and move strings.* to bytes.*
func part2(line []byte) (count int) {
	displays := processInput(line)
	for _, d := range displays {
		correctDigitLetters := make(map[int][]byte, 10)
		//find top:
		top := removeChars(d.lenThree, (d.lenTwo))
		applyLetters(correctDigitLetters, top[0], []int{0, 2, 3, 5, 6, 7, 8, 9})

		otherCharsinFour := removeChars(d.lenFour, d.lenTwo)

		// find three
		var three []byte
		for _, f := range d.lenFive {
			if bytes.Contains(f, []byte{d.lenTwo[0]}) && bytes.Contains(f, []byte{d.lenTwo[1]}) {
				three = f
				break
			}
		}
		// find middle
		var middle []byte
		if bytes.Contains(three, []byte{otherCharsinFour[0]}) {
			middle = []byte{otherCharsinFour[0]}
		} else {
			middle = []byte{otherCharsinFour[1]}
		}
		applyLetters(correctDigitLetters, middle[0], []int{2, 3, 4, 5, 6, 8, 9})

		// top-left must be remaining char in 4:
		topleft := removeChars(otherCharsinFour, middle)
		applyLetters(correctDigitLetters, topleft[0], []int{0, 4, 5, 6, 8, 9})

		// find char left after removing 7 and middle from three, must be bottom.
		bottom := removeChars(three, d.lenThree)
		bottom = removeChars(bottom, middle)

		applyLetters(correctDigitLetters, bottom[0], []int{0, 2, 3, 5, 6, 8, 9})

		// bottom left must be 8 less 3, less topleft
		bottomleft := removeChars(d.lenSeven, three)
		bottomleft = removeChars(bottomleft, topleft)
		applyLetters(correctDigitLetters, bottomleft[0], []int{0, 2, 6, 8})

		//find top right and bottom right -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
		var topright []byte
		for _, f := range d.lenSix {
			if !(bytes.Contains(f, []byte{d.lenTwo[0]}) && bytes.Contains(f, []byte{d.lenTwo[1]})) {
				if bytes.Contains(f, []byte{d.lenTwo[0]}) {
					topright = []byte{d.lenTwo[1]}
				} else {
					topright = []byte{d.lenTwo[0]}
				}
				break
			}
		}
		applyLetters(correctDigitLetters, topright[0], []int{0, 1, 2, 3, 4, 7, 8, 9})

		//bottom right remaining segment of 1
		bottomright := removeChars(d.lenTwo, topright)
		applyLetters(correctDigitLetters, bottomright[0], []int{0, 1, 3, 4, 5, 6, 7, 8, 9})

		// now match the examples
		count += findActualDigits(correctDigitLetters, d.example)
	}

	return
}

func findActualDigits(correctDigitLetters map[int][]byte, examples [][]byte) (value int) {
	actualDigits := make([]int, 0, 4)
	for _, ex := range examples {
		matchedInt := 0
		switch len(ex) {
		case 2:
			matchedInt = 1
		case 3:
			matchedInt = 7
		case 4:
			matchedInt = 4
		case 7:
			matchedInt = 8
		case 5:
			for _, k := range []int{2, 3, 5} {
				match := true
				for _, c := range ex {
					if !bytes.ContainsRune(correctDigitLetters[k], rune(c)) {
						match = false
						break
					}
				}
				if match {
					matchedInt = k
					break
				}
			}
		case 6:
			for _, k := range []int{0, 6, 9} {
				match := true
				for _, c := range ex {
					if !bytes.ContainsRune(correctDigitLetters[k], rune(c)) {
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

func applyLetters(correctDigitLetters map[int][]byte, ch byte, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = append(correctDigitLetters[s], ch)
	}
}

func removeChars(s []byte, remove []byte) (changed []byte) {
	filter := func(r rune) rune {
		if bytes.ContainsRune(remove, r) {
			return -1
		}
		return r
	}
	return bytes.Map(filter, s)
}

type display struct {
	lenTwo   []byte
	lenThree []byte
	lenFour  []byte
	lenFive  [][]byte
	lenSix   [][]byte
	lenSeven []byte

	example [][]byte
}
