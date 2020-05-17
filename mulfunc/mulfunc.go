package mulfunc

import (
	"errors"
	"log"
	"math/big"
)

// TODO: Give functions their proper names

// CountFractions count amount of fractions of a specific number
func CountFractions(input *big.Int) (count *big.Int) {
	num, pow, count, zero, buf, one := big.NewInt(0).Set(input), big.NewInt(0), big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)

	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		pow.Set(zero)
		for buf.Mod(num, p).Cmp(zero) == 0 {
			num.Div(num, p)
			pow.Add(pow, one)
		}
		count.Mul(count, buf.Add(pow, one))
	}

	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			pow.Set(zero)

			for buf.Mod(num, p).Cmp(zero) == 0 {
				num.Div(num, p)
				pow.Add(pow, one)
			}
			count.Mul(count, buf.Add(pow, one))
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}
	}

	if num.Cmp(one) > 0 {
		count.Mul(count, big.NewInt(2))
	}

	return count
}

// SumFractions ...
func SumFractions(input *big.Int) (sum *big.Int) {
	num, pow, sum, zero, buf, one := big.NewInt(0).Set(input), big.NewInt(2), big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)

	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			pow.Set(p)

			for buf.Mod(num, p).Cmp(zero) == 0 {
				num.Div(num, p)
				pow.Mul(pow, p)
			}
			sum.Mul(sum, buf.Div(pow.Sub(pow, one), p.Sub(p, one)))
		}
	}

	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			pow.Set(p)
			//FIX! loops endlessly here..
			for buf.Mod(num, p).Cmp(zero) == 0 {
				num.Div(num, p)
				pow.Mul(pow, p)
				//fmt.Println(p, num)
			}
			sum.Mul(sum, buf.Div(pow.Sub(pow, one), p.Sub(p, one)))
			//fmt.Println(p)
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}

		if num.Cmp(one) > 0 {
			sum.Mul(sum, buf.Div(num.Mul(num, num).Sub(num, one), buf.Sub(input, one)))
		}
	}

	return sum
}

// Euler calculates Euler totient function
func Euler(input *big.Int) (res *big.Int) {
	num, res, zero, buf, one := big.NewInt(0).Set(input), big.NewInt(0).Set(input), big.NewInt(0), big.NewInt(0), big.NewInt(1)

	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			for buf.Mod(num, p).Cmp(zero) == 0 {
				num.Div(num, p)
			}
			res.Sub(res, buf.Div(res, p))
		}
	}

	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			for buf.Mod(num, p).Cmp(zero) == 0 {
				num.Div(num, p)
			}
			res.Sub(res, buf.Div(res, p))
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}
	}

	if num.Cmp(one) > 0 {
		res.Sub(res, buf.Div(res, num))
	}

	return res
}

// Factorization - factorizes integers
func Factorization(input *big.Int) []*big.Int {
	var (
		num  = big.NewInt(0).Set(input)
		zero = big.NewInt(0)
		buf  = big.NewInt(0)
		one  = big.NewInt(1)
		res  = []*big.Int{}
	)

	if num.Cmp(zero) == 0 {
		log.Fatal(errors.New("Input can't be zero"))
	}

	if num.Cmp(one) == 0 {
		return append(res, one)
	}

	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		for buf.Mod(num, p).Cmp(zero) == 0 {
			res = append(res, big.NewInt(0).Set(p))
			num.Div(num, p)
		}
	}

	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		for buf.Mod(num, p).Cmp(zero) == 0 {
			res = append(res, big.NewInt(0).Set(p))
			num.Div(num, p)
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}
	}

	if num.Cmp(one) > 0 {
		res = append(res, num)
	}

	return res
}

// Uniquefactor - factorizes integers, but only unique integers
func Uniquefactor(input *big.Int) []*big.Int {
	var (
		num  = big.NewInt(0).Set(input)
		zero = big.NewInt(0)
		buf  = big.NewInt(0)
		one  = big.NewInt(1)
		res  = []*big.Int{}
	)

	if num.Cmp(zero) == 0 {
		log.Fatal(errors.New("Input can't be zero"))
	}

	if num.Cmp(one) == 0 {
		return append(res, one)
	}

	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			res = append(res, big.NewInt(0).Set(p))
		}
		for buf.Mod(num, p).Cmp(zero) == 0 {
			num.Div(num, p)
		}
	}

	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			res = append(res, big.NewInt(0).Set(p))
		}
		for buf.Mod(num, p).Cmp(zero) == 0 {
			num.Div(num, p)
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}
	}

	if num.Cmp(one) > 0 {
		res = append(res, num)
	}

	return res
}
