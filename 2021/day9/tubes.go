package main

import (
	"fmt"
	"sort"

	"github.com/thisisdavidbell/adventofcode/utils"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type lowPoint struct {
	x, y, risk int
}

func main() {
	answer, _, _ := part1All("test-input.txt")
	fmt.Printf("Test Part 1: %v\n", answer)
	answer, _, _ = part1All("real-input.txt")
	fmt.Printf("Real Part 1: %v\n\n", answer)

	fmt.Printf("Test Part 2: %v\n", part2("test-input.txt"))

	fmt.Printf("Real Part 2: %v\n", part2("real-input.txt"))
}

func part1All(filename string) (answer int, lowPoints []lowPoint, units [][]int) {
	units = utils.ReadFileToUnits2DSlice(filename)
	answer, lowPoints = part1(units)
	return answer, lowPoints, units
}

/* plan
- read in 2d slice
- for each entry:
    - check if up,down,left,right is more (use single function and case on UP DOWN LEFT RIGHT to encapsulate edge checking)
	- lazy evaluate. If all true, found a low point.
	- store co-ordinates and height (for now as we don't know what we need later) in another slice
- calculate sum of (each low point +1)
*/
func part1(units [][]int) (int, []lowPoint) {
	lowPointRiskLevels := make([]lowPoint, 0) // for now just save what i need - the actual height +1 of lowpoints (can when we see part2 save co-ords, or even just sum as we go)
	for y, line := range units {
		for x, _ := range line {
			if isPossibleLowPoint(x, y, units, Up) && isPossibleLowPoint(x, y, units, Down) && isPossibleLowPoint(x, y, units, Left) && isPossibleLowPoint(x, y, units, Right) {
				//fmt.Printf("Found Lowpoint. x %v, y=%v, height=%v\n", x, y, height)
				lowPointRiskLevels = append(lowPointRiskLevels, lowPoint{x, y, units[y][x] + 1})
			}
		}
	}
	return sumRiskLevels(lowPointRiskLevels), lowPointRiskLevels
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

func sumRiskLevels(lowPointRiskLevels []lowPoint) (sum int) {
	for _, aLowPoint := range lowPointRiskLevels {
		sum += aLowPoint.risk
	}
	return sum
}

func checkLocation(x, y int, units [][]int) (count int) {

	count += checkAdjacent(x, y, units, Up)
	count += checkAdjacent(x, y, units, Down)
	count += checkAdjacent(x, y, units, Left)
	count += checkAdjacent(x, y, units, Right)

	return count
}

func checkAdjacent(x, y int, units [][]int, dir Direction) (count int) {
	var adjX, adjY int
	switch dir {
	case (Up):
		if y == 0 { // up would be off top edge of grid
			return 0
		}
		adjX = x
		adjY = y - 1

	case (Down):
		if y == len(units)-1 {
			return 0
		}
		adjX = x
		adjY = y + 1

	case (Left):
		if x == 0 {
			return 0
		}
		adjX = x - 1
		adjY = y

	case (Right):
		if x == len(units[0])-1 {
			return 0
		}
		adjX = x + 1
		adjY = y
	}

	if units[adjY][adjX] >= 9 { // found top, or an already processed location, so return
		return 0
	}
	// remove 10 added to ensure location to considered in future
	if units[y][x]-10 < units[adjY][adjX] { // adj is uphill - part of basin. +1 and recurse
		//fmt.Printf("adj found. x:%v, y:%v\n", adjX, adjY)
		count++
		units[adjY][adjX] += 10 //ensure this location is never counted again
		count += checkLocation(adjX, adjY, units)
	}
	// effectice else lower, return count which still = 0
	return count

}

/* 1st plan doesnt work - some locations would be counted twice with recursion:
keep it simple - 2 pass throughs for now
- 1. find low points - part 1
  	- convert captured points to be x, y points only
- 2. for each low point:
	- getBasinSize(x,y)
		- checkAdjacent(Up)
		- checkAdjacent(Down)
		- checkAdjacent(Left)
		- checkAdjacent(Right)
		CheckAdjacent:
			- switch case for each (possiby just changing the compare to x and y values)
			- if edge, as before
			- if direction is lower, stop and return - we are only looking for uphill (if it goes down to ours it would be found on another oute. If to another, dont count it. If down and up, thats a diff lowpoint)
			- if direction is 9, stop and return - reached the top
			- if direction is higher:
				- basinSize++
				- basinSize += getBasinSize(compareX, compareY)
			- return basinSize

- find 3 largest basins. product!

Avoiding duplicates:

Notes:
- you have to always do up/down/left/right - passages tubes could be any shape - only way to find them all.
   - so the only way to overcome recursion problem of duplicates is to compare the processed co-ords, either
      - at the time, which cuts off duplicate branches, but needs careful locking/ordering
	  - remove duplicates at the end before counting basin size

Possible fixes to above recursive method:
- set processed location to 9 when its processed - no race condition as we remain single threaded.
  - wont work - value needs to be used when comparing adjacents after recurse.
  - we can however add 10, and remove it. current location when checking adjacents will always have been processed and +10'ed, so we can always remove 10 from currnet location.
    - We can check if adj >= 9, so its eitehr a top, or an already processed location
- have returned item be the set of co-ordinates in the basin - so can remove dups - but we will repeat processing nunecesarily
-
*/

func part2(filename string) (answer int) {
	_, lowPoints, units := part1All(filename)
	//fmt.Printf("Lowpoints: %v\n", lowPoints)
	basinSizes := make([]int, len(lowPoints))
	for i, aLowPoint := range lowPoints {
		//fmt.Printf("\n\nProcessing lowpoint: %v: %+v\n", i, aLowPoint)
		basinSizes[i]++                       // count lowpoint
		units[aLowPoint.y][aLowPoint.x] += 10 //ensure this location is never counted again
		basinSizes[i] += checkLocation(aLowPoint.x, aLowPoint.y, units)
	}
	//fmt.Printf("Basin sizes: %v\n", basinSizes)
	sort.Ints(basinSizes)
	//fmt.Printf("Sorted basin sizes: %v\n", basinSizes)
	answer = basinSizes[(len(basinSizes)-1)] * basinSizes[(len(basinSizes)-2)] * basinSizes[(len(basinSizes)-3)]
	return answer
}
