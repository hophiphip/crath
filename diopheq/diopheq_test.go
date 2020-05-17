package diopheq

import (
	"fmt"
	"math/big"
	"testing"
)

type testPair struct {
	input  [2]*big.Int
	result [2]*big.Int
}

type testPairGen struct {
	input  [3]*big.Int
	result [4]*big.Int
}

// TODO: Add more tests
var tests = []testPair{
	{[2]*big.Int{big.NewInt(45), big.NewInt(29)}, [2]*big.Int{big.NewInt(-9), big.NewInt(14)}},
	{[2]*big.Int{big.NewInt(46), big.NewInt(-17)}, [2]*big.Int{big.NewInt(-7), big.NewInt(-19)}},
	{[2]*big.Int{big.NewInt(48), big.NewInt(-17)}, [2]*big.Int{big.NewInt(-6), big.NewInt(-17)}},
	{[2]*big.Int{big.NewInt(41), big.NewInt(23)}, [2]*big.Int{big.NewInt(9), big.NewInt(-16)}},
	{[2]*big.Int{big.NewInt(44), big.NewInt(27)}, [2]*big.Int{big.NewInt(8), big.NewInt(-13)}},
	{[2]*big.Int{big.NewInt(43), big.NewInt(28)}, [2]*big.Int{big.NewInt(-13), big.NewInt(20)}},
}

var testsGen = []testPairGen{
	{[3]*big.Int{big.NewInt(43), big.NewInt(18), big.NewInt(4)}, [4]*big.Int{big.NewInt(-20), big.NewInt(-18), big.NewInt(48), big.NewInt(43)}},
	{[3]*big.Int{big.NewInt(43), big.NewInt(28), big.NewInt(5)}, [4]*big.Int{big.NewInt(-65), big.NewInt(-28), big.NewInt(100), big.NewInt(43)}},
}

func TestSimple(t *testing.T) {
	fmt.Println("TestSimple:")

	for _, pair := range tests {
		fmt.Println("	Testing pair:", pair.input)

		retx, rety := Simple(pair.input[0], pair.input[1])
		if retx.Cmp(pair.result[0]) != 0 || rety.Cmp(pair.result[1]) != 0 {
			t.Error(
				"For", pair.input,
				"expected", pair.result,
				"got", retx,
				"and", rety,
			)
		}
	}
}

func TestGeneral(t *testing.T) {
	fmt.Println("TestGeneral:")

	// Why did I name it 'trio' ? I dunno...
	for _, trio := range testsGen {
		fmt.Println("	Testing input:", trio.input)

		xl, xr, yl, yr := General(trio.input[0], trio.input[1], trio.input[2])
		if xl.Cmp(trio.result[0]) != 0 || xr.Cmp(trio.result[1]) != 0 || yl.Cmp(trio.result[2]) != 0 || yr.Cmp(trio.result[3]) != 0 {
			t.Error(
				"For", trio.input,
				"expected", trio.result,
				"got", xl,
				",", xr,
				",", yl,
				"and", yr,
			)
		}
	}
}
