package asymenc

import (
	"crypto/rand"
	"math.io/crath/coprime"
	"math.io/crath/primegen"
	"math/big"
)

type RsaContext struct {
	// p & q
	secretPrimeP *big.Int
	secretPrimeQ *big.Int
	// p & q bit size
	secretBitSize int

	// Euler function of n:
	// euler(n) = (p - 1)(q - 1)
	pairEuler *big.Int

	// e
	fixedCoPrime *big.Int

	// d
	modSolution *big.Int

	// n
	PairMultiplication *big.Int
}

func RsaSign(message string) error {
	context := RsaContext{
		secretPrimeP:       nil,
		secretPrimeQ:       nil,
		secretBitSize:     2048,
		pairEuler:          nil,
		fixedCoPrime:       nil,
		modSolution:        nil,
		PairMultiplication: nil,
	}

	// Init p
	secret , err := primegen.Primegen(rand.Reader, context.secretBitSize)
	if err != nil {
		return err
	}
	context.secretPrimeP = secret

	// Init q
	secret , err = primegen.Primegen(rand.Reader, context.secretBitSize)
	if err != nil {
		return err
	}
	context.secretPrimeQ = secret

	// Init n
	context.PairMultiplication = big.NewInt(0).Mul(
		context.secretPrimeP,
		context.secretPrimeQ)

	// Init euler
	one := big.NewInt(1)
	context.pairEuler = big.NewInt(0).Mul(
		big.NewInt(0).Sub(context.secretPrimeQ, one),
		big.NewInt(0).Sub(context.secretPrimeP, one))

	// Find fixed e
	context.fixedCoPrime = coprime.GetCoPrimeSimple(
		context.pairEuler)

	// Find d
	context.modSolution =

	return nil
}