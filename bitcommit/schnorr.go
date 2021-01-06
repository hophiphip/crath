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
	// w size
	secretKeySize = 128
	// random size
	// r size
	rSize = 100
	// e size
	eSize = 72

	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

// SchnorrContext - contains public & private values for keys
type SchnorrContext struct {
	q *big.Int
	// necessary for calculating p
	qp *big.Int

	p *big.Int

	// necessary for calculating g
	a *big.Int
	g *big.Int

	y *big.Int

	// private key
	w *big.Int
}

// PublicPart  - values of public key
type PublicPart struct {
	p, q, g, y *big.Int
}

// PrivatePart - values of private key
type PrivatePart struct {
	w *big.Int
}

// RandomPart - client random number generation
type RandomPart struct {
	r, x *big.Int
}

// Init - define all values for context
func (ctx *SchnorrContext) Init() error {
	// Will check whether number is a prime one with certainty:  1 - 1/2*certainty
	certainty := 100

	// Initialize q
	_q, err := primegen.Primegen(rand.Reader, SmallKeySize)
	if err != nil {
		return err
	}

	ctx.q.Set(_q)
	ctx.qp.Set(one)

	// Calculate p
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

	// Calculate g
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

		// a^ga mod p = 1
		ctx.g.Exp(ctx.a, ga, ctx.p)

		// g != 1
		if ctx.g.Cmp(one) != 0 {
			break
		}
	}

	// Calculate w
	for {
		_w, err := primegen.Primegen(rand.Reader, secretKeySize)
		if err != nil {
			return err
		}
		ctx.w.Set(_w)

		if ctx.w.Cmp(ctx.q) < 0 {
			break
		}
	}

	// Calculate y
	_negw := big.NewInt(0).Neg(ctx.w)
	ctx.y.Exp(ctx.g, _negw, ctx.p)

	return nil
}

// GetPublicPart - basically returns public key
func (ctx *SchnorrContext) GetPublicPart() *PublicPart {
	return &PublicPart{
		p: ctx.p,
		q: ctx.q,
		g: ctx.g,
		y: ctx.y,
	}
}

// GetPrivatePart - basically returns private key
func (ctx *SchnorrContext) GetPrivatePart() *PrivatePart {
	return &PrivatePart{
		w: ctx.w,
	}
}

// FindR - random r, generated on a client
func (public *PublicPart) FindR() (*RandomPart, error) {
	random := &RandomPart{
		r: big.NewInt(0),
		x: big.NewInt(0),
	}

	// Find r
	for {
		_r, err := primegen.Primegen(rand.Reader, rSize)
		if err != nil {
			return nil, err
		}
		if _r.Cmp(public.q) < 0 {
			random.r.Set(_r)
			break
		}
	}

	// Find x
	random.x.Exp(public.g, random.r, public.p)

	return random, nil
}

// GenE - generates e (B response)
func GenE() (*big.Int, error) {
	e, err := primegen.Primegen(rand.Reader, eSize)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// FindS - gneerates s (to send to B)
func (ctx *SchnorrContext) FindS(e *big.Int, random *RandomPart) *big.Int {
	s := big.NewInt(0)
	s.Mul(ctx.w, e)
	s.Add(s, random.r)
	s.Mod(s, ctx.q)

	return s
}

// IsApproved - B approve function
func IsApproved(x, s, e *big.Int, public *PublicPart) bool {
	_x1 := big.NewInt(0).Exp(public.y, e, public.p)
	_x2 := big.NewInt(0).Exp(public.g, s, public.p)
	_x1.Mul(_x1, _x2)
	_x1.Mod(_x1, public.p)

	return x.Cmp(_x1) == 0
}
