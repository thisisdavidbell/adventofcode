package main

import (
	"testing"
)

func BenchmarkPart2Only(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		part2(realSlice)
	}
}
