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
BenchmarkPart2All-16                 472           2,421,530 ns/op
BenchmarkPart2-16                    510           2,315,802 ns/op
```

### Perf Thoughts
- was never going to perform loop
- found maths function for nth triangular number

```

### Profiling
```
         .          .     55:func part2(crabs []int, min int, max int) (bestSoFar int) {
         .          .     56:
         .       20ms     57:   for pos := min; pos <= max; pos++ {
         .          .     58:           fuelUsed := 0
     440ms      440ms     59:           for _, crab := range crabs {
     380ms      610ms     60:                   numMoves := utils.IntAbs(pos - crab)
     170ms      810ms     61:                   fuelUsed += nthTriangularNumber(numMoves)
     150ms      150ms     62:                   if fuelUsed > bestSoFar && pos != min {
         .          .     63:                           continue
         .          .     64:                   }
         .          .     65:           }
         .          .     66:           if pos == min || fuelUsed < bestSoFar {
      60ms       60ms     67:                   bestSoFar = fuelUsed
         .          .     68:           }
         .          .     69:   }
         .          .     70:   return
         .          .     71:}
```