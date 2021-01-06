package coinflip

import (
	"math/big"
	"testing"
)

func TestCoinflip(t *testing.T) {
	ctx := &Context{
		q:  big.NewInt(0),
		qp: big.NewInt(0),
		p:  big.NewInt(0),
		a:  big.NewInt(0),
		g:  big.NewInt(0),
	}

	if err := ctx.Init(); err != nil {
		t.Error("Init failed", err)
	}

	init1Res, err := InitiatorStep1(ctx)
	if err != nil {
		t.Error("Failed to initiate step 1", err)
	}

	res1Res, err := ReceiverStep1(ctx, init1Res)
	if err != nil {
		t.Error("Failed to perform step 1 on receiver side", err)
	}

	b, err := InitiatorStep2()
	if err != nil {
		t.Error("Failed to perform step 2 on receiver side", err)
	}

	if !InitiatorStepCheck(ctx, res1Res, init1Res) {
		t.Error("Verification has failed. That not supposed to happen")
	}

	// Final test. Just to verify that XOR is ok.
	result := big.NewInt(0).Xor(res1Res.aBit, b)
	if result.Cmp(big.NewInt(0)) != 0 && result.Cmp(big.NewInt(1)) != 0 {
		t.Error("Result is faulty", result.String())
	}

}
