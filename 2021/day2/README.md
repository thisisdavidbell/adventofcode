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
BenchmarkReadCommandToSlice-16    	    1368	    860013 ns/op
BenchmarkPart2Only-16             	  775762	      1480 ns/op

BenchmarkAll2-16                  	    1398	    858110 ns/op
```

go 1.13.10
```
BenchmarkPart2Only-16    	  640623	      1692 ns/op
```
