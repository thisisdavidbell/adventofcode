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

func BenchmarkPart2OnePass(b *testing.B) {
	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2OnePass(realStringMap)
	}
}

func BenchmarkPart2Loops(b *testing.B) {
	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2Loops(realStringMap)
	}
}

func BenchmarkPart2Delete(b *testing.B) {
	realStringMap := utils.ImportFileToStringMap("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2Delete(realStringMap)
	}
}

func BenchmarkPart2Slice(b *testing.B) {
	realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2Slice(realStringSlice)
	}
}

func BenchmarkPart2SliceReuse(b *testing.B) {
	realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2SliceReuse(realStringSlice)
	}
}

func BenchmarkPart2OnePassAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringMap := utils.ImportFileToStringMap("real-input.txt")
		part2OnePass(realStringMap)
	}
}

func BenchmarkPart2LoopsAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringMap := utils.ImportFileToStringMap("real-input.txt")
		part2Loops(realStringMap)
	}
}

func BenchmarkPart2DeleteAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringMap := utils.ImportFileToStringMap("real-input.txt")
		part2Delete(realStringMap)
	}
}

func BenchmarkPart2SliceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
		part2Slice(realStringSlice)
	}
}

func BenchmarkPart2SliceReuseAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realStringSlice := utils.ReadFileToByteSliceSlice("real-input.txt")
		part2SliceReuse(realStringSlice)
	}
}
