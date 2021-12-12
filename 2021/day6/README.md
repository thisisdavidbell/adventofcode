# To use

### To test:
- `go test`
- `go test -v`

### To run:
- `go run .`

### To get benchmarks:
- `go test -run=XXX -bench .`

### Profiling:
go pprof link: https://pkg.go.dev/runtime/pprof#hdr-Profiling_a_Go_program

- `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
- `go tool pprof cpu.prof`
  - `top`
  - `list part2`


# Latest Analysis

### Benchmark results:
Successful solve:
```
BenchmarkSolveAll-16    	   28670	     40285 ns/op
BenchmarkSolve-16       	 1200590	       992.5 ns/op
```

Note:
- I did confirm my suspicion that copying the slice using [1:] was slower, and it was. Reuse same slice much quicker.
```
BenchmarkSolve-16       	  664946	      1816 ns/op < using slice copy with append [1:]

```

### Profiling

```
.          .        40:func solve(theCount []int, numDays int) (count int) {
.          .        41:
.          .        42:	for d := 0; d < numDays; d++ {
40ms       40ms     43:		day0 := theCount[0]
530ms      530ms    44:		for i := 1; i < 9; i++ {
290ms      290ms    45:			theCount[i-1] = theCount[i]
.          .        46:		}
50ms       50ms     47:		theCount[6] = theCount[6] + day0
30ms       30ms     48:		theCount[8] = day0
.          .        49:	}
.          .        50:	for c := 0; c < 9; c++ {
.          .        51:		count = count + theCount[c]
.          .        52:	}
.          .        53:	return
```

### Orig Perf Thoughts
After initial Pass which failed to complete as expected:
- part2 test didn't even complete in reasonable time - slice would have to become the answer (26984457539) long/
- thoughts:
  - store child fish in orig fish object - but some will be 6 and some 8, so tricker.
  - instead of storing fish, store a count of how many fish are on each day from 0 to 8. And have a process for each day...
