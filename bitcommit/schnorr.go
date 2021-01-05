package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"math.io/crath/primegen"
)

var (
	// p size
	keySize = 1024
	// q size
	smallKeySize = 160

	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	six   = big.NewInt(6)
)

func findQ(prime *big.Int) *big.Int {
	// buffer for calculations
	buf := big.NewInt(0)
	// result value - Q
	ret := big.NewInt(0)

	// We need: p - 1
	num := big.NewInt(0).Sub(prime, one)

	for _, p := range [3]*big.Int{two, three, five} {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			ret.Set(p)
			num.Div(num, p)
		}
	}

	iteration := 0
	for p, i, jumps := big.NewInt(7), -1, [8]*big.Int{four, two, four, two, four, six, two, six}; buf.Mul(p, p).Cmp(num) <= 0; p.Add(p, jumps[i]) {
		if buf.Mod(num, p).Cmp(zero) == 0 {
			ret.Set(p)
			num.Div(num, p)
		}

		if len(ret.Bytes()) >= smallKeySize {
			goto final
		}

		// Set next jump index
		if i == 7 {
			i = 0
		} else {
			i++
		}

		// DEBUG
		if iteration%100000 == 0 {
			fmt.Println("[DEBUG] ITERATION:[", iteration, "] Q:[", ret, "] SIZE:[", len(ret.Bytes()), "]")
		}
		// DEBUG
		iteration++
	}

final:

	return ret
}

func main() {
	buf := big.NewInt(0)

	// Generate prime number p
	p, err := primegen.Primegen(rand.Reader, keySize)
	if err != nil {
		log.Fatal("failed to generate key (p)")
	}

	// D
	fmt.Println("P size:", len(p.Bytes()))
	// D

	// Choose q: p - 1 = 0 (mod q) , q is 160 bits long
	q := findQ(p)

	// D
	fmt.Println("Mod:", buf.Mod(buf.Sub(p, one), q))
	fmt.Println("Q size:", len(q.Bytes()))
	fmt.Println("P size:", len(p.Bytes()))
	// D
}
