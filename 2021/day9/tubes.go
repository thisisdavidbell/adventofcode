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

/* part 1 plan
- read in 2d slice
- for each entry:
    - check if up,down,left,right is more (use single function and case on UP DOWN LEFT RIGHT to encapsulate edge checking)
	- lazy evaluate. If all true, found a low point.
	- store only risk for now - min we need. We may need co-ordinates later, but easy to add if so
- calculate sum of (each low point +1)
*/
func part1(units [][]int) (int, []lowPoint) {
	lowPointRiskLevels := make([]lowPoint, 0)
	for y, line := range units {
		for x, _ := range line {
			if isLowPoint(x, y, units) {
				lowPointRiskLevels = append(lowPointRiskLevels, lowPoint{x, y, units[y][x] + 1})
			}
		}
	}
	return sumRiskLevels(lowPointRiskLevels), lowPointRiskLevels
}

func isLowPoint(x, y int, units [][]int) bool {
	if y == 0 || units[y][x] < units[y-1][x] { // Up
		if y == len(units)-1 || units[y][x] < units[y+1][x] { // Down
			if x == 0 || units[y][x] < units[y][x-1] { // Left
				if x == len(units[0])-1 || units[y][x] < units[y][x+1] { // Right
					return true
				}
			}
		}
	}
	return false

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
		if y == len(units)-1 { // down would be off bottom edge
			return 0
		}
		adjX = x
		adjY = y + 1

	case (Left):
		if x == 0 { // left would be off left edge
			return 0
		}
		adjX = x - 1
		adjY = y

	case (Right):
		if x == len(units[0])-1 { // right would be off right edge
			return 0
		}
		adjX = x + 1
		adjY = y
	}

	if units[adjY][adjX] >= 9 { // found top, or an already processed location, so return
		return 0
	}
	// remove 10 added to ensure location not considered in future
	if units[y][x]-10 < units[adjY][adjX] { // adj is uphill - part of basin. +1 and recurse
		count++
		units[adjY][adjX] += 10 //ensure this location is not processed again
		count += checkLocation(adjX, adjY, units)
	}
	// effectice else lower, return count which is still = 0
	return count

}

/* Plan
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

Clean that above will have dups.
Avoiding duplicates:

Notes:
- you have to always do up/down/left/right - passages tubes could be any shape - only way to find them all.
   - so the only way to overcome recursion problem of duplicates is to compare the processed co-ords, either
      - at the time, which cuts off duplicate branches <- clear winner
	  - remove duplicates at the end before counting basin size - very inefficient

Solution:
- set processed location to greater than 9 when its processed. (no race conditions as we will remain single threaded, and order is deterministic.)
  - note just setting to 9 wont work - current location height needed comparing adjacents after recurse.
  - we can however add 10, and remove it. current location when checking adjacents will always have been processed and +10'ed, so we can always remove 10 before compare.
    - We can check if adj >= 9, so its either a top, or an already processed location
*/

func part2(filename string) (answer int) {
	_, lowPoints, units := part1All(filename)
	basinSizes := make([]int, len(lowPoints))
	for i, aLowPoint := range lowPoints {
		basinSizes[i]++                       // count lowpoint
		units[aLowPoint.y][aLowPoint.x] += 10 // ensure this location is never counted again
		basinSizes[i] += checkLocation(aLowPoint.x, aLowPoint.y, units)
	}
	sort.Ints(basinSizes)
	answer = basinSizes[(len(basinSizes)-1)] * basinSizes[(len(basinSizes)-2)] * basinSizes[(len(basinSizes)-3)]
	return answer
}
