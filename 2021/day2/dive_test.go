package main

import (
	"testing"
)

func BenchmarkReadCommandToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readCommandToSlice("real-input.txt")
	}
}

func BenchmarkPart2All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2("real-input.txt")
	}
}
