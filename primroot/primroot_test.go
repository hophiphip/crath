package primroot

import (
	"fmt"
	"math/big"
	"sort"
	"testing"
)

type testValues struct {
	input   *big.Int
	results []*big.Int
}

type testValuesPn struct {
	p, n    *big.Int
	results []*big.Int
}

// just a help function to conver int array to big.Int array
func simpleSliceOfBigInt(values []int) []*big.Int {
	res := []*big.Int{}
	for _, val := range values {
		res = append(res, big.NewInt(int64(val)))
	}
	return res
}

var testsP = []testValues{
	//{big.NewInt(2), []*big.Int{big.NewInt(1)}}, <-- it affects everything, as 2 is not checked by default
	{big.NewInt(3), []*big.Int{big.NewInt(2)}},
	{big.NewInt(5), []*big.Int{big.NewInt(2), big.NewInt(3)}},
	{big.NewInt(7), []*big.Int{big.NewInt(3), big.NewInt(5)}},
	{big.NewInt(11), simpleSliceOfBigInt([]int{2, 6, 7, 8})},
	{big.NewInt(13), simpleSliceOfBigInt([]int{2, 6, 7, 11})},
	{big.NewInt(17), simpleSliceOfBigInt([]int{3, 5, 6, 7, 10, 11, 12, 14})},
	{big.NewInt(19), simpleSliceOfBigInt([]int{2, 3, 10, 13, 14, 15})},
	{big.NewInt(23), simpleSliceOfBigInt([]int{5, 7, 10, 11, 14, 15, 17, 19, 20, 21})},
	{big.NewInt(29), simpleSliceOfBigInt([]int{2, 3, 8, 10, 11, 14, 15, 18, 19, 21, 26, 27})},
	{big.NewInt(31), simpleSliceOfBigInt([]int{3, 11, 12, 13, 17, 21, 22, 24})},
	{big.NewInt(37), simpleSliceOfBigInt([]int{2, 5, 13, 15, 17, 18, 19, 20, 22, 24, 32, 35})},
	{big.NewInt(41), simpleSliceOfBigInt([]int{6, 7, 11, 12, 13, 15, 17, 19, 22, 24, 26, 28, 29, 30, 34, 35})},
	{big.NewInt(43), simpleSliceOfBigInt([]int{3, 5, 12, 18, 19, 20, 26, 28, 29, 30, 33, 34})},
	{big.NewInt(47), simpleSliceOfBigInt([]int{5, 10, 11, 13, 15, 19, 20, 22, 23, 26, 29, 30, 31, 33, 35, 38, 39, 40, 41, 43, 44, 45})},
	{big.NewInt(53), simpleSliceOfBigInt([]int{2, 3, 5, 8, 12, 14, 18, 19, 20, 21, 22, 26, 27, 31, 32, 33, 34, 35, 39, 41, 45, 48, 50, 51})},
	{big.NewInt(59), simpleSliceOfBigInt([]int{2, 6, 8, 10, 11, 13, 14, 18, 23, 24, 30, 31, 32, 33, 34, 37, 38, 39, 40, 42, 43, 44, 47, 50, 52, 54, 55, 56})},
	{big.NewInt(61), simpleSliceOfBigInt([]int{2, 6, 7, 10, 17, 18, 26, 30, 31, 35, 43, 44, 51, 54, 55, 59})},
	{big.NewInt(67), simpleSliceOfBigInt([]int{2, 7, 11, 12, 13, 18, 20, 28, 31, 32, 34, 41, 44, 46, 48, 50, 51, 57, 61, 63})},
	{big.NewInt(71), simpleSliceOfBigInt([]int{7, 11, 13, 21, 22, 28, 31, 33, 35, 42, 44, 47, 52, 53, 55, 56, 59, 61, 62, 63, 65, 67, 68, 69})},
}

var testsPn = []testValuesPn{
	{big.NewInt(3), big.NewInt(2), simpleSliceOfBigInt([]int{2, 5})},
	{big.NewInt(3), big.NewInt(3), simpleSliceOfBigInt([]int{2, 5, 11, 14, 20, 23})},
	{big.NewInt(5), big.NewInt(2), simpleSliceOfBigInt([]int{2, 3, 8, 12, 13, 17, 22, 23})},
}

func TestPrimrootp(t *testing.T) {
	fmt.Println("Test for TestPrimrootp")

	for _, test := range testsP {
		res := Primrootp(test.input)

		// A fancy way to sort a slice
		// ..because of using random big.Int
		//   the order of the result may differ
		sort.Slice(res, func(i, j int) bool {
			return res[i].Cmp(res[j]) < 0
		})

		if len(res) != len(test.results) {
			t.Error(
				"For", test.input,
				"expected", test.results,
				"got", res,
			)
		}

		for i, r := range res {
			if r.Cmp(test.results[i]) != 0 {
				t.Error(
					"For", test.input,
					"expected", test.results,
					"got", res,
				)
			}
		}
	}
}

func TestPrimrootpn(t *testing.T) {
	fmt.Println("Test for TestPrimrootpn")

	for _, test := range testsPn {
		res := Primrootpn(test.p, test.n)

		// A fancy way to sort a slice
		// ..because of using random big.Int
		//   the order of the result may differ
		sort.Slice(res, func(i, j int) bool {
			return res[i].Cmp(res[j]) < 0
		})

		if len(res) != len(test.results) {
			t.Error(
				"For", test.p, "and", test.n,
				"expected", test.results,
				"got", res,
			)
		}

		for i, r := range res {
			if r.Cmp(test.results[i]) != 0 {
				t.Error(
					"For", test.p, "and", test.n,
					"expected", test.results,
					"got", res,
				)
			}
		}
	}
}
