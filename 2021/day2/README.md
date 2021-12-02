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
go 1.16.10
```
BenchmarkReadCommandToSlice-16    	    1386	    808222 ns/op

BenchmarkPart2Only-16             	  785119	      1406 ns/op   <=== first pass
BenchmarkPart2LetterOnly-16       	  792643	      1480 ns/op   <=== using single read and process function

BenchmarkAll2-16                  	    1506	    784907 ns/op   <=== first pass
BenchmarkPerfPart2-16             	    1420	    770575 ns/op   <=== using single read and process function
BenchmarkPerfPart2Cast-16         	    1593	    744403 ns/op   <=== use atrconv.Atoi for cast not Sscanf
```

go 1.13.10
```
BenchmarkPart2Only-16    	  640623	      1692 ns/op
```
