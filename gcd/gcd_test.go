package gcd

import (
	"fmt"
	"math/big"
	"testing"
)

type testPair struct {
	input  [2]*big.Int
	result *big.Int
}

// TODO: Add more tests
var tests = []testPair{
	{[2]*big.Int{big.NewInt(0), big.NewInt(0)}, big.NewInt(-1)},
	{[2]*big.Int{big.NewInt(0), big.NewInt(123)}, big.NewInt(123)},
	{[2]*big.Int{big.NewInt(123), big.NewInt(0)}, big.NewInt(123)},
	{[2]*big.Int{big.NewInt(123), big.NewInt(123)}, big.NewInt(123)},
	{[2]*big.Int{big.NewInt(-123), big.NewInt(123)}, big.NewInt(123)},
	{[2]*big.Int{big.NewInt(123), big.NewInt(-123)}, big.NewInt(123)},
	{[2]*big.Int{big.NewInt(4), big.NewInt(2)}, big.NewInt(2)},
	{[2]*big.Int{big.NewInt(2), big.NewInt(4)}, big.NewInt(2)},
	{[2]*big.Int{big.NewInt(3), big.NewInt(136)}, big.NewInt(1)},
	{[2]*big.Int{big.NewInt(222), big.NewInt(122)}, big.NewInt(2)},
	{[2]*big.Int{big.NewInt(4444), big.NewInt(2222)}, big.NewInt(2222)},
	{[2]*big.Int{big.NewInt(1060), big.NewInt(212)}, big.NewInt(212)},
	{[2]*big.Int{big.NewInt(40), big.NewInt(20)}, big.NewInt(20)},
}

func TestGcd(t *testing.T) {
	fmt.Println("TestGcd:")

	for _, pair := range tests {
		fmt.Println("	Testing pair:", pair.input)

		ret := Gcd(pair.input[0], pair.input[1])
		if ret.Cmp(pair.result) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", ret,
			)
		}
	}
}

func TestBgcd(t *testing.T) {
	fmt.Println("TestBgcd:")

	for _, pair := range tests {
		fmt.Println("	Testing pair:", pair.input)

		ret := Bgcd(pair.input[0], pair.input[1])
		if ret.Cmp(pair.result) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", ret,
			)
		}
	}
}
