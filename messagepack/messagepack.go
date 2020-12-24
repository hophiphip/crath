package messagepack

import (
	"math/big"
)

func MessageToBigInt(message string) *big.Int {
	return big.NewInt(0).SetBytes([]byte(message))
}

func BigIntToMessage(bigNum *big.Int) string {
	return string(bigNum.Bytes())
}
