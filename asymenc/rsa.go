package asymenc

import (
	"crypto/rand"
	"math/big"

	"math.io/crath/coprime"
	"math.io/crath/diopheq"
	"math.io/crath/primegen"
)

// RsaPublic - public shareable part of rsa context
type RsaPublic struct {
	// e
	FixedCoPrime *big.Int

	// n
	PairMultiplication *big.Int
}

// RsaSecret  - secret part of rsa context
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

// RsaContext full context with it's secret and public part
type RsaContext struct {
	secret *RsaSecret
	Public *RsaPublic
}

// RsaPrivateKey - number big.Int that is a secret
type RsaPrivateKey struct {
	// d
	key *big.Int
}

// RsaPublicKey - two numbers that are a public key
type RsaPublicKey struct {
	// e
	fixed *big.Int

	// n
	multiplication *big.Int
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

	// Find fixed e
	ctx.Public.FixedCoPrime = coprime.GetCoPrimeSimple(
		ctx.secret.pairEuler)

	// Find d
	_, ctx.secret.modSolution = diopheq.Simple(ctx.secret.pairEuler, ctx.Public.FixedCoPrime)

	return nil
}

// GetPrivateKey - returns private key(*big.Int) in a RSA context
func (secret *RsaSecret) GetPrivateKey() *RsaPrivateKey {
	return &RsaPrivateKey{
		secret.modSolution,
	}
}

// GetPublicKey - returns public key (two numers)
func (public *RsaPublic) GetPublicKey() *RsaPublicKey {
	return &RsaPublicKey{
		public.FixedCoPrime,
		public.PairMultiplication,
	}
}

// Encrypt - encrypts message with public key
// (message must be smaller than bitsize of n)
func (public *RsaPublicKey) Encrypt(message string) *big.Int {
	mesBigInt := big.NewInt(0).SetBytes([]byte(message))
	return big.NewInt(0).Exp(
		mesBigInt,
		public.fixed,
		public.multiplication,
	)
}

// Decrypt - decrypts a message with private key
func (private *RsaPrivateKey) Decrypt(encrypted *big.Int, publicKey *RsaPublicKey) string {
	decrMesBigInt := big.NewInt(0).Exp(
		encrypted,
		private.key,
		publicKey.multiplication,
	)

	return string(decrMesBigInt.Bytes())
}
