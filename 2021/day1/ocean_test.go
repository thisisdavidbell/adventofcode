package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkPart2countSumIncrease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realIntSlice := utils.ImportFileToIntSlice("real-input.txt")
		part2countSumIncrease(realIntSlice)
	}
}
