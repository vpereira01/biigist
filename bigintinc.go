package biigist

import (
	"math"
	"math/big"
)

var (
	one = big.NewInt(1)
	two = big.NewInt(2)
)

// bigIntInc speeds up z.Add(z, one) operations when z > 0
//
// When z.Sign() == 1 && z.Bits()[0] < math.MaxUint an increment is performed without
// calling z.Add(z, one), otherwise z.Add(z, one) is called.
//
// Incrementing/adding one with z.Add(z, y) is slower because Int.Add(x, y) and nat.add(), called by Int.Add(x, y),
// are implemented and optimized to add similar length values and store the result in another value. When just
// adding a small value to a bigger value an overhead is suffered.
//
// Performance improvement on average of around 8x, as in, time(bigIntInc(z)) = time(z.Add(z, one)) / 8
func bigIntInc(z *big.Int) {
	zBits := z.Bits()
	// big.NewInt( 9).Sign() == 1
	// For speed and simplicity this only tries to increment the least significant digits.
	// Using a base 10 example: 1000+1 only affects the least significant digit.
	// Only when the least significant digit is 9, 1009+1, then more significant digits are affected.
	if z.Sign() > 0 && zBits[0] < math.MaxUint {
		// overflow does not happen due to zBits[0] < math.MaxUint check
		// uint(zBits[0])++ does not work
		zBits[0] = big.Word(uint(zBits[0]) + 1)
	} else {
		z.Add(z, one)
	}
}
