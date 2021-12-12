package main

import (
	"testing"
)

func BenchmarkPart2All(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		part2All("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := readInputs("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(lines, 989, 988)
	}
}
