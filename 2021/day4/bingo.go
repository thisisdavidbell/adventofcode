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
	fmt.Printf("Real Part 2: %v\n", allPart2("real-input.txt"))

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
	for _, n := range nums {
		numbersMap[n] = struct{}{}
		for boardKey, b := range boards {
			if checkBoardWins(numbersMap, b) {
				winningBoards[boardKey] = struct{}{}
				if len(winningBoards) == len(boards) {
					return calculateAnswer(b, numbersMap, n)
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
