package main

import (
	"testing"
)

func BenchmarkReadBingoInputToSlices(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		readBingoInputToSlices("real-input.txt")
	}
}

func BenchmarkReadBingoInputToMaps(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		readBingoInputToMaps("real-input.txt")
	}
}

func BenchmarkAllPart2(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		allPart2("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	nums, boards := readBingoInputToMaps("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(nums, boards)
	}
}
