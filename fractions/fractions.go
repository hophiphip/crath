package fractions

import (
	"math/big"
)

func _Calculate(num *big.Int, den *big.Int) (quo *big.Int) {
	quo, div, mod := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	if den.Cmp(big.NewInt(0)) != 0 {
		quo, _ = div.DivMod(num, den, mod)
		num.Set(den)
		den.Set(mod)
	}

	return quo
}

// Continuous converts normal fraction to continuous fraction
// num / den => [a , b, c, d, ...]
func Continuous(num *big.Int, den *big.Int) (fraction []*big.Int) {
	zero := big.NewInt(0)

	for den.Cmp(zero) != 0 {
		fraction = append(fraction, _Calculate(num, den))
	}

	return fraction
}

// Normal converts continuous fraction to normal fraction
// [a, b, c, d, ...] => num / den
func Normal(fraction []*big.Int) (num *big.Int, den *big.Int) {
	box := [2][2]*big.Int{
		{big.NewInt(1), fraction[0]},
		{big.NewInt(0), big.NewInt(1)},
	}
	tmp1, tmp2, num, den := big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)

	if len(fraction) < 2 {
		return num.Set(fraction[0]), den.SetUint64(1)
	}

	for _, f := range fraction {
		tmp1.Add(box[0][0], tmp1.Mul(box[0][1], f))
		tmp2.Add(box[1][0], tmp2.Mul(box[1][1], f))
		box[0][0].Set(box[0][1])
		box[0][1].Set(tmp1)
		box[1][0].Set(box[1][1])
		box[1][1].Set(tmp2)
	}

	return num.Set(box[0][1]), den.Set(box[1][1])
}
