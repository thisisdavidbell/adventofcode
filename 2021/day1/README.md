To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Initial Benchmark results:

Part 2 "quick and dirty" - processing twice (three times with read from file) :
```
BenchmarkPart2countSumIncrease-16    	    9385	    136144 ns/op
```