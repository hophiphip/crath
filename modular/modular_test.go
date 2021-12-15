package modular

import (
	"fmt"
	"math/big"
	"testing"
)

type testBinaryModulo struct {
	result, val, pow, mod *big.Int
}

type testCase struct {
	result  []*big.Int
	a, b, m *big.Int
}

var tests = []testCase{}

var testsBinaryModulo = []testBinaryModulo{
	// result -- value -- power -- module
	{big.NewInt(4), big.NewInt(5), big.NewInt(14), big.NewInt(7)},
	{big.NewInt(4), big.NewInt(24), big.NewInt(16), big.NewInt(7)},
	{big.NewInt(72), big.NewInt(4), big.NewInt(114), big.NewInt(92)},
	{big.NewInt(8), big.NewInt(6), big.NewInt(75), big.NewInt(26)},
	{big.NewInt(1), big.NewInt(3), big.NewInt(75), big.NewInt(26)},
	{big.NewInt(4), big.NewInt(99), big.NewInt(402), big.NewInt(101)},
	{big.NewInt(14), big.NewInt(4298), big.NewInt(33), big.NewInt(17)},
}

type ModularInverseCase struct {
	element, module, result *big.Int
}

var testsModularInverse = []ModularInverseCase{
	{
		element: big.NewInt(2),
		module:  big.NewInt(7),
		result:  big.NewInt(4),
	},
	{
		element: big.NewInt(123),
		module:  big.NewInt(4567),
		result:  big.NewInt(854),
	},
	{
		element: big.NewInt(12),
		module:  big.NewInt(119),
		result:  big.NewInt(10),
	},
}

func TestBinaryModulo(t *testing.T) {
	fmt.Println("BinaryModulo test")

	for _, test := range testsBinaryModulo {
		res := BinaryModulo(test.val, test.pow, test.mod)
		if test.result.Cmp(res) != 0 {
			t.Error(
				"For", test.val, test.pow, test.mod,
				"Expected", test.result,
				"Got", res,
			)
		}
	}
}

func TestModularInverse(t *testing.T) {
	buf := big.NewInt(0)
	for _, testCase := range testsModularInverse {
		buf.Set(ModularInverse(testCase.element, testCase.module))
		if testCase.result.Cmp(buf) != 0 {
			t.Error(
				"For element: ", testCase.element, " and module: ", testCase.module,
				"expected: ", testCase.result, " ,but got: ", buf,
			)
		}
	}
}
