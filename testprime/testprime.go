package testprime

import (
	"errors"
	"math/big"
)

// SimpleTest checks whether input argument is a prime number
// returns: (true, nil) if number is prime
//			(false,nil) if number is not prime
//			(false, error) in case of 1 and 0, as they are neither prime nor not prime
func SimpleTest(n *big.Int) (bool, error) {
	var (
		zero = big.NewInt(0)
		buf  = big.NewInt(0)
		one  = big.NewInt(1)
	)

	if n.Cmp(zero) == 0 {
		return false, errors.New("Function argument can not be 'zero'")
	}

	if n.Cmp(one) == 0 {
		return false, errors.New("Function argument can not be 'one'")
	}

	// Checks for 2,3,5
	for _, p := range [3]*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)} {
		if n.Cmp(p) == 0 {
			return true, nil
		}
		if buf.Mod(n, p).Cmp(zero) == 0 {
			return false, nil
		}
	}
	// Specifically check if input is equal 7
	if n.Cmp(big.NewInt(7)) == 0 {
		return true, nil
	}

	// Use jump table to test input further, starting with 7
	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(2), big.NewInt(4), big.NewInt(6), big.NewInt(2), big.NewInt(6)}; buf.Mul(p, p).Cmp(n) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(n, p).Cmp(zero) == 0 {
			return false, nil
		}
		if i == 7 {
			i = 0
		} else {
			i++
		}
	}

	return true, nil
}

// SoloveyShtrassenTest - checks whether input argument is prime number
func SoloveyShtrassenTest(n *big.Int) {

}

// MillerRabinTest - checks whether input argument is prime number
func MillerRabinTest(n *big.Int) {

}
