package jacobi

import (
	"log"
	"math/big"
)

// TODO: log.Fatal is not a good way to handle errors

// Jacobi - calculates jacobi symbol, returns `int64` not `big.Int`
func Jacobi(a, m *big.Int) int64 {
	if big.NewInt(0).Mod(m, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Second input parameter can't be even")
	}
	var (
		result = int64(1)
		buf    = big.NewInt(0)
		num    = big.NewInt(0).Set(a)
		den    = big.NewInt(0).Set(m)
		one    = big.NewInt(1)
		zero   = big.NewInt(0)
		three  = big.NewInt(3)
		four   = big.NewInt(4)
		five   = big.NewInt(5)
		eight  = big.NewInt(8)
	)

	for num.Cmp(zero) != 0 {
		for buf.And(num, one).Cmp(zero) == 0 {
			num.Rsh(num, 1)
			buf.Mod(den, eight)
			if buf.Cmp(three) == 0 || buf.Cmp(five) == 0 {
				result *= -1
			}
		}

		buf.Set(num)
		num.Set(den)
		den.Set(buf)

		if buf.Mod(num, four).Cmp(three) == 0 && buf.Mod(den, four).Cmp(three) == 0 {
			result *= -1
		}

		num.Mod(num, den)
	}

	if den.Cmp(one) == 0 {
		return result
	}

	return 0
}
