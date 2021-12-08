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


Benchmark results:

```
BenchmarkAllPart2-16    	     100	  10038436 ns/op
BenchmarkPart2-16       	     123	   9615733 ns/op

```

Quick perf test to prove map quicker than for loop and slice. From commit: 30e7b98d6382bac2a1e369baa35a5f7ae00bbe77
```
BenchmarkAllPart2-16    	      28	  37767681 ns/op
BenchmarkPart2-16       	      28	  37486417 ns/op
```
