package primegen

import (
	"crypto/rand"
	"fmt"
	"testing"
)

var (
	roundCount = 100
	bitLength  = 2048
)

func TestPrimegen(t *testing.T) {
	fmt.Println("Testing 'Primegen'")

	for i := 10; i < roundCount; i++ {
		newPrime, err := Primegen(rand.Reader, bitLength)
		if err != nil {
			t.Error("For bit length", bitLength,
				"random generator returned", err,
			)
		}

		fmt.Println("Random value:", newPrime)

		/*
			Not necessary, as Primegen checks the value before returning it
			Uncomment it in case further tests are needed
		*/
		//isPrime, err := testprime.ProbablyMillerRabin(newPrime, 20)
		//if err != nil {
		//	t.Error("For bit length", i,
		//		"and random generated value", newPrime,
		//		"prime test returned error",
		//	)
		//}
		//if isPrime == false {
		//	t.Error("For bit length", i,
		//		"and random generated value", newPrime,
		//		"prime test returned false",
		//	)
		//}
	}
}
