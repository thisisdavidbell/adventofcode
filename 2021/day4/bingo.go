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

	for n := 0; n < len(nums); n++ {
		for _, b := range boards {
			if checkBoardWins(nums[:n+1], b) {
				return calculateAnswer(b, nums[:n+1], n)
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
	//numbersMap := make(map[string]struct{})
	winningBoards := make(map[int]struct{})
	for n := 0; n < len(nums); n++ {
		for boardKey, b := range boards {
			if checkBoardWins(nums[:n+1], b) {
				winningBoards[boardKey] = struct{}{}
				if len(winningBoards) == len(boards) {
					return calculateAnswer(b, nums[:n+1], n)
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

func checkBoardWins(numbers []string, board [][]string) bool {

	matched := false

	//check row
	for r := 0; r < len(board); r++ {
		matched = true
		for c := 0; c < len(board[0]); c++ {
			if !numInNumbers(board[r][c], numbers) {
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
			if !numInNumbers(board[r][c], numbers) {
				matched = false
			}
		}
		if matched {
			return true
		}
	}
	return false

}

func numInNumbers(num string, numbers []string) bool {
	//	fmt.Printf("Checking num %v in numbers %v\n", num, numbers)
	for _, n := range numbers {
		if n == num {
			//		fmt.Println("FOUND")
			return true
		}
	}

	return false
}
func calculateAnswer(board [][]string, numbers []string, lastNum int) (answer int) {
	//	fmt.Printf("FOUND WINNER: board: %v, numbers: %v, lastNum: %v\n", board, numbers, lastNum)
	total := 0
	for r := range board {
		for c := range board[r] {
			if !numInNumbers(board[r][c], numbers) {
				theInt, _ := strconv.Atoi(board[r][c])
				//				fmt.Printf("NOT MATCHED %v, adding to total: %v\n", board[r][c], total)
				total += theInt
			}
		}
	}
	//	fmt.Printf("lastnum: %v\n", lastNum)
	lastNumInt, _ := strconv.Atoi(numbers[lastNum])
	return total * lastNumInt
}
