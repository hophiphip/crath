package bitcommit

// Definetly not the best and not the fastest solution
// For better example look at ECDSA implementation

import (
	"crypto/rand"
	"math/big"

	"math.io/crath/primegen"
)

var (
	// KeySize - big-key bit size
	// p size
	KeySize = 1024
	// SmallKeySize - small-key bit size
	// q size
	SmallKeySize = 160

	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	five  = big.NewInt(5)
	six   = big.NewInt(6)
)

// NOTE: this imlpementation doesn't bother about 160 bits small-key size, but uses the first found one

// SchnorrContext - contains public & private values for keys
type SchnorrContext struct {
	q  *big.Int
	qp *big.Int
	p  *big.Int
	a  *big.Int
	g  *big.Int
}

// Init - define all values for context
func (ctx *SchnorrContext) Init() error {
	// Will check whether number is a prime one with certainty:  1 - 1/2*certainty
	certainty := 100

	_q, err := primegen.Primegen(rand.Reader, KeySize)
	if err != nil {
		return err
	}

	ctx.q.Set(_q)
	ctx.qp.Set(one)

	for {
		// p := 2*(q * qp) + 1
		ctx.p.Mul(ctx.q, ctx.qp)
		ctx.p.Mul(ctx.p, two)
		ctx.p.Add(ctx.p, one)

		if ctx.p.ProbablyPrime(certainty) {
			break
		}

		ctx.qp.Add(ctx.qp, one)
	}

	for {
		// (2+a) mod p
		_a, err := primegen.Primegen(rand.Reader, SmallKeySize)
		if err != nil {
			return err
		}
		_a.Add(_a, two)
		_a.Mod(_a, ctx.p)
		ctx.a.Set(_a)

		// (p - 1)/q
		ga := big.NewInt(0).Sub(ctx.p, one)
		ga.Div(ga, ctx.q)

		// a^ga mod p
		ctx.g.Exp(ctx.a, ga, ctx.p)

		// g != 1
		if ctx.g.Cmp(one) != 0 {
			break
		}
	}

	return nil
}

// Test ...
func Test() {

}
