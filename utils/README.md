README.md

To get benchmarks:
- `go test -run=XXX -bench .`


Benchmark results:

```
BenchmarkImportFileToByteSliceSlice-16    	  218572	      4756 ns/op  <--- scanner.Bytes()
BenchmarkImportFileToStringSlice-16       	  258112	      4662 ns/op  <--- scanner.Text()

BenchmarkReadFileToByteSlice-16           	  340419	      3733 ns/op  <--- os.ReadFile()
BenchmarkReadFileToByteSliceSlice-16      	  320428	      3970 ns/op  <--- os.ReadFile() bytes.Split()

BenchmarkReadFileToString-16              	  311913	      4281 ns/op  <--- string(os.ReadFile())
BenchmarkReadFileToStringSlice-16         	  288110	      4041 ns/op  <--- string(os.ReadFile()) bytes.Split()
```
