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
	fmt.Printf("Part 1 test count: %v\n", countTrees(lineSlice, 3, 1))

	realLineSlice, _ := importFileTo2DSlice("real-input.txt")
	fmt.Printf("Part 1 real count: %v\n", countTrees(realLineSlice, 3, 1))

	//quick and dirty
	fmt.Printf("\nTest Part 2 count: %v\n",
		countTrees(lineSlice, 1, 1)*
			countTrees(lineSlice, 3, 1)*
			countTrees(lineSlice, 5, 1)*
			countTrees(lineSlice, 7, 1)*
			countTrees(lineSlice, 1, 2))

	fmt.Printf("Part 2 test count: %v\n",
		countTrees(realLineSlice, 1, 1)*
			countTrees(realLineSlice, 3, 1)*
			countTrees(realLineSlice, 5, 1)*
			countTrees(realLineSlice, 7, 1)*
			countTrees(realLineSlice, 1, 2))
}

func importFileTo2DSlice(filename string) (lineSlice []string, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	utils.CheckErr("Open", err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineSlice = append(lineSlice, scanner.Text())
	}
	return lineSlice, err
}

func countTrees(grid []string, right int, down int) (count int) {
	for d, r := down, right; d < len(grid); d, r = d+down, (r+right)%len(grid[0]) {
		if grid[d][r] == byte('#') {
			count++
		}
	}
	return count
}
