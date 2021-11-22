To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Initial Benchmark results:
```
BenchmarkImportFileToSlice-16         	   30366	     39738 ns/op
BenchmarkSolveItWithNestedLoops-16    	   22587	     47424 ns/op

BenchmarkImportFileToMap-16           	   23570	     52513 ns/op
BenchmarkSolveItWithMap-16            	   20175	     55447 ns/op
```
Conclusions:
- processing Map way quicker than nested loop
- BUT creation of Map way more expensive than creation of slice. When one time use - slice+nested loop quicker!
- Creating the map at a specified size is much faster, but no efficient way to count the number of lines ahead of time.
