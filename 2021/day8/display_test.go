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

func SkipBenchmarkApplyLettersString(b *testing.B) {
	correctDigitLetters := make(map[int]string, 10)
	ch := "a"
	segments := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		applyLettersString(correctDigitLetters, ch, segments)
	}
}

func SkipBenchmarkApplyLettersByteSlice(b *testing.B) {
	correctDigitLetters := make(map[int][]byte, 10)
	ch := byte('a')
	segments := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		applyLettersByteSlice(correctDigitLetters, ch, segments)
	}
}

func SkipBenchmarkApplyLettersByteRune(b *testing.B) {
	correctDigitLetters := make(map[int][]rune, 10)
	ch := 'a'
	segments := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		applyLettersRuneSlice(correctDigitLetters, ch, segments)
	}
}
