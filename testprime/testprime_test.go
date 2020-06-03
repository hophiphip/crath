package testprime

import (
	"fmt"
	"math/big"
	"testing"
)

type tests struct {
	n      *big.Int
	result bool
}

type testsBig struct {
	n      string
	result bool
}

var testValues = []tests{
	{big.NewInt(0), false},
	{big.NewInt(1), false},
	{big.NewInt(2), true},
	{big.NewInt(3), true},
	{big.NewInt(4), false},
	{big.NewInt(5), true},
	{big.NewInt(6), false},
	{big.NewInt(7), true},
	{big.NewInt(8), false},
	{big.NewInt(9), false},
	{big.NewInt(10), false},
	{big.NewInt(11), true},
	{big.NewInt(12), false},
	{big.NewInt(13), true},
	{big.NewInt(14), false},
}

var testValuesBig = []testsBig{
	{"19238129048039483389012830912839012830928390128390123890128390128309128418917489127389122", false},
	{"284840938409084394823948230948029480923840338423241194819238129048039483389012830912839012830928390128390123890128390128309128418917489127389122", false},
	{"2848409384090843948239482309480294809238403384232411948192381290480394833890128309128390128309283901283901238901283901283091284189174891273891", false},
}

func TestPrimetest(t *testing.T) {
	fmt.Println("Test for 'SimplePrimetest'")

	for _, test := range testValues {
		res, err := SimpleTest(test.n)

		if err != nil {
			if test.n.Cmp(big.NewInt(0)) == 0 || test.n.Cmp(big.NewInt(1)) == 0 {
				continue
			} else {
				t.Error("For", test.n,
					"unexpected error", err,
				)
			}
		} else {
			if res != test.result {
				t.Error("For", test.n,
					"result must be", test.result,
					"but was", res,
				)
			}
		}

	}

	fmt.Println("Test for 'SimplePrimetest' with big values")

	bigTest := big.NewInt(0)
	for _, test := range testValuesBig {
		// In this case 'err' is bool
		bigTest, err := bigTest.SetString(test.n, 10)
		if !err {
			t.Error("For", test.n,
				"can't convert to big.Int",
			)
		}

		// In this case 'errt' is type of error
		res, errt := SimpleTest(bigTest)
		if errt != nil {
			t.Error("For", test.n,
				"function returned error", errt,
			)
		}

		if res != test.result {
			t.Error("For", test.n,
				"expected", test.result,
				"got", res,
			)
		}
	}
}
