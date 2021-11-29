To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Initial Benchmark results:

Part 2 "quick and dirty":
With bool 2d slice (at commit: af3e13124aadcd02c84165264b0833b8a3ccce66):
```
BenchmarkPart2EndToEnv-16    	    6090	    187499 ns/op
```

Converted to string slice:
```
BenchmarkPart2EndToEnv-16    	   15817	     75542 ns/op
```

