package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

// ImportFileToMap - read numbers from file into slice
func ImportFileToMap(filename string) (map[int]struct{}, error) {
	numberMap := make(map[int]struct{})
	f, err := os.Open(filename)
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
func FindAnswerUsingMap(numberMap map[int]struct{}) (int, error) {
	for i1, _ := range numberMap {
		if _, ok := numberMap[TARGET-i1]; ok {
			return i1 * (TARGET - i1), nil
		}
	}
	return 0, errors.New("Not valid values found")
}

// SolveItWithMap find the matching numbers
func SolveItWithMap(filename string) (answer int, err error) {
	numberMap, err := ImportFileToMap(filename)
	checkErr("SolveItWithMap.ImportFileToMap", err)

	answer, err = FindAnswerUsingMap(numberMap)
	checkErr("SolveItWithMAp.FindAnswerUsingMap", err)
	return
}
