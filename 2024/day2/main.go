package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	answer1test, answer2test := calculateAnswers("test-input.txt")
	answer1real, answer2real := calculateAnswers("real-input.txt")

	fmt.Printf("Test answer 1: %v\n", answer1test)
	fmt.Printf("Test answer 2: %v\n", answer2test)

	fmt.Printf("\nReal answer 1: %v\n", answer1real)
	fmt.Printf("Real answer 2: %v\n", answer2real)
}

func calculateAnswers(filename string) (answer1, answer2 int) {

	/* part 1 plan:
	- [x] read whole file into slice of slices
	- [x] create function to compare 2 numbers
	- [x] create function which checks if a line safe
	- [x] check each line and count number of safe lines
	*/

	data := utils.ReadSpacesFileToIntSliceofSlices(filename)

	for _, report := range data {

		// part 1
		if isReportSafe(report) {
			answer1++
		}

		/* Part 2 plan:
		- for each report,
		  - create len-1 sub-slices with 1 part missing
		  - if any of these are safe, return safe(true), else return unsafe(false)
		*/

		// part 2
		reportStatus := false
		for i := 0; i < len(report); i++ {
			subReport := utils.RemoveIntSliceElement(report, i)
			if isReportSafe(subReport) {
				reportStatus = true
				break
			}
		}
		if reportStatus {
			answer2++
		}

	}
	return answer1, answer2
}

func isReportSafe(report []int) bool {

	/*
		plan:
		- process first pair
			have a mode/increasing or decreasing
		- work through each subsequent pair checking for:
			- not same
			- increasing/decreasing matching mode
			- by 1,2 or 3 only
	*/

	increasing := false // default decreasing
	if report[0] == report[1] {
		return false
	}
	if report[0] < report[1] {
		increasing = true
	}

	for i := 0; i+1 < len(report); i++ {
		if report[i] == report[i+1] {
			return false
		}
		absDiff := utils.IntAbs(report[i] - report[i+1])
		if absDiff < 1 || absDiff > 3 {
			// fmt.Println("diff to next too big/small")
			return false
		}
		if increasing {
			if report[i] > report[i+1] {
				// fmt.Println("now decreasing")
				return false
			}
		} else {
			if report[i] < report[i+1] {
				// fmt.Println("now increasing")
				return false
			}
		}
	}

	return true
}
