package main

import (
	"testing"
)

func BenchmarkPart2Only(b *testing.B) {
	realSlice := readCommandToSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(realSlice)
	}
}
