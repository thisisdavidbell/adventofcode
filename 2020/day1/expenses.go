package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

// ImportFileToMap - read numbers from file into slice
func ImportFileToMap(filename string) (map[int]struct{}, error) {
	//numberMap := make(map[int]struct{}, 200) // would make for a much faster method, but no efficient way to count lines of file...
	numberMap := make(map[int]struct{})
	f, err := os.Open(filename)
	defer f.Close()
	checkErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		theInt, err := strconv.Atoi(scanner.Text())
		checkErr("ImportFileToSlice.Atoi", err)
		numberMap[theInt] = struct{}{}
	}
	return numberMap, err
}

// FindAnswerUsingMap - use maps to avoid nasty nested loops
func FindAnswerUsingMap(numberMap map[int]struct{}) (answer1 int, answer2 int, err error) {
	for i1 := range numberMap {
		//part1
		if _, ok := numberMap[TARGET-i1]; ok {
			answer1 = i1 * (TARGET - i1)
		}
		delete(numberMap, i1) // ToDo more efficient way to sub map
		//part2
		for j2 := range numberMap {
			if _, ok := numberMap[TARGET-i1-j2]; ok {
				answer2 = i1 * j2 * (TARGET - i1 - j2)
			}
		}
	}
	if answer1 == 0 || answer2 == 0 {
		err = errors.New("No valid values found")
	}
	return answer1, answer2, err
}

// SolveItWithMap find the matching numbers
func SolveItWithMap(filename string) (answer int, answer2 int, err error) {
	numberMap, err := ImportFileToMap(filename)
	checkErr("SolveItWithMap.ImportFileToMap", err)

	answer, answer2, err = FindAnswerUsingMap(numberMap)
	checkErr("SolveItWithMAp.FindAnswerUsingMap", err)
	return
}
