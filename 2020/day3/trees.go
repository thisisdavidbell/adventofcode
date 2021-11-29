package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/thisisdavidbell/adventofcode/utils"
)

// TREE - latter that represents a tree in input data
const TREE string = "#"

func main() {
	lineSlice, _ := importFileTo2DSlice("test-input.txt")
	fmt.Printf("Test count: %v\n", countTrees(lineSlice))

	realLineSlice, _ := importFileTo2DSlice("real-input.txt")
	fmt.Printf("Real count: %v\n", countTrees(realLineSlice))
}

func importFileTo2DSlice(filename string) (lineSlice [][]bool, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	utils.CheckErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineSlice = append(lineSlice, convertLineToSliceOfTrees(scanner.Text()))
	}
	return lineSlice, err
}

func convertLineToSliceOfTrees(line string) (slice []bool) {
	for _, letter := range line {
		if string(letter) == TREE {
			slice = append(slice, true)
		} else {
			slice = append(slice, false)
		}
	}
	return slice
}

func countTrees(grid [][]bool) (count int) {
	DOWN := 1
	RIGHT := 3

	for d, r := DOWN, RIGHT; d < len(grid); d, r = d+DOWN, (r+RIGHT)%len(grid[0]) {
		//fmt.Printf("d: %v, r: %v, tree: %v\n", d, r, grid[d][r])
		if grid[d][r] {
			count++
		}
	}
	return count
}
