package main

import (
	"testing"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func BenchmarkImportFileToByteSliceSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		utils.ImportFileToByteSliceSlice("real-input-txt")
	}
}

func BenchmarkImportFileToStringSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		utils.ImportFileToStringSlice("real-input-txt")
	}
}

func BenchmarkReadFileToByteSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		utils.ReadFileToByteSlice("real-input-txt")
	}
}

func BenchmarkReadFileToByteByteSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		utils.ReadFileToByteSliceSlice("real-input-txt")
	}
}
