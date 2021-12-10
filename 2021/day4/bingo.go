package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {

	fmt.Printf("Test Part 1: %v\n", allPart1("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", allPart2("test-input.txt"))
	fmt.Printf("Test Part 2: %v\n\n", allPart2Perf("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", allPart1("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n\n", allPart2("real-input.txt"))
	fmt.Printf("Real Perf Part 2: %v\n\n", allPart2Perf("real-input.txt"))

}

var boardSize int = 5

func allPart1(filename string) int {
	nums, boards := readBingoInputToSlices(filename)
	return part1(nums, boards)
}

func part1(nums []string, boards [][][]string) int {

	numbersMap := make(map[string]struct{})
	for _, n := range nums {
		numbersMap[n] = struct{}{}
		for _, b := range boards {
			if checkBoardWins(numbersMap, b) {
				return calculateAnswer(b, numbersMap, n)
			}
		}
	}

	return 0
}

func allPart2(filename string) int {
	nums, boards := readBingoInputToSlices(filename)
	return part2(nums, boards)
}

func part2(nums []string, boards [][][]string) int {
	numbersMap := make(map[string]struct{})
	winningBoards := make(map[int]struct{})
	for key, n := range nums {
		numbersMap[n] = struct{}{}
		for boardKey, b := range boards {
			//change to only take in the latest number - board object replacement must maintain state
			if key < 4 {
				continue
			}
			if checkBoardWins(numbersMap, b) {
				winningBoards[boardKey] = struct{}{}
				if len(winningBoards) == len(boards) {
					// WRONG - unchecked numbers are what is left !! orig: need a way (if we removed numbers during checkBoardWins) to have full board for score...
					return calculateAnswer(b, numbersMap, n)
				}
			}
		}
	}
	return 0
}

func allPart2Perf(filename string) int {
	nums, boards := readBingoInputToMaps(filename)
	return part2Perf(nums, boards)
}

func part2Perf(nums []string, boards [][]map[string]struct{}) int {
	winningBoards := make(map[int]struct{})
	for _, n := range nums {
		for boardKey, b := range boards {
			if checkBoardWinsPerf(n, b) {
				winningBoards[boardKey] = struct{}{}
				if len(winningBoards) == len(boards) {
					return calculateAnswerPerf(n, b)
				}
			}
		}
	}
	return 0
}

func readBingoInputToSlices(filename string) (numbers []string, boards [][][]string) {
	numbers = make([]string, 0)
	boards = make([][][]string, 0)
	input := utils.ReadFileToStringSlice(filename)
	numbers = strings.Split(input[0], ",")

	for i := 2; i < len(input); i++ {
		aBoard := make([][]string, 0)
		for ; i < len(input) && input[i] != ""; i++ {
			aBoard = append(aBoard, strings.Fields(input[i]))
		}
		boards = append(boards, aBoard)
	}
	return
}

// a board is just a set of winnable Maps (i.e. a row or a column) - Maps so can remove
func readBingoInputToMaps(filename string) (numbers []string, boards [][]map[string]struct{}) {
	numbers = make([]string, 0)
	boards = make([][]map[string]struct{}, 0)
	input := utils.ReadFileToStringSlice(filename)
	numbers = strings.Split(input[0], ",")

	for i := 2; i < len(input); i++ {
		setsSlice := make([]map[string]struct{}, 0)
		colsSlice := make([]map[string]struct{}, 0)
		for j := 0; j < boardSize; j++ {
			colsSlice = append(colsSlice, make(map[string]struct{}))
		}
		//for each board
		for ; i < len(input) && input[i] != ""; i++ {
			//found rows of a board
			setMap := make(map[string]struct{}, 0)

			// for each line which is a row, build a map
			for key, e := range strings.Fields(input[i]) {
				// for each column
				setMap[e] = struct{}{}
				colsSlice[key][e] = struct{}{}
			}
			setsSlice = append(setsSlice, setMap)
		}
		setsSlice = append(setsSlice, colsSlice...)
		boards = append(boards, setsSlice)
	}
	return
}

func checkBoardWins(numbers map[string]struct{}, board [][]string) bool {

	matched := false

	//check row
	for _, r := range board {
		matched = true
		for _, c := range r {
			if _, ok := numbers[c]; !ok {
				matched = false
			}
		}
		if matched {
			return true
		}
	}

	// check columns
	matched = false
	for c := 0; c < len(board[0]); c++ {
		matched = true
		for r := 0; r < len(board); r++ {
			if _, ok := numbers[board[r][c]]; !ok {
				matched = false
			}
		}
		if matched {
			return true
		}
	}
	return false

}

func checkBoardWinsPerf(number string, board []map[string]struct{}) bool {

	//check row
	won := false
	for _, s := range board {
		for c := range s {
			if c == number {
				delete(s, c)
				if len(s) == 0 {
					won = true
				}
			}
		}
	}
	return won
}

func calculateAnswer(board [][]string, numbersMap map[string]struct{}, lastNum string) (answer int) {
	total := 0
	for r := range board {
		for c := range board[r] {
			if _, ok := numbersMap[board[r][c]]; !ok {
				theInt, _ := strconv.Atoi(board[r][c])
				total += theInt
			}
		}
	}
	lastNumInt, _ := strconv.Atoi(lastNum)
	return total * lastNumInt
}

func calculateAnswerPerf(number string, board []map[string]struct{}) (answer int) {
	total := 0
	for s := range board {
		for c := range board[s] {
			theInt, _ := strconv.Atoi(c)
			total += theInt
		}
	}
	numberInt, _ := strconv.Atoi(number)
	return (total * numberInt) / 2
}
