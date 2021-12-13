package main

import (
	"github.com/thisisdavidbell/adventofcode/utils"
	"testing"
)

func BenchmarkPart2All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2All("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	input, min, max := utils.ReadFileOfCommaSeperatedIntsToSliceWithMinMax("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(input, min, max)
	}
}
