package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkPart1String(b *testing.B) {
	realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part1String(realStringSlice)
	}
}

func BenchmarkPart1StringAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
		part1String(realStringSlice)
	}
}

func BenchmarkPart2SliceReuse(b *testing.B) {
	realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2SliceReuse(realStringSlice)
	}
}

func BenchmarkPart2SliceReuseAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
		part2SliceReuse(realStringSlice)
	}
}
