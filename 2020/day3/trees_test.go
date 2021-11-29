package main

import (
	"reflect"
	"testing"
)

func TestImportFileTo2DSlice(t *testing.T) {
	/*  ..##.......
	    #...#...#..
	    .#....#..#.
	    ..#.#...#.#
	    .#...##..#.
	    ..#.##.....
	    .#.#.#....#
	    .#........#
	    #.##...#...
	    #...##....#
	    .#..#...#.#
	*/
	expected2DSlice := [][]bool{
		{false, false, true, true, false, false, false, false, false, false, false},
		{true, false, false, false, true, false, false, false, true, false, false},
	}

	got2DSlice, _ := importFileTo2DSlice("short-test-input.txt")
	if !reflect.DeepEqual(expected2DSlice, got2DSlice) {
		t.Fatalf("importFileTo2DSlice didnt return correct 2d slice. Expected: %v, got: %v", expected2DSlice, got2DSlice)
	}
}

func TestConvertLineToSlice(t *testing.T) {
	fileLine := "#...#...#.."
	expectedSlice := []bool{true, false, false, false, true, false, false, false, true, false, false}
	gotSlice := convertLineToSliceOfTrees(fileLine)

	if !reflect.DeepEqual(expectedSlice, gotSlice) {
		t.Fatalf("converLineToSliceOfTrees didnt return correct slice. Expected: %v, got: %v", expectedSlice, gotSlice)
	}
}
