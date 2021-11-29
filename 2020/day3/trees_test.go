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

func BenchmarkPart2EndToEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		realLineSlice, _ := importFileTo2DSlice("real-input.txt")
		_ = countTrees(realLineSlice, 1, 1) * countTrees(realLineSlice, 3, 1) * countTrees(realLineSlice, 5, 1) * countTrees(realLineSlice, 7, 1) * countTrees(realLineSlice, 1, 2)
	}
}
