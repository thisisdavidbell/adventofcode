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
```
BenchmarkPart2All-16    	     318	   3491361 ns/op
BenchmarkPart2-16       	     507	   2186348 ns/op
```

### Profile Data
top:
```
820ms 21.69% 21.69%      820ms 21.69%  runtime.kevent
460ms 12.17% 33.86%      490ms 12.96%  github.com/thisisdavidbell/adventofcode/2021/day5.applyCoords
410ms 10.85% 44.71%      410ms 10.85%  runtime.pthread_cond_wait
```

list part2:
```
.          .     33:func part2(lines []line, maxX int, maxY int) int {
.          .     34:	count := 0
.      470ms     35:	grid := createGrid(maxX, maxY)
.          .     36:	for _, aLine := range lines {
.      490ms     37:		count = applyCoords(aLine, grid, count)
.          .     38:	}
.          .     39:	return count //countNumIntersects(grid)
```

readInputs line:
```
.      530ms     29:	lines, maxX, maxY := readInputs(filename)
```

# First Analysis
Orig 1st pass benchmark, including looking up result across whole grid:

### Benchmark Results:

```
BenchmarkPart2All-16    	     240	   5049924 ns/op
BenchmarkPart2-16       	     298	   4212361 ns/op

```

### Profile Data:
top
```
Showing top 10 nodes out of 95
      flat  flat%   sum%        cum   cum%
     650ms 32.02% 32.02%      650ms 32.02%  github.com/thisisdavidbell/adventofcode/2021/day5.countNumIntersects
     190ms  9.36% 41.38%      190ms  9.36%  runtime.kevent
     190ms  9.36% 50.74%      190ms  9.36%  runtime.pthread_cond_wait
     180ms  8.87% 59.61%      180ms  8.87%  github.com/thisisdavidbell/adventofcode/2021/day5.applyCoords
     150ms  7.39% 67.00%      150ms  7.39%  syscall.syscall
      70ms  3.45% 70.44%       70ms  3.45%  runtime.procyield
      70ms  3.45% 73.89%       70ms  3.45%  runtime.pthread_kill
      50ms  2.46% 76.35%      200ms  9.85%  os.ReadFile
      30ms  1.48% 77.83%       70ms  3.45%  fmt.(*ss).ReadRune
      30ms  1.48% 79.31%       30ms  1.48%  runtime.memclrNoHeapPointers
```
list part2:
```
.          .     28:func part2(filename string) (count int) {
.      360ms     29:	lines, maxX, maxY := readInputs(filename)
.      170ms     30:	grid := createGrid(maxX, maxY)
.          .     31:	for _, aLine := range lines {
.      180ms     32:		applyCoords(aLine, grid)
.          .     33:	}
10ms   660ms     34:	return countNumIntersects(grid)
.          .     35:}
.          .     36:
```
Perf thoughts:
- creation of grid fairly expensive single action
- applying lines we do 500 times, but each action is not so bad, as you only work with few points you need
- checking result is expensive as it traverses whole grid 900x900 grid
Plan:
- count each intersection when we find it, so dont need to do countNumIntersects
  - count each time a point becomes 2