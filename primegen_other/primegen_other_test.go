package primegen_other

import (
	"fmt"
	"math.io/crath/testprime"
	"testing"
)

var (
	roundCount = 10
	bitLength  = 2048
)

func TestBigPrimeGen(t *testing.T) {
	fmt.Println("Testing 'TestBigPrimeGen'")

	for iter := 0; iter < roundCount; iter += 1 {
		newPrime := BigPrimeGen(int64(bitLength))
		isPrime, err := testprime.ProbablyMillerRabin(newPrime, 20)

		if err != nil {
			t.Error("For iteration", iter,
				"generation has failed",
			)
		}

		if isPrime == false {
			t.Error("For bit length", iter,
				"and random generated value", newPrime,
				"prime test returned false",
			)
		}
	}
}
