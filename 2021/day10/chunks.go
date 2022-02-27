package main

import (
	"errors"
	"fmt"
	"strings"
)

//import (
//	"github.com/thisisdavidbell/adventofcode/utils"
//)

func main() {
	fmt.Printf("Test Part 1: %v\n", part1All("test-input.txt"))
	//	fmt.Printf("Real Part 1: %v\n\n", part1All("test-input.txt"))

	//	fmt.Printf("Test Part 2: %v\n", part2All("test-input.txt"))
	//	fmt.Printf("Real Part 2: %v\n", part2All("test-input.txt"))
}

/*
Plan
  - create 8 iota entries for 4 x open/close pairs for readability
  - FILO queue
   - add if opening
   - remove if closing, checking its a matching pair
   - corrupted if not matching pair
   - corrupt if queue empty but more input - add score for incorrect item we found (NOT the expected item)
   - incomplete if get to end and queue not empty - ignore incomplete
*/

type lifo struct {
	stack []rune
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
		return popped, nil
	} else {
		return ' ', errors.New("cannot pop empty lifo")
	}

}

func (l *lifo) Push(r rune) {
	l.stack = append(l.stack, r)
}

func part1All(filename string) (answer int) {
	stack := NewLifo()
	stack.Push('a')
	stack.Push('b')
	stack.Push('c')
	aRune, err := stack.Pop()
	if err != nil {
		fmt.Printf("ERROR: pop returned error: %v", err)
	}
	fmt.Printf("aRune: %c\n", aRune)
	fmt.Printf("stack: %v\n", stack)

	return answer
}

/*
func part1(?) (answer int) {

	return answer
}
*/

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
