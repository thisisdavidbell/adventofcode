package main

import (
	"fmt"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	//fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", part2All("test-input.txt"))

	//fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt"))
	//	fmt.Printf("Real Part 2: %v\n", part2("test-input.txt"))
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
	//fmt.Printf("displays: %+v", displays)
	count = part1(displays)
	return
}

//drb: make an actual plan
//  - dont second guess part 2 - just resolve part1 and rewrite when its clear whats needed!
//  -

/*
Unique:

num	| number of segments
------------------------
1	 2
4	 4
7	 3
8	 7
*/

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

/*
  0:6     1:2     2:5     3:5     4:4
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:5     6:6     7:3     8:7     9:6
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

/*
num	| number of segments
------------------------
0	6
1	2 <==
2	5
3	5
4	4 <==
5	5
6	6
7	3 <==
8	7 <==
9	6
*/

/*
plan:
 - do just need an approach:
Manual:
latters to match: d
matched: a b c e f g

    7     5     5     5     3    6     6     4     6    2
acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf

=========
working solution 1:

ab = 1
dab = 7
diff = d = top

4 contains ab from 1 and middle and top left, so can find these 2 letters

the only 5 with a and b in (from 1) must be 3 - which only contains one of e/f (from 4) - which ever it contains must be middle! leaving other e/f as top-left
so f middle
e top left

remaining char in 3 must be bottom?
bottom = c

final char in 4 must be top-left

char not used must (not in 3 or 4) be bottom left
bottom- left - g

only 1 of the 3 sixes is missing top right left, so whichever of a and b is missing from 1 of the sixes, is the top right.
a = top right

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc

 ============


*/

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

/*
plan
- work out an algorithm: DONE
- goal: have a map of number being displayed, to letters that make it up... map[int]string? // for now
- fill in following algorthm - just code the logic.
- with completed map:
  - for each example digit in display, for each entry in map: if len(digit) = entry length
    - for char in entry, if string.contains // could try man of maps later map[int]map[string]bool(trade off of create time versus lookup time.)
	   - break on match
	   - method to convert 4 vals to a 4 digiit number
prep:
- how do I want input data stored?
  - map[length][]strings
*/

func part2(displays []displayPart2) (count int) {
	for _, d := range displays { // each line in input
		correctDigitLetters := make(map[int]string, 0)
		//find top:
		top := strings.ReplaceAll(d.digits[3][0], string(d.digits[2][0][0]), "")
		top = strings.ReplaceAll(top, string(d.digits[2][0][1]), "")
		//fmt.Printf("top: %v\n", top)

		applyLetters(correctDigitLetters, top, []int{0, 2, 3, 5, 6, 7, 8, 9})

		//the only len 5 with a and b in (from 1) must be 3 - which only contains one of e/f - which ever it contains must be middle! leaving other e/f as top-left
		// find other chars in four
		otherCharsinFour := strings.ReplaceAll(d.digits[4][0], string(d.digits[2][0][0]), "")
		otherCharsinFour = strings.ReplaceAll(otherCharsinFour, string(d.digits[2][0][1]), "")

		// find three
		three := ""
		for _, f := range d.digits[5] {
			if strings.Contains(f, string(d.digits[2][0][0])) && strings.Contains(f, string(d.digits[2][0][1])) {
				three = f
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
		//fmt.Printf("middle: %v\n", middle)
		applyLetters(correctDigitLetters, middle, []int{2, 3, 4, 5, 6, 8, 9})

		// top-left must be remaining char in 4:
		topleft := strings.ReplaceAll(otherCharsinFour, middle, "")
		//fmt.Printf("topleft: %v\n", topleft)
		applyLetters(correctDigitLetters, topleft, []int{0, 4, 5, 6, 8, 9})

		//fmt.Printf("three: %v\n", three)
		//fmt.Printf("seven: %v\n", d.digits[3][0])
		//fmt.Printf("removing: %v\n", string(d.digits[3][0][0]))
		bottom := strings.ReplaceAll(three, string(d.digits[3][0][0]), "")
		bottom = strings.ReplaceAll(bottom, string(d.digits[3][0][1]), "")
		bottom = strings.ReplaceAll(bottom, string(d.digits[3][0][2]), "")
		//fmt.Printf("three less 7: %v\n", bottom)

		// find char left after removing 7 and middle from three, must be bottom.
		bottom = strings.ReplaceAll(bottom, string(middle), "")

		//fmt.Printf("bottom: %v\n", bottom)
		applyLetters(correctDigitLetters, bottom, []int{0, 2, 3, 5, 6, 8, 9})

		// bottom left must be 8 less 3, less topleft

		bottomleft := strings.ReplaceAll(d.digits[7][0], string(three[0]), "")
		bottomleft = strings.ReplaceAll(bottomleft, string(three[1]), "")
		bottomleft = strings.ReplaceAll(bottomleft, string(three[2]), "")
		bottomleft = strings.ReplaceAll(bottomleft, string(three[3]), "")
		bottomleft = strings.ReplaceAll(bottomleft, string(three[4]), "")
		bottomleft = strings.ReplaceAll(bottomleft, topleft, "")

		//fmt.Printf("bottomleft: %v\n", bottomleft)
		applyLetters(correctDigitLetters, bottomleft, []int{0, 2, 6, 8})

		//find topleft and bottom left -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
		topright := ""
		for _, f := range d.digits[6] {
			if !(strings.Contains(f, string(d.digits[2][0][0])) && strings.Contains(f, string(d.digits[2][0][1]))) {
				if strings.Contains(three, string(d.digits[2][0][0])) {
					topright = string(d.digits[2][0][0])
				} else {
					topright = string(d.digits[2][0][1])
				}
			}
		}
		//fmt.Printf("topright: %v\n", topright)
		applyLetters(correctDigitLetters, topright, []int{0, 1, 2, 3, 4, 7, 8, 9})

		// bottom right remaining segment of 1
		bottomright := strings.ReplaceAll(d.digits[2][0], topright, "")
		//fmt.Printf("bottomright: %v\n", bottomright)
		applyLetters(correctDigitLetters, bottomright, []int{0, 1, 3, 4, 5, 6, 7, 8, 9})

		// now match the examples
		actualDigits := make([]int, 0)
		for _, ex := range d.example {
			//fmt.Printf("Example: %v\n", ex)
			matchedInt := 0
			for k, s := range correctDigitLetters {
				//fmt.Printf("check each correctDigitLetter: %v\n", s)
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
			//fmt.Printf("found digit: %v\n", matchedInt)
			actualDigits = append(actualDigits, matchedInt)

		}
		value := actualDigits[0]*1000 + actualDigits[1]*100 + actualDigits[2]*10 + actualDigits[3]
		fmt.Println(value)
		//exit so only process 1st line
		//fmt.Printf("correctDigitLetters: %v\n", correctDigitLetters)
		//os.Exit(1)
	}

	return
}

/*
func part2All() () {

	return
}
*/

type display struct {
	digits  []string
	example []string
}

func applyLetters(correctDigitLetters map[int]string, ch string, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = correctDigitLetters[s] + ch
	}
}

/*
func makeCorrectDigitLettersMap() (correctDigitsMap map[int]string) {

}
*/
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
