package mulfunc

import (
	"fmt"
	"math/big"
	"testing"
)

type TestPair struct {
	input, resultCount, resultSum, resultEuler *big.Int
}

type FactorPair struct {
	value  *big.Int
	result []*big.Int
}

var tests = []TestPair{
	//{big.NewInt(1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9), big.NewInt(169), big.NewInt(255 * 121 * 6 * 8), big.NewInt(64 * 54 * 4 * 6)},
	{big.NewInt(144), big.NewInt(15), big.NewInt(403), big.NewInt(48)},
	{big.NewInt(5), big.NewInt(2), big.NewInt(6), big.NewInt(4)},
	{big.NewInt(7), big.NewInt(2), big.NewInt(8), big.NewInt(6)},
	{big.NewInt(11), big.NewInt(10), big.NewInt(12), big.NewInt(10)},
	{big.NewInt(13), big.NewInt(2), big.NewInt(14), big.NewInt(12)},
	{big.NewInt(17), big.NewInt(2), big.NewInt(8), big.NewInt(16)},
	{big.NewInt(1234), big.NewInt(4), big.NewInt(1854), big.NewInt(616)},
	{big.NewInt(12345), big.NewInt(8), big.NewInt(19776), big.NewInt(6576)},
	{big.NewInt(1234567), big.NewInt(4), big.NewInt(1244416), big.NewInt(1224720)},
}

var testsFactorization = []FactorPair{
	{big.NewInt(1), []*big.Int{big.NewInt(1)}},
	{big.NewInt(2), []*big.Int{big.NewInt(2)}},
	{big.NewInt(3), []*big.Int{big.NewInt(3)}},
	{big.NewInt(4), []*big.Int{big.NewInt(2), big.NewInt(2)}},
}

var testsUniquefact = []FactorPair{
	{big.NewInt(1), []*big.Int{big.NewInt(1)}},
	{big.NewInt(2), []*big.Int{big.NewInt(2)}},
	{big.NewInt(3), []*big.Int{big.NewInt(3)}},
	{big.NewInt(4), []*big.Int{big.NewInt(2)}},
	{big.NewInt(1 * 2 * 3 * 4 * 5 * 6), []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)}},
	{big.NewInt(7623), []*big.Int{big.NewInt(3), big.NewInt(7), big.NewInt(11)}},
	{big.NewInt(1768), []*big.Int{big.NewInt(2), big.NewInt(13), big.NewInt(17)}},
	{big.NewInt(495), []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(11)}},
	{big.NewInt(4500), []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5)}},
	{big.NewInt(1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16), []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13)}},
	{big.NewInt(1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17), []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13), big.NewInt(17)}},
	{big.NewInt(1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18), []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13), big.NewInt(17)}},
}

func TestCountFractions(t *testing.T) {
	fmt.Println("TestCountFractions:")

	for _, pair := range tests {
		fmt.Println("	Testing input:", pair.input)

		ret := CountFractions(pair.input)
		if ret.Cmp(pair.resultCount) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.resultCount,
				"got", ret,
			)
		}
	}
}

func TestSumFractions(t *testing.T) {
	fmt.Println("TestSumFractions:")

	for _, pair := range tests {
		fmt.Println("	Testing input:", pair.input)

		ret := SumFractions(pair.input)
		if ret.Cmp(pair.resultSum) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.resultSum,
				"got", ret,
			)
		}
	}
}

func TestEuler(t *testing.T) {
	fmt.Println("TestEuler:")

	for _, pair := range tests {
		fmt.Println("	Testing input:", pair.input)

		ret := Euler(pair.input)
		if ret.Cmp(pair.resultEuler) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.resultEuler,
				"got", ret,
			)
		}
	}

}

func TestFactorization(t *testing.T) {
	fmt.Println("TestFactorization:")

	for _, test := range testsFactorization {
		fmt.Println("	Testing input:", test.value)

		result := Factorization(test.value)

		if len(result) != len(test.result) {
			t.Error(
				"For", test.value,
				"expected", test.result,
				"got", result,
			)
		}

		for i, val := range result {
			if test.result[i].Cmp(val) != 0 {
				t.Error(
					"For", test.value,
					"expected", test.result,
					"got", result,
				)
			}
		}
	}
}

func TestUniquefactor(t *testing.T) {
	fmt.Println("TestUniquefactor:")

	for _, test := range testsUniquefact {
		fmt.Println("	Testing input:", test.value)

		result := Uniquefactor(test.value)

		if len(result) != len(test.result) {
			t.Error(
				"For", test.value,
				"expected", test.result,
				"got", result,
			)
		}

		for i, val := range result {
			if test.result[i].Cmp(val) != 0 {
				t.Error(
					"For", test.value,
					"expected", test.result,
					"got", result,
				)
			}
		}
	}
}
