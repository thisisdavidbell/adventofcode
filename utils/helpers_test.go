package utils

import (
	"testing"
)

func BenchmarkImportFileToByteSliceSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ImportFileToByteSliceSlice("real-input-txt")
	}
}

func BenchmarkImportFileToStringSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ImportFileToStringSlice("real-input-txt")
	}
}

func BenchmarkReadFileToByteSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ReadFileToByteSlice("real-input-txt")
	}
}

func BenchmarkReadFileToByteSliceSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ReadFileToByteSliceSlice("real-input-txt")
	}
}

func BenchmarkReadFileToString(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ReadFileToString("real-input-txt")
	}
}

func BenchmarkReadFileToStringSlice(b *testing.B) {
	// input
	for i := 0; i < b.N; i++ {
		ReadFileToStringSlice("real-input-txt")
	}
}
