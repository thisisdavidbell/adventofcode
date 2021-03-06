To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`

Profiling:
go pprof link: https://pkg.go.dev/runtime/pprof#hdr-Profiling_a_Go_program

- `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
- `go tool pprof cpu.prof`
  - `top`
  - `list part2`


Perf observations:
- sheer number of permutations meant reducing number of scans of each card valuable - huge win to only check each number once
- removal of winning board did not save time:
  - deleting board from slice - slice processing more expensive than time saved
  - skipping boards that have already won - slightly worse - cost of operation to check board each time more than saving it seems.

Benchmark results:

```
BenchmarkReadBingoInputToSlices-16    	    8286	    142926 ns/op
BenchmarkAllPart2-16    	                   100	  10038436 ns/op
BenchmarkPart2-16       	                   123	   9615733 ns/op <- scanning whole board for all numbers every time

BenchmarkReadBingoInputToMaps-16      	    3255	    374239 ns/op
BenchmarkAllPart2-16                  	     163	   7116813 ns/op  <- profiling bug? is 7x the sum of the 2 it ctually runs...
BenchmarkPart2-16                     	    2613	    448427 ns/op  <- removing matches numbers from map, so scanning remaining board for latest number only

BenchmarkPart2-16                     	    1755	    630866 ns/op <= deleting board when it wins <- overhead of slice manipulation clearly way more than gain
```

Profile:
```
170ms      1.51s (flat, cum) 78.24% of Total
    .          .     80:func checkBoardWins(numbers map[string]struct{}, board [][]string) bool {
    .          .     81:
    .          .     82:	matched := false
    .          .     83:
    .          .     84:	//check row
 40ms       40ms     85:	for _, r := range board {
    .          .     86:		matched = true
 20ms       20ms     87:		for _, c := range r {
 20ms      860ms     88:			if _, ok := numbers[c]; !ok {
    .          .     89:				matched = false
    .          .     90:			}
    .          .     91:		}
 30ms       30ms     92:		if matched {
    .          .     93:			return true
    .          .     94:		}
    .          .     95:	}
    .          .     96:
    .          .     97:	// check columns
    .          .     98:	matched = false
    .          .     99:	for c := 0; c < len(board[0]); c++ {
    .          .    100:		matched = true
 20ms       20ms    101:		for r := 0; r < len(board); r++ {
 40ms      540ms    102:			if _, ok := numbers[board[r][c]]; !ok {
    .          .    103:				matched = false
    .          .    104:			}
    .          .    105:		}
    .          .    106:		if matched {
    .          .    107:			return true
```
Perf plans:
 - skip first 4 numbers
 - board maintains state, so dont do all numbers every runtime
 
 - remove boards already won/

Quick perf test to prove map quicker than for loop and slice. From commit: 30e7b98d6382bac2a1e369baa35a5f7ae00bbe77
```
BenchmarkAllPart2-16    	      28	  37767681 ns/op
BenchmarkPart2-16       	      28	  37486417 ns/op
```

Quick test to see what skipping first 4 numbers with inefficient solution saved:
```
BenchmarkAllPart2-16    	     124	   9407986 ns/op
BenchmarkPart2-16       	     128	   9245565 ns/op
```
