To test:
- `go test`
- `go test -v`

To run:
- `go run .`

To get benchmarks:
- `go test -run=XXX -bench .`


Benchmark results:

```
BenchmarkPart1Bytes-16        	   24948	     48536 ns/op
BenchmarkPart1BytesAll-16     	   10999	    110089 ns/op

BenchmarkPart1String-16       	   24666	     49134 ns/op
BenchmarkPart1StringAll-16    	   10795	    110767 ns/op


BenchmarkPart2-16             	    2079	    553291 ns/op
BenchmarkPart2All-16          	    1552	    744512 ns/op
```
