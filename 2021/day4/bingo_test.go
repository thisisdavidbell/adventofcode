package main

import (
	"testing"
)

func BenchmarkAllPart2(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		allPart2("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	nums, boards := readBingoInputToSlices("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(nums, boards)
	}
}

func BenchmarkPart2Perf(b *testing.B) {
	nums, boards := readBingoInputToMaps("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2Perf(nums, boards)
	}
}
