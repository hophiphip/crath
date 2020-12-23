package diffie_hellman

import (
	"crypto/rand"
	"math.io/crath/primegen"
	"math/big"
)

// NOTE: key is just a big number(big.Int)

// PublicValues stores public accessible keys
type PublicValues struct {
	SmallPrimeBitSize int `default:"256"`
	SmallPrime        *big.Int

	BigPrimeBitSize int `default:"2048"`
	BigPrime        *big.Int
}

// SharedSecret - stores value
// ready to be exported
type SharedSecret struct {
	PersonalSecret  *big.Int
	ExchangedSecret *big.Int
}

// PrivateValues stores private key info
type PrivateValues struct {
	initialSecretBitSize int `default:"1024"`
	initialSecret        *big.Int

	finalSecret *big.Int

	SharedSecret *SharedSecret
}

// InitPublicStep1 - init public known values: q, n
func (publicVal *PublicValues) InitPublicStep1() error {
	p, err := primegen.Primegen(rand.Reader, publicVal.SmallPrimeBitSize)
	if err != nil {
		return err
	}
	publicVal.SmallPrime = p

	p, err = primegen.Primegen(rand.Reader, publicVal.BigPrimeBitSize)
	if err != nil {
		return err
	}
	publicVal.BigPrime = p

	return nil
}

// InitClientInternalStep1 - init private values
// 		For example: Alice's - a
//		..or Bob's - b
func (privateVal *PrivateValues) InitClientGenerateInternalStep1() error {
	p, err := primegen.Primegen(rand.Reader, privateVal.initialSecretBitSize)
	if err != nil {
		return err
	}
	privateVal.initialSecret = p

	privateVal.SharedSecret = &SharedSecret{
		PersonalSecret:  nil,
		ExchangedSecret: nil,
	}

	return nil
}

// InitClientSharedStep2 - shared secret generation
func (privateVal *PrivateValues) InitClientGenerateSharedStep2(publicVal *PublicValues) error {
	privateVal.SharedSecret.PersonalSecret = big.NewInt(0).Exp(
		publicVal.SmallPrime,
		privateVal.initialSecret,
		publicVal.BigPrime)

	return nil
}

// ExchangeClientSharedStep3 - exchange shared secrets between 2 clients
func (sharedOne *SharedSecret) ExchangeClientSharedStep3(sharedTwo *SharedSecret) error {
	sharedOne.ExchangedSecret = big.NewInt(0).Set(sharedTwo.PersonalSecret)
	sharedTwo.ExchangedSecret = big.NewInt(0).Set(sharedOne.PersonalSecret)

	return nil
}

// InitClientGenerateFinalStep4 - init final secret keys
// that are equal between 2 clients
func (privateVal *PrivateValues) InitClientGenerateFinalStep4(publicVal *PublicValues) error {
	privateVal.finalSecret = big.NewInt(0).Exp(
		privateVal.SharedSecret.ExchangedSecret,
		privateVal.initialSecret,
		publicVal.BigPrime)

	return nil
}
