package main

import (
	"fmt"
)

// TARGET - valid values should sum to this
var TARGET int = 2020

func main() {
	//test answer
	nestedLoopTestAnswer, err := SolveItWithNestedLoops("test-input.txt")
	checkErr("Test main.SolveItWithNestedLoops", err)
	fmt.Printf("\nNestedLoop test answer part1 : %v\n", nestedLoopTestAnswer)

	mapTestAnswer, mapTestAnswer2, err := SolveItWithMap("test-input.txt")
	checkErr("Test main.SolveItWithMap", err)
	fmt.Printf("Map test answer part1: %v\n", mapTestAnswer)
	fmt.Printf("Map test answer part2: %v\n\n", mapTestAnswer2)

	//real answer
	nestedLoopRealAnswer, err := SolveItWithNestedLoops("real-input.txt")
	checkErr("Real main.SolveItWithNestedLoops", err)
	fmt.Printf("NestedLoop real answer part 1: %v\n", nestedLoopRealAnswer)

	mapRealAnswer, mapRealAnswer2, err := SolveItWithMap("real-input.txt")
	checkErr("Real main.SolveItWithMap", err)
	fmt.Printf("\nNestedLoop real answer part 1: %v\n", mapRealAnswer)
	fmt.Printf("NestedLoop real answer part 2: %v\n", mapRealAnswer2)
}
