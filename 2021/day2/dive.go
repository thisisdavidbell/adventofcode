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
	fmt.Printf("Part2 real: %v\n\n", part2(realSlice))

	fmt.Printf("Pref Part2 real: %v\n", perfPart2("real-input.txt"))
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
	}
	return
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

/////////////////// perf comparisons ///////////////////

func part2Letter(commandSlice []command) (sum int) {
	x, depth, aim := 0, 0, 0
	for _, com := range commandSlice {
		switch com.direction[0] {
		case 'f':
			x = x + com.units
			depth = depth + (aim * com.units)
		case 'd':
			aim = aim + com.units
		case 'u':
			aim = aim - com.units
		}
	}
	return x * depth
}

func perfPart2(filename string) (sum int) {
	f, _ := os.Open(filename)
	defer f.Close()
	var direction string
	var units int
	x, depth, aim := 0, 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &units)
		switch direction[0] {
		case 'f':
			x = x + units
			depth = depth + (aim * units)
		case 'd':
			aim = aim + units
		case 'u':
			aim = aim - units
		}
	}
	return x * depth
}

func perfPart2Cast(filename string) (sum int) {
	f, _ := os.Open(filename)
	defer f.Close()
	var direction, unitsStr string
	var units int
	x, depth, aim := 0, 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %s", &direction, &unitsStr)
		units, _ = strconv.Atoi(unitsStr)
		switch direction[0] {
		case 'f':
			x = x + units
			depth = depth + (aim * units)
		case 'd':
			aim = aim + units
		case 'u':
			aim = aim - units
		}
	}
	return x * depth
}

type command struct {
	direction string
	units     int
}
