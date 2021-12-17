package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CheckErr -  check for errors
func CheckErr(name string, err error) {
	if err != nil {
		fmt.Printf("Error found at %v. error: %v", name, err)
		os.Exit(1)
	}
}

// ReadFileToByteSlice - read file into a single long byte slice for whole file
func ReadFileToByteSlice(filename string) (bytes []byte) {
	bytes, _ = os.ReadFile(filename)
	return
}

// readFileToByteSliceSlice - read file with os.ReadFile then create slice
func ReadFileToByteSliceSlice(filename string) (byteSlice [][]byte) {
	theBytes := ReadFileToByteSlice(filename)
	return bytes.Split(theBytes, []byte("\n"))
}

// ReadFileToString - read file into a single long byte slice for whole file
func ReadFileToString(filename string) string {
	str, _ := os.ReadFile(filename)
	return string(str)
}

// ReadFileToStringSlice - read file with os.ReadFile then create slice
func ReadFileToStringSlice(filename string) (stringSlice []string) {
	str := ReadFileToString(filename)
	return strings.Split(str, "\n")
}

// ReadFileOfCommaSeperatedIntsToSlice - read ints from single line of file
func ReadFileOfCommaSeperatedIntsToSlice(filename string) (theInts []int) {
	str := ReadFileToString(filename)
	theStrings := strings.Split(str, ",")
	theInts = make([]int, 0, len(theStrings))
	for _, theString := range theStrings {
		theInt, _ := strconv.Atoi(theString)
		theInts = append(theInts, theInt)
	}
	return
}

func ReadFileOfCommaSeperatedIntsToSliceWithMinMax(filename string) (theInts []int, min int, max int) {
	str := ReadFileToString(filename)
	theStrings := strings.Split(str, ",")
	theInts = make([]int, 0, len(theStrings))
	for i, theString := range theStrings {
		theInt, _ := strconv.Atoi(theString)
		theInts = append(theInts, theInt)
		if i == 0 || theInt < min {
			min = theInt
		}
		if i == 0 || theInt > max {
			max = theInt
		}
	}
	return
}

/////// less efficient below this line that alternatives above

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

// ImportFileToByteSliceSlice - read file
func ImportFileToByteSliceSlice(filename string) (byteSlice [][]byte) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scannedBytes := scanner.Bytes()
		length := len(scannedBytes)
		theBytes := make([]byte, length, length) // https://pkg.go.dev/bufio#Scanner.Bytes states: The underlying array may point to data that will be overwritten by a subsequent call to Scan. It does no allocation.
		copy(theBytes, scanner.Bytes())
		byteSlice = append(byteSlice, theBytes)
	}
	return byteSlice
}

// ImportFileToStringSlice - read file
func ImportFileToStringSlice(filename string) (stringSlice []string) {
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		stringSlice = append(stringSlice, scanner.Text())
	}
	return
}

// ImportFileToStringMap - read file
func ImportFileToStringMap(filename string) (stringMap map[int]string) {
	stringMap = make(map[int]string)
	f, _ := os.Open(filename)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		stringMap[i] = scanner.Text()
		i++
	}
	return
}
