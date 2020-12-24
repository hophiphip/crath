package asymenc

import "testing"

var (
	mesText = []string{
		"aah",
		"bobby",
		"coccyx",
		"diddled",
		"epee",
		"faff",
		"Shh",
	}
)

func TestRsa(t *testing.T) {
	ctx := RsaContext{
		secret: &RsaSecret{
			secretPrimeP:  nil,
			secretPrimeQ:  nil,
			secretBitSize: 2048,
			modSolution:   nil,
			pairEuler:     nil,
		},
		Public: &RsaPublic{
			FixedCoPrime:       nil,
			PairMultiplication: nil,
		},
	}

	for i, text := range mesText {
		if ctx.Init() != nil {
			t.Error("Failed to init context on iteration:", i)
		}

		smes, err := ctx.Sign(text)
		if err != nil {
			t.Error("Failed to sign message:", text, ",on iteration:", i)
		}

		if smes == nil {
			t.Error("Signed message is nil:", text, ",on iteration:", i)
		} else {
			if !smes.Verify(ctx.Public) {
				t.Error("Signature rejected on message:", text, ",on iteration", i,
					"sign:")
			}
		}
	}
}
