package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkPart1(b *testing.B) {
	realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part1(realStringSlice)
	}
}

func BenchmarkPart1All(b *testing.B) {

	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
		part1(realStringSlice)
	}
}

func BenchmarkPart2(b *testing.B) {
	realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(realStringSlice)
	}
}

func BenchmarkPart2All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
		part2(realStringSlice)
	}
}
