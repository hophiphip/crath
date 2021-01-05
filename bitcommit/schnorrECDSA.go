package bitcommit

// Implementation that is based on eliptic curve cryptography (ECC): ECDSA

import (
	"fmt"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

var (
	curve      = edwards25519.NewBlakeSHA256Ed25519()
	hashSha256 = curve.Hash()
)

// SchnorrContext - public & private keys
type SchnorrContext struct {
	privateKey kyber.Scalar
	PublicKey  kyber.Point
}

// Signature - info on signature
type Signature struct {
	r kyber.Point
	s kyber.Scalar
}

// hash - an example hash function
func hash(s string) kyber.Scalar {
	hashSha256.Reset()
	hashSha256.Write([]byte(s))

	return curve.Scalar().SetBytes(hashSha256.Sum(nil))
}

// NewCtx - generates schnorr context with public & private keys
func NewCtx() SchnorrContext {
	private := curve.Scalar().Pick(curve.RandomStream())
	return SchnorrContext{
		privateKey: private,
		PublicKey:  curve.Point().Mul(private, curve.Point().Base()),
	}
}

// ComparePublic - compares public key with another
func (context SchnorrContext) ComparePublic(derivedPublicKey kyber.Point) bool {
	return context.PublicKey.Equal(derivedPublicKey)
}

// Sign - signs message
func (context SchnorrContext) Sign(message string) Signature {
	// Base of the curve
	g := curve.Point().Base()

	// Pick a random 'k' from allowed set
	k := curve.Scalar().Pick(curve.RandomStream())

	// r = k * G (same as r = g^k)
	r := curve.Point().Mul(k, g)

	// hash of m || r
	e := hash(message + r.String())

	// s = k - e * x
	s := curve.Scalar().Sub(k, curve.Scalar().Mul(e, context.privateKey))

	return Signature{
		r: r,
		s: s,
	}
}

// PublicKey - generates derived public key
func PublicKey(message string, signature Signature) kyber.Point {
	// create a new generator
	g := curve.Point().Base()

	// e = hash of m || r
	e := hash(message + signature.r.String())

	// y = (r - s * G) * (1 / e)
	y := curve.Point().Sub(signature.r, curve.Point().Mul(signature.s, g))
	y = curve.Point().Mul(curve.Scalar().Div(curve.Scalar().One(), e), y)

	return y
}

// Verify - verifies message signature
func Verify(message string, signature Signature, publicKey kyber.Point) bool {
	// create a new generator
	g := curve.Point().Base()

	// e = hash of m || r
	e := hash(message + signature.r.String())

	// Try to reconstruct 's * G'
	// s * G = r - e * y
	sGvalidate := curve.Point().Sub(signature.r, curve.Point().Mul(e, publicKey))

	// Actual 's * G'
	sG := curve.Point().Mul(signature.s, g)

	// Check equality
	return sG.Equal(sGvalidate)
}

func (s Signature) String() string {
	return fmt.Sprintf("(r=%s, s=%s)", s.r, s.s)
}
