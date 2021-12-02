package main

import (
	"testing"
)

func BenchmarkReadCommandToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readCommandToSlice("real-input.txt")
	}
}

func BenchmarkPart2Only(b *testing.B) {
	realSlice := readCommandToSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(realSlice)
	}
}

func BenchmarkPart2LetterOnly(b *testing.B) {
	realSlice := readCommandToSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2Letter(realSlice)
	}
}

func BenchmarkAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realSlice := readCommandToSlice("real-input.txt")
		part2(realSlice)
	}
}

func BenchmarkPerfPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		perfPart2("real-input.txt")
	}
}

func BenchmarkPerfPart2Cast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		perfPart2Cast("real-input.txt")
	}
}
