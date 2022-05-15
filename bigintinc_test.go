package biigist

import (
	"math"
	"math/big"
	"testing"
)

func TestBigIntInc(t *testing.T) {
	// validate zero handling and size increase from 0
	a := big.NewInt(0)
	b := big.NewInt(0)
	for i := 0; i < 4094; i++ {
		a.Add(a, one)
		bigIntInc(b)
		if b.Cmp(a) != 0 {
			t.Fatalf("a != b; %v != %v\n", a, b)
		}
	}

	// validate negative values handling
	a = big.NewInt(-8)
	b = big.NewInt(-8)
	for i := 0; i < 16; i++ {
		a.Add(a, one)
		bigIntInc(b)
		if b.Cmp(a) != 0 {
			t.Fatalf("a != b; %v != %v\n", a, b)
		}
	}

	// validate around MaxUint values that should increase required bit length
	a = big.NewInt(0).SetUint64(math.MaxUint - 2)
	b = big.NewInt(0).SetUint64(math.MaxUint - 2)
	for i := 0; i < 6; i++ {
		a.Add(a, one)
		bigIntInc(b)
		if b.Cmp(a) != 0 {
			t.Fatalf("a != b; %v != %v\n", a, b)
		}
	}

	// validate big values
	a = big.NewInt(1)
	b = big.NewInt(1)
	for i := 0; i < 23; i++ {
		a.Add(a, one)
		bigIntInc(b)
		if b.Cmp(a) != 0 {
			t.Fatalf("a != b; %v != %v\n", a, b)
		}
		a.Exp(a, two, nil)
		b.Set(a)
	}
}
