package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", part2All("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n", part2All("real-input.txt"))
}

func part1(filename string) (count int) {
	lines, maxX, maxY := readInputs(filename)
	grid := createGrid(maxX, maxY)
	for _, aLine := range lines {
		if isHorizVertLine(aLine) {
			applyCoords(aLine, grid, 0)
		}
	}
	return countNumIntersects(grid)
}

func part2All(filename string) int {
	lines, maxX, maxY := readInputs(filename)
	return part2(lines, maxX, maxY)
}

func part2(lines []line, maxX int, maxY int) int {
	count := 0
	grid := createGrid(maxX, maxY)
	for _, aLine := range lines {
		count = applyCoords(aLine, grid, count)
	}
	return count //countNumIntersects(grid)
}

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

func applyCoords(aLine line, grid [][]int, count int) int {
	incX, stopX := 0, 0
	incY, stopY := 0, 0
	if aLine.x1 == aLine.x2 {
		incX = 0
		stopX = aLine.x1 + 1
	} else if aLine.x1 < aLine.x2 {
		incX = 1
		stopX = aLine.x2 + 1
	} else { //(aLine.x1 > aLine.x2)
		incX = -1
		stopX = aLine.x2 - 1
	}
	if aLine.y1 == aLine.y2 {
		incY = 0
		stopY = aLine.y1 + 1
	} else if aLine.y1 < aLine.y2 {
		incY = 1
		stopY = aLine.y2 + 1
	} else { //(aLine.y1 > aLine.y2)
		incY = -1
		stopY = aLine.y2 - 1
	}

	for x, y := aLine.x1, aLine.y1; x != stopX && y != stopY; x, y = x+incX, y+incY {
		grid[x][y]++
		if grid[x][y] == 2 {
			count++
		}
	}
	return count
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

type line struct {
	x1, y1, x2, y2 int
}
