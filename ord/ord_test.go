package ord

import (
	"math/big"
	"testing"
)

type testCase struct {
	groupElement *big.Int
	module       *big.Int
	order        *big.Int // expected runction result
}

var tests = []testCase{
	{big.NewInt(2), big.NewInt(11), big.NewInt(10)},
	{big.NewInt(3), big.NewInt(11), big.NewInt(5)}, // 10
}

func TestOrd(t *testing.T) {
	for _, test := range tests {
		val := Ord(test.groupElement, test.module)
		if val.Cmp(test.order) != 0 {
			t.Error(
				"For element: ", test.groupElement.String(),
				" and module: ", test.module.String(),
				" expected: ", test.order.String(),
				" got: ", val.String(),
			)
		}
	}
}
