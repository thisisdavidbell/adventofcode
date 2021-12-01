To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Benchmark results:

Part 2 "quick and dirty" - processing twice (three times with read from file) :
```
BenchmarkImportFileToIntSlice-16          9032      118193 ns/op

BenchmarkPart2countSumIncreaseOnly-16   138136	      8879 ns/op   <== quick and dirty twice processing from commit ddfebad

BenchmarkPart2Only-16               	 1000000	      1006 ns/op

```
