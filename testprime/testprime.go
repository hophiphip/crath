package testprime

import (
	"crypto/rand"
	"errors"
	"math/big"

	"math.io/crath/gcd"
	"math.io/crath/jacobi"
	"math.io/crath/modular"
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

// ProbablySoloveyShtrassen - checks whether input argument is prime number
// func args:
//			n          - the number we need to test
//			iterations - optional argument (by default is 20),
//						 amount of itearations/tests that will be performed on input
//
// returns: true  - input is probably prime
//			false - input is not prime
func ProbablySoloveyShtrassen(n *big.Int, iterations ...int) (bool, error) {
	iterCount := 20
	if len(iterations) > 0 {
		iterCount = iterations[0]
	}

	var (
		two    = big.NewInt(2)
		one    = big.NewInt(1)
		buf    = big.NewInt(0)
		zero   = big.NewInt(0)
		jacbig *big.Int
	)

	// Check if number is 2
	if n.Cmp(two) == 0 {
		return true, nil
	}

	// Check if number is even
	// before passing it to Jacobi funcion
	if buf.Mod(n, two).Cmp(zero) == 0 {
		return false, nil
	}

	for i := 0; i < iterCount; i++ {
		a, err := rand.Int(rand.Reader, n)
		if err != nil {
			return true, errors.New("Error in random number generation")
		}
		if a.Cmp(zero) == 0 {
			a.Add(a, one)
		}

		if gcd.Gcd(a, n).Cmp(one) > 0 {
			return false, nil
		}

		modulo := modular.BinaryModulo(a, buf.Sub(n, one).Div(buf, two), n)
		jac, errj := jacobi.Jacobi(a, n)
		if errj != nil {
			return true, errors.New("Error in jacobi symbol calculation")
		}

		if jac == -1 {
			jacbig = big.NewInt(0).Sub(n, one)
		} else {
			jacbig = big.NewInt(1)
		}

		if jacbig.Cmp(modulo) != 0 {
			return false, nil
		}
	}

	return true, nil
}

// ProbablyMillerRabin - checks whether input argument is prime number
// func args:
//			n          - the number we need to test
//
// returns: true  - input is probably prime
//			false - input is not prime
func ProbablyMillerRabin(n *big.Int, iterations ...int) (bool, error) {
	iterCount := 6
	if len(iterations) > 0 {
		iterCount = iterations[0]
	}

	var (
		two    = big.NewInt(2)
		one    = big.NewInt(1)
		buf    = big.NewInt(0)
		zero   = big.NewInt(0)
		s      = 0
		q      = big.NewInt(0).Sub(n, one)
		modulo *big.Int
	)

	if n.Cmp(two) == 0 {
		return true, nil
	}

	if buf.Mod(n, two).Cmp(zero) == 0 {
		return false, nil
	}

	for {
		if buf.And(q, one).Cmp(zero) == 0 {
			goto NextRandom
		} else {
			q.Rsh(q, 1)
			s++
		}
	}

NextRandom:
	for i := 0; i < iterCount; i++ {
		a, err := rand.Int(rand.Reader, n)
		if err != nil {
			return true, errors.New("Error in random number generation")
		}
		// Random mustn't be 0
		if a.Cmp(zero) == 0 {
			a.Add(a, one)
		}

		if gcd.Gcd(a, n).Cmp(one) > 0 {
			return false, nil
		}

		modulo = modular.BinaryModulo(a, q, n)

		if modulo.Cmp(one) == 0 || modulo.Cmp(buf.Sub(n, one)) == 0 {
			continue
		}

		for j := 1; j < s; j++ {
			modulo = modular.BinaryModulo(modulo, two, n)
			if modulo.Cmp(buf.Sub(n, one)) == 0 {
				goto NextRandom
			}
			if modulo.Cmp(one) == 0 {
				return false, nil
			}
		}
		// Only can be reached if the inner loop was passed and
		// modulo didn't reach necessay criteria
		return false, nil
	}

	return true, nil

}
