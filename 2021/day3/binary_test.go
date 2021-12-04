package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkPart1Bytes(b *testing.B) {
	realByteSliceSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part1Bytes(realByteSliceSlice)
	}
}
func BenchmarkPart1BytesAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		realByteSliceSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
		part1Bytes(realByteSliceSlice)
	}
}

func BenchmarkPart1String(b *testing.B) {
	realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part1String(realStringSlice)
	}
}

func BenchmarkPart1StringAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToStringSlice("real-input.txt")
		part1String(realStringSlice)
	}
}

func BenchmarkPart2(b *testing.B) {
	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(realStringMap)
	}
}

func BenchmarkPart2All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringMap := utils.ImportFileToStringMap("real-input.txt")
		part2(realStringMap)
	}
}
