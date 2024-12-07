package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	//answer1, answer2 := calculateAnswers("test-input.txt")
	answer1, answer2 := calculateAnswers("real-input.txt")

	fmt.Printf("Answer 1: %v\n", answer1)
	fmt.Printf("Answer 2: %v\n", answer2)
}

func calculateAnswers(filename string) (answer1, answer2 int) {

	lines := utils.ReadFileToStringSlice(filename)
	length := len(lines)
	firstLocations := make([]int, 0, length)
	secondLocations := make([]int, 0, length)

	for _, line := range lines {
		locationIDs := strings.Fields(line)
		firstLocations = append(firstLocations, utils.IgnoreError(strconv.Atoi(locationIDs[0])))
		secondLocations = append(secondLocations, utils.IgnoreError(strconv.Atoi(locationIDs[1])))
	}

	slices.Sort(firstLocations)
	slices.Sort(secondLocations)

	occurrences := make(map[int]int)

	for i := 0; i < len(secondLocations); i++ {
		// answer 1
		answer1 += int(utils.IntAbs(firstLocations[i] - secondLocations[i]))

		//answer 2
		_, ok := occurrences[secondLocations[i]]
		if ok {
			occurrences[secondLocations[i]]++
		} else {
			occurrences[secondLocations[i]] = 1
		}
	}

	for _, l := range firstLocations {
		answer2 += l * occurrences[l]
	}

	return answer1, answer2
}
