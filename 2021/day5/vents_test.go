package main

import (
	"testing"
)

func BenchmarkPart2(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		part2("real-input.txt")
	}
}
