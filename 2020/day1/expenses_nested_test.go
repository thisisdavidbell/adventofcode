package main

import (
	"reflect"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestImportFileToSlice(t *testing.T) {
	var expectedSlice = []int{1721, 979, 366, 299, 675, 1456}
	gotSlice, err := ImportFileToSlice("test-input.txt")

	if err != nil {
		t.Fatalf("Test failed with error: %v", err)
	}

	if !reflect.DeepEqual(expectedSlice, gotSlice) {
		t.Fatalf("Slices are not equal.\nExpected: %v, Got: %v", expectedSlice, gotSlice)
	}

}

func TestFindAnswerUsingNestedLoops(t *testing.T) {
	var numberSlice = []int{1721, 979, 366, 299, 675, 1456}

	expectedProd := 1721 * 299

	gotProd, err := FindAnswerUsingNestedLoops(numberSlice)

	if err != nil {
		t.Fatalf("Test failed, error returned from Find2ValuesUsingNestedLoops. Error: %v", err)
	}

	if !(expectedProd == gotProd) {
		t.Fatalf("Find2ValuesUsingNestedLoops didn't find correct values. Exp: %v, Got: %v", expectedProd, gotProd)
	}
}

func TestFindAnswerUsingNestedLoops_NoMatch(t *testing.T) {
	var numberSlice = []int{1721, 979, 366, 675, 1456}
	expectedErrorString := "Not valid values found"

	_, err := FindAnswerUsingNestedLoops(numberSlice)

	if err == nil {
		t.Fatal("Find2ValuesUsingNestedLoops did not return error, when no matches.\n")
	} else {
		if err.Error() != expectedErrorString {
			t.Fatalf("Find2ValuesUsingNestedLoops error was not the expected error when no matching values in slice. Expected Error: %v. Error: %v", expectedErrorString, err)
		}
	}
}

func TestSolveItWithNestedLoops(t *testing.T) {
	expectedAnswer := 1721 * 299
	answer, err := SolveItWithNestedLoops("test-input.txt")
	if err != nil {
		t.Fatalf("SolveItWithNestedLoops failed with error: %v", err)
	}
	if answer != expectedAnswer {
		t.Fatalf("SolveItWithNestedLoops didnt returned correct answer. Expected: %v, got: %v", expectedAnswer, answer)
	}
}
