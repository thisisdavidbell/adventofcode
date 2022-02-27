package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/thisisdavidbell/adventofcode/utils"
)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	fmt.Printf("Real Part 1: %v\n\n", part1All("real-input.txt"))

	//	fmt.Printf("Test Part 2: %v\n", part2All("test-input.txt"))
	//	fmt.Printf("Real Part 2: %v\n", part2All("test-input.txt"))
}

/*
Plan
- Implement basic FILO stack
- add if opening
- remove if closing, checking its a matching pair
- corrupted if not matching pair
- corrupt if stack empty but next item closing
- incomplete if get to end and stack not empty - ignore incomplete as instructed
  - for now ensure new line, then can look for \n and doc check. Could check for last char of input
*/

type lifo struct {
	stack []rune
	size  int
}

func NewLifo() *lifo {
	return &lifo{}
}

func (l *lifo) String() string {
	var sb strings.Builder
	for i, r := range l.stack {
		sb.WriteString(fmt.Sprintf("%v:%c, ", i, r))
	}
	sb.WriteString("\n")

	return sb.String()
}

func (l *lifo) Pop() (rune, error) {
	if len(l.stack) > 0 {
		n := len(l.stack) - 1
		popped := l.stack[n]
		l.stack = l.stack[:(n)]
		l.size--
		return popped, nil
	} else {
		return ' ', errors.New("cannot pop empty lifo")
	}

}

func (l *lifo) Push(r rune) {
	l.stack = append(l.stack, r)
	l.size++
}

func part1All(filename string) (answer int) {
	input := utils.ReadFileToByteSlice(filename)
	input = append(input, '\n') // ensure /n at end of input - tidy up later.
	answer = part1(input)
	return answer
}

func sumPoints(corruptChars []rune) (count int) {
	for _, r := range corruptChars {
		switch r {
		case ')':
			count += 3
		case ']':
			count += 57
		case '}':
			count += 1197
		case '>':
			count += 25137
		}
	}
	return count
}

func part1(input []byte) (answer int) {
	stack := NewLifo()
	//utils.PrintByteSlice(input)
	var corruptChars []rune
	runes := bytes.Runes(input)
	for _, r := range runes {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(rune(r))
			continue
		case '\n':
			if stack.size != 0 {
				// found incomplete line - no-op for now
			}
			// successfully finished line.
			stack = NewLifo()
			continue
		}
		// assuming no illegal chars for now, so must be a close
		open, err := stack.Pop()

		if err != nil && strings.Contains(err.Error(), "cannot pop empty lifo") {
			//corrupt - closing when can't have an open
			corruptChars = append(corruptChars, r)
			continue
		}

		switch r {
		case ')':
			if open != '(' {
				corruptChars = append(corruptChars, r)
				continue
			}
		case ']':
			if open != '[' {
				corruptChars = append(corruptChars, r)
				continue
			}
		case '}':
			if open != '{' {
				corruptChars = append(corruptChars, r)
				continue
			}
		case '>':
			if open != '<' {
				corruptChars = append(corruptChars, r)
				continue
			}
		}
	}
	//utils.PrintRuneSlice(corruptChars)
	answer = sumPoints(corruptChars)
	return answer
}

/*
func part2() () {

	return
}
*/

/*
func part2All() () {

	return
}
*/
