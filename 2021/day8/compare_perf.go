package main

func applyLettersString(correctDigitLetters map[int]string, ch string, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = correctDigitLetters[s] + ch
		//correctDigitLetters[s] = strings.Join([]string{correctDigitLetters[s], ch}, "")
	}
}

func applyLettersByteSlice(correctDigitLetters map[int][]byte, ch byte, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = append(correctDigitLetters[s], ch)
	}
}

func applyLettersRuneSlice(correctDigitLetters map[int][]rune, ch rune, segments []int) {
	for _, s := range segments {
		correctDigitLetters[s] = append(correctDigitLetters[s], ch)
	}
}
