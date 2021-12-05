package utils

import "fmt"

func PrintByteSliceSlice(theSlice [][]byte) {
	fmt.Print("[")
	for _, bytes := range theSlice {
		fmt.Printf(" %s", bytes)
	}
	fmt.Print(" ]\n")
}
