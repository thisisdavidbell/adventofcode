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
	lines := readInputs(filename)
	grid := createGrid(989, 988)
	for _, aLine := range lines {
		if isHorizVertLine(aLine) {
			count = applyCoords(aLine, grid, count)
		}
	}
	return count
}

func part2All(filename string) int {
	lines := readInputs(filename)
	return part2(lines, 989, 988)
}

func part2(lines []line, maxX int, maxY int) int {
	count := 0
	grid := createGrid(maxX, maxY)
	for _, aLine := range lines {
		count = applyCoords(aLine, grid, count)
	}
	return count
}

func readInputs(filename string) (linesCoords []line) {
	fileLines := utils.ReadFileToStringSlice(filename)
	linesCoords = make([]line, 0, len(fileLines))
	x1, y1, x2, y2 := 0, 0, 0, 0
	for _, fileLine := range fileLines {
		//x1, y1, x2, y2 = 0, 0, 0, 0
		fmt.Sscanf(fileLine, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		linesCoords = append(linesCoords, line{p1: createComposite(x1, y1), p2: createComposite(x2, y2)})
	}
	return
}

// create an initialise 2D grid of ints
func createGrid(maxX int, maxY int) (grid []int) {
	val := ((maxX + 1) * 1000) + maxY + 1
	grid = make([]int, val, val)
	//	for xPos := 0; xPos < maxX+1; xPos++ {
	//		grid = append(grid, make([]int, maxY+1, maxY+1))
	//	}
	return
}

// is data a horiz/vert line?
func isHorizVertLine(aLine line) bool {
	if getX(aLine.p1) == getX(aLine.p2) || getY(aLine.p1) == getY(aLine.p2) {
		return true
	}
	return false
}

func applyCoords(aLine line, grid []int, count int) int {
	incX, stopX := 0, 0
	incY, stopY := 0, 0
	x1 := getX(aLine.p1)
	y1 := getY(aLine.p1)

	x2 := getX(aLine.p2)
	y2 := getY(aLine.p2)

	if x1 == x2 {
		incX = 0
		stopX = x1 + 1
	} else if x1 < x2 {
		incX = 1
		stopX = x2 + 1
	} else { //(aLine.x1 > aLine.x2)
		incX = -1
		stopX = x2 - 1
	}
	if y1 == y2 {
		incY = 0
		stopY = y1 + 1
	} else if y1 < y2 {
		incY = 1
		stopY = y2 + 1
	} else { //(aLine.y1 > aLine.y2)
		incY = -1
		stopY = y2 - 1
	}

	for x, y := x1, y1; x != stopX && y != stopY; x, y = x+incX, y+incY {
		grid[createComposite(x, y)]++
		if grid[createComposite(x, y)] == 2 {
			count++
		}
	}
	return count
}

/*
// print grid (with dots) for debug
func printGrid(grid [][]int) {
	// orig createGrid:
	grid = make([][]int, 0, maxX+1)
	for xPos := 0; xPos < maxX+1; xPos++ {
		grid = append(grid, make([]int, maxY+1, maxY+1))
	}
	return

	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			fmt.Printf("%v\t", grid[x][y])
		}
		fmt.Println()
	}

}
*/

func createComposite(x int, y int) int {
	return (x * 1000) + y
}

func getX(comp int) int {
	return comp % 1000
}

func getY(comp int) int {
	return (comp - (comp % 1000)) / 1000
}

type line struct {
	p1, p2 int
}
