package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	answer1test := calculateAnswerQ1("test-input.txt")
	answer1real := calculateAnswerQ1("real-input.txt")

	answer2test := calculateAnswerQ2("test-input-q2.txt")
	answer2real := calculateAnswerQ2("real-input.txt")

	fmt.Printf("Test answer 1: %v\n", answer1test)
	fmt.Printf("Test answer 2: %v\n", answer2test)

	fmt.Printf("\nReal answer 1: %v\n", answer1real)
	fmt.Printf("Real answer 2: %v\n", answer2real)
}

func calculateAnswerQ1(filename string) int {

	fileAsString := utils.ReadFileToString(filename)
	return findAndMultiplyMuls(fileAsString)

}

func calculateAnswerQ2(filename string) int {

	fileAsByteSlice := utils.ReadFileToByteSlice(filename)
	theStringWithoutDonts := findAndRemoveDontSections(fileAsByteSlice)
	return findAndMultiplyMuls(theStringWithoutDonts)

}

// findAndMultiplyMuls - find all mul(x,y), multiply x and y and sum total
func findAndMultiplyMuls(theString string) int {

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(theString, -1)

	total := 0
	for _, element := range matches {
		re2 := regexp.MustCompile(`\d{1,3}`)
		nums := re2.FindAllString(element, -1)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		total = total + (num1 * num2)
	}
	return total
}

// findAndRemoveDontSections -  find all strings starting with "don't()", and remove the string upto "do()" or the end of the string
func findAndRemoveDontSections(theString []byte) string {

	foundEnd := false
	for !foundEnd {
		pattern := regexp.MustCompile(`don\'t\(\)`)
		dontLoc := pattern.FindIndex(theString)
		if dontLoc == nil {
			break
		}
		startDont := dontLoc[0]
		endDont := dontLoc[1]

		dopattern := regexp.MustCompile(`do\(\)`)
		doLoc := dopattern.FindIndex(theString[endDont:])
		doLocEnd := 0
		if doLoc == nil {
			doLocEnd = len(theString)
			fmt.Printf("doLocEnd=%v\n", doLocEnd)
			foundEnd = true
		} else {
			doLocEnd = endDont + doLoc[1]
		}
		theString = append(theString[:startDont], theString[doLocEnd:]...)
	}
	return string(theString)
}
