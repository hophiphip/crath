package gcd

import (
	"math/big"
)

// Gcd calculates Greatest common divider
func Gcd(a *big.Int, b *big.Int) *big.Int {
	x, y, zero := big.NewInt(0).Abs(a), big.NewInt(0).Abs(b), big.NewInt(0)

	if x.Cmp(zero) == 0 && y.Cmp(zero) == 0 {
		return big.NewInt(-1)
	}

	sum := big.NewInt(0)

	if x.Cmp(zero) == 0 || y.Cmp(zero) == 0 {
		return sum.Add(x, y)
	}

	for x.Cmp(zero) != 0 && y.Cmp(zero) != 0 {
		if x.Cmp(y) > 0 {
			x.Mod(x, y)
		} else {
			y.Mod(y, x)
		}
	}

	return sum.Add(x, y)
}

// Bgcd calculates Greatest common divider using binary algorithm
func Bgcd(a *big.Int, b *big.Int) *big.Int {
	var shift uint
	x, y, zero, one := big.NewInt(0).Abs(a), big.NewInt(0).Abs(b), big.NewInt(0), big.NewInt(1)

	if x.Cmp(zero) == 0 && y.Cmp(zero) == 0 {
		return big.NewInt(-1)
	}

	buf := big.NewInt(0)

	if x.Cmp(zero) == 0 || y.Cmp(zero) == 0 {
		return buf.Add(x, y)
	}

	for shift = 0; buf.Or(x, y).And(buf, one).Cmp(zero) == 0; shift++ {
		x.Rsh(x, 1)
		y.Rsh(y, 1)
	}

	for buf.And(x, one).Cmp(zero) == 0 {
		x.Rsh(x, 1)
	}

	for {
		for buf.And(y, one).Cmp(zero) == 0 {
			y.Rsh(y, 1)
		}

		if x.Cmp(y) > 0 {
			buf.Add(x, y)
			y.Set(x)
			x.Sub(buf, x)
		}

		y.Sub(y, x)

		if y.Cmp(zero) == 0 {
			break
		}
	}

	return buf.Add(x, y).Lsh(buf, shift)
}
