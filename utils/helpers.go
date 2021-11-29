package utils

import (
	"fmt"
	"os"
)

// CheckErr -  check for errors
func CheckErr(name string, err error) {
	if err != nil {
		fmt.Printf("Error found at %v. error: %v", name, err)
		os.Exit(1)
	}
}
