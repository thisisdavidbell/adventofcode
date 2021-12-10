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

	fmt.Printf("Real Part 1: %v\n", allPart1("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n\n", allPart2("real-input.txt"))

}

var boardSize int = 5

func allPart1(filename string) int {
	nums, boards := readBingoInputToMaps(filename)
	return part1(nums, boards)
}

func part1(nums []string, boards [][]map[string]struct{}) int {

	for _, n := range nums {
		for _, b := range boards {
			if checkBoardWins(n, b) {
				return calculateAnswer(n, b)
			}
		}
	}

	return 0
}

func allPart2(filename string) int {
	nums, boards := readBingoInputToMaps(filename)
	return part2(nums, boards)
}

func part2(nums []string, boards [][]map[string]struct{}) int {

	winningBoards := make(map[int]struct{})
	for _, n := range nums {
		for boardKey := 0; boardKey < len(boards); boardKey++ {
			if checkBoardWins(n, boards[boardKey]) {
				winningBoards[boardKey] = struct{}{}
				if len(winningBoards) == len(boards) {
					//if len(boards) == 1 {
					return calculateAnswer(n, boards[boardKey])
				}
				//boards = removeBoard(boards, boardKey)
				//boardKey--
			}
		}
	}
	return 0
}

func removeBoard(boards [][]map[string]struct{}, index int) [][]map[string]struct{} {
	boards[index] = boards[0]
	return boards[1:]
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

func checkBoardWins(number string, board []map[string]struct{}) bool {

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

func calculateAnswer(number string, board []map[string]struct{}) (answer int) {
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
