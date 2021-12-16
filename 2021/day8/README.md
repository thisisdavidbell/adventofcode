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

Move all strings to []bytes:
### Benchmark results:
```
BenchmarkPart2All-16                        1360            852809 ns/op
BenchmarkPart2-16                           1928            553152 ns/op
```

### Profiling
```
      20ms      1.37s (flat, cum)  5.43% of Total
         .          .     70:   return
         .          .     71:}
         .          .     72:
         .          .     73:// todo: update to remove string cassts - use byte everywhere we can... and move strings.* to bytes.*
         .          .     74:func part2(displays []display) (count int) {
         .       20ms     75:   for _, d := range displays {
         .      100ms     76:           correctDigitLetters := make(map[int][]byte, 10)
         .          .     77:           //find top:
         .       50ms     78:           top := removeChars(d.lenThree, (d.lenTwo))
         .      110ms     79:           applyLettersByteSlice(correctDigitLetters, top[0], []int{0, 2, 3, 5, 6, 7, 8, 9})
         .          .     80:
         .       40ms     81:           otherCharsinFour := removeChars(d.lenFour, d.lenTwo)
         .          .     82:
         .          .     83:           // find three
         .          .     84:           var three []byte
         .          .     85:           for _, f := range d.lenFive {
         .          .     86:                   if bytes.Contains(f, []byte{d.lenTwo[0]}) && bytes.Contains(f, []byte{d.lenTwo[1]}) {
         .          .     87:                           three = f
         .          .     88:                           break
         .          .     89:                   }
         .          .     90:           }
         .          .     91:           // find middle
         .          .     92:           var middle []byte
         .       10ms     93:           if bytes.Contains(three, []byte{otherCharsinFour[0]}) {
         .          .     94:                   middle = []byte{otherCharsinFour[0]}
         .          .     95:           } else {
         .          .     96:                   middle = []byte{otherCharsinFour[1]}
         .          .     97:           }
         .      110ms     98:           applyLettersByteSlice(correctDigitLetters, middle[0], []int{2, 3, 4, 5, 6, 8, 9})
         .          .     99:
         .          .    100:           // top-left must be remaining char in 4:
         .       10ms    101:           topleft := removeChars(otherCharsinFour, middle)
         .       20ms    102:           applyLettersByteSlice(correctDigitLetters, topleft[0], []int{0, 4, 5, 6, 8, 9})
         .          .    103:
         .          .    104:           // find char left after removing 7 and middle from three, must be bottom.
      10ms      100ms    105:           bottom := removeChars(three, d.lenThree)
         .       10ms    106:           bottom = removeChars(bottom, middle)
         .          .    107:
         .       30ms    108:           applyLettersByteSlice(correctDigitLetters, bottom[0], []int{0, 2, 3, 5, 6, 8, 9})
         .          .    109:
         .          .    110:           // bottom left must be 8 less 3, less topleft
         .       30ms    111:           bottomleft := removeChars(d.lenSeven, three)
         .       20ms    112:           bottomleft = removeChars(bottomleft, topleft)
         .       10ms    113:           applyLettersByteSlice(correctDigitLetters, bottomleft[0], []int{0, 2, 6, 8})
         .          .    114:
         .          .    115:           //find top right and bottom right -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
         .          .    116:           var topright []byte
      10ms       10ms    117:           for _, f := range d.lenSix {
         .       10ms    118:                   if !(bytes.Contains(f, []byte{d.lenTwo[0]}) && bytes.Contains(f, []byte{d.lenTwo[1]})) {
         .          .    119:                           if bytes.Contains(f, []byte{d.lenTwo[0]}) {
         .          .    120:                                   topright = []byte{d.lenTwo[1]}
         .          .    121:                           } else {
         .          .    122:                                   topright = []byte{d.lenTwo[0]}
         .          .    123:                           }
         .          .    124:                           break
         .          .    125:                   }
         .          .    126:           }
         .      100ms    127:           applyLettersByteSlice(correctDigitLetters, topright[0], []int{0, 1, 2, 3, 4, 7, 8, 9})
         .          .    128:
         .          .    129:           //bottom right remaining segment of 1
         .       40ms    130:           bottomright := removeChars(d.lenTwo, topright)
         .      140ms    131:           applyLettersByteSlice(correctDigitLetters, bottomright[0], []int{0, 1, 3, 4, 5, 6, 7, 8, 9})
         .          .    132:
         .          .    133:           // now match the examples
         .      400ms    134:           count += findActualDigits(correctDigitLetters, d.example)
         .          .    135:   }
         .          .    136:
         .          .    137:   return
         .          .    138:}
```


# Older analysis:

Removed string concat for decent gain, after proving it was as bad as I suspected.
(Acknowledged that more gain to be had by pushing whole letters to correctDigitLetters, but little to learn there.)

### Benchmark results:
```
BenchmarkPart2All-16                        1365            844,213 ns/op
BenchmarkPart2-16                           1786            587,570 ns/op
BenchmarkApplyLettersString-16            125299              78459 ns/op
BenchmarkApplyLettersByteSlice-16        8846604               131.0 ns/op
BenchmarkApplyLettersByteRune-16         7789998               179.0 ns/op
```

### Older Benchmarks
Only minor improvements from tweaks:
```
BenchmarkPart2All-16                 969           1178825 ns/op
BenchmarkPart2-16                   1298            920174 ns/op
```

### Profiling
list part2
```
     10ms      1.38s (flat, cum) 54.76% of Total
         .          .     69:   return
         .          .     70:}
         .          .     71:
         .          .     72:func part2(displays []display) (count int) {
         .          .     73:   for _, d := range displays {
         .       40ms     74:           correctDigitLetters := make(map[int]string, 10)
         .          .     75:           //find top:
         .       30ms     76:           top := removeChars(d.lenThree, (d.lenTwo))
         .       90ms     77:           applyLetters(correctDigitLetters, top, []int{0, 2, 3, 5, 6, 7, 8, 9})
         .          .     78:
         .       10ms     79:           otherCharsinFour := removeChars(d.lenFour, d.lenTwo)
         .          .     80:
         .          .     81:           // find three
         .          .     82:           three := ""
      10ms       10ms     83:           for _, f := range d.lenFive {
         .       10ms     84:                   if strings.Contains(f, string(d.lenTwo[0])) && strings.Contains(f, string(d.lenTwo[1])) {
         .          .     85:                           three = f
         .          .     86:                           break
         .          .     87:                   }
         .          .     88:           }
         .          .     89:           // find middle
         .          .     90:           middle := ""
         .          .     91:           if strings.Contains(three, string(otherCharsinFour[0])) {
         .          .     92:                   middle = string(otherCharsinFour[0])
         .          .     93:           } else {
         .          .     94:                   middle = string(otherCharsinFour[1])
         .          .     95:           }
         .      180ms     96:           applyLetters(correctDigitLetters, middle, []int{2, 3, 4, 5, 6, 8, 9})
         .          .     97:
         .          .     98:           // top-left must be remaining char in 4:
         .          .     99:           topleft := removeChars(otherCharsinFour, middle)
         .       60ms    100:           applyLetters(correctDigitLetters, topleft, []int{0, 4, 5, 6, 8, 9})
         .          .    101:
         .          .    102:           // find char left after removing 7 and middle from three, must be bottom.
         .       20ms    103:           bottom := removeChars(three, d.lenThree)
         .       20ms    104:           bottom = removeChars(bottom, middle)
         .          .    105:
         .      200ms    106:           applyLetters(correctDigitLetters, bottom, []int{0, 2, 3, 5, 6, 8, 9})
         .          .    107:
         .          .    108:           // bottom left must be 8 less 3, less topleft
         .       30ms    109:           bottomleft := removeChars(d.lenSeven, three)
         .       30ms    110:           bottomleft = removeChars(bottomleft, topleft)
         .       40ms    111:           applyLetters(correctDigitLetters, bottomleft, []int{0, 2, 6, 8})
         .          .    112:
         .          .    113:           //find top right and bottom right -only 1 of sixes is missing eitehr segment of 1, so find which it is, and thats top right
         .          .    114:           topright := ""
         .          .    115:           for _, f := range d.lenSix {
         .       10ms    116:                   if !(strings.Contains(f, string(d.lenTwo[0])) && strings.Contains(f, string(d.lenTwo[1]))) {
         .       10ms    117:                           if strings.Contains(f, string(d.lenTwo[0])) {
         .          .    118:                                   topright = string(d.lenTwo[1])
         .          .    119:                           } else {
         .       10ms    120:                                   topright = string(d.lenTwo[0])
         .          .    121:                           }
         .          .    122:                           break
         .          .    123:                   }
         .          .    124:           }
         .      150ms    125:           applyLetters(correctDigitLetters, topright, []int{0, 1, 2, 3, 4, 7, 8, 9})
         .          .    126:
         .          .    127:           //bottom right remaining segment of 1
         .          .    128:           bottomright := removeChars(d.lenTwo, topright)
         .      240ms    129:           applyLetters(correctDigitLetters, bottomright, []int{0, 1, 3, 4, 5, 6, 7, 8, 9})
         .          .    130:
         .          .    131:           // now match the examples
         .      190ms    132:           count += findActualDigits(correctDigitLetters, d.example)
         .          .    133:   }
         .          .    134:
         .          .    135:   return
         .          .    136:}
```


# Older Benchmark
Quick and very dirty results - using strings, and finding all segments: from commit: 31f35b5c91ae25330779ba49d7ed97c9c365b56b
```
BenchmarkPart2All-16                 834           1,330,706 ns/op
BenchmarkPart2-16                   1171            957,445 ns/op
```

