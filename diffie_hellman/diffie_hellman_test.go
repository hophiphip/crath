package diffie_hellman

import "testing"

var (
	testCount = 10
)

func TestExchange(t *testing.T) {
	for iter := 0; iter < testCount; iter++ {
		var (
			public  *PublicValues
			client1 *PrivateValues
			client2 *PrivateValues
		)

		public = &PublicValues{
			SmallPrimeBitSize: 256,
			SmallPrime:        nil,
			BigPrimeBitSize:   2048,
			BigPrime:          nil,
		}

		client1 = &PrivateValues{
			initialSecretBitSize: 1024,
			initialSecret:        nil,
			finalSecret:          nil,
			SharedSecret:         nil,
		}

		client2 = &PrivateValues{
			initialSecretBitSize: 1024,
			initialSecret:        nil,
			finalSecret:          nil,
			SharedSecret:         nil,
		}

		err := public.InitPublicStep1()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate public keys",
			)
		}

		err = client1.InitClientGenerateInternalStep1()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate private keys",
			)
		}

		err = client2.InitClientGenerateInternalStep1()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate private keys",
			)
		}

		err = client1.InitClientGenerateSharedStep2(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to init shared keys",
			)
		}

		err = client2.InitClientGenerateSharedStep2(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to init shared keys",
			)
		}

		if client1.SharedSecret == nil || client2.SharedSecret == nil {
			t.Error("For iteration", iter,
				"shared secrets are 'nil'",
			)
		} else {
			err = client1.SharedSecret.ExchangeClientSharedStep3(client2.SharedSecret)
		}

		err = client1.InitClientGenerateFinalStep4(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate final key",
			)
		}

		err = client2.InitClientGenerateFinalStep4(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate final key",
			)
		}

		if client1.finalSecret.Cmp(client2.finalSecret) != 0 {
			t.Error("For iteration", iter,
				"keys differ",
				"client1 key", client1.finalSecret,
				"client2 key", client2.finalSecret,
			)
		}
	}
}
