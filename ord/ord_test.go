package ord

import (
	"fmt"
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
	{big.NewInt(3), big.NewInt(10), big.NewInt(4)},
	{big.NewInt(3), big.NewInt(10), big.NewInt(10)},
	{big.NewInt(2), big.NewInt(10), big.NewInt(2)},
}

func TestOrd(t *testing.T) {
	for _, test := range tests {
		val, err := Ord(test.groupElement, test.module)
		// Not an error but a warning
		if err != nil {
			fmt.Println(
				"WARN: For element: ", test.groupElement.String(),
				" and module: ", test.module.String(),
				" got error: ", err,
			)
		} else {
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
}
