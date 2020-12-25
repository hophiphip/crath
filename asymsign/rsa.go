package asymsign

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"

	"math.io/crath/coprime"
	"math.io/crath/diopheq"
	"math.io/crath/primegen"
)

// RsaSecret - secret part of rsa context
// ,must be cept in a secret
type RsaSecret struct {
	// p & q
	secretPrimeP *big.Int
	secretPrimeQ *big.Int

	// p & q bit size
	secretBitSize int

	// d
	modSolution *big.Int

	// Euler function of n:
	// euler(n) = (p - 1)(q - 1)
	pairEuler *big.Int
}

// RsaPublic - public shareable part of rsa context
type RsaPublic struct {
	// e
	FixedCoPrime *big.Int

	// n
	PairMultiplication *big.Int
}

// RsaContext full context with it's secret and public part
type RsaContext struct {
	secret *RsaSecret
	Public *RsaPublic
}

// SignedMessage contains message & sign
type SignedMessage struct {
	message string
	sign    *big.Int
}

// Init - initializes RSA context
func (ctx *RsaContext) Init() error {
	// Init p
	secret, err := primegen.Primegen(rand.Reader, ctx.secret.secretBitSize)
	if err != nil {
		return err
	}
	ctx.secret.secretPrimeP = secret

	// Init q
	secret, err = primegen.Primegen(rand.Reader, ctx.secret.secretBitSize)
	if err != nil {
		return err
	}
	ctx.secret.secretPrimeQ = secret

	// Init n
	ctx.Public.PairMultiplication = big.NewInt(0).Mul(
		ctx.secret.secretPrimeP,
		ctx.secret.secretPrimeQ)

	// Init euler
	one := big.NewInt(1)
	ctx.secret.pairEuler = big.NewInt(0).Mul(
		big.NewInt(0).Sub(ctx.secret.secretPrimeQ, one),
		big.NewInt(0).Sub(ctx.secret.secretPrimeP, one))
	// TODO: Remove debug test
	// sub one test
	// fmt.Println("Sub one:")
	// fmt.Println(ctx.secret.secretPrimeQ, " ", big.NewInt(0).Sub(ctx.secret.secretPrimeQ, one))
	// fmt.Println(ctx.secret.secretPrimeP, " ", big.NewInt(0).Sub(ctx.secret.secretPrimeP, one))
	// fmt.Println("")

	// Find fixed e
	ctx.Public.FixedCoPrime = coprime.GetCoPrimeSimple(
		ctx.secret.pairEuler)

	// Find d
	_, ctx.secret.modSolution = diopheq.Simple(ctx.secret.pairEuler, ctx.Public.FixedCoPrime)

	// TODO: Remove
	// Debug Test
	// test := big.NewInt(0).Mul(ctx.secret.modSolution, ctx.Public.FixedCoPrime)
	// test.Mod(test, ctx.secret.pairEuler)
	// if test.Cmp(one) != 0 {
	// 	log.Fatal("Modular failure")
	// }

	return nil
}

// Sign - signs message
func (ctx *RsaContext) Sign(message string) (*SignedMessage, error) {
	smes := &SignedMessage{
		message: "",
		sign:    nil,
	}

	hash := fixedHash(smes.message)

	messageSign := big.NewInt(0).Exp(hash, ctx.secret.modSolution, ctx.Public.PairMultiplication)

	smes.message = message
	smes.sign = messageSign

	return smes, nil
}

// Verify - verifies message sign
func (smes *SignedMessage) Verify(pub *RsaPublic) bool {
	hash := big.NewInt(0).Mod(fixedHash(smes.message), pub.PairMultiplication)

	mod := big.NewInt(0).Exp(
		smes.sign,
		pub.FixedCoPrime,
		pub.PairMultiplication)

	// TODO: remove debug print
	fmt.Println(smes.message)
	fmt.Println(smes.sign)
	fmt.Println(fixedHash(smes.message))
	fmt.Println(hash)
	fmt.Println(mod)

	return mod.Cmp(hash) == 0
}

func fixedHash(message string) *big.Int {
	// Converts md5 byte[16] result value to byte[] (slice)
	fixedHash := md5.Sum([]byte(message))
	return big.NewInt(0).SetBytes(fixedHash[:])
}
