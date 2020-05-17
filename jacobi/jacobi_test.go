package jacobi

import (
	"fmt"
	"math/big"
	"testing"
)

type testValue struct {
	a, m   *big.Int
	jacobi int64
}

var testValues = []testValue{
	{big.NewInt(24), big.NewInt(47), 1},
	{big.NewInt(42), big.NewInt(47), 1},
	{big.NewInt(29), big.NewInt(53), 1},
	{big.NewInt(46), big.NewInt(97), -1},
	{big.NewInt(45), big.NewInt(101), 1},
	{big.NewInt(69), big.NewInt(107), 1},
	{big.NewInt(180), big.NewInt(307), -1},
	{big.NewInt(328), big.NewInt(421), 1},
	{big.NewInt(572), big.NewInt(971), -1},
	{big.NewInt(582), big.NewInt(983), -1},
	{big.NewInt(524), big.NewInt(727), 1},
	{big.NewInt(724), big.NewInt(1031), 1},
}

func TestJacobi(t *testing.T) {
	fmt.Println("Test for Jacobi")

	for _, test := range testValues {
		res := Jacobi(test.a, test.m)

		if res != test.jacobi {
			t.Error("For", test.a, "and", test.m,
				"expected", test.jacobi,
				"got", res,
			)
		}
	}
}
