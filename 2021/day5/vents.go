package main

import (
	"fmt"
	"os"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1("real-input.txt"))

	//	fmt.Printf("Real Part 1: %v\n", part1(realSlice))
	//	fmt.Printf("Real Part 2: %v\n", part2(realslice))
}

func part1(filename string) (count int) {
	lines, maxX, maxY := readInputs(filename)
	grid := createGrid(maxX, maxY)
	for _, aLine := range lines {
		if isHorizVertLine(aLine) {
			applyCoords(aLine, grid)
		}
	}
	return countNumIntersects(grid)
}

/*
func part2() () {

	return
}
*/

// read input into ?
func readInputs(filename string) (linesCoords []line, maxX int, maxY int) {
	fileLines := utils.ReadFileToStringSlice(filename)
	linesCoords = make([]line, 0, len(fileLines))
	for _, fileLine := range fileLines {
		lineCoord := line{}
		fmt.Sscanf(fileLine, "%d,%d -> %d,%d", &lineCoord.x1, &lineCoord.y1, &lineCoord.x2, &lineCoord.y2)
		linesCoords = append(linesCoords, lineCoord)
		if lineCoord.x1 > maxX {
			maxX = lineCoord.x1
		}
		if lineCoord.x2 > maxX {
			maxX = lineCoord.x2
		}
		if lineCoord.y1 > maxY {
			maxY = lineCoord.y1
		}
		if lineCoord.y2 > maxY {
			maxY = lineCoord.y2
		}

	}
	return
}

// create an initialise 2D grid of ints
func createGrid(maxX int, maxY int) (grid [][]int) {
	grid = make([][]int, 0, maxX)
	for xPos := 0; xPos < maxX+1; xPos++ {
		grid = append(grid, make([]int, maxY+1, maxY+1))
	}
	return
}

// is data a horiz/vert line?
func isHorizVertLine(aLine line) bool {
	if aLine.x1 == aLine.x2 || aLine.y1 == aLine.y2 {
		return true
	}
	return false
}

// apply co-ord to grid
// drb: consider if code works with a single point (i.e. x1,y1=x2,y2)
func applyCoords(aLine line, grid [][]int) {
	if aLine.x1 == aLine.x2 {
		if aLine.y1 >= aLine.y2 {
			for i := aLine.y2; i != aLine.y1+1; i++ {
				grid[aLine.x1][i]++
			}
		} else {
			for i := aLine.y2; i != aLine.y1-1; i-- {
				grid[aLine.x1][i]++
			}
		}
	} else if aLine.y1 == aLine.y2 {
		if aLine.x1 >= aLine.x2 {
			for i := aLine.x2; i != aLine.x1+1; i++ {
				grid[i][aLine.y1]++
			}
		} else {
			for i := aLine.x2; i != aLine.x1-1; i-- {
				grid[i][aLine.y1]++
			}
		}
	} else {
		fmt.Errorf("ERROR: we should never get here")
		os.Exit(1)
	}
}

// countNumIntersects in grid
func countNumIntersects(grid [][]int) (count int) {
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

// print grid (with dots) for debug
func printGrid(grid [][]int) {
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			fmt.Printf("%v\t", grid[x][y])
		}
		fmt.Println()
	}

}

//struct for point? so []point
type line struct {
	x1, y1, x2, y2 int
}
