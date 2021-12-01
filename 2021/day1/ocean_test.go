package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkImportFileToIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.ImportFileToIntSlice("real-input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realIntSlice := utils.ImportFileToIntSlice("real-input.txt")
		part2(realIntSlice)
	}
}

func BenchmarkPart2Only(b *testing.B) {
	realIntSlice := utils.ImportFileToIntSlice("real-input.txt")
	for i := 0; i < b.N; i++ {
		part2(realIntSlice)
	}
}
