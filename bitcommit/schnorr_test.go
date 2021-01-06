package bitcommit

import (
	"math/big"
	"testing"
)

// It works, but requires refactoring

func TestSchnorrSignature(t *testing.T) {
	ctx := &SchnorrContext{
		q:  big.NewInt(0),
		qp: big.NewInt(0),
		p:  big.NewInt(0),
		a:  big.NewInt(0),
		g:  big.NewInt(0),
		y:  big.NewInt(0),
		w:  big.NewInt(0),
	}

	err := ctx.Init()
	if err != nil {
		t.Error("Context initialization has failed.")
	}

	//private := ctx.GetPrivatePart()
	Public := ctx.GetPublicPart()

	clientRandom, err := Public.FindR()
	if err != nil {
		t.Error("Failed to generate random values for client (r, x).")
	}

	// A sends x to B
	x := clientRandom.x

	// B sends random e
	e, err := GenE()
	if err != nil {
		t.Error("failed to generate B's response (e).")
	}

	// A calculates: s = r + we (mod q) and sends it to B
	s := ctx.FindS(e, clientRandom)

	// B checks
	if !IsApproved(x, s, e, Public) {
		t.Error("Approve has failed. That is not supposed to happen")
	}
}
