package diopheq

import (
	"math/big"

	"math.io/crath/gcd"
)

// All pre-checks are done by the caller
// TODO: Add pre-checks

// Simple solves equation: ax + by = 1,
// returns: (x , y) -- solutions
func Simple(a *big.Int, b *big.Int) (*big.Int, *big.Int) {
	var (
		box = [2][2]*big.Int{
			{big.NewInt(1), big.NewInt(0)},
			{big.NewInt(0), big.NewInt(1)},
		}
		x, y, quo, mod, zero, buf = big.NewInt(0).Set(a), big.NewInt(0).Set(b), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)
	)

	for quo.DivMod(x, y, mod); mod.Cmp(zero) != 0; quo.DivMod(x, y, mod) {
		buf.Set(box[0][1])
		box[0][1].Mul(quo, box[0][1]).Sub(box[0][0], box[0][1])
		box[0][0].Set(buf)

		buf.Set(box[1][1])
		box[1][1].Mul(quo, box[1][1]).Sub(box[1][0], box[1][1])
		box[1][0].Set(buf)

		x.Set(y)
		y.Set(mod)
	}

	return box[0][1], box[1][1]
}

// General solves equation: ax + by = dq
// X and Y are general solutions
// Where: X = xl - xr * t, t is any Integer
// And: Y = yl + yr * t
func General(a *big.Int, b *big.Int, dq *big.Int) (xl *big.Int, xr *big.Int, yl *big.Int, yr *big.Int) {
	var (
		x, y       = big.NewInt(0).Set(a), big.NewInt(0).Set(b)
		resl, resr = Simple(x, y)
		gcdab      = gcd.Gcd(x, y)
		q          = big.NewInt(0).Div(dq, gcdab)
	)
	return big.NewInt(0).Mul(resl, q), big.NewInt(0).Div(y.Neg(y), gcdab), big.NewInt(0).Mul(resr, q), big.NewInt(0).Div(x, gcdab)
}
