package bitcommit

import (
	"testing"
)

var (
	testSamples = []string{
		"aaaaa",
		"hello",
		"asfsd,glsdgs",
		"fsdsdfsdsfsdf",
	}
)

func TestSchnorrECDSASignature(t *testing.T) {
	schnorrCtx := NewCtx()

	for i, message := range testSamples {
		signature := schnorrCtx.Sign(message)
		derivedPublicKey := PublicKey(message, signature)

		if !schnorrCtx.ComparePublic(derivedPublicKey) {
			t.Error("For iteration", i, "sample:", message,
				"public key comparison has failed",
				"Public Key:", schnorrCtx.PublicKey,
				"Derived Public Key:", derivedPublicKey,
			)
		}

		if !Verify(message, signature, schnorrCtx.PublicKey) {
			t.Error("For iteration", i, "sample:", message,
				"verification has failed",
			)
		}
	}
}
