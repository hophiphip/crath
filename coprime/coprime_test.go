package coprime

import (
	"crypto/rand"
	"math.io/crath/gcd"
	"math.io/crath/primegen"
	"math/big"
	"testing"
)

var (
	testCount  = 100
	numBitSize = 1024
)

func TestGetCoPrimeSimple(t *testing.T) {
	for i := 0; i < testCount; i++ {
		p, err := primegen.Primegen(rand.Reader, numBitSize)
		if err != nil {
			t.Error("Failed to generate prime number", i)
		}

		cp := GetCoPrimeSimple(p)
		if cp == nil {
			t.Error("Failed to generate correct prime", i,
				"Generation fail")
		}

		if gcd.Gcd(p, cp).Cmp(big.NewInt(1)) != 0 {
			t.Error("Failed to generate correct prime", i,
				"Gcd fail")
		}
	}
}
