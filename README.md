# Faster increment of big.Int bench tests

## How to run

1. `go generate` to generate gnark finite field arithmetic package
2. `go test . -run=nothing -bench=. -benchtime=5s`

## My results

```
user@host$ go test . -run=nothing -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: example.com/biigist
cpu: Intel(R) Core(TM) i5-4590T CPU @ 2.00GHz
BenchmarkBigIntInc/512bits_NewInt-4             94064070                57.74 ns/op
BenchmarkBigIntInc/512bits_NewInt_Add-4         64224325                92.77 ns/op
BenchmarkBigIntInc/512bits_NewInt_Inc-4         66054766                92.00 ns/op
BenchmarkBigIntInc/513bits_BIAdd-4              246136504               24.44 ns/op
BenchmarkBigIntInc/513bits_GFAdd-4              447096817               11.57 ns/op
BenchmarkBigIntInc/513bits_Inc-4                1000000000               3.913 ns/op
BenchmarkBigIntInc/1025bits_BIAdd-4             208855520               29.20 ns/op
BenchmarkBigIntInc/1025bits_GFAdd-4             514192068               11.43 ns/op
BenchmarkBigIntInc/1025bits_Inc-4               1000000000               3.905 ns/op
PASS
ok      example.com/biigist     57.451s
```

## Context 

https://vpereira01.com/posts/go-bigint-perf/