package asymenc

import "testing"

var (
	// Test sample
	mesText = []string{
		"aah",
		"bobby",
		"coccyx",
		"diddled",
		"epee",
		"faff",
		"Shh",
	}

	// Initialize new RSA context every iteration
	initEveryTime = false
)

func TestRsaEncryption(t *testing.T) {
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
		if initEveryTime || i == 0 {
			if ctx.Init() != nil {
				t.Error("Failed to init context on iteration:", i)
			}
		}

		publicKey := ctx.Public.GetPublicKey()
		privateKey := ctx.secret.GetPrivateKey()

		encrypted := publicKey.Encrypt(text)
		decrypted := privateKey.Decrypt(encrypted, publicKey)

		if decrypted != text {
			t.Error(
				"For iteration", i,
				"decrypted message:", decrypted,
				"not equal to:", text,
			)
		}
	}
}
