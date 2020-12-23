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

// Init - init public known values: q, n
// Step1
func (publicVal *PublicValues) Init() error {
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
// Step1
func (privateVal *PrivateValues) InitInternal() error {
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
// Step2
func (privateVal *PrivateValues) InitShared(publicVal *PublicValues) error {
	privateVal.SharedSecret.PersonalSecret = big.NewInt(0).Exp(
		publicVal.SmallPrime,
		privateVal.initialSecret,
		publicVal.BigPrime)

	return nil
}

// Exchange - exchange shared secrets between 2 clients
// Step3
func (sharedOne *SharedSecret) Exchange(sharedTwo *SharedSecret) error {
	sharedOne.ExchangedSecret = big.NewInt(0).Set(sharedTwo.PersonalSecret)
	sharedTwo.ExchangedSecret = big.NewInt(0).Set(sharedOne.PersonalSecret)

	return nil
}

// Finalize - init final secret keys
// that are equal between 2 clients
// Step4
func (privateVal *PrivateValues) Finalize(publicVal *PublicValues) error {
	privateVal.finalSecret = big.NewInt(0).Exp(
		privateVal.SharedSecret.ExchangedSecret,
		privateVal.initialSecret,
		publicVal.BigPrime)

	return nil
}
