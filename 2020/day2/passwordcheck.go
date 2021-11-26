package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count, count2 := solveItForFile("test-input.txt")
	fmt.Printf("Part 1 test count: %v\n", count)
	fmt.Printf("Part 2 test count: %v\n\n", count2)

	countReal, countReal2 := solveItForFile("real-input.txt")
	fmt.Printf("Part 1 real count: %v\n", countReal)
	fmt.Printf("Part 2 real count: %v\n", countReal2)

}

// PWData - struct for key values from each line of file
type PWData struct {
	min      int
	max      int
	letter   string
	password string
}

// checkError -  check for errors
func checkErr(name string, err error) {
	if err != nil {
		fmt.Printf("Error found at %v. error: %v", name, err)
		os.Exit(1)
	}
}

// ImportFileToSlice - read each line into a slice
func ImportFileToSlice(filename string) (lineSlice []string, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	checkErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineSlice = append(lineSlice, scanner.Text())
	}
	return lineSlice, err
}

// SliceLinesToSliceStruct convert slice of lines to slice of structs
func SliceLinesToSliceStruct(lineSlice []string) (pwSlice []PWData, err error) {
	for _, line := range lineSlice {
		pwData := PWData{}
		fmt.Sscanf(line, "%d-%d %1s: %s", &pwData.min, &pwData.max, &pwData.letter, &pwData.password)
		pwSlice = append(pwSlice, pwData)
	}
	return
}

func checkValidPasswordPart1(pwData PWData) bool {
	count := strings.Count(pwData.password, pwData.letter)
	return count >= pwData.min && count <= pwData.max
}

func checkValidPasswordPart2(pwData PWData) bool {
	match1 := string(pwData.password[pwData.min-1]) == pwData.letter
	match2 := string(pwData.password[pwData.max-1]) == pwData.letter

	return (match1 && !match2) || (!match1 && match2)
}

func countValidPWsInPWDataPart1(pwDataSlice []PWData) (count int) {
	for _, pwData := range pwDataSlice {
		if checkValidPasswordPart1(pwData) {
			count++
		}
	}
	return
}

func countValidPWsInPWDataPart2(pwDataSlice []PWData) (count int) {
	for _, pwData := range pwDataSlice {
		if checkValidPasswordPart2(pwData) {
			count++
		}
	}
	return
}

func solveItForFile(inputfile string) (count int, count2 int) {
	lineSlice, _ := ImportFileToSlice(inputfile)
	structSlice, _ := SliceLinesToSliceStruct(lineSlice)
	count = countValidPWsInPWDataPart1(structSlice)
	count2 = countValidPWsInPWDataPart2(structSlice)
	return
}
