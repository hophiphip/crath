package fractions

import (
	"fmt"
	"math/big"
	"testing"
)

type testType struct {
	num, den *big.Int
	fraction []*big.Int
}

var tests = []testType{
	{big.NewInt(84), big.NewInt(30), []*big.Int{big.NewInt(2), big.NewInt(1), big.NewInt(4)}},
	{big.NewInt(76), big.NewInt(53), []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(2)}},
	{big.NewInt(1394), big.NewInt(1107), []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(1), big.NewInt(6)}},
	{big.NewInt(129), big.NewInt(53), []*big.Int{big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(2)}},
	{big.NewInt(117), big.NewInt(34), []*big.Int{big.NewInt(3), big.NewInt(2), big.NewInt(3), big.NewInt(1), big.NewInt(3)}},
	{big.NewInt(97), big.NewInt(61), []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(1), big.NewInt(2)}},
	{big.NewInt(198), big.NewInt(53), []*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(2), big.NewInt(1), big.NewInt(3), big.NewInt(1), big.NewInt(2)}},
	{big.NewInt(189), big.NewInt(134), []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(2), big.NewInt(3)}},
	{big.NewInt(132), big.NewInt(38), []*big.Int{big.NewInt(3), big.NewInt(2), big.NewInt(9)}},
	{big.NewInt(1961), big.NewInt(1537), []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(2)}},
	{big.NewInt(1376), big.NewInt(1505), []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(10), big.NewInt(1), big.NewInt(2)}},
}

func TestContinuous(t *testing.T) {
	fmt.Println("TestContinuous:")

	for _, test := range tests {
		fmt.Println("Testing pair:", test.num, test.den)

		result := Continuous(big.NewInt(0).Set(test.num), big.NewInt(0).Set(test.den))

		if len(result) == len(test.fraction) {
			for i, r := range result {
				if r.Cmp(test.fraction[i]) != 0 {
					t.Error(
						"For", test.num,
						"/", test.den,
						"Expected", test.fraction,
						"but got", result,
					)
				}
			}
		} else {
			t.Error(
				"For", test.num,
				"/", test.den,
				"Expected", test.fraction,
				"but got", result,
			)
		}
	}
}

func TestNormal(t *testing.T) {
	fmt.Println("TestNormal:")

	for _, test := range tests {
		fmt.Println("Testing fraction:", test.fraction)

		num, den := Normal(test.fraction)

		if num.Div(num, test.num).Cmp(den.Div(den, test.den)) != 0 {
			t.Error(
				"For", test.fraction,
				"Expected", test.num,
				"and", test.den,
				"but got", num,
				"and", den,
			)
		}
	}
}
