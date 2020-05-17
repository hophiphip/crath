package modular

import (
	"math/big"

	"math.io/crath/diopheq"
	"math.io/crath/fractions"
	"math.io/crath/gcd"
	"math.io/crath/mulfunc"
)

// BinaryExponention is used to quickly calculate power of a number
func BinaryExponention(value, power *big.Int) *big.Int {
	var (
		result = big.NewInt(1)
		zero   = big.NewInt(0)
		buf    = big.NewInt(0)
		two    = big.NewInt(2)
		pow    = big.NewInt(0).Set(power)
		val    = big.NewInt(0).Set(value)
	)
	for pow.Cmp(zero) > 0 {
		if buf.Mod(pow, two).Cmp(zero) != 0 {
			result.Mul(result, val)
		}
		val.Mul(val, val)
		pow.Rsh(pow, 1)
	}

	return result
}

// TODO!!!: BinaryModulo can be done without 'buf' .. well mostly

// BinaryModulo is used to efficiently calculate modulo of a large number
// is same as: value^power % mod
func BinaryModulo(value, power, mod *big.Int) *big.Int {
	var (
		val    = big.NewInt(0).Set(value)
		pow    = big.NewInt(0).Set(power)
		zero   = big.NewInt(0)
		one    = big.NewInt(1)
		result = big.NewInt(1)
		buf    = big.NewInt(0)
	)

	val.Mod(val, mod)
	for pow.Cmp(zero) > 0 {
		if buf.And(pow, one).Cmp(zero) != 0 {
			result.Mod(buf.Mul(result, val), mod) // here -- just do it in order res.Mul().Mod() ...
		}
		val.Mod(buf.Mul(val, val), mod) // here
		pow.Rsh(pow, 1)
	}

	return result
}

// modularFunction is a generic type for different approaches to calculate modular component
type modularFunction func(*big.Int, *big.Int, *big.Int) *big.Int

// Modulareuler - calculates modular component using BinaryModulo algorithm
func Modulareuler(a, b, m *big.Int) *big.Int {
	buf := big.NewInt(0)
	return big.NewInt(0).Mul(b, BinaryModulo(a, buf.Sub(mulfunc.Euler(m), big.NewInt(1)), m))
}

// Modulardioph - calculates modular component using diophantine equations
func Modulardioph(a, b, m *big.Int) *big.Int {
	_, right := diopheq.Simple(m, a)
	return big.NewInt(0).Mul(b, right)
}

// FIX: Incorrect return result

// Modularfract - calculates modular component using continuous fractions
func Modularfract(a, b, m *big.Int) *big.Int {
	buf := big.NewInt(0)
	fraction := fractions.Continuous(m, a)
	size := len(fraction)
	coeff := 0

	// we remove last element
	if size > 0 {
		fraction = fraction[:size-1]
	}

	prnum, _ := fractions.Normal(fraction)

	if (size-1)%2 != 0 {
		coeff = -1
	} else {
		coeff = 1
	}

	return buf.Mul(big.NewInt(int64(coeff)), buf.Mul(b, prnum))
}

// GetSolution - returns solution for linear congruence equation
// @modfunc - ModularFucntion which you want to use to calculate the result
func GetSolution(modfunc modularFunction, a, b, m *big.Int) []*big.Int {
	var (
		abuf     = big.NewInt(0).Set(a)
		bbuf     = big.NewInt(0).Set(b)
		mbuf     = big.NewInt(0).Set(m)
		x        = big.NewInt(0)
		zero     = big.NewInt(0)
		one      = big.NewInt(1)
		gcd      = gcd.Gcd(a, m)
		buf      = big.NewInt(0)
		solution []*big.Int
	)

	if buf.Mod(b, gcd).Cmp(zero) != 0 || m.Cmp(one) < 0 {
		return solution
	}

	abuf.Div(abuf, gcd)
	bbuf.Div(bbuf, gcd)
	mbuf.Div(mbuf, gcd)
	x.Mod(BinaryModulo(modfunc(abuf, bbuf, mbuf), one, mbuf), mbuf)

	for gcd.Cmp(zero) > 0 {
		solution = append(solution, x)
		x.Add(x, mbuf)
		gcd.Sub(gcd, one)
	}

	return solution
}
