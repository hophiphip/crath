package ord

import (
	"math/big"

	"math.io/crath/mulfunc"
)

// Pair - pair of factorization elements (factor and degree)
type Pair struct {
	first, second *big.Int
}

// coolFactorization: [2 2 3] => [(2, 2), (3, 1)]
// Returns array of pairs: (factor, degree)
func coolFactorization(value *big.Int) []Pair {
	factorization := mulfunc.Factorization(value)
	temp := make(map[string]*big.Int)
	var res []Pair

	one := big.NewInt(1)
	for _, fact := range factorization {
		if _, exists := temp[fact.String()]; exists {
			temp[fact.String()].Add(temp[fact.String()], one)
		} else {
			temp[fact.String()] = big.NewInt(1)
			res = append(res, Pair{first: fact, second: big.NewInt(0)})
		}
	}

	for _, fact := range res {
		if _, exists := temp[fact.first.String()]; exists {
			fact.second.Set(temp[fact.first.String()])
		}
	}

	return res
}

// Ord - Element order function
// g - element
// m - module
func Ord(g, m *big.Int) *big.Int {
	N := mulfunc.Euler(m)
	n := coolFactorization(N)

	d := N

	one := big.NewInt(1)
	for _, factor := range n {
		d.Div(d, big.NewInt(0).Exp(factor.first, factor.second, m))
		b := big.NewInt(0).Exp(g, d, m)

		for one.Cmp(b) != 0 {
			b.Exp(b, factor.first, m)
			d.Mul(d, factor.first)
		}
	}

	return d
}
