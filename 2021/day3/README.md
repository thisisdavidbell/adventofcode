To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Benchmark results:
(method all available at commit: c7c8d66b4e41d2925689c952f2ee8c21ef8e165f)

```
BenchmarkPart1Bytes-16            	   23989	     48132 ns/op
BenchmarkPart1BytesAll-16         	   10484	    113937 ns/op

BenchmarkPart1String-16           	   24606	     47878 ns/op - NOW CALLED part1
BenchmarkPart1StringAll-16        	    8949	    115268 ns/op - NOW CALLED part1


BenchmarkPart2OnePass-16          	    2196	    542226 ns/op
BenchmarkPart2Loops-16            	    2743	    413442 ns/op
BenchmarkPart2Delete-16           	  452187	      2825 ns/op (looks like a benchmark bug)
BenchmarkPart2Slice-16            	   32884	     36105 ns/op
BenchmarkPart2SliceReuse-16       	   37765	     31998 ns/op - NOW CALLED part2

BenchmarkPart2OnePassAll-16       	    1593	    783321 ns/op
BenchmarkPart2LoopsAll-16         	    1915	    619785 ns/op
BenchmarkPart2DeleteAll-16        	    2065	    582093 ns/op
BenchmarkPart2SliceAll-16         	   10000	    102572 ns/op
BenchmarkPart2SliceReuseAll-16    	   13892	     85806 ns/op - NOW CALLED part2
```

Conclusions:
- only use Maps if there is a big benefit (creation and copy expensive, would need lots of access)
- pre-allocating slice length saves a lot of time (66660 ns/op -> 36105 ns/op)
- extra looping cheaper than creating extra slices/maps
- reading from input file very expensive (extra file read adds 40000ns/op over a slice copy)
