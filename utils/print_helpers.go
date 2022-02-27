package utils

import "fmt"

func PrintByteSliceSlice(theSlice [][]byte) {
	fmt.Print("[")
	for _, bytes := range theSlice {
		fmt.Printf(" %s,", bytes)
	}
	fmt.Print(" ]\n")
}

func PrintByteSlice(theSlice []byte) {
	for _, bytes := range theSlice {
		fmt.Printf(" %c,", bytes)
	}
}

func PrintRuneSlice(theSlice []rune) {
	for _, bytes := range theSlice {
		fmt.Printf(" %c,", bytes)
	}
}
