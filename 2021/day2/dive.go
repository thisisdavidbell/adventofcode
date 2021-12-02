package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {

	testSlice := readCommandToSlice("test-input.txt")
	fmt.Printf("Part1 test: %v\n", part1(testSlice))
	fmt.Printf("Part2 test: %v\n\n", part2(testSlice))

	realSlice := readCommandToSlice("real-input.txt")
	fmt.Printf("Part1 real: %v\n", part1(realSlice))
	fmt.Printf("Part1 real: %v\n", part2(realSlice))

}

func part1(commandSlice []command) (sum int) {
	x, depth := 0, 0
	for _, com := range commandSlice {
		switch com.direction {
		case "forward":
			x = x + com.units
		case "down":
			depth = depth + com.units
		case "up":
			depth = depth - com.units
		}
	}
	return x * depth
}

func part2(commandSlice []command) (sum int) {
	x, depth, aim := 0, 0, 0
	for _, com := range commandSlice {
		switch com.direction {
		case "forward":
			x = x + com.units
			depth = depth + (aim * com.units)
		case "down":
			aim = aim + com.units
		case "up":
			aim = aim - com.units
		}
	}
	return x * depth
}

// readCommandToSlice - read string and int into slice from file
func readCommandToSlice(filename string) (commandSlice []command) {
	f, err := os.Open(filename)
	defer f.Close()
	utils.CheckErr("Open", err)
	com := command{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "%s %d", &com.direction, &com.units)
		utils.CheckErr("fmtSscanf", err)
		commandSlice = append(commandSlice, com)
		strconv.Atoi(scanner.Text())
	}
	return
}

type command struct {
	direction string
	units     int
}
