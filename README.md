# Faster increment of big.Int bench tests

## How to run

1. `go generate` to generate gnark finite field arithmetic package
2. `go test . -run=nothing -bench=. -benchtime=5s`
