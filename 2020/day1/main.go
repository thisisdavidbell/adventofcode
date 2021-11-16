package main

import (
	"fmt"
)

// TARGET - valid values should sum to this
var TARGET int = 2020

func main() {
	nestedLoopTestAnswer, err := SolveItWithNestedLoops("test-input.txt")
	checkErr("Test main.SolveItWithNestedLoops", err)
	fmt.Printf("NestedLoop test answer: %v\n", nestedLoopTestAnswer)

	nestedLoopRealAnswer, err := SolveItWithNestedLoops("real-input.txt")
	checkErr("Real main.SolveItWithNestedLoops", err)
	fmt.Printf("NestedLoop test answer: %v\n", nestedLoopRealAnswer)

	mapTestAnswer, err := SolveItWithMap("test-input.txt")
	checkErr("Test main.SolveItWithMap", err)
	fmt.Printf("Map test answer: %v\n", mapTestAnswer)

	mapRealAnswer, err := SolveItWithMap("real-input.txt")
	checkErr("Real main.SolveItWithMap", err)
	fmt.Printf("NestedLoop test answer: %v\n", mapRealAnswer)
}
