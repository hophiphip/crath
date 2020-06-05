package primegen

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestPrimegen(t *testing.T) {
	fmt.Println("Testing 'Primegen'")

	for i := 10; i < 256; i++ {
		randp, err := Primegen(rand.Reader, i)
		if err != nil {
			t.Error("For bit length", i,
				"random genereator returned", err,
			)
		}

		fmt.Println("Random value:", randp)

		/*
			Not necessay, as Primegen checks the value before returning it
			Uncomment it in case further tests are needed
		*/
		//isPrime, errp := testprime.ProbablyMillerRabin(randp, 20)
		//if errp != nil {
		//	t.Error("For bit length", i,
		//		"and random generated value", randp,
		//		"prime test returned error",
		//	)
		//}
		//if isPrime == false {
		//	t.Error("For bit length", i,
		//		"and random generated value", randp,
		//		"prime test returned false",
		//	)
		//}
	}
}
