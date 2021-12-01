package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CheckErr -  check for errors
func CheckErr(name string, err error) {
	if err != nil {
		fmt.Printf("Error found at %v. error: %v", name, err)
		os.Exit(1)
	}
}

// ImportFileToIntSlice - read ints from file into slice
func ImportFileToIntSlice(filename string) (numberSlice []int) {
	f, err := os.Open(filename)
	defer f.Close()
	CheckErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		theInt, err := strconv.Atoi(scanner.Text())
		CheckErr("ImportFileToSlice.Atoi", err)
		numberSlice = append(numberSlice, theInt)
	}
	return numberSlice
}
