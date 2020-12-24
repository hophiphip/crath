package messagepack

import (
	"math/big"
)

func StringToBigInt(message string) *big.Int {
	return big.NewInt(0).SetBytes([]byte(message))
}

func BigIntToString(bigNum *big.Int) string {
	return string(bigNum.Bytes())
}
