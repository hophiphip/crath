package coinflip

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

	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

// Context - contains all necessary values
type Context struct {
	q *big.Int
	// necessary for calculating p
	qp *big.Int

	p *big.Int

	// necessary for calculating g
	a *big.Int
	g *big.Int
}

// Init - define all values for context
func (ctx *Context) Init() error {
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

	return nil
}

func generateBit() (*big.Int, error) {
	// two is max value:
	// so the only possible results are: {0, 1}
	r, err := rand.Int(rand.Reader, two)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// The precedure is represented with 2 sides: Initiator & Receiver
// .. with two steps each

// InitiatorStep1Results - data from first step
type InitiatorStep1Results struct {
	x, y *big.Int
}

// ReceiverStep1Results - data from first step
type ReceiverStep1Results struct {
	k, aBit, r *big.Int
}

// InitiatorStep1 ...
// B-1
func InitiatorStep1(ctx *Context) (*InitiatorStep1Results, error) {
	// B: generates x (1 <= x <= q - 1)
	var x *big.Int
	subq := big.NewInt(0).Sub(ctx.q, big.NewInt(1))
	for {
		var err error
		x, err = primegen.Primegen(rand.Reader, SmallKeySize/2)
		if err != nil {
			return nil, err
		}
		if x.Cmp(subq) <= 0 {
			break
		}
	}

	// B: generates y
	y := big.NewInt(0).Exp(ctx.g, x, ctx.p)

	return &InitiatorStep1Results{
		x: x,
		y: y,
	}, nil
}

// ReceiverStep1 ...
func ReceiverStep1(ctx *Context, res1 *InitiatorStep1Results) (*ReceiverStep1Results, error) {
	// A: generates k (1 <= k <= q - 1)
	var k *big.Int
	subq := big.NewInt(0).Sub(ctx.q, big.NewInt(1))
	for {
		var err error
		k, err = primegen.Primegen(rand.Reader, SmallKeySize/2)
		if err != nil {
			return nil, err
		}
		if res1.x.Cmp(subq) <= 0 {
			break
		}
	}

	// A: generates a
	max := big.NewInt(2)
	a, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}

	// A: calculates r
	r := big.NewInt(0).Exp(ctx.g, k, ctx.p)
	_r := big.NewInt(0).Exp(res1.y, a, ctx.p)
	r.Mul(r, _r)
	r.Mod(r, ctx.p)

	return &ReceiverStep1Results{
		k:    k,
		aBit: a,
		r:    r,
	}, nil
}

// InitiatorStep2 ...
func InitiatorStep2() (b *big.Int, e error) {
	b, err := rand.Int(rand.Reader, two)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// InitiatorStepCheck ...
func InitiatorStepCheck(ctx *Context, res1 *ReceiverStep1Results, init1 *InitiatorStep1Results) bool {
	// B checks
	left := big.NewInt(0).Exp(ctx.g, res1.k, ctx.p)
	_left := big.NewInt(0).Exp(init1.y, res1.aBit, ctx.p)
	left.Mul(left, _left)
	left.Mod(left, ctx.p)

	right := big.NewInt(0).Mod(res1.r, ctx.p)

	// Compare keys
	return right.Cmp(left) == 0
}
