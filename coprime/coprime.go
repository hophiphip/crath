package coprime

import (
	"crypto/rand"
	"log"
	"math.io/crath/primegen"
	"math/big"
)

var one = big.NewInt(1)
var coPrimeBitSize = 32

// areCoPrime asserts that the highest common divisor belonging to two integers is one.
func areCoPrime(x *big.Int, y *big.Int) bool {
	return x.Cmp(one) == 1 && x.Cmp(y) < 0 && new(big.Int).GCD(nil, nil, x, y).Cmp(one) == 0
}

// GetCoPrimeSimple - bad solution for generation co-prime number
func GetCoPrimeSimple(num *big.Int) *big.Int {
	var e *big.Int
	var err error
	for {
		e, err = primegen.Primegen(rand.Reader, coPrimeBitSize)

		if err != nil {
			log.Fatalf("error generating prime: %q", err)
		}

		if areCoPrime(e, num) {
			break
		}
	}

	return e
}
