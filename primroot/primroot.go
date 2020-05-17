package primroot

import (
	"crypto/rand"
	"log"
	"math/big"

	"math.io/crath/gcd"
	"math.io/crath/modular"
	"math.io/crath/mulfunc"
)

// TODO: Add some theory and explanations

// Some interesting fact 2 and 3 functioins search only for 1 prime root ... this is dull

// p is a modulo of a prime root

// !!! Whether input is prime is not checked, which is bad ... but it is done for performance
// TODO!!!!: As for the rest numbers use phi(Euler totient function) .
// .. which I uncounciously actually have done already ... wait ... need some tests to prove it works with every number

// ..and actually not only for performance -- look at the nest functions and their inputs .. get it now ?

// Primrootp - prim root of p (where p is prime number by default) and p > 2
func Primrootp(p *big.Int) []*big.Int {
	var (
		input   = big.NewInt(0).Set(p)
		buf     = big.NewInt(0)
		one     = big.NewInt(1)
		euler   = mulfunc.Euler(input) //  == input - 1
		factors = mulfunc.Uniquefactor(euler)
		size    = big.NewInt(int64(len(factors)))

		// the only int64 here
		// lets us know total amount of roots there is
		total = mulfunc.Euler(euler).Int64()

		result = []*big.Int{}
	)

OuterLoop:
	for {
		// The rand approach...
		// It might be not that great with arbitrary arithmetic
		// For security purpose use crypto/rand.Int
		a, err := rand.Int(rand.Reader, input)
		if err != nil {
			log.Fatal("Error in random big integer generation")
		}
		if a.Cmp(big.NewInt(0)) == 0 {
			a.Add(a, one)
		}

		for i := big.NewInt(1); i.Cmp(size) <= 0; i.Add(i, one) {
			// TODO: There exists an alternative to BinaryModulo and you know it...
			if modular.BinaryModulo(a, buf.Div(euler, factors[i.Int64()-1]), input).Cmp(one) == 0 {
				continue OuterLoop
			}
		}

		result = append(result, a)
		break OuterLoop
	}

	for i := big.NewInt(2); i.Cmp(euler) <= 0; i.Add(i, one) {
		if gcd.Gcd(i, euler).Cmp(one) == 0 {
			result = append(result, modular.BinaryModulo(result[0], i, input))
		}

		// we break out of the loop when total amount is reached
		// might worsen the speed
		//TODO: add more tests to check it
		if int64(len(result)) == total {
			break
		}
	}

	return result
}

// p - is a prime number and p > 2, n - is power of p.

// Primrootpn calculates primroot of p^n - only one - but I'll leave an array as a return type
// just in case I'll change it later
func Primrootpn(p, n *big.Int) []*big.Int {
	var (
		primrootsp = Primrootp(p)
		result     = []*big.Int{}
		//mod        = modular.BinaryExponention(p, n)
		one  = big.NewInt(1)
		bufa = big.NewInt(0)
		bufb = big.NewInt(0)
	)

	// TODO: we dont really need buf we can do: big.NewInt(0).Sub(p, one) for that ..
	if modular.BinaryModulo(primrootsp[0], bufa.Sub(p, one), bufb.Mul(p, p)).Cmp(one) != 0 {
		result = append(result, primrootsp[0])
	} else {
		result = append(result, big.NewInt(0).Add(primrootsp[0], p))
	}

	// TODO: One day I'll be back and change it for the better
	//bufa.Sub(mod, one)
	//for i := big.NewInt(2); i.Cmp(mod) < 0; i.Add(i, one) {
	//	if gcd.Gcd(i, bufa).Cmp(one) == 0 {
	//		result = append(result, modular.BinaryModulo(result[0], i, mod))
	//	}
	//}

	return result
}

// Primroot2pn calculates primroot (only one) of value 2*p^n
func Primroot2pn(p, n *big.Int) []*big.Int {
	var (
		buf        = big.NewInt(0)
		one        = big.NewInt(1)
		zero       = big.NewInt(0)
		result     = []*big.Int{}
		primrootpn = Primrootpn(p, n)
	)

	if buf.And(primrootpn[0], one).Cmp(zero) != 0 {
		result = append(result, big.NewInt(0).Set(primrootpn[0]))
	} else {
		result = append(result, big.NewInt(0).Add(primrootpn[0], modular.BinaryExponention(p, n)))
	}

	return result
}
