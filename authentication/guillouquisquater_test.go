package authentication

import (
	"math/big"
	"testing"
)

func TestGQ(t *testing.T) {
	ctx := &GqCtx{
		Public: &PublicPart{
			e: big.NewInt(0),
			n: big.NewInt(0),
		},
		private: &privatePart{
			p:       big.NewInt(0),
			q:       big.NewInt(0),
			pqEuler: big.NewInt(0),
		},
	}

	clientCtx := &GqClientCtx{
		x: big.NewInt(0),
		Y: big.NewInt(0),
		k: big.NewInt(0),
		r: big.NewInt(0),
		s: big.NewInt(0),
	}

	authCtx := &GqAuthenticatorCtx{
		a: big.NewInt(0),
	}

	if err := ctx.Init(); err != nil {
		t.Error("Failed to init ctx", err)
	}

	if err := clientCtx.Step1(ctx.Public); err != nil {
		t.Error("Failed to perform Step1", err)
	}

	if err := authCtx.Step2(ctx.Public); err != nil {
		t.Error("Failed to perfrom Step2", err)
	}

	if err := clientCtx.Step3(ctx.Public, authCtx.a); err != nil {
		t.Error("Failed tpo perform Step3", err)
	}

	if !authCtx.Step4(ctx.Public, clientCtx) {
		t.Error("Authenticatioin has failed. That was not supposed to happen.")
	}
}
