package main

import (
	"reflect"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestImportFileToMap(t *testing.T) {
	expectedMap := map[int]struct{}{
		1721: struct{}{},
		979:  struct{}{},
		366:  struct{}{},
		299:  struct{}{},
		675:  struct{}{},
		1456: struct{}{},
	}
	gotMap, err := ImportFileToMap("test-input.txt")
	if err != nil {
		t.Fatalf("Test failed with error: %v", err)
	}

	if !reflect.DeepEqual(expectedMap, gotMap) {
		t.Fatalf("Slices are not equal.\nExpected: %v, Got: %v", expectedMap, gotMap)
	}

}

func TestFindAnswerUsingMAp(t *testing.T) {
	numberMap := map[int]struct{}{
		1721: struct{}{},
		979:  struct{}{},
		366:  struct{}{},
		299:  struct{}{},
		675:  struct{}{},
		1456: struct{}{},
	}

	expectedProd := 1721 * 299
	expectedProd2 := 979 * 366 * 675

	gotProd, gotProd2, err := FindAnswerUsingMap(numberMap)
	if err != nil {
		t.Fatalf("Test failed, error returned from Find2ValuesUsingNestedLoops. Error: %v", err)
	}

	if !(expectedProd == gotProd) {
		t.Fatalf("Find2ValuesUsingMap didn't find correct values part1. Exp: %v, Got: %v", expectedProd, gotProd)
	}
	if !(expectedProd2 == gotProd2) {
		t.Fatalf("Find2ValuesUsingMap didn't find correct values part2. Exp: %v, Got: %v", expectedProd2, gotProd2)
	}
}

/*
func skipTestFindAnswerUsingMap_NoMatch(t *testing.T) {
	var numberSlice = []int{1721, 979, 366, 675, 1456}
	expectedErrorString := "Not valid values found"

	_, err := FindAnswerUsingMap(numberSlice)

	if err == nil {
		t.Fatal("Find2ValuesUsingNestedLoops did not return error, when no matches.\n")
	} else {
		if err.Error() != expectedErrorString {
			t.Fatalf("Find2ValuesUsingNestedLoops error was not the expected error when no matching values in slice. Expected Error: %v. Error: %v", expectedErrorString, err)
		}
	}
}
*/
// TestSolveItWithMap - perform solution using map
func TestSolveItWithMap(t *testing.T) {
	expectedAnswer := 1721 * 299
	expectedAnswer2 := 979 * 366 * 675
	answer, answer2, err := SolveItWithMap("test-input.txt")
	if err != nil {
		t.Fatalf("SolveItWithNestedLoops failed with error: %v", err)
	}
	if answer != expectedAnswer {
		t.Fatalf("SolveItWithNestedLoops didnt returned correct answer1. Expected: %v, got: %v", expectedAnswer, answer)
	}
	if answer2 != expectedAnswer2 {
		t.Fatalf("SolveItWithNestedLoops didnt returned correct answer2. Expected: %v, got: %v", expectedAnswer2, answer2)
	}
}

func BenchmarkImportFileToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImportFileToMap("real-input.txt")
	}
}

func BenchmarkSolveItWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveItWithMap("real-input.txt")
	}
}
