package main

import (
	"fmt"

	"github.com/thisisdavidbell/adventofcode/utils"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	//	fmt.Printf("Test Part 2: %v\n\n", part2("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1All("real-input.txt"))
	//	fmt.Printf("Real Part 2: %v\n", part2("test-input.txt"))
}

func part1All(filename string) (answer int) {
	units := utils.ReadFileToUnits2DSlice(filename)
	answer = part1(units)
	return answer
}

/* plan
- read in 2d slice
- for each entry:
    - check if up,down,left,right is more (use single function and case on UP DOWN LEFT RIGHT to encapsulate edge checking)
	- lazy evaluate. If all true, found a low point.
	- store co-ordinates and height (for now as we don't know what we need later) in another slice
- calculate sum of (each low point +1)
*/
func part1(units [][]int) int {
	lowPointRiskLevels := make([]int, 0, 0) // for now just save what i need - the actual height +1 of lowpoints (can when we see part2 save co-ords, or even just sum as we go)
	for y, line := range units {
		for x, _ := range line {
			if isPossibleLowPoint(x, y, units, Up) && isPossibleLowPoint(x, y, units, Down) && isPossibleLowPoint(x, y, units, Left) && isPossibleLowPoint(x, y, units, Right) {
				//fmt.Printf("Found Lowpoint. x %v, y=%v, height=%v\n", x, y, height)
				lowPointRiskLevels = append(lowPointRiskLevels, units[y][x]+1)
			}
		}
	}
	return sumRiskLevels(lowPointRiskLevels)
}

func isPossibleLowPoint(x, y int, units [][]int, dir Direction) (possibleLowPoint bool) {
	switch dir {
	case (Up):
		if y == 0 {
			return true
		}
		if units[y][x] < units[y-1][x] {
			return true
		} else {
			return false
		}

	case (Down):
		if y == len(units)-1 {
			return true
		}
		if units[y][x] < units[y+1][x] {
			return true
		} else {
			return false
		}

	case (Left):
		if x == 0 {
			return true
		}
		if units[y][x] < units[y][x-1] {
			return true
		} else {
			return false
		}

	case (Right):
		if x == len(units[0])-1 {
			return true
		}
		if units[y][x] < units[y][x+1] {
			return true
		} else {
			return false
		}
	}
	return possibleLowPoint

}

func sumRiskLevels(lowPointRiskLevels []int) (sum int) {
	for _, risk := range lowPointRiskLevels {
		sum += risk
	}
	return sum
}

/*
func part2() () {

	return
}
*/

/*
func part2All() () {

	return
}
*/
