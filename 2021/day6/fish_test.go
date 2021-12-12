package main

import (
	"testing"
)

func BenchmarkSolveAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solveAll("real-input.txt", 256)
	}
}

func BenchmarkSolve(b *testing.B) {
	theCount := readInput("real-input,txt")
	for i := 0; i < b.N; i++ {
		solve(theCount, 265)
	}
}
