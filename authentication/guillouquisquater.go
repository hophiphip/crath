package authentication

import (
	"crypto/rand"
	"math/big"

	"math.io/crath/gcd"

	"math.io/crath/primegen"
)

var (
	primeNumBitSize = 1024
	eBitSize        = 512
	xBitSize        = 512
	constKBitSize   = 512
	aBitSize        = 256

	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

// GqCtx - contains necessary values for calculation
type GqCtx struct {
	Public  *PublicPart
	private *privatePart
}

// PublicPart - public part of GQ context
type PublicPart struct {
	n, e *big.Int
}

// privatePart - private part of GQ context
type privatePart struct {
	p, q    *big.Int
	pqEuler *big.Int
}

// GqClientCtx - client constants
type GqClientCtx struct {
	x, Y    *big.Int
	k, r, s *big.Int
}

// GqAuthenticatorCtx - authenticator constants
type GqAuthenticatorCtx struct {
	a *big.Int
}

// Init - init context
func (ctx *GqCtx) Init() error {
	var err error

	// Init p
	ctx.private.p, err = primegen.Primegen(rand.Reader, primeNumBitSize)
	if err != nil {
		return err
	}

	// Init q
	ctx.private.q, err = primegen.Primegen(rand.Reader, primeNumBitSize)
	if err != nil {
		return err
	}

	// Init n
	ctx.Public.n.Mul(ctx.private.p, ctx.private.q)

	// Init pqEuler
	_q := big.NewInt(0).Sub(ctx.private.q, one)
	_p := big.NewInt(0).Sub(ctx.private.p, one)
	ctx.private.pqEuler.Mul(_q, _p)

	// Init e
	for {
		ctx.Public.e, err = primegen.Primegen(rand.Reader, eBitSize)
		if err != nil {
			return err
		}

		if ctx.Public.e.Cmp(ctx.private.pqEuler) < 0 && ctx.Public.e.Cmp(two) != 0 {
			if gcd.Gcd(ctx.Public.e, ctx.private.pqEuler).Cmp(one) == 0 {
				break
			}
		}
	}

	return nil
}

// Init client constants
func (client *GqClientCtx) Init(public *PublicPart) error {
	var err error

	// Init x - secret
	for {
		client.x, err = primegen.Primegen(rand.Reader, xBitSize)
		if err != nil {
			return err
		}

		if gcd.Gcd(client.x, public.n).Cmp(one) == 0 {
			break
		}
	}

	// Init Y - public
	negE := big.NewInt(0).Neg(public.e)
	client.Y.Exp(client.x, negE, public.n)

	return nil
}

// Step1 - first step of authentication
func (client *GqClientCtx) Step1(public *PublicPart) error {
	var err error
	top := big.NewInt(0).Sub(public.n, two)

	for {
		client.k, err = primegen.Primegen(rand.Reader, constKBitSize)
		if err != nil {
			return nil
		}

		if client.k.Cmp(top) <= 0 && client.k.Cmp(one) >= 0 {
			break
		}
	}

	client.r.Exp(client.k, public.e, public.n)

	return nil
}

// Step2 - second step (first step for authenticator)
func (auth *GqAuthenticatorCtx) Step2(public *PublicPart) error {
	var err error
	for {
		auth.a, err = primegen.Primegen(rand.Reader, aBitSize)
		if err != nil {
			return err
		}

		if auth.a.Cmp(public.e) < 0 && auth.a.Cmp(zero) >= 0 {
			break
		}
	}

	return nil
}

// Step3 - respons to step2
func (client *GqClientCtx) Step3(public *PublicPart, a *big.Int) error {
	client.s.Exp(client.x, a, public.n)
	client.s.Mul(client.s, client.k)
	client.s.Mod(client.s, public.n)

	return nil
}

// Step4 - final step(verification)
func (auth *GqAuthenticatorCtx) Step4(public *PublicPart, client *GqClientCtx) bool {
	_y := big.NewInt(0).Exp(client.Y, auth.a, public.n)
	_s := big.NewInt(0).Exp(client.s, public.e, public.n)
	_y.Mul(_y, _s)
	_y.Mod(_y, public.n)

	return client.r.Cmp(_y) == 0
}
