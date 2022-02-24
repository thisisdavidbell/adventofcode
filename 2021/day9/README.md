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