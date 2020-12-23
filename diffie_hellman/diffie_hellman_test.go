package diffie_hellman

import "testing"

var (
	testCount = 100
)

func TestExchange(t *testing.T) {
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

	for iter := 0; iter < testCount; iter++ {
		err := public.Init()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate public keys",
			)
		}

		err = client1.InitInternal()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate private keys",
			)
		}

		err = client2.InitInternal()
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate private keys",
			)
		}

		err = client1.InitShared(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to init shared keys",
			)
		}

		err = client2.InitShared(public)
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
			err = client1.SharedSecret.Exchange(client2.SharedSecret)
		}

		err = client1.Finalize(public)
		if err != nil {
			t.Error("For iteration", iter,
				"failed to generate final key",
			)
		}

		err = client2.Finalize(public)
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
