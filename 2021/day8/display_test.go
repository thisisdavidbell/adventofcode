package main

import (
	"testing"
)

func BenchmarkPart2All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2All("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	displays := readInput("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(displays)
	}
}
