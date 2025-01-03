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

/*
func ReadFileToLinesOfRunes(filename string) (runes [][]rune) {
	lines  := ReadFileToByteSliceSlice(filename)


	return runes
}
*/

// readFileToByteSliceSlice - read file with os.ReadFile then create slice
func ReadFileToUnits2DSlice(filename string) (units2DSlice [][]int) {
	theBytes := ReadFileToByteSlice(filename)
	lines := bytes.Split(theBytes, []byte("\n"))
	width := len(lines[0])
	units2DSlice = make([][]int, 0, len(lines))
	for _, line := range lines {
		unitsLine := make([]int, width)
		for i, char := range line {
			unitsLine[i], _ = strconv.Atoi(string(char))
		}
		units2DSlice = append(units2DSlice, unitsLine)
	}
	return units2DSlice
}

// readFileToByteSliceSlice - read file with os.ReadFile then create slice
func ReadSpacesFileToIntSliceofSlices(filename string) (intSliceofSlices [][]int) {
	theBytes := ReadFileToByteSlice(filename)
	lines := bytes.Split(theBytes, []byte("\n"))
	intSliceofSlices = make([][]int, 0, len(lines))
	for _, line := range lines {
		lineSlice := bytes.Fields(line)
		intLine := make([]int, len(lineSlice))
		for i, num := range lineSlice {
			intLine[i], _ = strconv.Atoi(string(num))
		}
		intSliceofSlices = append(intSliceofSlices, intLine)
	}
	return intSliceofSlices
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

// ReadFileToStringSlice - read file with os.ReadFile then create slice
func ReadFileToStringSliceofSlices(filename string) (stringSliceOfSlices [][]string) {
	lines := ReadFileToStringSlice(filename)
	stringSliceOfSlices = make([][]string, 0, len(lines))
	for _, line := range lines {
		lineSlice := strings.Fields(line)
		stringSliceOfSlices = append(stringSliceOfSlices, lineSlice)
	}
	return stringSliceOfSlices
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

/////// less efficient below this line than alternatives above

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
