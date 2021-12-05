package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {

	fmt.Printf("Test Part 1: %v\n", part1("test-input.txt"))
	fmt.Printf("Test Part 2 : %v\n\n", part2("test-input.txt"))

	fmt.Printf("Real Part 1: %v\n", part1("real-input.txt"))
	fmt.Printf("Real Part 2: %v\n\n", part2("real-input.txt"))

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

func part1(filename string) (sum int) {
	commandSlice := readCommandToSlice(filename)
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

// renames part 2
func part2(filename string) (sum int) {
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
