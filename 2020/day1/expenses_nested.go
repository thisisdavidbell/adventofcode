package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func checkErr(name string, err error) {
	if err != nil {
		fmt.Printf("Error found at %v. error: %v", name, err)
		os.Exit(1)
	}
}

// ImportFileToSlice - read numbers from file into slice
func ImportFileToSlice(filename string) (numberSlice []int, err error) {
	f, err := os.Open(filename)
	checkErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		theInt, err := strconv.Atoi(scanner.Text())
		checkErr("ImportFileToSlice.Atoi", err)
		numberSlice = append(numberSlice, theInt)
	}
	return numberSlice, err
}

// FindAnswerUsingNestedLoops - find the 2 values that match 2020 in the provided int slice, using basic nested loop approach (will test later that this is slower than Map approach)
func FindAnswerUsingNestedLoops(numberSlice []int) (int, error) {

	for i1, v1 := range numberSlice {
		for _, v2 := range numberSlice[i1+1:] {
			if v1+v2 == TARGET {
				return v1 * v2, nil
			}
		}
	}
	return 0, errors.New("Not valid values found")
}

// SolveItWithNestedLoops find the matching numbers
func SolveItWithNestedLoops(filename string) (answer int, err error) {
	numberSlice, err := ImportFileToSlice(filename)
	checkErr("SolveItWithNestedLoops.ImportFileToSlice", err)
	answer, err = FindAnswerUsingNestedLoops(numberSlice)
	checkErr("SolveItWithNestedLoops.Find2ValuesUsingNestedLoops", err)
	return
}
